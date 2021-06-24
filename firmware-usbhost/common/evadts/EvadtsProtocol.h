#ifndef COMMON_CONFIG_EVADTSPROTOCOL_H_
#define COMMON_CONFIG_EVADTSPROTOCOL_H_

#include <stdint.h>

namespace Evadts {

#define EvadtsPaymentDeviceSize 2 // �� EVADTS 2 ������� AN
#define EvadtsProductIdSize 4     // �� EVADTS 8 �������� AN
#define EvadtsUint32Undefined 0xFFFFFFFF

/*
 * ����������� ����� �������� ������:
 * EVADTS 20
 * ���������� 128
 * ���������� 65
 * PayKiosk�� 40 ��� 1-9999
 * TLV1030 128
 */
#define ConfigProductNameSize 50
#define CONFIG_INDEX_UNDEFINED 0xFFFF
#define LOYALITY_CODE_SIZE 128

enum Result {
	Result_OK = 0,
	Result_Busy,
	Result_RomReadError,
	Result_RomWriteError,
	Result_WrongCrc,
	Result_PriceListNumberNotEqual,
	Result_PriceListNotFound,
	Result_ProductNumberNotEqual,
	Result_ProductNotFound,
	Result_CashlessIdNotEqual
};

}

#endif
