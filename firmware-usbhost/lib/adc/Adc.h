#ifndef __ADC_H
#define __ADC_H

/*
 * ���: ���������
 * ��������: ��� ������ ������������� ����������� ����������� �������������� ������� ��� � ������ � �������� 8 � DMA.
 * �� ������� ADC_ExternalTrigConv_T8_TRGO ���������� ������� ������� ��� ���� ������ (Rank). ��������� ���� ���������
 * �� 1 ������ DMA ���������� � ������ adcResultValues. ������� �����������.
 *
 *
 * ����������� � ���:
 * VCC3  : A1 (1k/1k)
 * V��24 : A4 (10k/1k) (10k/620)
 * VCC5  : A5 (2k/1k)
 * VCC_CHARGE : A6 (1k/1k)
 * VCC_EXT_TEMP: B1 -
 */

#include "defines.h"
#include "Adc_ext.h"
#include "common/utils/include/FastAverage.h"

#define ADC_REG_CHANNELS 8

class Adc
{
private:
	volatile int index;
	uint16_t regularValues[ADC_REG_CHANNELS];

	FastAverage<uint16_t> **avg;


	void DMA_Config();
	void ADC_Config(void);
	void TIM8_Config();
	void NVIC_Configuration(void);

public:
	static Adc *get();

	Adc();
	void placeRegularValues();

	enum Rank {
		VCC_3, 			// ADC_Channel_1,
		VCC_24, 		// ADC_Channel_4,
		VCC_5, 			// ADC_Channel_5,
		VCC_BAT1,		// ADC_Channel_6,
		VCC_EXT_TEMP,	// ADC_Channel_9,
		VCC_REF,		// ADC_Channel_Vref
		TEMP_SENSOR,	// ADC_Channel_TempSensor,
		VCC_BAT2		// ADC_Channel_Vbat
	};

	uint32_t read(Rank rank);

	// ���������� �������� ����������� CPU � ��������
	uint32_t getCpuTemp();

	// ���������� ������� �������� ��������� ���������� � ���
	uint32_t getInputVoltage();


private:

};
#endif
