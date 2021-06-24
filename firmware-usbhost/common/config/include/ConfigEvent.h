#ifndef COMMON_CONFIG_CONFIGEVENT_H_
#define COMMON_CONFIG_CONFIGEVENT_H_

#if 0
#include "config/v3/event/Config3EventData.h"
#include "timer/include/RealTime.h"
#include "utils/include/StringBuilder.h"
#include "memory/include/Memory.h"

#define CONFIG_EVENT_UNSET 0xFFFF

class ConfigEvent {
public:
	enum Code : uint16_t {
		Type_None					= 0xFFFF,
		Type_OnlineStart			= 0x0000, // ����� �����������
		Type_OnlineEnd				= 0x0001, // ����� ��������
		Type_OnlineLast				= 0x0002, // ����� �������
		Type_Sale					= 0x0003, // �������
		Type_PaymentBlocked			= 0x0004, // ������� ��������� (�����������, ���� ������� ������������� ������� �����)
		Type_PaymentUnblocked		= 0x0005, // ������� ��������
		Type_PowerUp				= 0x0006, // ������� �������
		Type_PowerDown				= 0x0007, // ������� ��������
		Type_BadSignal				= 0x0008, // ������ ������
		Type_CashlessIdNotFound		= 0x0009, // �� �������� ������������ ����� �������� (STRING:<cashlessId>)
		Type_PriceListNotFound		= 0x000A, // ����� ���� �� ������ (STRING:<deviceId><priceListNumber>)
		Type_SyncConfigError		= 0x000B, // �� ������������ (�������� �������������)
		Type_PriceNotEqual			= 0x000C, // �� �������� ������������ ���� �������� (STRING:<selectId>*<expectedPrice>*<actualPrice>)
		Type_SaleDisabled			= 0x000D, // ������� ��������� ��������� ���������� �����
		Type_SaleEnabled			= 0x000E, // ������� �������� ���������
		Type_ConfigEdited			= 0x000F, // ������������ �������� ��������
		Type_ConfigLoaded			= 0x0010, // ������������ ��������� � �������
		Type_ConfigLoadFailed		= 0x0011, // ������ �������� ������������
		Type_ConfigParseFailed		= 0x0012, // ������ ������� ������������
		Type_ConfigUnknowProduct	= 0x0013, // ����������� ����� �������� (STRING:<selectId>)
		Type_ConfigUnknowPriceList	= 0x0014, // ����������� �����-���� (STRING:<deviceId><priceListNumber>)
		Type_ModemReboot			= 0x0015, // ���������� ������ (STRING:<rebootReason>)
		Type_CashCanceled			= 0x0016, // ������ ��������� �������� ���������
		Type_SaleFailed				= 0x0017, // ������ ������� (STRING:<selectId>)
		Type_WaterOverflow			= 0x0018, // ������������ ����� ������ �������
		Type_FiscalUnknownError		= 0x0300, // ���������������� ������ �� (STRING:<���-������-��>)
		Type_FiscalLogicError		= 0x0301, // ����������� ��������� �� (STRING:<������-�-�����>)
		Type_FiscalConnectError		= 0x0302, // ��� ����� � ��
		Type_FiscalPassword			= 0x0303, // ������������ ������ ��
		Type_PrinterNotFound		= 0x0304, // ������� �� ������
		Type_PrinterNoPaper			= 0x0305, // � �������� ����������� ������
		Type_FiscalNotInited		= 0x0306, // ��� �� ���������������
		Type_WrongResponse			= 0x0307, // ������������ ������ ������
		Type_BrokenResponse			= 0x0308, // ������������ �����
		Type_FiscalCompleteNoData	= 0x0309, // ��� ������, �� ��������� �� ��������
		Type_BillIn					= 0x0401, // ������� ������ (STRING:<nominal>)
		Type_BillUnwaitedPacket		= 0x0402, // ������ ��������������� (����������� �����)
		Type_CoinIn					= 0x0501, // ������� ������ (STRING:<nominal>)
		Type_ChangeOut				= 0x0502, // ������ ����� (STRING:<sum>)
		Type_CoinUnwaitedPacket		= 0x0503, // ������ ������������� (����������� �����)
		Type_CashlessCanceled		= 0x0601, // ����������� ������ �������� ��������� (STRING:<selectId><timeout>)
		Type_CashlessDenied			= 0x0602, // ������ ������ ��������� ���������� (STRING:<selectId><timeout>)
		Type_SessionClosedByMaster	= 0x0603, // �������� ����������� ������ �������� ��������� (STRING:<timeout>)
		Type_SessionClosedByTimeout	= 0x0604, // �������� ����������� ������ �������� �� �������� (STRING:<timeout>)
		Type_SessionClosedByTerminal= 0x0605, // �������� ����������� ������ �������� ���������� (STRING:<timeout>)
		Type_EventReadError			= 0xFF01, // ������ ������ ������� �������
		Type_WatchDog				= 0xFF02, // �������� WatchDog
		Type_MdbUnwaitedPacket		= 0xFF03, // ����������� ����� � MDB-���������
	};

	ConfigEvent();

	void set(DateTime *datetime, ConfigEvent::Code code);
	void set(DateTime *datetime, ConfigEvent::Code code, uint32_t param);
	void set(DateTime *datetime, ConfigEvent::Code code, uint32_t param, const char *str);
	void set(DateTime *datetime, uint16_t code, const char *str);
//	void set(DateTime *datetime, ConfigEventSale *data);
//	void set(ConfigEventStruct *data);
	DateTime *getDate() { return &data.date; }
	uint16_t getCode() { return data.code; }
	uint32_t getNumber() { return data.data.number; }
	const char *getString() { return data.data.string.get(); }
//	ConfigEventSale *getSale() { return &(data.sale); }
//	ConfigEventStruct *getData() { return &data; }

	static uint32_t getDataSize();
	static const char *getEventName(ConfigEvent *event);
	static void getEventDescription(ConfigEvent *event, StringBuilder *buf);

private:
	Config3EventStruct data;

	static void getEventSaleDescription(ConfigEvent *event, StringBuilder *buf);
	static void getEventPriceNotEqualDescription(ConfigEvent *event, StringBuilder *buf);
	static const char *paymentDeviceToString(const char *device);
};
#else
#include "config.h"
#ifndef NEW_CONFIG
#include "config/v3/event/Config3Event.h"

typedef Config3Event ConfigEvent;
typedef Config3EventSale ConfigEventSale;
#else
#include "config/v4/event/Config4Event.h"

typedef Config4Event ConfigEvent;
typedef Config4EventSale ConfigEventSale;
#endif
#endif

#endif
