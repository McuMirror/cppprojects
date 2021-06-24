#include "ShtrihmReceiver.h"

#include "fiscal_register/shtrih-m/include/ShtrihM.h"
#include "fiscal_register/shtrih-m/ShtrihmProtocol.h"
#include "logger/include/Logger.h"

/*
��������� ������:
  ���� 0: ������� ������ ��������� STX;
  ���� 1: ����� ��������� (N) � �������� �����. � ����� ��������� �� ���������� ����� 0, LRC � ���� ����;
  ���� 2: ��� ������� ��� ������ � �������� �����;
  ����� 3...(N+1): ���������, ��������� �� ������� (����� �������������);
  ���� N+2 � ����������� ����� ��������� � ���� LRC � ����������� ����������� ��������� (XOR) ���� ������ ��������� (����� ����� 0).

������� ������ ���������� �������:
  ��������: Y:\modem\KKT\paykiosk\������������\�����������_��_2.0(c24).pdf
  �����: ���������� 3 ������������� ��������� ��������� ������ ������������ ������� ������ �� ������� ��

TODO: 1. �������� ��� �������� �� �������. ��������� ����� ��� � ���� ��� �� ������ ���?
*/
ShtrihmReceiver::ShtrihmReceiver(TimerEngine *timers, AbstractUart *uart, ShtrihM *observer) : timers(timers), uart(uart), observer(observer), state(State_Idle) {
	this->recvBuf = new Buffer(SHM_PACKET_MAX_SIZE);
	this->timer = timers->addTimer<ShtrihmReceiver, &ShtrihmReceiver::procTimer>(this);
	this->uart->setReceiveHandler(this);
}

ShtrihmReceiver::~ShtrihmReceiver() {
	this->timers->deleteTimer(this->timer);
	delete recvBuf;
}

bool ShtrihmReceiver::sendPacket(Buffer *data) {
	LOG_DEBUG(LOG_FR, "sendPacket");
	if(state != State_Idle && state != State_NoConnection) {
		LOG_ERROR(LOG_FR, "Wrong state " << state);
		return false;
	}
	this->sendBuf = data;
	gotoStateENQ();
	return true;
}

void ShtrihmReceiver::handle() {
	while(uart->isEmptyReceiveBuffer() == false) {
		LOG_TRACE(LOG_FR, "have data " << state);
		switch(state) {
		case State_ENQ: stateENQRecv(); break;
		case State_Confirm: stateConfirmRecv(); break;
		case State_STX: stateSTXRecv(); break;
		case State_Length: stateLengthRecv(); break;
		case State_Data: stateDataRecv(); break;
		case State_CRC: stateCRCRecv(); break;
		case State_SkipSTX: stateSkipSTXRecv(); break;
		case State_SkipLength: stateSkipLengthRecv(); break;
		case State_SkipData: stateSkipDataRecv(); break;
		case State_SkipCRC: stateSkipCRCRecv(); break;
		default: LOG_ERROR(LOG_FR, "Unwaited data state=" << state << ", byte=" << uart->receive());
		}
	}
}

void ShtrihmReceiver::procTimer() {
	LOG_DEBUG(LOG_FR, "procTimer " << state);
	switch(state) {
		case State_ENQ: stateENQTimeout(); break;
		case State_Confirm: gotoStateENQ(); break;
		case State_STX: gotoStateNoConnection(); break;
		case State_Length: gotoStateNoConnection(); break;
		case State_Data: gotoStateNoConnection(); break;
		case State_CRC: gotoStateNoConnection(); break;
		case State_Answer: stateAnswerTimeout(); break;
		case State_SkipSTX: gotoStateNoConnection(); break;
		case State_SkipLength: gotoStateNoConnection(); break;
		case State_SkipData: gotoStateNoConnection(); break;
		case State_SkipCRC: gotoStateNoConnection(); break;
		case State_SkipPause: gotoStateENQ(); break;
		case State_SkipAnswer: stateSkipAnswerTimeout(); break;
		default: LOG_ERROR(LOG_FR, "Unwaited timeout state=" << state);
	}
}

void ShtrihmReceiver::gotoStateENQ() {
	LOG_DEBUG(LOG_FR, "gotoStateENQ");
	this->tryNumber = 0;
	this->uart->send(ShmControl_ENQ);
	this->timer->start(SHM_ENQ_ANSWER_TIMER);
	this->state = State_ENQ;
}

void ShtrihmReceiver::stateENQRecv() {
	LOG_DEBUG(LOG_FR, "stateENQRecv");
	uint8_t b1 = uart->receive();
	if(b1 == ShmControl_UNKNOWN) {
		LOG_ERROR(LOG_FR, "Ignore 0xFF");
	} else if(b1 == ShmControl_ACK) {
		gotoStateSkipSTX();
	} else if(b1 == ShmControl_NAK) {
		gotoStateConfirm();
	} else {
		LOG_ERROR(LOG_FR, "Unwaited answer " << b1);
		gotoStateNoConnection();
		return;
	}
}

void ShtrihmReceiver::stateENQTimeout() {
	if(tryNumber >= SHM_ENQ_TRY_MAX_NUMBER) {
		gotoStateNoConnection();
		return;
	}
	this->tryNumber++;
	this->uart->send(ShmControl_ENQ);
	this->timer->start(SHM_ENQ_ANSWER_TIMER);
}

void ShtrihmReceiver::gotoStateConfirm() {
	uint8_t dataLen = sendBuf->getLen();
	LOG_DEBUG(LOG_FR, "gotoStateConfirm " << dataLen);
	LOG_TRACE_HEX(LOG_FR, sendBuf->getData(), sendBuf->getLen());

	uart->send(ShmControl_STX);
	uart->send(dataLen);
	for(uint8_t i = 0; i < dataLen; i++) {
		uart->send((*sendBuf)[i]);
	}
	uart->send(ShtrihmCrc::calc(sendBuf));

	timer->start(SHM_REQUEST_CONFIRM_TIMER);
	state = State_Confirm;
}

void ShtrihmReceiver::stateConfirmRecv() {
	LOG_DEBUG(LOG_FR, "stateConfirmRecv");
	uint8_t b1 = uart->receive();
	if(b1 == ShmControl_ACK) {
		gotoStateSTX();
		return;
	} else if(b1 == ShmControl_NAK) {
		//todo: ��� ����������� 10 ������� ����������� �������
		gotoStateNoConnection();
		return;
	} else {
		LOG_ERROR(LOG_FR, "Unwaited answer " << b1);
		gotoStateNoConnection();
		return;
	}
}

void ShtrihmReceiver::gotoStateSTX() {
	LOG_DEBUG(LOG_FR, "gotoStateSTX");
	timer->start(SHM_WAIT_STX_TIMER);
	state = State_STX;
}

void ShtrihmReceiver::stateSTXRecv() {
	LOG_DEBUG(LOG_FR, "stateSTXRecv");
	uint8_t b1 = uart->receive();
	if(b1 == ShmControl_STX) {
		gotoStateLenght();
	} else {
		LOG_ERROR(LOG_FR, "Unwaited answer " << b1);
		gotoStateNoConnection();
		return;
	}
}

void ShtrihmReceiver::gotoStateLenght() {
	LOG_DEBUG(LOG_FR, "gotoStateLenght");
	timer->start(SHM_WAIT_LENGHT_TIMER);
	recvBuf->clear();
	state = State_Length;
}

void ShtrihmReceiver::stateLengthRecv() {
	LOG_DEBUG(LOG_FR, "stateLengthRecv");
	recvLength = uart->receive();
	LOG_DEBUG(LOG_FR, "responseLength=" << recvLength);
	timer->start(SHM_WAIT_DATA_TIMER);
	state = State_Data;
}

void ShtrihmReceiver::stateDataRecv() {
	LOG_DEBUG(LOG_FR, "stateDataRecv");
	recvBuf->addUint8(uart->receive());
	if(recvBuf->getLen() >= recvLength) {
		state = State_CRC;
	}
}

void ShtrihmReceiver::stateCRCRecv() {
	LOG_DEBUG(LOG_FR, "stateCRCRecv");
	uint8_t recvCrc = uart->receive();
	uint8_t calcCrc = ShtrihmCrc::calc(recvBuf);
	if(recvCrc == calcCrc) {
		uart->send(ShmControl_ACK);
		gotoStateAnswer();
		return;
	} else {
		uart->send(ShmControl_NAK);
		LOG_ERROR(LOG_FR, "Wrong CRC: recv=" << recvCrc << ", calc=" << calcCrc);
		gotoStateNoConnection();
		return;
	}
}

void ShtrihmReceiver::gotoStateAnswer() {
	LOG_DEBUG(LOG_FR, "gotoStateAnswer");
	timer->start(SHM_ANSWER_TIMER);
	state = State_Answer;
}

void ShtrihmReceiver::stateAnswerTimeout() {
	LOG_DEBUG(LOG_FR, "stateAnswerTimeout");
	state = State_Idle;
	LOG_DEBUG(LOG_FR, "Recv answer");
	observer->procRecvData(recvBuf->getData(), recvBuf->getLen());
}

void ShtrihmReceiver::gotoStateNoConnection() {
	LOG_DEBUG(LOG_FR, "gotoStateNoConnection");
	this->timer->stop();
	this->state = State_NoConnection;
	this->observer->procRecvError();
}

void ShtrihmReceiver::gotoStateSkipSTX() {
	LOG_DEBUG(LOG_FR, "gotoStateSkipSTX");
	this->timer->start(SHM_SKIP_DATA_TIMER);
	this->state = State_SkipSTX;
}

void ShtrihmReceiver::stateSkipSTXRecv() {
	LOG_DEBUG(LOG_FR, "stateSkipSTXRecv");
	uint8_t b1 = uart->receive();
	if(b1 == ShmControl_STX) {
		gotoStateSkipLenght();
	} else {
		LOG_ERROR(LOG_FR, "Unwaited answer " << b1);
		gotoStateNoConnection();
		return;
	}
}

void ShtrihmReceiver::gotoStateSkipLenght() {
	LOG_DEBUG(LOG_FR, "gotoStateSkipLenght");
	recvBuf->clear();
	timer->start(SHM_WAIT_LENGHT_TIMER);
	state = State_SkipLength;
}

void ShtrihmReceiver::stateSkipLengthRecv() {
	LOG_DEBUG(LOG_FR, "stateSkipLengthRecv");
	recvLength = uart->receive();
	LOG_DEBUG(LOG_FR, "responseLength=" << recvLength);
	timer->start(SHM_WAIT_DATA_TIMER);
	state = State_SkipData;
}

void ShtrihmReceiver::stateSkipDataRecv() {
	LOG_DEBUG(LOG_FR, "stateSkipDataRecv");
	recvBuf->addUint8(uart->receive());
	if(recvBuf->getLen() >= recvLength) {
		state = State_SkipCRC;
	}
}

void ShtrihmReceiver::stateSkipCRCRecv() {
	LOG_DEBUG(LOG_FR, "stateSkipCRCRecv");
	uint8_t crc = uart->receive();
	LOG_TRACE_HEX(LOG_FR, recvBuf->getData(), recvBuf->getLen());
	LOG_TRACE(LOG_FR, "CRC=" << crc);
	gotoStateSkipAnswer();
}

void ShtrihmReceiver::gotoStateSkipAnswer() {
	LOG_DEBUG(LOG_FR, "gotoStateSkipAnswer");
	timer->start(SHM_ANSWER_TIMER);
	state = State_SkipAnswer;
}

void ShtrihmReceiver::stateSkipAnswerTimeout() {
	LOG_DEBUG(LOG_FR, "stateSkipAnswerTimeout");
	uart->send(ShmControl_ACK);
	timer->start(SHM_SKIP_DATA_TIMER);
	state = State_SkipPause;
}
