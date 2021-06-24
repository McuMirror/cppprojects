#ifndef SHTRIHMPROTOCOL_H
#define SHTRIHMPROTOCOL_H

#include "utils/include/NetworkProtocol.h"

/*
---------------------
����������� ���������
---------------------
����������� �������� �������:
  ��� ����� � ������ ������� � ����� ��������, ��������� � ���Ż.
  ��� � ����������� �������� �������. � 01.01.1998 � ���������� ��������� 1 ��� �����
  1 ������� (�� 01.01.1998 1 ��� ���� ����� 1 �����).
������ �������� ��������:
  ��� �������� �������� ���������� � �������� �������, ���� �� ������� ������.
  ������ ���������� ����� ������� ����, ��������� ����� ������� ����.
  ��� �������� ���� (3 �����) ������� ��������� ����� (1 ���� � ��),
  ����� ����� (2 ����� � ��), � ��������� � ��� (1 ���� � ��).
  ��� �������� ������� (3 �����) ������ ������ ���������� ���� (1 ���� � ��),
  ����� ������ (2 ����� � ��), � ���������� ���������� ������� (1 ���� � ��).
������ � ���� ������:
  �������� ��������� �������� ���������� ����������, ���� ��� ������
  (������ ���� � �������� ���������) 0.
  ���� ��� ������ �� 0, ���������� ������ ��� ������� � ��� ������ � 2 �����.
 */

#define SHM_ENQ_ANSWER_TIMER		1000
#define SHM_ENQ_TRY_MAX_NUMBER		3
#define SHM_REQUEST_CONFIRM_TIMER	1000
#define SHM_PACKET_MAX_SIZE			256
#define SHM_WAIT_STX_TIMER			5000 // ����������� �������, ��� �� ����� ������������ �� 5 ������
#define SHM_WAIT_LENGHT_TIMER		1000
#define SHM_WAIT_DATA_TIMER			1000
#define SHM_ANSWER_TIMER			50
#define SHM_SKIP_DATA_TIMER			300

enum ShmControl {
	ShmControl_STX = 0x02,
	ShmControl_ENQ = 0x05,
	ShmControl_ACK = 0x06,
	ShmControl_NAK = 0x15,
	ShmControl_UNKNOWN = 0xFF,
};

enum ShmCommand {
	ShmCommand_ShortStatus		= 0x10,
	ShmCommand_PrintContinue	= 0xB0,
	ShmCommand_OpenShift		= 0xE0,
	ShmCommand_CheckOpen		= 0x8D,
	ShmCommand_CheckAddSale		= 0x80,
	ShmCommand_CheckClose		= 0x85,
	ShmCommand_CheckReset		= 0x88,
	ShmCommand_ShiftClose		= 0x41,
};

enum ShmError {
	ShmError_OK				= 0x00, // ��� ������
	ShmError_ShiftClosed	= 0x16,	// ��� �������� �����
	ShmError_Shift24More	= 0x4E, // ������ ������� ������ 24 �����
	ShmError_Printing		= 0x50, // ������ ����������� ���������
	ShmError_WaitPrint		= 0x58, // �������� ������� ����������� ������ (ShmCommand_PrintContinue)
};

enum ShmDocumentType {
	ShmDocumentType_Sale		 = 0,
	ShmDocumentType_Buy			 = 1,
	ShmDocumentType_SaleReturn	 = 2,
	ShmDocumentType_BuyReturn	 = 3,
};

enum ShmKKTMode {
	ShmKKTMode_Mask				 = 0x0F,
	ShmKKTMode_DataOut			 = 1,	// ������ ������
	ShmKKTMode_ShiftOpened24Less = 2,	// �������� �����, 24 ���� �� ���������
	ShmKKTMode_ShiftOpened24More = 3,	// �������� �����, 24 ���� ���������
	ShmKKTMode_ShiftClosed		 = 4,	// �������� �����
	ShmKKTMode_BlockedByPassword = 5,	// ���������� �� ������������� ������ ���������� ����������
	ShmKKTMode_WaitDateConfirm	 = 6,	// �������� ������������� ����� ����
	ShmKKTMode_DecimalPlaces	 = 7,	// ���������� ��������� ��������� ���������� �����1
	ShmKKTMode_Document			 = 8,	// �������� ��������
	ShmKKTMode_TechnicalZero	 = 9,	// ����� ���������� ���������������� ���������. � ���� ����� ��� ��������� �� ��������� �������, ���� ����������� ���������� � ����������������� ��� ���.
	ShmKKTMode_Test				 = 10,	// �������� ������
	ShmKKTMode_PrintFiscalReport = 11,	// ������ ������� ����������� ������
	ShmKKTMode_PrintFNReport	 = 12,	// ������ ������ ����
	ShmKKTMode_Document1		 = 13,	// ������ � ���������� ���������� ����������1
	ShmKKTMode_PrintDocument1	 = 14,	// ������ ����������� ���������1
	ShmKKTMode_Document1Complete = 15,	// ���������� ���������� �������� �����������1
};

#pragma pack(push,1)
/*
����������� ������� ��� �������������� ����������.
 */
struct ShmRequest {
	uint8_t command;
	BEUint4 password;
};

/*
����� �� ����� ������� � ������ ������ ��� ����� ��� �������������� ����������.
 */
struct ShmResponse {
	uint8_t command;
	uint8_t errorCode;
};

struct ShmShortStatusResponse {
	uint8_t command;
	uint8_t errorCode;
	uint8_t operatorNumber;
	uint8_t flags1;
	uint8_t flags2;
	uint8_t mode;
	uint8_t submode;
	uint8_t operNumberLow;
	uint8_t reserveVoltage;
	uint8_t voltage;
	uint8_t frErrorCode;
	uint8_t fnErrorCode;
	uint8_t operNumberHigh;
	uint8_t reserved1;
	uint8_t reserved2;
	uint8_t reserved3;

	uint16_t getOperNumber() {
		uint16_t value = (operNumberHigh << 8) | operNumberLow;
		return value;
	}
};

struct ShmPrintContinueResponse {
	uint8_t command;
	uint8_t errorCode;
	uint8_t operatorNumber;
};

struct ShmCheckAddSaleRequest {
	uint8_t command;
	BEUint4 password;
	BEUint5 number;
	BEUint5 price;
	uint8_t department;
	BEUint4 tax;
	uint8_t name[40];

	void setName(const char *value) {
		uint8_t nameMaxLen = 39;
		uint8_t nameBufSize = 40;
		uint8_t i = 0;
		for(; i < nameMaxLen; i++) {
			name[i] = value[i];
			if(value[i] == '\0') {
				break;
			}
		}
		for(; i < nameBufSize; i++) {
			name[i] = '\0';
		}
	}
};

struct ShmCheckCloseRequest {
	uint8_t command;   // 85H. ����� ���������: 71 ��� 40+Y1,2 ����.
	BEUint4 password;  // ������ ��������� (4 �����)
	BEUint5 cash;      // ����� �������� (5 ����) 0000000000�9999999999
	BEUint5 value2;    // ����� ���� ������ 2 (5 ����) 0000000000�9999999999
	BEUint5 value3;    // ����� ���� ������ 3 (5 ����) 0000000000�9999999999
	BEUint5 value4;    // ����� ���� ������ 4 (5 ����) 0000000000�9999999999
	BEUint2 discont;   // ������/��������(� ������ �������������� ��������) � % �� ��� �� 0 �� 99,99 % (2 ����� �� ������) -9999�9999
	uint8_t tax1;      // ����� 1 (1 ����) �0� � ���, �1���4� � ��������� ������
	uint8_t tax2;      // ����� 2 (1 ����) �0� � ���, �1���4� � ��������� ������
	uint8_t tax3;      // ����� 3 (1 ����) �0� � ���, �1���4� � ��������� ������
	uint8_t tax4;      // ����� 4 (1 ����) �0� � ���, �1���4� � ��������� ������
	uint8_t name[40];  // �����3,4,5,6 (40 ��� �� Y1,2 ����)

	void setName(const char *value) {
		uint8_t nameMaxLen = 39;
		uint8_t nameBufSize = 40;
		uint8_t i = 0;
		for(; i < nameMaxLen; i++) {
			name[i] = value[i];
			if(value[i] == '\0') {
				break;
			}
		}
		for(; i < nameBufSize; i++) {
			name[i] = '\0';
		}
	}
};

struct ShmCheckCloseResponse {
	uint8_t command;
	uint8_t errorCode;
	uint8_t operatorNumber;
	uint8_t change;
	char url[0];
};
#pragma pack(pop)

class ShtrihmCrc {
public:
	void start(uint8_t byte) { crc = byte; }
	void add(uint8_t byte) { crc = crc ^ byte; }
	uint8_t getCrc() { return crc; }
	static uint8_t calc(const Buffer *data) {
		uint8_t dataLen = data->getLen();
		ShtrihmCrc crc;
		crc.start(dataLen);
		for(uint8_t i = 0; i < dataLen; i++) {
			crc.add((*data)[i]);
		}
		return crc.getCrc();
	}

private:
	uint8_t crc;
};

#endif
