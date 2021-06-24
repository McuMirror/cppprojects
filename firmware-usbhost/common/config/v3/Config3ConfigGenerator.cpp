#include "Config3ConfigGenerator.h"
#include "dex/DexProtocol.h"
#include "tcpip/include/TcpIpUtils.h"
#include "utils/include/Hex.h"
#include "logger/include/Logger.h"

Config3ConfigGenerator::Config3ConfigGenerator(Config3Modem *config) :
	EvadtsGenerator(config->getAutomat()->getCommunictionId()),
	config(config),
	cert(CERT_SIZE, CERT_SIZE),
	parser("")
{
	this->product = config->getAutomat()->createIterator();
}

Config3ConfigGenerator::~Config3ConfigGenerator() {
	delete this->product;
}

void Config3ConfigGenerator::reset() {
	state = State_Header;
	next();
}

void Config3ConfigGenerator::next()  {
	switch(state) {
	case State_Header: {
		generateHeader();
		state = State_Main;
		break;
	}
	case State_Main: generateMain(); break;
	case State_FiscalSection: generateFiscalSection(); break;
	case State_AuthPublicKey: generateAuthPublicKey(); break;
	case State_AuthPrivateKey: generateAuthPrivateKey(); break;
	case State_SignPrivateKey: generateSignPrivateKey(); break;
	case State_PriceLists: generatePriceLists(); break;
	case State_Products: generateProducts(); break;
	case State_Custom: generateCustom(); break;
	case State_Footer: {
		generateFooter();
		state = State_Complete;
		break;
	}
	case State_Complete: break;
	}
}

bool Config3ConfigGenerator::isLast() {
	return state == State_Complete;
}

/*
Example of configuration data:
DXS*XYZ1234567*VA*V0/6*1
ST*001*0001
// AC2 - ��������� ����������
// 1 - IMEI (�������� ������)
// 2 - GPRS APN
// 3 - GPRS �����
// 4 - GPRS ������
// 5 - ��������� ���� (MDB/EXECUTIVE)
AC2*123456789012345*internet*gdata*gdata*1
// AC3 - ��������� ����������� ������������
// 1 - ������������� ���
// 2 - ��������� ���
// 3 - IP-����� ���
// 4 - ���� ���
// 5 - IP-����� ���
// 6 - ���� ���
AC3*1*1*192.168.1.201*5555*91.107.67.212*7790
// IC1 - �������� ��������
// 1 - �������� �����
// 6 - ����� ��
IC1*EPHOR00001*135
// IC4 - �������� ������ (����������)
// 1 - ��������� �����
// 2 - ��� ������
// 3 - �������� ������
IC4*2*0000
// PC1 - �������� ��������
// 1 - ������������� �������� (������)
// 2 - ���� ��������
// 3 - ������������� ������
// 4 - ������������ �������� ������
// 5 - Standard Filling Level
// 6 - Standard Dispensed Quantity
// 7 - Selection status (0 - ��������, 1 - ���������)
PC1*10******1
PC1*11**�����****0
// PC7 - ���� ��������
// 1 - ������������� �������� (������)
// 2 - ������ ������ (CA - ��������, DA - ��������� 1, DB - ��������� 2)
// 3 - ����� ����� ����� � ������� ������
// 4 - �������� ����
PC7*11*CA*0*3500
PC7*11*DA*1*3500
PC7*11*DA*2*3500
PC7*11*DA*3*3500
G85*1234 (example G85 CRC is 1234)
SE*9*0001
DXE*1*1
*/
void Config3ConfigGenerator::generateMain() {
	Config1Boot *boot = config->getBoot();
	Config2Fiscal *fiscal = config->getFiscal();
	Config3Automat *automat = config->getAutomat();

	startBlock();
	*str << "IC1******" << automat->getAutomatId(); finishLine();
	*str << "IC4*" << automat->getDecimalPoint()
		 << "*" << automat->getCurrency()
		 << "**" << automat->getTaxSystem()
		 << "*" << automat->getMaxCredit()
		 << "*" << automat->getExt1Device()
		 << "*" << automat->getEvadts()
		 << "*" << automat->getCashlessNumber()
		 << "*" << automat->getScaleFactor()
		 << "*" << automat->getCategoryMoney()
		 << "*" << automat->getShowChange()
		 << "*" << automat->getPriceHolding()
		 << "*" << automat->getMultiVend()
		 << "*" << automat->getCreditHolding()
		 << "*" << automat->getCashless2Click()
		 << "*" << automat->getFiscalPrinter()
		 << "*" << automat->getCashlessMaxLevel(); finishLine();
	*str << "AC2*" << boot->getImei()
		 << "*" << boot->getGprsApn()
		 << "*" << boot->getGprsUsername()
		 << "*" << boot->getGprsPassword()
		 << "*" << automat->getPaymentBus()
		 << "*" << boot->getHardwareVersion()
		 << "*" << boot->getFirmwareVersion()
		 << "*" << boot->getFirmwareRelease()
		 << "*" << boot->getGsmFirmwareVersion(); finishLine();
	*str << "FC1*" << fiscal->getKkt()
		 << "*" << fiscal->getKktInterface()
		 << "*" << fiscal->getKktAddr()
		 << "*" << fiscal->getKktPort()
		 << "*" << fiscal->getOfdAddr()
		 << "*" << fiscal->getOfdPort(); finishLine();
	finishBlock();

	state = State_FiscalSection;
}

void Config3ConfigGenerator::generateFiscalSection() {
	Config2Fiscal *fiscal = config->getFiscal();
	Config3Automat *automat = config->getAutomat();

	startBlock();
	*str << "FC2*" << fiscal->getINN()
		 << "*" << automat->getAutomatId()
		 << "*" << fiscal->getPointName()
		 << "*" << fiscal->getPointAddr(); finishLine();
	finishBlock();

	gotoStateAuthPublicKey();
}

void Config3ConfigGenerator::gotoStateAuthPublicKey() {
	config->getFiscal()->getAuthPublicKey()->load(&cert);
	parser.init(cert.getString(), cert.getLen());
	state = State_AuthPublicKey;
}

void Config3ConfigGenerator::generateAuthPublicKey() {
	char buf[200];
	uint16_t len = parser.getValue("\r\n", buf, sizeof(buf));
	parser.skipEqual("\r\n\t ");
	startBlock();
	*str << "FC3*"; str->addStr(buf, len); finishLine();
	finishBlock();

	if(parser.hasUnparsed() == false) {
		gotoStateAuthPrivateKey();
	}
}

void Config3ConfigGenerator::gotoStateAuthPrivateKey() {
	config->getFiscal()->getAuthPrivateKey()->load(&cert);
	parser.init(cert.getString(), cert.getLen());
	state = State_AuthPrivateKey;
}

void Config3ConfigGenerator::generateAuthPrivateKey() {
	char buf[200];
	uint16_t len = parser.getValue("\r\n", buf, sizeof(buf));
	parser.skipEqual("\r\n\t ");
	startBlock();
	*str << "FC4*"; str->addStr(buf, len); finishLine();
	finishBlock();

	if(parser.hasUnparsed() == false) {
		gotoStateSignPrivateKey();
	}
}

void Config3ConfigGenerator::gotoStateSignPrivateKey() {
	config->getFiscal()->getSignPrivateKey()->load(&cert);
	parser.init(cert.getString(), cert.getLen());
	state = State_SignPrivateKey;
}

void Config3ConfigGenerator::generateSignPrivateKey() {
	char buf[200];
	uint16_t len = parser.getValue("\r\n", buf, sizeof(buf));
	parser.skipEqual("\r\n\t ");
	startBlock();
	*str << "FC5*"; str->addStr(buf, len); finishLine();
	finishBlock();

	if(parser.hasUnparsed() == false) {
		gotoStateGeneratePriceLists();
	}
}

void Config3ConfigGenerator::gotoStateGeneratePriceLists() {
	for(uint16_t i = 0; i < config->getAutomat()->getPriceListNum(); i++) {
		Config3PriceIndex *index = product->getPriceIdByIndex(i);
		if(index->type == Config3PriceIndexType_Time) {
			state = State_PriceLists;
			return;
		}
	}

	if(product->first() == false) {
		state = State_Footer;
		return;
	} else {
		state = State_Products;
		return;
	}
}

void Config3ConfigGenerator::generatePriceLists() {
	startBlock();
	for(uint16_t i = 0; i < config->getAutomat()->getPriceListNum(); i++) {
		Config3PriceIndex *index = product->getPriceIdByIndex(i);
		if(index->type == Config3PriceIndexType_Time) {
			WeekTable *week = index->timeTable.getWeek();
			TimeInterval *interval = index->timeTable.getInterval();
			*str << "LC2*" << index->device.get() << "*" << index->number << "*" << week->getValue()
				 << "*" << interval->getTime()->getHour() << ":" << interval->getTime()->getMinute() << ":" << interval->getTime()->getSecond()
				 << "*" << interval->getInterval(); finishLine();
		}
	}
	finishBlock();

	if(product->first() == false) {
		state = State_Footer;
		return;
	} else {
		state = State_Products;
		return;
	}
}

void Config3ConfigGenerator::generateProducts() {
	startBlock();
	*str << "PC1*" << product->getId() << "*" << product->getPrice("CA", 0)->getPrice() << "*" << product->getName(); finishLine();
	*str << "PC9*" << product->getId() << "*";
	if(product->getCashlessId() != Config3ProductIndexList::UndefinedIndex) { *str << product->getCashlessId(); }
	*str << "*" << product->getTaxRate() << "*";
	if(product->getWareId() > 0) { *str << product->getWareId(); }
	finishLine();
	generateProductPrices();
	finishBlock();

	if(product->next() == false) {
		state = State_Custom;
	}
}

void Config3ConfigGenerator::generateProductPrices() {
	for(uint16_t i = 0; i < config->getAutomat()->getPriceListNum(); i++) {
		Config3PriceIndex *index = product->getPriceIdByIndex(i);
		if(index->type == Config3PriceIndexType_None) {
			continue;
		}
		Config3Price *price = product->getPriceByIndex(i);
		*str << "PC7*" << product->getId();
		*str << "*" << index->device.get() << "*" << index->number << "*";
		*str << price->getPrice();
		finishLine();
	}
}

void Config3ConfigGenerator::generateCustom() {
	Config3Automat *automat = config->getAutomat();

	startBlock();
	*str << "MC5*0*INTERNET*" << automat->getInternetDevice(); finishLine();
	*str << "MC5*1*EXT1*" << automat->getExt1Device(); finishLine();
	*str << "MC5*2*EXT2*" << automat->getExt2Device(); finishLine();
	*str << "MC5*3*USB1*" << automat->getUsb1Device(); finishLine();
	*str << "MC5*4*QRTYPE*" << automat->getQrType(); finishLine();
	*str << "MC5*5*ETH1MAC*"; dataToHex(automat->getEthMac(), automat->getEthMacLen(), str); finishLine();
	*str << "MC5*6*ETH1ADDR*" << LOG_IPADDR(automat->getEthAddr()); finishLine();
	*str << "MC5*7*ETH1MASK*" << LOG_IPADDR(automat->getEthMask()); finishLine();
	*str << "MC5*8*ETH1GW*" << LOG_IPADDR(automat->getEthGateway()); finishLine();
	finishBlock();

	state = State_Footer;
}
