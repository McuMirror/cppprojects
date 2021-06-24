#ifndef COMMON_ECP_PROTOCOL_H
#define COMMON_ECP_PROTOCOL_H

#include "utils/include/NetworkProtocol.h"
#include "utils/include/Buffer.h"

namespace Ecp {

/*
---------------------
����������� ���������
---------------------
������ 1.0.0

��������� ����� ������� � ������:
	<STX><length><command><data><CRC>, ���
	<length> - ����� �����. � ����� �� ������ STX, <length> � <crc> �����;
	<command> - ��� �������;
	<CRC> - ���� ����������� �����

� �������� ETHERNET �� ������������ ����� STX � <crc>, �.�. ��� �������
�� ����������� ����������� � ������������ ������� ����� �� ���� ���� TCP/IP.
����� ������ ��� �� �� ���������� ������������� ACK, NAK, ENQ.
���� ������ ���������� ���������� � ����������� �� 3333 �����, ���������
� ���� ��������� � ������� ������ �� ���������� ���������� �������,
��� �� ���������� ������ ��������� ������ ������������ ��������.
����� ��������� ������ ���� ������ ������� ����������.
*/

#define ECP_PACKET_MAX_SIZE		200		// ��������
#define ECP_CONNECT_TIMEOUT		100
#define ECP_CONNECT_TRIES		500
#define ECP_DISCONNECT_TIMEOUT  100
#define ECP_KEEP_ALIVE_PERIOD	500
#define ECP_KEEP_ALIVE_TIMEOUT	ECP_KEEP_ALIVE_PERIOD * 20 // ������� ��������� 2, ����� ECP ����� ���������� � �������� �������
#define ECP_PACKET_TIMEOUT		1000	// ����� �������� ������ ������
#define ECP_CONFIRM_TIMEOUT		10000	// ����� �������� ������ (���������� ��-�� ������� ����-������, ���������)
#define ECP_BUSY_TIMEOUT        10
#define ECP_DELIVER_TIMEOUT		1

enum Control {
	Control_STX  = 0x02,
	Control_EOT  = 0x04,
	Control_ENQ  = 0x05,
	Control_ACK  = 0x06,
	Control_NAK  = 0x15,
};

enum Command {
	Command_Setup			= 0x01, // ������������ ������ ����������
	Command_UploadStart		= 0x02, // ������ �� ������ �������� ������
	Command_UploadData		= 0x03, // �������� ���������� �����
	Command_UploadEnd		= 0x04, // ���������� ��������
	Command_UploadCancel	= 0x05, // ������ �������� ������
	Command_DownloadStart	= 0x06, // ������ �� ������ �������� ������
	Command_DownloadData	= 0x07, // �������� ���������� �����
	Command_DownloadCancel	= 0x08, // ������ �������� ������
	Command_TableInfo		= 0x09, // �������� ���������� � �������
	Command_TableEntry		= 0x0A, // �������� ������ �������
	Command_DateTime		= 0x0B, // ������ ���� � ������� ������
	Command_ConfigReset		= 0x0C, // ����� ������������ ������
};

enum Error {
	Error_OK				= 0x00,  // ��� ������
	Error_Busy              = 0x01,
	Error_Disconnect        = 0x02,
	Error_TooManyTries		= 0x03,
	Error_Timeout			= 0x04,
	Error_WrongPacketSize   = 0x05,
	Error_WrongDestination	= 0x06,
	Error_WrongSource		= 0x07,
	Error_ServerError       = 0x08,
	Error_TableNotFound     = 0x09,
	Error_EntryNotFound     = 0x0A,
	Error_EndOfFile			= 0x0B,
};

class TableProcessor {
public:
	virtual bool isTableExist(uint16_t tableId) = 0;
	virtual uint32_t getTableSize(uint16_t tableId) = 0;
	virtual uint16_t getTableEntry(uint16_t tableId, uint32_t entryIndex, uint8_t *buf, uint16_t bufSize) = 0;
	virtual uint16_t getDateTime(uint8_t *buf, uint16_t bufSize) = 0;
};

class ClientPacketLayerInterface {
public:
	class Observer {
	public:
		virtual void procConnect() = 0;
		virtual void procRecvData(const uint8_t *data, uint16_t dataLen) = 0;
		virtual void procRecvError(Error error) = 0;
		virtual void procDisconnect() = 0;
	};

	virtual void setObserver(Observer *observer) = 0;
	virtual bool connect() = 0;
	virtual void disconnect() = 0;
	virtual bool sendData(const Buffer *data) = 0;
};

class ServerPacketLayerInterface {
public:
	class Observer {
	public:
		virtual void procConnect() = 0;
		virtual void procRecvData(const uint8_t *data, uint16_t dataLen) = 0;
		virtual void procError(uint8_t code) = 0;
		virtual void procDisconnect() = 0;
	};

	virtual void setObserver(Observer *observer) = 0;
	virtual void reset() = 0;
	virtual void shutdown() = 0;
	virtual void disconnect() = 0;
	virtual bool sendData(const Buffer *data) = 0;
};

class Crc {
public:
	void start(uint8_t byte) { crc = byte; }
	void add(uint8_t byte) { crc = crc ^ byte; }
	uint8_t getCrc() { return crc; }
	static uint8_t calc(const Buffer *data) {
		uint8_t dataLen = data->getLen();
		Crc crc;
		crc.start(dataLen);
		for(uint8_t i = 0; i < dataLen; i++) {
			crc.add((*data)[i]);
		}
		return crc.getCrc();
	}

private:
	uint8_t crc;
};

#pragma pack(push,1)
/*
 * ����������� ������� ��� �������������� ����������.
 */
struct Request {
	uint8_t command;
};

struct Response {
	uint8_t command;
	uint8_t errorCode;  // ��� ������
};

enum Destination {
	Destination_FirmwareGsm		= 0x01,
	Destination_FirmwareModem	= 0x02,
	Destination_Config			= 0x03,
	Destination_FirmwareScreen	= 0x04,
};

struct UploadStartRequest {
	uint8_t command;
	uint8_t destination;
	uint32_t dataSize;
};

struct UploadDataRequest {
	uint8_t command;
	uint8_t data[0];
};

enum Source {
	Source_Audit  = 0x01,
	Source_Config = 0x02,
};

struct DownloadStartRequest {
	uint8_t command;
	uint8_t source;
};

struct DownloadStartResponse {
	uint8_t command;
	uint8_t errorCode;
	uint32_t dataSize;
};

struct DownloadDataResponse {
	uint8_t command;
	uint8_t errorCode;
	uint8_t data[0];
};

enum Table {
	Table_Event = 0x0001,
};

struct TableInfoRequest {
	uint8_t command;
	uint16_t tableId;
};

struct TableInfoResponse {
	uint8_t command;
	uint8_t errorCode;
	uint32_t size;
};

struct TableInfo {
	uint32_t size;
};

struct TableEntryRequest {
	uint8_t command;
	uint16_t tableId;
	uint32_t entryIndex;
};

struct TableEntryResponse {
	uint8_t command;
	uint8_t errorCode;
	uint8_t data[0];
};

struct DateTimeResponse {
	uint8_t command;
	uint8_t errorCode;
	uint8_t year;
	uint8_t month;
	uint8_t day;
	uint8_t hour;
	uint8_t minute;
	uint8_t second;
};

#pragma pack(pop)

}

#endif
