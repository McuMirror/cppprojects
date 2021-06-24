#include <stdio.h>
#include <string.h>

#include "stm32f4xx_conf.h"
#include "SD.h"
#include "extern/ub_fatfs/stm32_ub_fatfs.h"

#include "common/logger/include/Logger.h"
#include "common/timer/stm32/include/SystemTimer.h"
#include "common/timer/stm32/include/RealTimeControl.h"

#define SD_CARD_DETECT_PORT		GPIOE
#define SD_CARD_DETECT_PIN		GPIO_Pin_7

static SD *INSTANCE	= NULL;

SD *SD::get() {
	if(INSTANCE == NULL) { INSTANCE = new SD(); }
	return INSTANCE;
}

SD::SD() {
	LOG_DEBUG(LOG_SD, "�������� ������ SD");

	GPIO_InitTypeDef gpio;
	GPIO_StructInit(&gpio);
	gpio.GPIO_Pin = SD_CARD_DETECT_PIN;
	gpio.GPIO_Mode = GPIO_Mode_IN;
	gpio.GPIO_Speed = GPIO_Speed_2MHz;
	gpio.GPIO_OType = GPIO_OType_OD;
	gpio.GPIO_PuPd = GPIO_PuPd_UP;
	GPIO_Init(SD_CARD_DETECT_PORT, &gpio);

	// ������������� ������� FATFS
	UB_Fatfs_Init();
}

bool SD::test() {
	PROBE_INIT(B, 12, GPIO_OType_PP);
	PROBE_INIT(E, 15, GPIO_OType_PP);

	PROBE_OFF(B, 12);
	PROBE_OFF(E, 15);


	LOG_DEBUG(LOG_SD, "���� �������� �����");
	FIL myFile;   // ������� �����

	// �������� ������� �����
	if(UB_Fatfs_CheckMedia(MMC_0) != FATFS_OK) {
		LOG_ERROR(LOG_SD, "���� �� �������");
		return false;
	}

	// ������������ �����
	if(UB_Fatfs_Mount(MMC_0) != FATFS_OK) {
		LOG_ERROR(LOG_SD, "����� �� ������������");
		return false;
	}

	// �������� � ���� � �������� ��������
	if(UB_Fatfs_OpenFile(&myFile, "0:/test.txt", F_WR_CLEAR) != FATFS_OK) {
		LOG_ERROR(LOG_SD, "�� ������� ������� ���� 0:/test.txt");
		UB_Fatfs_UnMount(MMC_0);
		return false;
	}

	//�������� ��������� ����� � ����
	UB_Fatfs_WriteString(&myFile, "English string test");
	UB_Fatfs_WriteString(&myFile, "������ � �������� �������");
	UB_Fatfs_WriteString(&myFile, "����� ����� 4 �������� ������");

	const char *buf = "00000000001111111111222222222233333333334444444444555555555566666666667777777777888888888899999999991234567890123456789012345670000000000111111111122222222223333333333444444444455555555556666666666777777777788888888889999999999123456789012345678901234567\r\n";
	int len = strlen(buf);

	uint32_t seconds = RealTimeControl::get()->getUnixTimestamp();
	uint32_t size = 1024 * 1024 * 4;
	uint32_t steps = size / len;

	uint32_t wr;


//	PROBE_OFF(B, 12);
//	PROBE_OFF(E, 15);

	//TODO: SD, 512 ���� ������� �������� �� 2 ���.

	while(steps--) {
//		PROBE_ON(E, 15);
		UB_Fatfs_WriteBlock(&myFile, (unsigned char *) buf, len, &wr);
//		PROBE_OFF(E, 15);

		IWDG_ReloadCounter();
	}

	// (1024*1024*4/256)*18= 294 ���

	seconds = RealTimeControl::get()->getUnixTimestamp() - seconds;

	UB_Fatfs_WriteString(&myFile, "����� �����");

	LOG_DEBUG(LOG_SD, "�������� �����");
	// ������� ����
	UB_Fatfs_CloseFile(&myFile);

	LOG_DEBUG(LOG_SD, "���������� �����");
	// ��������� �����
	UB_Fatfs_UnMount(MMC_0);

	LOG_DEBUG(LOG_SD, "������� ���� test.txt, ������ �����: " << len << ", ������ ����� " << size << " ����, ����� ������ " << seconds << " �������, �������� ������: " << (size / seconds) << " ����/���");
	return true;
}

bool SD::hasCard() {
	return UB_Fatfs_CheckMedia(MMC_0) == FATFS_OK;
}

