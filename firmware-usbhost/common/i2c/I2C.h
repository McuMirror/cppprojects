#ifndef __I2C_H
#define __I2C_H

/*
 * ���: ���������
 * ��������: ������, ��� ������ � I2C � ���������� � ����������� ������ � ��������� ������ ����� DMA
 *
 * ������ ������ � DMA ����: http://letanphuc.net/2014/06/stm32-mpu6050-dma-i2c/
 * ������� ������ ����������� ������ ������ I2C ����� ����������: https://electronix.ru/forum/index.php?showtopic=93745
 */

#include "defines.h"
#include "common/utils/include/Event.h"
#include "common/config/include/StatStorage.h"

enum enI2C {
	I2C_1,
	I2C_2,
	I2C_3
};

class I2C
{
private:
	enum enI2C index;
	volatile bool busy;
	volatile uint32_t lastEvent;
	StatStorage *stat;
	I2C_InitTypeDef i2cDef;
	int retryPause;


	void DMA_Config();
	void DMA_RxConfig();
	void DMA_TxConfig();
	void DMA_Config(DMA_Stream_TypeDef* dmaTypeDef, uint8_t *data, int dataSize);
	void I2C_Config(void);
	void NVIC_Config(void);

	bool startTransmission(const uint8_t deviceAddress, const uint8_t transmissionDirection);
	void stopTransmission();
	bool writeData(uint8_t data);
	void reinit();
	uint32_t getOperationTimeDiff(uint32_t lastTimeMs);
	bool checkAndClearErrorFlags(uint32_t flag);


	bool syncWriteDataImpl(const uint8_t deviceAddress, const uint16_t writeAddress, const uint8_t addressSize, uint8_t *data, const int len);
	bool syncReadDataImpl(const uint8_t deviceAddress, const uint16_t readAddress, const uint8_t addressSize, uint8_t *data, const int len);

public:
	static I2C *get(enI2C index);

	I2C_TypeDef *i2c;

	EventObserver *eventObserver;

	enum Mode {
		Sync,
		Async,
		Indefined
	};

	struct AsyncData {
//		// ����� ����������
//		uint8_t deviceAddress;
//		// ����� ��������
//		uint8_t rwAddress;
//		// ��������� �� ������
//		uint8_t *data;
//		// ������ ������
//		int dataLen;
//		// ����� ������ ?
//		bool readMode;
		// ���, ��������, ��� ������ ��������� � �������� ������/������ ������
		volatile bool isProcess;
		// ���� ������
		bool error;
//		// ������� ����� ����� � ����������, ��������������� ����������. ������ ��� ������������ ������ ������ - ��������� � ����� ������� I2C_EVENT_MASTER_MODE_SELECT
//		int index;
		// ��� �������, ������� ����� ������������� �� ��������� ����������
		uint16_t eventType;
	} asyncData;

	I2C(enI2C index, uint32_t speed = I2C_DEFAULT_SPEED);
	void setObserver(EventObserver *eventObserver);
	void setMode(enum Mode mode);
	bool syncWriteData(const uint8_t deviceAddress, const uint16_t writeAddress, const uint8_t addressSize, uint8_t *data, const int len, uint32_t timeout = 0);
	bool syncReadData(const uint8_t deviceAddress, const uint16_t readAddress, const uint8_t addressSize,  uint8_t *data, const int len, uint32_t timeout = 0);

	enum Mode getMode();

	bool asyncWriteData(const uint8_t deviceAddress, const uint16_t writeAddress, const uint8_t addressSize, uint8_t *data, const int len, const uint16_t eventType);
	bool asyncReadData(const uint8_t deviceAddress, const uint16_t writeAddress, const uint8_t addressSize, uint8_t *data, const int len, const uint16_t eventType);
	volatile bool isAsyncCompleted() {
		return !asyncData.isProcess;
	}
	void setStatStorage(StatStorage *stat);
	void setRetryPause(int mls);

private:
	enum Mode mode;

};

#endif
