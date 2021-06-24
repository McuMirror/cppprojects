#ifndef COMMON_FISCALREGISTER_RPSYSTEM1FAPROTOCOL_H
#define COMMON_FISCALREGISTER_RPSYSTEM1FAPROTOCOL_H

#include "utils/include/NetworkProtocol.h"

namespace RpSystem1Fa {

/*
---------------------
����������� ���������
---------------------
������ 2.2.11

��������� ����� ������� � ������:
	STX <length> <command> <data> <crc>, ���
	STX � ���� ������ ����� �������;
	<length> - ����� �����. � ����� �� ������ STX, <length> � <crc> �����;
	<command> - ��� �������;
	<crc> - ���� ����������� �����

� �������� ETHERNET �� ������������ ����� STX � <crc>, �.�. ��� �������
�� ����������� ����������� � ������������ ������� ����� �� ���� ���� TCP/IP.
����� ������ ��� �� �� ���������� ������������� ACK, NAK, ENQ.
���� ������ ���������� ���������� � ����������� �� 3333 �����, ���������
� ���� ��������� � ������� ������ �� ���������� ���������� �������,
��� �� ���������� ������ ��������� ������ ������������ ��������.
����� ��������� ������ ���� ������ ������� ����������.

-------------------------
���� �������� ����� � ���
-------------------------
*/

#define RPSYSTEM1FA_PACKET_MAX_SIZE		100 // ��������
#define RPSYSTEM1FA_CONFIRM_TIMEOUT		1000
#define RPSYSTEM1FA_RESPONSE_TIMEOUT	10000 // ����� ���������� �������
#define RPSYSTEM1FA_DELIVER_TIMEOUT		10

enum Control {
	Control_STX  = 0x02,
	Control_ENQ  = 0x05,
	Control_ACK  = 0x06,
	Control_NAK  = 0x15,
};

enum Command {
	Command_Status			= 0x11, // ������ ��������� ���
	Command_ShiftOpen		= 0xE0, // ������� �����
	Command_CheckOpen		= 0x8D, // ������� ���
	Command_CheckAddSale	= 0x80, // �������� �������
	Command_CheckClose		= 0x85, // ������� ���
	Command_CheckCancel		= 0x88, // ������ ����
	Command_ShiftClose		= 0x41, // ������� �����
};

enum Error {
	Error_OK				= 0,  // ��� ������
	Error_Shift24More		= 22, // ����������������� ����� ����� 24 �����
	Error_ShiftClosed		= 61, // ����� �� �������
	Error_DocumentOpened	= 74, // �������� ��� ������
};

enum FrState {
	FrState_Idle				 = 0x00, // ��������� ���������� ����� ������ (���������� ����� ������������� ��� ��������� ��� 0x04 ������ �������)
	FrState_ShiftOpened			 = 0x02, // ����� �������
	FrState_Shift24More			 = 0x03, // ����� ������� ����� 24 �����
	FrState_ShiftClosed			 = 0x04, // ����� �������
	FrState_WaitDateConfirm		 = 0x06, // ������� ������������� ����
	FrState_DocumentIn			 = 0x08, // ������ �������� �������
	FrState_DocumentOut			 = 0x18, // ������ �������� �������
	FrState_DocumentReturnIn	 = 0x28, // ������ �������� �������� �������
	FrState_DocumentReturnOut	 = 0x38, // ������ �������� �������� �������
	FrState_FatalError			 = 0xFF, // ��������� ������ ����������
};

enum ModeFlag {
	ModeFlag_Encoding	 = 0x01, // ����������
	ModeFlag_Autonomous	 = 0x02, // ���������� �����
	ModeFlag_Automatic	 = 0x04, // �������������� �����
	ModeFlag_Service	 = 0x08, // ���������� � ����� �����
	ModeFlag_BSO		 = 0x10, // ����� ���
	ModeFlag_Internet	 = 0x20, // ��� ������������� ��� �������� � ���� ��������
};

enum TaxFlag {
	TaxFlag_OSNO	 = 0x01, // �����
	TaxFlag_USNO6	 = 0x02, // ���������� �����
	TaxFlag_USNO15	 = 0x04, // ���������� ����� ����� ������
	TaxFlag_ENVD	 = 0x08, // ������ ����� �� �������� �����
	TaxFlag_ESHN	 = 0x10, // ������ �������������������� �����
	TaxFlag_PSNO	 = 0x20, // ��������� ������� ���������������
};

enum LifePhase {
	LifePhase_Setup			 = 0, // ���������
	LifePhase_ReadyToFiscal	 = 1, // ���������� � ������������
	LifePhase_Fiscal		 = 3, // ���������� �����
	LifePhase_PostFiscal	 = 7, // ����-���������� �����, ��� �������� �� � ���
	LifePhase_FnReading		 = 15, // ������ ������ �� ������ ��
};

enum PrinterState {
	PrinterState_Idle		 = 0, // ������� � �������� (����� ��� ������?)
	PrinterState_PaperOff	 = 1, // ����������� ������
	PrinterState_PaperLess	 = 2, // ������ ����� ����������
	PrinterState_Open		 = 3, // ������� ������ ��������
	PrinterState_Error1		 = 4, // ������������� ������ ��������
	PrinterState_Error2		 = 5, // ��������������� ������ ��������
	PrinterState_Error3		 = 6, // ������� ������� ������������ �����
};

enum FnFlag {
	FnFlag_3DaysLeft	 = 0x01, // ������� ������ �� (�� ��������� ����� �������� 3 ���)
	FnFlag_30DaysLeft	 = 0x02, // ���������� ������� �� (�� ��������� ����� �������� 30 ����)
	FnFlag_MemoryEnds	 = 0x04, // ������������ ������ �� (����� �� �������� �� 90 %)
	FnFlag_Timeout		 = 0x08, // ��������� ����� �������� ������ ���
	FnFlag_CriticalError = 0x80, // ����������� ������ ��
};

enum DocumentType {
	DocumentType_In			 = 0, // ������
	DocumentType_Out		 = 1, // ������
	DocumentType_ReturnIn	 = 2, // ������� �������
	DocumentType_ReturnOut	 = 3, // ������� �������
};

enum PaymentType {
	PaymentType_FullPayBeforeTake			= 1, // ������ ��������������� ������ �� ������� �������� �������� �������
	PaymentType_PartialPayBeforeTake		= 2, // ��������� ��������������� ������ �� ������� �������� �������� �������
	PaymentType_Prepayment					= 3, // �����
	PaymentType_FullPayAndTake				= 4, // ������ ������, � ��� ����� � ������ ������ (��������������� ������) � ������ �������� �������� �������
	PaymentType_PartialPayAndTakeInCredit	= 5, // ��������� ������ �������� ������� � ������ ��� �������� � ����������� ������� � ������
	PaymentType_TakeInCredit				= 6, // �������� �������� ������� ��� ��� ������ � ������ ��� �������� � ����������� ������� � ������
	PaymentType_CreditPayment				= 7, // ������ �������� ������� ����� ��� �������� � ������� � ������ (������ �������)
};

/*
enum TaxType {
	TaxType_OSNO	 = 1, // �����
	TaxType_USNO6	 = 2, // ���������� �����
	TaxType_USNO15	 = 3, // ���������� ����� ����� ������
	TaxType_ENVD	 = 4, // ������ ����� �� �������� �����
	TaxType_ESHN	 = 5, // ������ �������������������� �����
	TaxType_PSNO	 = 6, // ��������� ������� ���������������
};*/

#pragma pack(push,1)
/*
 * ����������� ������� ��� �������������� ����������.
 */
struct Request {
	uint8_t command;
	BEUint4 password;
};

struct Response {
	uint8_t command;
	uint8_t errorCode;  // ��� ������
};

struct StatusResponse {
	uint8_t command;
	uint8_t errorCode;
	uint8_t softwareVersion[2];	// ������ ��
	BEUint2 softwareBuild;		// ����� ������ ��
	uint8_t softwareDate[3];	// ���� ������ ��
	uint8_t reserved1[5];
	uint8_t state;				// ������� ��������� �� (������ FrState)
	uint8_t lifePhase;			// ������ LifePhase
	uint8_t printerLastState;	// ��������� �������� ����� ������ ���������� ���������
	uint8_t fnDate[3];			// ���� �������� ��
	uint8_t fnFlags;			// ����� ��������� �� (������ FnFlag)
	uint8_t printerState;		// ������� ��������� �������� (������ PrinterState)
	uint8_t reserved2[2];
	uint8_t date[3];			// ������� ����
	uint8_t time[3];			// ������� �����
	uint8_t reserved3[37];
};

struct CheckOpenRequest {
	uint8_t command;
	BEUint4 password;
	uint8_t documentType;
};

struct CheckAddRequest {
	uint8_t command;
	BEUint4 password;
	BEUint5 number;			// ����������
	BEUint5 price;			// ����
	uint8_t paymentType;	// ������� ������� ������� (�������� PaymentType)
	uint8_t taxType;		// ��� ������ WTF?
	uint8_t reserved1[3];
	uint8_t name[0];		// ��������, ����� �� ����� ������ ������ ����������� ����������

	void setName(const char *str, uint16_t strLen) {
		uint16_t i = 0;
		for(; i < strLen && str[i] != '\0'; i++) {
			name[i] = str[i];
		}
	}
};

struct CheckCloseRequest {
	uint8_t command;
	BEUint4 password;
	BEUint5 cash;			// ����� ���������
	BEUint5 cashless1;		// ����� ������������ 1-�� ����
	BEUint5 cashless2;		// ����� ������������ 2-�� ����
	BEUint5 cashless3;		// ����� ������������ 3-�� ����
	uint8_t reserved1[6];
	uint8_t text[0];		// �����, ����� �� ����� ������ ������ ����������� ����������

	void setText(const char *str, uint16_t strLen) {
		uint16_t i = 0;
		for(; i < strLen && str[i] != '\0'; i++) {
			text[i] = str[i];
		}
	}
};

#pragma pack(pop)

}

#endif
