#ifndef COMMON_CONFIG_CONFIGEVENT1_H_
#define COMMON_CONFIG_CONFIGEVENT1_H_

#include "timer/include/RealTime.h"
#include "utils/include/StringBuilder.h"
#include "utils/include/NetworkProtocol.h"
#include "memory/include/Memory.h"

#define EvadtsPaymentDeviceSize1 2 // �� EVADTS 2 ������� AN
#define EvadtsProductIdSize1 4     // �� EVADTS 8 �������� AN
#define EvadtsProductNameSize1 20  // �� EVADTS 20 �������� AN
#define EvadtsUint32Undefined1 0xFFFFFFFF

#pragma pack(push,1)
struct ConfigEventData1 {
	uint32_t number;
	StrParam<64> string;
};

struct ConfigEventSale1 {
	StrParam<EvadtsProductIdSize1> selectId;
	StrParam<EvadtsProductNameSize1> name;
	StrParam<EvadtsPaymentDeviceSize1> device;
	uint8_t priceList;
	uint32_t price;
	uint64_t fiscalRegister; // �� (����� ����������� ������������)
	uint64_t fiscalStorage;  // �� (����� ����������� ����������)
	uint32_t fiscalDocument; // �� (����� ����������� ���������)
	uint32_t fiscalSign;     // �� ��� ��� (���������� ������� ���������)

	void set(ConfigEventSale1 *entry);
};

struct ConfigEventStruct1 {
	DateTime date;
	uint16_t code;
	union {
		ConfigEventData1 data;
		ConfigEventSale1 sale;
	};

	ConfigEventStruct1() {}
};
#pragma pack(pop)

class Config1Event {
public:
	enum Code : uint16_t {
		Type_None					= 0xFFFF,
		Type_Sale					= 0x0003, // �������
	};

	ConfigEventStruct1 data;

	Config1Event();
	MemoryResult load(Memory *memory);
	MemoryResult save(Memory *memory);

	DateTime *getDate() { return &data.date; }
	uint16_t getCode() { return data.code; }
	uint32_t getNumber() { return data.data.number; }
	const char *getString() { return data.data.string.get(); }
	ConfigEventSale1 *getSale() { return &(data.sale); }
	ConfigEventStruct1 *getData() { return &data; }
};

#endif
