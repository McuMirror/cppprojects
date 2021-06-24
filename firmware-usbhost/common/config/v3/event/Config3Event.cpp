#include <string.h>
#include <strings.h>

#include "Config3EventList.h"
#include "memory/include/MemoryCrc.h"
#include "utils/include/StringParser.h"
#include "logger/include/Logger.h"
#include "platform/include/platform.h"


void Config3EventSale::set(Config3EventSale *data) {
	selectId.set(data->selectId.get());
	wareId = data->wareId;
	name.set(data->name.get());
	device.set(data->device.get());
	priceList = data->priceList;
	price = data->price;
	taxSystem = data->taxSystem;
	taxRate = data->taxRate;
	taxValue = data->taxValue;
	fiscalRegister = data->fiscalRegister;
	fiscalStorage = data->fiscalStorage;
	fiscalDocument = data->fiscalDocument;
	fiscalSign = data->fiscalSign;
}

Config3Event::Config3Event() : memory(NULL) {

}

void Config3Event::bind(Memory *memory) {
	this->memory = memory;
	this->address = memory->getAddress();
	memset(&data, 0, sizeof(data));
}

MemoryResult Config3Event::init(Memory *memory) {
	this->memory = memory;
	this->address = memory->getAddress();
	memset(&data, 0, sizeof(data));
	return save();
}

MemoryResult Config3Event::load(Memory *memory) {
	this->memory = memory;
	this->address = memory->getAddress();
	MemoryCrc crc(memory);
	return crc.readDataWithCrc(&data, sizeof(data));
}

MemoryResult Config3Event::save() {
	if(memory == NULL) {
		LOG_ERROR(LOG_CFG, "Memory not inited");
		return MemoryResult_WrongData;
	}
	memory->setAddress(address);
	MemoryCrc crc(memory);
	return crc.writeDataWithCrc(&data, sizeof(data));
}

void Config3Event::set(DateTime *datetime, uint16_t code, const char *str) {
	data.date.set(datetime);
	data.code = code;
	data.data.number = 0;
	data.data.string.set(str);
	save();
}

void Config3Event::set(DateTime *datetime, Fiscal::Sale *sale, uint16_t index) {
	data.date.set(datetime);
	data.code = Type_Sale;

	Fiscal::Product *product = sale->products.get(index);
	data.sale.selectId.set(product->selectId.get());
	data.sale.wareId = product->wareId;
	data.sale.name.set(product->name.get());
	data.sale.price = product->price;
	data.sale.taxRate = product->taxRate;

	data.sale.device.set(sale->device.get());
	data.sale.priceList = sale->priceList;
	data.sale.taxSystem = sale->taxSystem;
	data.sale.taxValue = sale->taxValue;
	data.sale.fiscalRegister = sale->fiscalRegister;
	data.sale.fiscalStorage = sale->fiscalStorage;
	data.sale.fiscalDocument = sale->fiscalDocument;
	data.sale.fiscalSign = sale->fiscalSign;
	save();
}

void Config3Event::set(Config3EventStruct *entry) {
	data.date.set(&entry->date);
	data.code = entry->code;
	if(data.code == Type_Sale) {
		data.sale.set(&entry->sale);
	} else {
		data.data.number = entry->data.number;
		data.data.string.set(entry->data.string.get());
	}
	save();
}

uint32_t Config3Event::getDataSize() {
	return (sizeof(Config3EventStruct));
}

const char *Config3Event::getEventName(Config3Event *event) {
	switch(event->getCode()) {
	case Type_OnlineStart: return "����� �����������";
	case Type_OnlineEnd: return "����� ��������";
	case Type_OnlineLast: return "����� �������";
	case Type_Sale: return "�������";
	case Type_PaymentBlocked: return "������� ���������";
	case Type_PaymentUnblocked: return "������� ��������";
	case Type_PowerUp: return "������� �������";
	case Type_PowerDown: return "������� ��������";
	case Type_BadSignal: return "������ ������";
	case Type_CashlessIdNotFound: return "������ ���������";
	case Type_PriceListNotFound: return "������ ���������";
	case Type_SyncConfigError: return "������ ���������";
	case Type_PriceNotEqual: return "������ ���������";
	case Type_SaleDisabled: return "������� �������������";
	case Type_SaleEnabled: return "������� ��������";
	case Type_ConfigEdited: return "������������ ���������";
	case Type_ConfigLoaded: return "������������ ���������";
	case Type_ConfigLoadFailed: return "������ ������������";
	case Type_ConfigParseFailed: return "������ ������������";
	case Type_ConfigUnknowProduct: return "������ ������������";
	case Type_ConfigUnknowPriceList: return "������ ������������";
	case Type_FiscalUnknownError: return "������ ��";
	case Type_FiscalLogicError: return "������ ��";
	case Type_FiscalConnectError: return "������ ��";
	case Type_FiscalPassword: return "������ ��";
	case Type_PrinterNotFound: return "������ ��������";
	case Type_PrinterNoPaper: return "������ ��������";
	case Type_EventReadError: return "������ ������";
	case Type_ModemReboot: return "����� ������������";
	case Type_CashCanceled: return "�������";
	case Type_SaleFailed: return "�������";
	case Type_BillIn: return "��������";
	case Type_CoinIn: return "��������";
	case Type_ChangeOut: return "��������";
	case Type_CashlessCanceled: return "��������";
	case Type_CashlessDenied: return "��������";
	default: return "Unknown";
	}
}

void Config3Event::getEventDescription(Config3Event *event, StringBuilder *buf) {
	buf->clear();
	switch(event->getCode()) {
	case Type_Sale: getEventSaleDescription(event, buf); return;
	case Type_CashlessIdNotFound: *buf << "�������� � ������� " << event->getString() << " ��� � �����������"; return;
	case Type_PriceListNotFound: *buf << "�����-����� " << event->getString() << " ��� � �����������"; return;
	case Type_SyncConfigError: *buf << "����������� �� ���������"; return;
	case Type_PriceNotEqual: getEventPriceNotEqualDescription(event, buf); return;
	case Type_ConfigEdited: *buf << "������������ �������� ��������"; return;
	case Type_ConfigLoaded: *buf << "������������ ��������� � �������"; return;
	case Type_ConfigLoadFailed: *buf << "������ �������� ������������"; return;
	case Type_ConfigParseFailed: *buf << "������ ������� ������������"; return;
	case Type_ConfigUnknowProduct: *buf << "����������� ����� �������� " << event->getString(); return;
	case Type_ConfigUnknowPriceList: *buf << "����������� �����-����" << event->getString(); return;
	case Type_FiscalUnknownError: *buf << "��� ������ " << event->getString(); return;
	case Type_FiscalLogicError: *buf << "������ ��������� " << event->getString(); return;
	case Type_FiscalConnectError: *buf << "��� ����� � ��"; return;
	case Type_FiscalPassword: *buf << "������������ ������ ��"; return;
	case Type_PrinterNotFound: *buf << "������� �� ������"; return;
	case Type_PrinterNoPaper: *buf << "� �������� ����������� ������"; return;
	case Type_EventReadError: *buf << "������ ������ ������� �������"; return;
	case Type_ModemReboot: *buf << "���������� ������"; return;
	case Type_CashCanceled: *buf << "������ ��������� �������� ���������"; return;
	case Type_SaleFailed: *buf << "������ �������"; return; // (STRING:<selectId>)
	case Type_WaterOverflow: *buf << "������������ ����� ������ �������"; return;
	case Type_FiscalNotInited: *buf << "��� �� ���������������"; return;
	case Type_WrongResponse: *buf << "������������ ������ ������"; return;
	case Type_BrokenResponse: *buf << "������������ �����"; return;
	case Type_FiscalCompleteNoData: *buf << "��� ������, �� ��������� �� ��������"; return;
	case Type_BillIn: *buf << "������� ������"; return; // (STRING:<nominal>)
	case Type_BillUnwaitedPacket: *buf << "������ ���������������"; return;
	case Type_CoinIn: *buf << "������� ������"; return; // (STRING:<nominal>)
	case Type_ChangeOut: *buf << "������ �����"; return; // (STRING:<sum>)
	case Type_CoinUnwaitedPacket: *buf << "������ �������������"; return; // (����������� �����)
	case Type_CashlessCanceled: *buf << "����������� ������ �������� ���������"; return;
	case Type_CashlessDenied: *buf << "������ ������ ��������� ����������"; return; // (STRING:<selectId>)
	default:;
	}
}

void Config3Event::getEventSaleDescription(Config3Event *event, StringBuilder *buf) {
	Config3EventSale *sale = event->getSale();
	*buf << "\"" << sale->name.get() << "\" �� " << sale->price << paymentDeviceToString(sale->device.get());
}

void Config3Event::getEventPriceNotEqualDescription(Config3Event *event, StringBuilder *buf) {
	const char *def = "���� �� ��������� � ������������";
	StringParser parser(event->getString());
	buf->clear();
	char selectId[8];
	if(parser.getValue("*", selectId, sizeof(selectId)) == 0) {
		*buf << def;
		return;
	}
	if(parser.compareAndSkip("*") == false) {
		*buf << def;
		return;
	}
	uint32_t expPrice = 0;
	if(parser.getNumber<uint32_t>(&expPrice) == false) {
		*buf << def;
		return;
	}
	if(parser.compareAndSkip("*") == false) {
		*buf << def;
		return;
	}
	uint32_t actPrice = 0;
	if(parser.getNumber<uint32_t>(&actPrice) == false) {
		*buf << def;
		return;
	}
	*buf << "���� �� ��������� � ������������ (������ " << selectId << ", ����������� " << expPrice << ", ������� " << actPrice << ")";
}

const char *Config3Event::paymentDeviceToString(const char *device) {
	if(strcasecmp("CA", device) == 0) {
		return " ���������";
	}
	if(strcasecmp("DA", device) == 0 || strcasecmp("DB", device) == 0) {
		return " ������������";
	}
	return "";
}
