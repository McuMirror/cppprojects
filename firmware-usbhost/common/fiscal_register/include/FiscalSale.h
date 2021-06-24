#ifndef COMMON_FISCAL_SALE_H
#define COMMON_FISCAL_SALE_H

#include "evadts/EvadtsProtocol.h"
#include "timer/include/DateTime.h"
#include "utils/include/NetworkProtocol.h"
#include "utils/include/List.h"
#include "config.h"

namespace Fiscal {

enum Status {
	Status_None		 = 0, // ������� �� 0 ������ ��� ����� ���������
	Status_Complete	 = 1, // ��� ������ �������, ��������� ��������
	Status_InQueue	 = 2, // ��� �������� � �������, ��������� �� ��������
	Status_Unknown	 = 3, // ��������� ���������� ���� � ������� �� ��������
	Status_Error	 = 4, // ������ �������� ����
	Status_Overflow	 = 5, // ������� ���������� ����������� �����������
	Status_Manual	 = 6, // ��� �������������� �������
	Status_Server	 = 7,
};

enum PaymentType {
	Payment_Cash	 = 0,
	Payment_Cashless = 1,
	Payment_Token	 = 2,
	Payment_Ephor	 = 3,
	Payment_Sberbank = 4,
};

enum TaxSystem {
	TaxSystem_None	 = 0x00,
	TaxSystem_OSN    = 0x01, // ����� ���
	TaxSystem_USND   = 0x02, // ���������� �����
	TaxSystem_USNDMR = 0x04, // ���������� ����� ����� ������
	TaxSystem_ENVD   = 0x08, // ������ ����� �� ��������� �����
	TaxSystem_ESN    = 0x10, // ������ �������������������� �����
	TaxSystem_Patent = 0x20, // ��������� ������� ���������������
};

enum TaxRate {
	TaxRate_NDSNone	 = 0,
	TaxRate_NDS0	 = 1,
	TaxRate_NDS10	 = 2,
	TaxRate_NDS20	 = 3,
};

class Product {
public:
	StrParam<EvadtsProductIdSize> selectId;
	uint32_t wareId;
	StrParam<ConfigProductNameSize> name;
	uint32_t price;
	uint8_t  taxRate;
	uint8_t  quantity;
	uint8_t  discountType;
	uint32_t discountValue;

	Product();
	void set(const char *selestId, uint32_t wareId, const char *name, uint32_t price, uint8_t taxRate, uint8_t quantity);
	uint32_t getTaxValue();
};

#ifndef FISCAL_PRODUCT_NUM
#define FISCAL_PRODUCT_NUM 1
#endif

class ProductList {
public:
	ProductList();
	void set(ProductList *list);
	void set(const char *selectId, uint32_t wareId, const char *name, uint32_t price, uint8_t taxRate, uint8_t quantity);
	void add(const char *selectId, uint32_t wareId, const char *name, uint32_t price, uint8_t taxRate, uint8_t quantity);
	void clear();
	uint16_t getNum();
	Product *get(uint16_t index);
	uint32_t getPrice();
	uint32_t getTaxValue();

private:
	List<Product> list;
};

/*
��������� ������ (EVADTS 5.2.3.4):
����: 20
���: 15
������: 5

PA1*5*20
PA2*X1+1*Y1+20*X2+1*Y2+20
PA7*5*CA*1*40*X3+1*Y3+15
PA7*5*DA*1*40*X4+1*Y4+5

������ �� �� ����� ���� ���������. ������� ������� �����������.
������ ������ ������ ����� ��������:
CA - ���
DA - ������
DB - ����� (��� DC?)
TA - ������
 */
class Sale {
public:
	DateTime datetime;
	ProductList products;
	StrParam<EvadtsPaymentDeviceSize> device;
	uint8_t  priceList;
	uint8_t  paymentType;
	uint32_t credit;
	uint8_t  taxSystem;
	uint32_t taxValue;
	uint8_t  loyalityType;
	BinParam<uint8_t, LOYALITY_CODE_SIZE> loyalityCode;
	DateTime fiscalDatetime; // ���� ������������ �� ����� �����
	uint64_t fiscalRegister; // �� (����� ����������� ������������)
	uint64_t fiscalStorage;  // �� (����� ����������� ����������)
	uint32_t fiscalDocument; // �� (����� ����������� ���������)
	uint32_t fiscalSign;     // �� ��� ��� (���������� ������� ���������)

	Sale();
	void set(Sale *sale);
	void setProduct(const char *selectId, uint32_t wareId, const char *name, uint32_t price, uint8_t taxRate, uint8_t quantity);
	void addProduct(const char *selectId, uint32_t wareId, const char *name, uint32_t price, uint8_t taxRate, uint8_t quantity);
	void clearProducts();
	uint16_t getProductNum();
	Product *getProduct(uint16_t index);
	uint32_t getPrice();
};

}

#endif
