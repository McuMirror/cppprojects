#include "test/include/Test.h"
#include "config/v4/event/Config4EventIterator.h"
#include "memory/include/RamMemory.h"
#include "timer/include/TestRealTime.h"
#include "logger/include/Logger.h"

class Config4EventListTest : public TestSet {
public:
	Config4EventListTest();
	bool test();
	bool testGetUnsync();
	bool testEventToHumanReadable();
};

TEST_SET_REGISTER(Config4EventListTest);

Config4EventListTest::Config4EventListTest() {
	TEST_CASE_REGISTER(Config4EventListTest, test);
	TEST_CASE_REGISTER(Config4EventListTest, testGetUnsync);
	TEST_CASE_REGISTER(Config4EventListTest, testEventToHumanReadable);
}

bool Config4EventListTest::test() {
	RamMemory memory(2048);
	TestRealTime realtime;
	Config4EventList list(&realtime);

	// init
	list.init(5, &memory);
	TEST_NUMBER_EQUAL(5, list.getSize());
	TEST_NUMBER_EQUAL(0, list.getLen());
	TEST_NUMBER_EQUAL(CONFIG4_EVENT_UNSET, list.getFirst());
	TEST_NUMBER_EQUAL(CONFIG4_EVENT_UNSET, list.getLast());
	TEST_NUMBER_EQUAL(CONFIG4_EVENT_UNSET, list.getUnsync());

	// add first
	list.add(Config4Event::Type_OnlineLast, "0"); // 0,0
	TEST_NUMBER_EQUAL(5, list.getSize());
	TEST_NUMBER_EQUAL(1, list.getLen());
	TEST_NUMBER_EQUAL(0, list.getFirst());
	TEST_NUMBER_EQUAL(0, list.getLast());
//	TEST_NUMBER_EQUAL(CONFIG_EVENT_UNSET, list.getSync());

	// overflow
	list.add(Config4Event::Type_OnlineLast, "1"); // 0,1
	list.add(Config4Event::Type_OnlineLast, "2"); // 0,2
	list.add(Config4Event::Type_OnlineLast, "3"); // 0,3
	list.add(Config4Event::Type_OnlineLast, "4"); // 0,4
	Fiscal::Sale sale1;
	sale1.setProduct("id1", 0, "name1", 1234, 0, 1);
	sale1.device.set("d1");
	sale1.priceList = 1;
	sale1.fiscalRegister = 11111;
	sale1.fiscalStorage = 22222;
	sale1.fiscalDocument = 33333;
	sale1.fiscalSign = 44444;
	list.add(&sale1);
	TEST_NUMBER_EQUAL(5, list.getSize());
	TEST_NUMBER_EQUAL(5, list.getLen());
	TEST_NUMBER_EQUAL(1, list.getFirst());
	TEST_NUMBER_EQUAL(0, list.getLast());
//	TEST_NUMBER_EQUAL(CONFIG_EVENT_UNSET, list.getSync());

	// foreach from first to last
	Config4EventIterator iterator1(&list);
	TEST_NUMBER_EQUAL(true, iterator1.first());
	TEST_NUMBER_EQUAL(Config4Event::Type_OnlineLast, iterator1.getCode());
	TEST_STRING_EQUAL("1", iterator1.getString());

	TEST_NUMBER_EQUAL(true, iterator1.next());
	TEST_NUMBER_EQUAL(Config4Event::Type_OnlineLast, iterator1.getCode());
	TEST_STRING_EQUAL("2", iterator1.getString());

	TEST_NUMBER_EQUAL(true, iterator1.next());
	TEST_NUMBER_EQUAL(Config4Event::Type_OnlineLast, iterator1.getCode());
	TEST_STRING_EQUAL("3", iterator1.getString());

	TEST_NUMBER_EQUAL(true, iterator1.next());
	TEST_NUMBER_EQUAL(Config4Event::Type_OnlineLast, iterator1.getCode());
	TEST_STRING_EQUAL("4", iterator1.getString());

	TEST_NUMBER_EQUAL(true, iterator1.next());
	TEST_NUMBER_EQUAL(Config4Event::Type_Sale, iterator1.getCode());
	Config4EventSale *sale2 = iterator1.getSale();
	TEST_STRING_EQUAL("id1", sale2->selectId.get());
	TEST_STRING_EQUAL("name1", sale2->name.get());
	TEST_STRING_EQUAL("d1", sale2->device.get());
	TEST_NUMBER_EQUAL(1, sale2->priceList);
	TEST_NUMBER_EQUAL(1234, sale2->price);
	TEST_NUMBER_EQUAL(11111, sale2->fiscalRegister);
	TEST_NUMBER_EQUAL(22222, sale2->fiscalStorage);
	TEST_NUMBER_EQUAL(33333, sale2->fiscalDocument);
	TEST_NUMBER_EQUAL(44444, sale2->fiscalSign);

	TEST_NUMBER_EQUAL(false, iterator1.next());

	// findByIndex
	iterator1.findByIndex(0);
	TEST_NUMBER_EQUAL(Config4Event::Type_Sale, iterator1.getCode());
	Config4EventSale *sale3 = iterator1.getSale();
	TEST_STRING_EQUAL("id1", sale3->selectId.get());
	TEST_STRING_EQUAL("name1", sale3->name.get());
	TEST_STRING_EQUAL("d1", sale3->device.get());
	TEST_NUMBER_EQUAL(1, sale3->priceList);
	TEST_NUMBER_EQUAL(1234, sale3->price);
	TEST_NUMBER_EQUAL(11111, sale3->fiscalRegister);
	TEST_NUMBER_EQUAL(22222, sale3->fiscalStorage);
	TEST_NUMBER_EQUAL(33333, sale3->fiscalDocument);
	TEST_NUMBER_EQUAL(44444, sale3->fiscalSign);

	// sync
	list.setSync(3);
	TEST_NUMBER_EQUAL(5, list.getSize());
	TEST_NUMBER_EQUAL(2, list.getLen());
	TEST_NUMBER_EQUAL(4, list.getFirst());
	TEST_NUMBER_EQUAL(0, list.getLast());
	TEST_NUMBER_EQUAL(4, list.getUnsync());

	// foreach from sync to last
	Config4EventIterator iterator2(&list);
	TEST_NUMBER_EQUAL(true, iterator2.unsync());
	TEST_NUMBER_EQUAL(Config4Event::Type_OnlineLast, iterator2.getCode());
	TEST_STRING_EQUAL("4", iterator2.getString());

	TEST_NUMBER_EQUAL(true, iterator2.next());
	TEST_NUMBER_EQUAL(Config4Event::Type_Sale, iterator2.getCode());

	TEST_NUMBER_EQUAL(false, iterator2.next());
//	list.setSync(iterator2.getIndex());
	TEST_NUMBER_EQUAL(5, list.getSize());
	TEST_NUMBER_EQUAL(4, list.getFirst());
	TEST_NUMBER_EQUAL(0, list.getLast());
	TEST_NUMBER_EQUAL(4, list.getUnsync());

	return true;
}

bool Config4EventListTest::testGetUnsync() {
	RamMemory memory(2048);
	TestRealTime realtime;
	Config4EventList list(&realtime);

	// init
	list.init(5, &memory);
	TEST_NUMBER_EQUAL(CONFIG4_EVENT_UNSET, list.getFirst());
	TEST_NUMBER_EQUAL(CONFIG4_EVENT_UNSET, list.getLast());
	TEST_NUMBER_EQUAL(CONFIG4_EVENT_UNSET, list.getUnsync());

	// first < last
	list.add(Config4Event::Type_OnlineLast, "0");
	list.add(Config4Event::Type_OnlineLast, "1");
	TEST_NUMBER_EQUAL(0, list.getFirst());
	TEST_NUMBER_EQUAL(1, list.getLast());
	TEST_NUMBER_EQUAL(0, list.getUnsync());

	// sync all
	list.setSync(list.getLast());
	TEST_NUMBER_EQUAL(CONFIG4_EVENT_UNSET, list.getFirst());
	TEST_NUMBER_EQUAL(1, list.getLast());
	TEST_NUMBER_EQUAL(CONFIG4_EVENT_UNSET, list.getUnsync());

	// last < first, unsync > last
	list.add(Config4Event::Type_OnlineLast, "2"); // 0,2
	list.add(Config4Event::Type_OnlineLast, "3"); // 0,3
	list.add(Config4Event::Type_OnlineLast, "4"); // 0,4
	list.add(Config4Event::Type_OnlineLast, "5"); // 1,0
	TEST_NUMBER_EQUAL(1, list.getFirst());
	TEST_NUMBER_EQUAL(0, list.getLast());
	TEST_NUMBER_EQUAL(1, list.getUnsync());

	list.add(Config4Event::Type_OnlineLast, "6");
	TEST_NUMBER_EQUAL(2, list.getFirst());
	TEST_NUMBER_EQUAL(1, list.getLast());
	TEST_NUMBER_EQUAL(2, list.getUnsync());

	// sync all
	list.setSync(list.getLast());
	TEST_NUMBER_EQUAL(CONFIG4_EVENT_UNSET, list.getFirst());
	TEST_NUMBER_EQUAL(1, list.getLast());
	TEST_NUMBER_EQUAL(CONFIG4_EVENT_UNSET, list.getUnsync());

	// last < first, unsync < last
	list.add(Config4Event::Type_OnlineLast, "7");
	list.add(Config4Event::Type_OnlineLast, "8");
	TEST_NUMBER_EQUAL(0, list.getFirst());
	TEST_NUMBER_EQUAL(3, list.getLast());
	TEST_NUMBER_EQUAL(0, list.getUnsync());
	return true;
}

bool Config4EventListTest::testEventToHumanReadable() {
	DateTime date;
	Config4Event event;
	StringBuilder str;

	Fiscal::Sale sale;
	sale.setProduct("id1", 0, "���������", 1500, 0, 1);
	sale.device.set("CA");
	sale.priceList = 0;
	event.set(&date, &sale, 0);
	TEST_STRING_EQUAL("�������", event.getEventName(&event));
	event.getEventDescription(&event, &str);
	TEST_STRING_EQUAL("\"���������\" �� 1500 ���������", str.getString());

	event.set(&date, Config4Event::Type_OnlineStart);
	TEST_STRING_EQUAL("����� �����������", event.getEventName(&event));

	event.set(&date, Config4Event::Type_CashlessIdNotFound, "123");
	TEST_STRING_EQUAL("������ ���������", event.getEventName(&event));
	event.getEventDescription(&event, &str);
	TEST_STRING_EQUAL("�������� � ������� 123 ��� � �����������", str.getString());

	event.set(&date, Config4Event::Type_PriceListNotFound, "CA0");
	TEST_STRING_EQUAL("������ ���������", event.getEventName(&event));
	event.getEventDescription(&event, &str);
	TEST_STRING_EQUAL("�����-����� CA0 ��� � �����������", str.getString());

	event.set(&date, Config4Event::Type_SyncConfigError);
	TEST_STRING_EQUAL("������ ���������", event.getEventName(&event));
	event.getEventDescription(&event, &str);
	TEST_STRING_EQUAL("����������� �� ���������", str.getString());

	event.set(&date, Config4Event::Type_PriceNotEqual, "123*456*789");
	TEST_STRING_EQUAL("������ ���������", event.getEventName(&event));
	event.getEventDescription(&event, &str);
	TEST_STRING_EQUAL("���� �� ��������� � ������������ (������ 123, ����������� 456, ������� 789)", str.getString());

	event.set(&date, Config4Event::Type_FiscalUnknownError, "123");
	TEST_STRING_EQUAL("������ ��", event.getEventName(&event));
	event.getEventDescription(&event, &str);
	TEST_STRING_EQUAL("��� ������ 123", str.getString());

	event.set(&date, Config4Event::Type_FiscalLogicError, "123");
	TEST_STRING_EQUAL("������ ��", event.getEventName(&event));
	event.getEventDescription(&event, &str);
	TEST_STRING_EQUAL("������ ��������� 123", str.getString());

	event.set(&date, Config4Event::Type_FiscalConnectError);
	TEST_STRING_EQUAL("������ ��", event.getEventName(&event));
	event.getEventDescription(&event, &str);
	TEST_STRING_EQUAL("��� ����� � ��", str.getString());

	event.set(&date, Config4Event::Type_FiscalPassword);
	TEST_STRING_EQUAL("������ ��", event.getEventName(&event));
	event.getEventDescription(&event, &str);
	TEST_STRING_EQUAL("������������ ������ ��", str.getString());

	event.set(&date, Config4Event::Type_PrinterNotFound);
	TEST_STRING_EQUAL("������ ��������", event.getEventName(&event));
	event.getEventDescription(&event, &str);
	TEST_STRING_EQUAL("������� �� ������", str.getString());

	event.set(&date, Config4Event::Type_PrinterNoPaper);
	TEST_STRING_EQUAL("������ ��������", event.getEventName(&event));
	event.getEventDescription(&event, &str);
	TEST_STRING_EQUAL("� �������� ����������� ������", str.getString());
	return true;
}
