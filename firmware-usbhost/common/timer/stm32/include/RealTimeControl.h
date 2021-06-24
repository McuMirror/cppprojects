#ifndef __RealTimeControl_H
#define __RealTimeControl_H

#include "timer/include/RealTime.h"

#include "stm32f4xx_rtc.h"
/*
 * ����������� VCC_BAT ��� RTC:
 *
 * V_BAT: 3.0 V
 * ��� �������������, ��� �������: 0.02 uA
 * � ��������, ��� �������������: 0.10 uA
 * � ��������, ����� �������������: 0.02 uA
 * ��� �������, ����� �������������: 1.0 uA
 *
 * ������ ������� ������, ��� ������� ��������� 35 ���/�
 * 35000/1 = 35000 � = 35000/24/365 = 4 ����.
 *
 * */

class RealTimeControl {
public:
	enum Mode {
		ExternalClock, // ���� �������� �� ���������.
		InternalClock // ���� �������� ������ �� ��������� �������.
	};

	static RealTimeControl *get();
	static RealTimeControl *get(enum Mode mode);
	RealTimeControl(enum Mode mode);

	bool setTime(RTC_TimeTypeDef *time);
	bool setDate(RTC_DateTypeDef *date);
	void getTime(RTC_TimeTypeDef *time);
	void getDate(RTC_DateTypeDef *date);
	uint32_t getTotalSeconds();
	uint32_t getSubSecond();
	// ���������� ���-�� ������ � 1970 ����. ��������: https://stm32f4-discovery.net/2014/07/library-19-use-internal-rtc-on-stm32f4xx-devices/
	uint32_t getUnixTimestamp();

private:
	void init(enum Mode mode);
	void setDefaultDateTime();
};

class RealTimeStm32 : public RealTimeInterface {
public:
	void init();
	virtual bool setDateTime(DateTime *datetime);
	virtual void getDateTime(DateTime *datetime);
	virtual uint32_t getUnixTimestamp();
};

#endif
