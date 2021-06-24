#include "include/TimeMeter.h"

#include "timer/stm32/include/SystemTimer.h"

MicroSecondMeter::MicroSecondMeter() {
	st = SystemTimer::get();
	maxValue = st->getUsMax();
	start();
}

void MicroSecondMeter::start() {
	from = st->getUs();
}

void MicroSecondMeter::meter() {
	to = st->getUs();
	if(from <= to) {
		dif = to - from;
	} else {
		dif = to + (maxValue - from);
		LOG_DEBUG(LOG_TM, ">>>TOOBIG=" << from << "," << to);
	}
	from = to;
}

/*
 * ��������� �������.
 * ���������:
 * cycleNumber - ���������� ���������, ������ ��� ����� �������� ������� ���������� � ����
 * ������������� �� ���������� �����, �� ���������� �����.
 * ������������ ����� ����� �������� start � meter ��� meter � meter �� ������ ���� ������
 * 0xFFFFFFFF / (SystemCoreClock/1000000). ��� ������� ����������� 168��� ��� �������� 25 ������.
 */
void MicroCycleMeter::start(uint32_t cycleNumber) {
	cntNumber = cycleNumber;
	cnt = 0;
	cycleMaxTime = 0;
	globalMaxTime = 0;
	meter.start();
}

/*
 * �������� ����� � ���-���������:
 * cmean - ������� �������������� ������ �����
 * ctime - ����� ����� ���������� ��������� ���������� ������
 * cmax - ������������ ����� ���������� ������ ����� � ����� ����� ���������
 * gmax - ������������ ����� ���������� ������ ����� �� ��� ����� ���������
 */
void MicroCycleMeter::cycle() {
	meter.meter();
	uint32_t dif = meter.getDif();
	if(dif > cycleMaxTime) { cycleMaxTime = dif; }
	if(dif > globalMaxTime) { globalMaxTime = dif; }
	time += dif;
	cnt++;
	if(cnt == cntNumber) {
		LOG_INFO(LOG_TM, ">>>cmean=" << time/cntNumber << ", ctime=" << time << ", cmax=" << cycleMaxTime << ", gmax=" << globalMaxTime);
		cycleMaxTime = 0;
		cnt = 0;
		time = 0;
	}
}


static MicroIntervalMeter *instance = 0;

MicroIntervalMeter *MicroIntervalMeter::get() {
	if(instance) return instance;
	return new MicroIntervalMeter();
}

MicroIntervalMeter::MicroIntervalMeter() {
	instance = this;
	start(0);
}

void MicroIntervalMeter::start(uint32_t maxIntervalSize) {
	this->maxIntervalSize = maxIntervalSize;
	meter.start();
}

bool MicroIntervalMeter::check() {
	meter.meter();
	return (meter.getDif() < maxIntervalSize);
}
