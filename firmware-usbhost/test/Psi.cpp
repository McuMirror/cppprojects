#include "config.h"

#ifdef PSI
#include <stdint.h>
#include <string.h>
#include <stdlib.h>
#include "Psi.h"
#include "lib/battery/Battery.h"

#include "common.h"
#include "common/memory/stm32/include/ExternalEeprom.h"
#include "common/utils/include/Version.h"
#include "common/timer/stm32/include/SystemTimer.h"
#include "common/timer/stm32/include/RealTimeControl.h"
#include "common/utils/include/StringBuilder.h"
#include "common/memory/stm32/include/ExternalEeprom.h"
#include "common/uart/include/UartStream.h"
#include "common/i2c/I2C.h"
#include "common/beeper/include/Gramophone.h"

#include "lib/adc/adc.h"
#include "lib/utils/stm32/buttons.h"
#include "lib/utils/stm32/Beeper.h"
#include "lib/mma7660/MMA7660.h"
#include "lib/network/Enc28j60.h"
#include "lib/sd/SD.h"
#include "lib/battery/Battery.h"
#include "lib/utils/stm32/HardwareUartForwardController.h"

#include <sys/types.h>
#include "cmsis_boot/stm32f4xx_conf.h"

extern void InitRele();

#define PSI_VersionH	0
#define PSI_VersionL	3

const char *sep1 = "-----";
const char *sep2 = "*****";
const char *success = "�������";
const char *error = "������";

void psiMain() {
	SystemCoreClockUpdate();

	Psi psi;
	psi.start();

	while(1) {
		IWDG_ReloadCounter();
	}
}

Psi::Psi(bool fastMode) {
	hwErrors = 0;
	this->fastMode = fastMode;

	stat = new StatStorage;
	initDebugUart();
	initTimers();
	buttons = new Buttons();
	buttons->init(timerEngine);
	led = Led::get();

	__enable_irq();

	hardware = new Gsm::Hardware;
	hardware->init();
	Battery::init();

	// ������������� ����������� �������� ������ UART.
	HardwareUartForwardController::init();
}

void Psi::initDebugUart() {
	debugUart = Uart::get(DEBUG_UART);
	debugUart->setup(256000, Uart::Parity_None, 0);
	debugUart->setTransmitBufferSize(10240);
	debugUart->setReceiveBufferSize(256);
	LogTargetUart *logTarget = new LogTargetUart(debugUart);
	Logger::get()->registerTarget(logTarget);
}

void Psi::initTimers() {
	this->timerEngine = new TimerEngine();
	Ticker *ticker = Ticker::get();
	ticker->registerConsumer(timerEngine);
}

void Psi::start() {
	*(Logger::get()) << "Firmware: hardware111 " << LOG_VERSION((uint32_t)&HardwareVersion) << ", software " << LOG_VERSION((uint32_t)&SoftwareVersion) << Logger::endl;
	LOG("PSI init...");

	testInit();

	Led::get()->setLed1(0,0)->setLed2(0,0)->setLed3(0,0)->setLed4(0,0);
	Led::get()->setPower(64);

	// ��������� ������ SIM.
	if(hardware->isStatusUp()) {
		printDetails("��������� ������ SIM");
		hardware->pressPowerButton();
		SystemTimer::get()->delay_ms(1500);
		hardware->releasePowerButton();
	}

//==================== ��������������� ����� ====================
//	testRfid();
//	testMMA7660();
	testRTC();
	testUarts();

#if (HW_VERSION >= HW_3_0_2)
	testExternalEeprom();
#else
#error "HW_VERSION must be defined in project settings"
#endif

//===================== ������������� ����� =====================
	testBeeper();
	testLeds();
	testSim900();
#if (HW_VERSION == HW_3_0_0)
	testEthernet();
	testSD();
#elif (HW_VERSION >= HW_3_0_2)
	testRele();
	testEthernet();
	testSD();
#else
#error "HW_VERSION must be defined in project settings"
#endif

//========================= ��������� ===========================
	StringBuilder str;
	str << "����� ���������: ";
	str << (hwErrors > 0 ? error : success) << ", ������: " << hwErrors;
	printMessage(str.getString());
}

void Psi::printHeader(const char *head) {
	LOG("\r\n" << sep1 << " " << head << " " << sep1);
}

void Psi::printResult(bool result, const char *info) {
	const char *res = result ? success : error;
	if (info)
		LOG(sep1 << " " << res << " " << sep1 << " -> " << info);
	else
		LOG(sep1 << " " << res << " " << sep1);

	if (!result) hwErrors++;
}

void Psi::printDetails(const char *info) {
	LOG(sep1 << " " << info);
}

void Psi::printMessage(const char *info) {
	LOG(sep2 << " " << info << " " << sep2);
}

bool Psi::printDialogAndWaitAnswer(const char *info, const char *waitString) {
	if(fastMode) return true;

	LOG(sep2 << " " << info << " " << sep2);
	String str;
	BYTE c = 0;

	do {
		while (debugUart->isEmptyReceiveBuffer()) {
			debugUart->execute();
			timerEngine->execute();
			IWDG_ReloadCounter();
		}
		c = debugUart->receive();
		if(c && c != '\n' && c != '\r') {
			str << (char) c;
		}
	} while ((c != '\n') && (c != '\r'));

	debugUart->clear();

	return (str == waitString);
}

bool Psi::sendToAndWaitAnswer(Uart *uart, const char *sendString, const char *waitString, uint32_t timeoutMls, bool printAnswer) {
	UartStream stream(uart);
	stream.send(sendString);

	String str;
	while (!str.contains(waitString) && timeoutMls) {
		while (!uart->isEmptyReceiveBuffer()) str << (char) uart->receive();

		uart->execute();
		IWDG_ReloadCounter();
		timeoutMls--;
		SystemTimer::get()->delay_ms(1);
	};

	if(str.contains(waitString) == false) {
		LOG("send: " << sendString);
		LOG("recv: " << str.getString());
		return false;
	}

	if (printAnswer) {
		LOG("recv: " << str.getString());
	}

	return true;
}

void Psi::testInit() {
	printHeader("���������� � �����");
	printDetails("1. ��������� ��� � ����� ��������� ������� SIM-�����");
	printDetails("2. ��������� �������� Rx � Tx ������� DEX (�������-�������)");

	printResult(printDialogAndWaitAnswer("�� ������ � ������ ������? -> ����", ""));
}

bool Psi::testRTC()
{
	printHeader("���� RTC");

	DateTime date;
	RealTimeStm32 realTime;
	realTime.getDateTime(&date);
	LOG(LOG_DATETIME(date));

	// ���������� �����
	RealTimeControl *rtc = RealTimeControl::get();
	uint32_t subSeconds = rtc->getSubSecond();
	// ���� 10 ���
	SystemTimer::get()->delay_ms(10);
	// ����������
	bool result = subSeconds != rtc->getSubSecond();
	printResult(result);

	return result;
}

void Psi::waitingPressButton(uint16_t buttonId)
{
	while(buttons->isPressed(buttonId) == false)
	{
		IWDG_ReloadCounter();
		timerEngine->execute();
	}

	printResult(buttons->isPressed(buttonId));
}

void Psi::testButtons() {
	printHeader("���� ������");

	// � ������������� 0.1 ��� ���������� 10 ��� ��� �������������� ������� �� ����
	SystemTimer::get()->delay_ms(20);

	bool error = false;
	buttons->check();
	if (buttons->isPressed(Buttons::Event_Button1)) { error = true; printResult(false, "������� ������ 1"); }
	if (buttons->isPressed(Buttons::Event_Button2)) { error = true; printResult(false, "������� ������ 2"); }
	if (buttons->isPressed(Buttons::Event_Button3)) { error = true; printResult(false, "������� ������ 3"); }
	if (error) return;

	printMessage("������� �� ����� ������ 1 ...");
	waitingPressButton(Buttons::Event_Button1);

	printMessage("������� �� ����� ������ 2 ...");
	waitingPressButton(Buttons::Event_Button2);

	printMessage("������� �� ����� ������ 3 ...");
	waitingPressButton(Buttons::Event_Button3);

	delete buttons;
}

void Psi::testLeds()
{
	printHeader("���� �����������");
	led->setPower(64);

	// �������� ��� ��������� �����������. ��������� ��� ���������� ������ �����, ��� ��������� � ����������������� ��������� SIM ������.
	const int delay = 150;

	SystemTimer::get()->delay_ms(250);

	// �������� ������� ����������
	led->setLed1(1, 0); SystemTimer::get()->delay_ms(delay);
	led->setLed2(1, 0); SystemTimer::get()->delay_ms(delay);
	led->setLed3(1, 0); SystemTimer::get()->delay_ms(delay);
	led->setLed4(1, 0); SystemTimer::get()->delay_ms(delay);
	printResult(printDialogAndWaitAnswer("������� ����� ������� �����������?", "4"));

	// �������� ������� ����������
	led->setLed1(0, 1); SystemTimer::get()->delay_ms(delay);
	led->setLed2(0, 1); SystemTimer::get()->delay_ms(delay);
	led->setLed3(0, 1); SystemTimer::get()->delay_ms(delay);
	led->setLed4(0, 1); SystemTimer::get()->delay_ms(delay);
	printResult(printDialogAndWaitAnswer("������� ����� ������� �����������?", "4"));

	// ����� ��� ����������
	led->setLed1(0, 0); SystemTimer::get()->delay_ms(delay);
	led->setLed2(0, 0); SystemTimer::get()->delay_ms(delay);
	led->setLed3(0, 0); SystemTimer::get()->delay_ms(delay);
	led->setLed4(0, 0); SystemTimer::get()->delay_ms(delay);
	printResult(printDialogAndWaitAnswer("������� ����� ����� �����������?", "0"));
}

void Psi::testUarts()
{
	testUart_MDB();
	testUart_DEX();
}

bool Psi::testUart_MDB()
{
	printHeader("�������� MDB ����������");

	return testUart_MDB(19200);
}

bool Psi::testUart_MDB(int speed)
{
	StringBuilder str;
	str << "�������� MDB �� ��������: " << speed;
	printHeader(str.getString());

	const int bufferSize = 1024;

	Uart *mdb_slave = Uart::get(MDB_UART);
	mdb_slave->setup(speed, Uart::Parity_None, 0, Uart::Mode_9Bit, GPIO_OType_OD);
	mdb_slave->setReceiveBufferSize(bufferSize);
	mdb_slave->setTransmitBufferSize(bufferSize);

	Uart *mdb_master = Uart::get(MDB_MASTER_UART);
	mdb_master->setup(speed, Uart::Parity_None, 0, Uart::Mode_9Bit, GPIO_OType_OD);
	mdb_master->setReceiveBufferSize(bufferSize);
	mdb_master->setTransmitBufferSize(bufferSize);

	const int timeout = 500;

	IWDG_ReloadCounter();

	SystemTimer::get()->delay_ms(timeout);

	mdb_slave->clear();
	mdb_master->clear();

	printDetails("��������� Master TX ---> Slave RX");

//	HardwareUartForwardController::start();

	for (uint8_t i = 0; i < 128; i++)
	{
		mdb_master->send(0);
		mdb_master->send(i);
		mdb_master->send(1);
		mdb_master->send(i);
	}

	SystemTimer::get()->delay_ms(timeout);

	bool result = true;
	str = "������ ����: ";
	for (uint8_t i = 0; i < 128; i++)
	{
		bool en = !mdb_slave->isEmptyReceiveBuffer();
		uint8_t v1 =  mdb_slave->receive();
		uint8_t v2 =  mdb_slave->receive();
		uint8_t v3 =  mdb_slave->receive();
		uint8_t v4 =  mdb_slave->receive();

		result = (en && v1 == 0 && v2 == i && v3 == 1 && v4 == i);
		if (!result)
		{
			str.clear();
			if (!en)
			{
				str << "��� ������, i: " << i;
				break;
			}

			str << "���������: 0 " << i << " 1 " << i << ", ��������: " << v1 << " " << v2 << " " << v3 << " "<< v4;
			break;
		}
	}
	if (result)
		printResult (result);
	else
	{
		printResult (result, str.getString());
		return false;
	}

	SystemTimer::get()->delay_ms(timeout);

	mdb_slave->clear();
	mdb_master->clear();

	printDetails("��������� Slave TX ---> Master RX");

	for (uint8_t i = 0; i < 64; i++)
	{
		mdb_slave->send(0);
		mdb_slave->send(i);
		mdb_slave->send(1);
		mdb_slave->send(i);
	}
	SystemTimer::get()->delay_ms(timeout);

	result = true;

	for (uint8_t i = 0; i < 64; i++)
	{
		bool en = !mdb_master->isEmptyReceiveBuffer();
		uint8_t v1 =  mdb_master->receive();
		uint8_t v2 =  mdb_master->receive();
		uint8_t v3 =  mdb_master->receive();
		uint8_t v4 =  mdb_master->receive();

		result = (en && v1 == 0 && v2 == i && v3 == 1 && v4 == i);
		if (!result)
		{
			str.clear();
			if (!en)
			{
				str << "��� ������, i: " << i;
				break;
			}

			str << "���������: 0 " << i << " 1 " << i << ", ��������: " << v1 << " " << v2 << " " << v3 << " "<< v4;
			break;
		}
	}

	if(result) {
		printResult (result);
	} else {
		printResult (result, str.getString());
		return false;
	}

	SystemTimer::get()->delay_ms(timeout);

	printDetails("��������� Master <---> Slave");
	mdb_slave->clear();
	mdb_master->clear();

	for (uint8_t i = 0; i < 64; i++) {
		mdb_slave->send((uint8_t) 0x00);
		mdb_slave->send(i);
		mdb_master->send((uint8_t) 0x00);
		mdb_master->send(i);
	}

	SystemTimer::get()->delay_ms(timeout);

	for (uint8_t i = 0; i < 64; i++) {
		if (mdb_master->isEmptyReceiveBuffer() || mdb_master->receive() != 0x00 || mdb_master->receive() != i) {
			str.clear();
			str << "Slave->Master, i: " << i;
			printResult(false, str.getString());
			return false;
		}
		if (mdb_slave->isEmptyReceiveBuffer() ||mdb_slave->receive() != 0x00 || mdb_slave->receive() != i) {
			str.clear();
			str << "Master->Slave, i: " << i;
			printResult(false, str.getString());
			return false;
		}
	}

//	// ��������� ���������� �������� ������
//	HardwareUartForwardController::stop();
//
//	printDetails("��������� ���� Master <-x-> Slave");
//
//	mdb_slave->clear();
//	mdb_master->clear();
//
//	for (uint8_t i = 0; i < 16; i++) {
//		mdb_slave->send((uint8_t) 0x01);
//		mdb_slave->send(i);
//		mdb_master->send((uint8_t) 0x01);
//		mdb_master->send(i);
//	}
//
//	SystemTimer::get()->delay_ms(timeout);
//
//	for (uint8_t i = 0; i < 16; i++) {
//		if (!mdb_master->isEmptyReceiveBuffer()) {
//			str.clear();
//			str << "���� Slave ---> Master ���������� ������, i: " << i;
//			printResult(false, str.getString());
//			return false;
//		}
//		if (!mdb_slave->isEmptyReceiveBuffer()) {
//			str.clear();
//			str << "����� Master ---> Slave ���������� ������, i: " << i;
//			printResult(false, str.getString());
//			return false;
//		}
//	}

	printResult(true);
	return true;
}

bool Psi::testUart_DEX() {
	String receiveString;
	StringBuilder testString = "Dex test string";
	int timeout = 200;

	printHeader("�������� DEX ����������. ����� ��������� �������� Rx � Tx �� �������");

	Uart *dex_uart = Uart::get(DEX_UART);
	dex_uart->setup(115200, Uart::Parity_None, 0);
	dex_uart->clear();
	UartStream stream(dex_uart);
	stream.send(testString);
	while (receiveString.getLen() < testString.getLen() && --timeout) {
		dex_uart->execute();
		if (!dex_uart->isEmptyReceiveBuffer()) receiveString << (char) dex_uart->receive();
		SystemTimer::get()->delay_ms(1);
	}

	bool result = timeout && receiveString == testString;

	printResult(result, receiveString.getString());
	printResult(timeout);

	return result;
}

bool Psi::testSim900() {
	printHeader("���� ������ SIM900");

	Uart *sim_uart = Uart::get(SIM900_UART);
	sim_uart->setup(115200, Uart::Parity_None, 0);
	sim_uart->setTransmitBufferSize(2048);
	sim_uart->setReceiveBufferSize(2048);

	if (hardware->isStatusUp()) {
		printDetails("������ ��� �������");
	} else {
		printDetails("�������� ������");
		hardware->pressPowerButton();
		SystemTimer::get()->delay_ms(1500);
		hardware->releasePowerButton();
		// ���� 7 ������
		for (int t = 0; t < 7; t++) {
			IWDG_ReloadCounter();
			SystemTimer::get()->delay_ms(1000);
		}
	}

	// ���������� � ������ ������� AT
	printResult(sendToAndWaitAnswer(sim_uart, "AT\r\n", "OK", 1000, false), "AT");

	// ������� IMEI
	bool result = sendToAndWaitAnswer(sim_uart, "AT+GSN\r\n", "OK", 1000, true);
	printResult(result, "AT+GSN");

	if (fastMode) return result;

	// ����������� ����������� � ����
	printResult(sendToAndWaitAnswer(sim_uart, "AT+CREG=1\r\n", "OK", 1000, false), "AT+CREG=1");
	// ���� 5 ������
	for (int t = 0; t < 5; t++) {
		IWDG_ReloadCounter();
		SystemTimer::get()->delay_ms(1000);
	}

	// ����������� ����������� � ����
	printResult(sendToAndWaitAnswer(sim_uart, "AT+CREG?\r\n", "+CREG: 1,1", 2000, false), "AT+CREG?");

	return result;
}

class GramophoneMaster : public EventObserver {
public:
	GramophoneMaster() : complete(false) {}
	virtual void proc(Event *event) { complete = true; }
	bool isComplete() { return complete; }
private:
	bool complete;
};

void Psi::testBeeper() {
	printHeader("���� ��������� �������");

	MelodyElochkaHalf melody;

	Gramophone *gramophone = new Gramophone(Beeper::get(), timerEngine);
	GramophoneMaster master;
	gramophone->play(&melody, &master);
	while(master.isComplete() == false) {
		IWDG_ReloadCounter();
		timerEngine->execute();
	}

	if (!fastMode)
		printResult(printDialogAndWaitAnswer("�� ������� �������� ������? (1=��, 0=���)", "1"));
}

void Psi::testRfid() {
	printHeader("���� ������ RFID �� ����������");
//	RFID *rfid = RFID::get();
//
//	StringBuilder str(128);
//
//	if (!rfid->PCD_PerformSelfTest()) {
//		printResult(false, "������ �� ��������");
//		return;
//	}
//
//	rfid->PCD_Init();
//
//	str << "������ RFID: ";
//	str.addHex(rfid->PCD_ReadRegister(MFRC522::VersionReg));
//	printDetails(str.getString());
//
//	printMessage("��������� � ������ RFID ����� ...");
//
//	while (!rfid->PICC_IsNewCardPresent()) IWDG_ReloadCounter();
//
//	printDetails("����� ����������");
//
//	if (rfid->PICC_ReadCardSerial()) {
//		// Show some details of the PICC (that is: the tag/card)
//		str = "����� �����: ";
//		for (int i = 0; i < rfid->uid.size; i++)
//			str.addHex(rfid->uid.uidByte[i]);
//
//		printDetails(str.getString());
//
//		MFRC522::PICC_Type piccType = rfid->PICC_GetType(rfid->uid.sak);
//		str = "��� �����: ";
//		str << rfid->PICC_GetTypeName(piccType);
//		printDetails(str.getString());
//
//		// Halt PICC
//		rfid->PICC_HaltA();
//		// Stop encryption on PCD
//		rfid->PCD_StopCrypto1();
//		printResult(true);
//		return;
//	}
//
//	printResult(false);
}

void Psi::testMMA7660() {
	printHeader("���� ������ 3-� ������� ������� ���������: MMA7660");

	MMA7660 *mma = MMA7660::get();
	SystemTimer::get()->delay_ms(250);
	MMA7660::Values values = mma->asyncReadValues();

	if (!mma->isInitialized() || (values.x == 0 && values.y == 0 && values.z == 0)) printResult(false, "������ �� ��������");
	else printResult(true);
}

bool Psi::testEthernet() {
	uint8_t hwaddr1[6]  = { 0x20, 0x89, 0x84, 0x6A, 0x96, 0x00 };

	printHeader("���� ������ Ethernet");

#if (HW_VERSION < HW_3_2_0)
	ENC28J60 *enc = new ENC28J60(SPI::get(SPI_3));
#elif (HW_VERSION >= HW_3_2_0)
	ENC28J60 *enc = new ENC28J60(SPI::get(SPI_2));
#else
	#error "HW_VERSION must be defined in project settings"
#endif

	enc->init(hwaddr1);

	bool result;
	if(!fastMode) {
		printDialogAndWaitAnswer("��������� ������ Ethernet � ������� ����", "");
		result = !enc->enc28j60linkup();
		printResult(result);

		if(result) {
			printDialogAndWaitAnswer("���������� ������ Ethernet � ������� ����", "");

			result = result && enc->enc28j60linkup();
			printResult(result);
		}
	} else {
		result = enc->checkAnswer();
	}

	delete enc;

	return result;
	//TODO: ��������, ����� ����� ������� ���� �� ping.
}

void Psi::testSD() {
	printHeader("���� ������ SD. ����������� ����� SD ��������, � �������� �������� FAT16");

	SD *sd = SD::get();

	printDialogAndWaitAnswer("������ SD �������� � ������� ����", "");
	printResult(!sd->hasCard());
	printDialogAndWaitAnswer("�������� SD �������� � ������� ����", "");
	printResult(sd->hasCard());
	if(sd->hasCard()) {
		printDetails("����� �� �������� �������� ����");
		printResult(sd->test());
		// FIXME: SD, �������� ���� ������ ����������� �����.
	}
}

void Psi::testRele() {
	printHeader("���� ����");

	// ������������� ����
	InitRele();
	printDialogAndWaitAnswer("���� ���������? (1=��, 0=���)", "1");
	RELE1_ON;
	printDialogAndWaitAnswer("���� ��������? (1=��, 0=���)", "1");
	RELE1_OFF;
}

bool Psi::testExternalEeprom() {
	String str;
	str << "���� ���������� ������, ������: " << (EE_MAX_SIZE/1024) << " ��";
	printHeader(str.getString());

	I2C *i2c = I2C::get(I2C_3);
	i2c->setStatStorage(stat);
	ExternalEeprom mem(i2c, EE_ADDRESS, EE_MAX_SIZE, EE_PAGE_SIZE);

	// TODO: ������� �� �������� ��������� ������ ������������ ����������!
	const int max = mem.getMaxSize() / 4; // ��������� 1/4 �� ������ ��� ��������� ��������.
	const int div = 4;
	bool result = true;

	uint8_t *source = new uint8_t[max/div];
	uint8_t *readData = new uint8_t[max/div];

	for(int i = 0; i < max/div; i++) {
		source[i] = i;
		readData[i] = 0xff;
	}

	mem.setAddress(0);

	if(mem.write(source, mem.getPageSize()) != MemoryResult_Ok) { result = false; }
	printResult(result, "����� �������� ������");
	mem.setAddress(0);

	if(mem.read(readData, mem.getPageSize()) != MemoryResult_Ok) { result = false; }
	printResult(result, "������");

	if(memcmp(source, readData, mem.getPageSize()) != 0) { result = false; }
	printResult(result, "����������");


	mem.setAddress(0);
	for(int i = 0; i < div; i++) {
		if(mem.write(source, max/div) != MemoryResult_Ok) { result = false; }
		printResult(result, "����� �� ��� ��������");
		IWDG_ReloadCounter();
	}

	mem.setAddress(0);
	for(int i = 0; i < div; i++) {
		if(mem.read(readData, max/div) != MemoryResult_Ok) { result = false; }
		printResult(result, "������");

		if(memcmp(source, readData, max/div) != 0) { result = false; }
		printResult(result, "����������");
		IWDG_ReloadCounter();
	}

	mem.setAddress(4);
	if(mem.write(source, 3) != MemoryResult_Ok) { result = false; }
	printResult(result, "����� �� ������ 4");
	mem.setAddress(1);

	if(mem.read(readData, 7) != MemoryResult_Ok) { result = false; }
	printResult(result, "������");
	const uint8_t randomRead[] = {1,2,3,0,1,2,7};

	if(memcmp(randomRead, readData, 7) != 0) { result = false; }
	printResult(result, "����������");

	mem.setAddress(640);
	source[0] = 54;

	if(mem.write(source, 1) != MemoryResult_Ok) { result = false; }
	printResult(result, "����� �� ������ 640, 1 ����");
	mem.setAddress(640);

	if(mem.read(readData, 1) != MemoryResult_Ok) { result = false; }
	printResult(result, "������ 1 ����");

	if(memcmp(source, readData, 1) != 0) { result = false; }
	printResult(result, "����������");
	readData[0] = 0;
	mem.setAddress(640);

	if(mem.read(readData, 1) != MemoryResult_Ok) { result = false; }
	printResult(result, "������ 1 ���� ��� ���");

	if(memcmp(source, readData, 1) != 0) { result = false; }
	printResult(result, "���������� ��� ��� ");

	for(int i = 0; i < max/div; i++) {
		source[i] = i+10;
		readData[i] = 0xff;
	}
	mem.setAddress(440);

	if(mem.write(source, 80) != MemoryResult_Ok) { result = false; }
	printResult(result, "����� �� ������ 440, 80 ����");
	mem.setAddress(440);

	if(mem.read(readData, 80) != MemoryResult_Ok) { result = false; }
	printResult(result, "������");

	if(memcmp(source, readData, 80) != 0) { result = false; }
	printResult(result, "����������");

	delete []source;
	delete []readData;

	return result;
}

#endif
