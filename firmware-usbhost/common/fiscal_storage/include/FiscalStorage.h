#ifndef COMMON_FISCAL_STORAGE_H
#define COMMON_FISCAL_STORAGE_H

#include "utils/include/CodePage.h"
#include "utils/include/NetworkProtocol.h"
#include "utils/include/Buffer.h"
#include "fiscal_register/include/FiscalRegister.h"

#include <stdint.h>

namespace FiscalStorage {
/*
 * ��������� ����� CP866
 *
 *
//----------------------------
// ����������� ����� ������ ��� ��������� ���� �� ���
//----------------------------
 * ��� �������� ���� ����� ����� ��������� ������:
 * - ���� � ����� �������;
 * - ��� ��������;
 * - ����� ����;
 * - ����� ����;
 * - ���������� ������� ���������.
 *
 * ������ ������:
 * - ���� � ����� �������;
 * - ��� ��������;
 * - ����� ����;
 * - ����� ����������� ����������
 * - ����� ���������
 * - ���������� ������� ���������.
 *
 * ������ ������:
 * - ����� ����������� ����������
 * - ����� ����������� ���������
 * - ���������� ������� ���������
 *
//----------------------------
// ������ ������ ���� � TLV
//----------------------------
// ����� �� (ASCII)
I Atol.cpp#743 tag=1041, len=16, data=x39;x39;x39;x39;x30;x37;x38;x39;x30;x30;x30;x30;x37;x33;x31;x37;
// ��������������� ����� ��� (ASCII)
I Atol.cpp#743 tag=1037, len=20, data=x30;x30;x30;x30;x30;x30;x30;x30;x30;x37;x30;x31;x32;x32;x37;x38;x20;x20;x20;x20;
// ��� ������������ (ASCII)
I Atol.cpp#743 tag=1018, len=12, data=x36;x39;x30;x32;x30;x39;x38;x31;x32;x37;x35;x32;
// ����� �� (uint32_t)
I Atol.cpp#743 tag=1040, len=4, data=x57;x00;x00;x00;
// ���� � ����� ������������ �� (������?)
I Atol.cpp#743 tag=1012, len=4, data=xFC;xEC;x8E;x5A;
// ���������� ������� ���������
I Atol.cpp#743 tag=1077, len=6, data=x31;x04;x66;x29;x76;x5F;
// ����� ����� (uint32_t)
I Atol.cpp#743 tag=1038, len=4, data=x0E;x00;x00;x00;
// ����� ���� �� ����� (uint32_t)
I Atol.cpp#743 tag=1042, len=4, data=x01;x00;x00;x00;
// ������� ������� (������, ������� � ��� �����)
I Atol.cpp#743 tag=1054, len=1, data=x01;
// ����� �������, ���������� � ���� (���) (uint16_t � ��������)
I Atol.cpp#743 tag=1020, len=2, data=xC4;x09;
// ������� ������� (���������?)
I Atol.cpp#743 tag=1059, len=41, data=x06;x04;x08;x00;x9D;xE1;xAF;xE0;xA5;xE1;xE1;xAE;x37;x04;x02;x00;xC4;x09;xFF;x03;x02;x00;x00;x01;x13;x04;x02;x00;xC4;x09;xAF;x04;x01;x00;x01;xB0;x04;x02;x00;x7D;x01;
// ����� �� ���� (���) ��������� (uint16_t � ��������)
I Atol.cpp#743 tag=1031, len=2, data=xC4;x09;
// ����� �� ���� (���) ������������ (uint8_t � ��������)
I Atol.cpp#743 tag=1081, len=1, data=x00;
// ����� �� ���� (���) ����������� (������� ������ � (���) ���������� ��������)
I Atol.cpp#743 tag=1215, len=1, data=x00;
// ����� �� ���� (���) ����������� (� ������)
I Atol.cpp#743 tag=1216, len=1, data=x00;
// ����� �� ���� (���) ��������� ���������������
I Atol.cpp#743 tag=1217, len=1, data=x00;
// ����������� ������� ���������������
I Atol.cpp#743 tag=1055, len=1, data=x08;
// ����� �������� (ASCII, "159")
I Atol.cpp#743 tag=1036, len=3, data=x31;x35;x39;
// ����� ��� ���� �� ������ 18% (uin16_t � ��������)
I Atol.cpp#743 tag=1102, len=2, data=x7D;x01;

//----------------------------
// ��� ��������� ��������� � �� ERP
//----------------------------
// ��������������� ����� ��� (ASCII)
I Atol.cpp#743 tag=1037, len=20, data=x30;x30;x30;x30;x30;x30;x30;x30;x30;x37;x30;x31;x32;x32;x37;x38;x20;x20;x20;x20;
// ����� �� (ASCII)
I Atol.cpp#743 tag=1041, len=16, data=x39;x39;x39;x39;x30;x37;x38;x39;x30;x30;x30;x30;x37;x33;x31;x37;
// ����� �� (uint32_t)
I Atol.cpp#743 tag=1040, len=4, data=x57;x00;x00;x00;
// ���������� ������� ���������
I Atol.cpp#743 tag=1077, len=6, data=x31;x04;x66;x29;x76;x5F;
// ���� � ����� ������������ �� (������?)
I Atol.cpp#743 tag=1012, len=4, data=xFC;xEC;x8E;x5A;
// ����� �������� (ASCII, "159")
I Atol.cpp#743 tag=1036, len=3, data=x31;x35;x39;
// ������� ������� (���������?)
I Atol.cpp#743 tag=1059, len=41, data=x06;x04;x08;x00;x9D;xE1;xAF;xE0;xA5;xE1;xE1;xAE;x37;x04;x02;x00;xC4;x09;xFF;x03;x02;x00;x00;x01;x13;x04;x02;x00;xC4;x09;xAF;x04;x01;x00;x01;xB0;x04;x02;x00;x7D;x01;
x06;x04;=1030 x08;x00; x9D;xE1;xAF;xE0;xA5;xE1;xE1;xAE;
x37;x04;=1079 x02;x00; xC4;x09;
xFF;x03;=1023 x02;x00; x00;x01;
x13;x04;=1043 x02;x00; xC4;x09;
xAF;x04;=1199 x01;x00; x01;
xB0;x04;=1200 x02;x00; x7D;x01;

// ����� �������, ���������� � ���� (���) (uint16_t � ��������)
I Atol.cpp#743 tag=1020, len=2, data=xC4;x09;
// ����������� ������� ���������������
I Atol.cpp#743 tag=1055, len=1, data=x08;
// ����� ��� ���� �� ������ 18% (uin16_t � ��������)
I Atol.cpp#743 tag=1102, len=2, data=x7D;x01;
 */

#define FN105_CHECK_POS_NAME_MAX_SIZE 128

enum Tag {
	Tag_FSNumber			 = 1041, // ����� ��
	Tag_FRRegistrationNumber = 1037, // ��������������� ����� ���
	Tag_UserINN				 = 1018, // ��� ������������
// ���� �� ����� - �������� ������������ � ��������
	Tag_Product				 = 1059,
	Tag_ProductName			 = 1030,
	Tag_ProductPrice		 = 1079,
	Tag_ProductNumber		 = 1023,
	Tag_TaxRate				 = 1199,
	Tag_PaymentMethod		 = 1214,

	Tag_TaxSystem			 = 1055, // ����� ��������������� BYTE
	Tag_PaymentCash			 = 1031, // ���������
	Tag_PaymentCashless		 = 1081, // ������������
	Tag_Prepayment			 = 1215, // ����������� (������� ������)
	Tag_Postpay				 = 1216, // ����������� (� ������)
	Tag_PaymentHZ			 = 1217, // ��������� ���������������
	Tag_ClientMail			 = 1008, // ����� ����������� ����� ������� ASCII
};

#pragma pack(push,1)
struct Header {
	BEUint2 tag;
	BEUint2 len;
	uint8_t data[0];
};
#pragma pack(pop)

enum TaxSystem {
	TaxSystem_OSN	 = 0x01, // ����� ���
	TaxSystem_USND	 = 0x02, // ���������� �����
	TaxSystem_USNDMR = 0x04, // ���������� ����� ����� ������
	TaxSystem_ENVD	 = 0x08, // ������ ����� �� ��������� �����
	TaxSystem_ESN	 = 0x10, // ������ �������������������� �����
	TaxSystem_Patent = 0x20, // ��������� ������� ���������������
};

extern TaxSystem convertTaxSystem2FN105(uint8_t taxSystem);

enum TaxRate {
	TaxRate_NDSNone = 6,
	TaxRate_NDS0	= 5,
	TaxRate_NDS10	= 2,
	TaxRate_NDS18	= 1,
};

extern TaxRate convertTaxRate2FN105(uint8_t taxRate);

enum PaymentMethod {
	PaymentMethod_Cash = 1,
	PaymentMethod_Cashless =2,
};

extern PaymentMethod convertPaymentMethod2FN105(uint8_t paymentMethod);

void addTlvHeader(FiscalStorage::Tag tag, uint16_t len, Buffer *buf);
void addTlvUint32(FiscalStorage::Tag tag, uint32_t value, Buffer *buf);
void addTlvFUint32(FiscalStorage::Tag tag, uint8_t dotPos, uint8_t number, Buffer *buf);
void addTlvString(FiscalStorage::Tag tag, const char *str, Buffer *buf);

}

#endif
