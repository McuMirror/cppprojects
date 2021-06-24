#ifndef LIB_FISCALREGISTER_NANOKASSA_PROTOCOL_H
#define LIB_FISCALREGISTER_NANOKASSA_PROTOCOL_H

#define NANOKASSA_MANUFACTURER "NKS"
#define NANOKASSA_MODEL "Nanokassa"
#define NANOKASSA_POLL_DELAY 5000
#define NANOKASSA_TRY_MAX 5
#define NANOKASSA_TRY_DELAY 30000

namespace Nanokassa {

enum TaxSystem {
	TaxSystem_OSN    = 0x01, // ����� ���
	TaxSystem_USND   = 0x02, // ���������� �����
	TaxSystem_USNDMR = 0x04, // ���������� ����� ����� ������
	TaxSystem_ENVD   = 0x08, // ������ ����� �� ��������� �����
	TaxSystem_ESN    = 0x10, // ������ �������������������� �����
	TaxSystem_Patent = 0x20, // ��������� ������� ���������������
};

enum TaxRate {
	TaxRate_NDSNone = 6,
	TaxRate_NDS0	= 5,
	TaxRate_NDS10	= 2,
	TaxRate_NDS20	= 1,
};

extern TaxRate convertTaxRate2Nanokassa(uint8_t taxRate);

}

#endif
