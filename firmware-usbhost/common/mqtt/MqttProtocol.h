#ifndef COMMON_MQTT_PROTOCOL_H
#define COMMON_MQTT_PROTOCOL_H

#include "utils/include/StringBuilder.h"
#include "utils/include/Buffer.h"
#include "utils/include/CodePage.h"
#include "utils/include/NetworkProtocol.h"

#include <stdint.h>

namespace Mqtt {

#define MQTT_SOCKET_TIMEOUT 5000
#define MQTT_PING_TIMEOUT 5000
#define MQTT_TRY_CONNECT_DELAY 10000
#define MQTT_TOPIC_SIZE 128
#define MQTT_PACKET_SIZE 512
#define MQTT_TOPIC_SUBSRIBERS 10

//-----------------------
// ����������� ���������
//-----------------------

// MQTT ���� �������
enum PacketType {
    PacketType_Reserved1Type  = 0x00,
    PacketType_Connect        = 0x10,
    PacketType_ConnAck        = 0x20,
    PacketType_Publish        = 0x30,
    PacketType_PubAck         = 0x40,
    PacketType_PubRec         = 0x50,
    PacketType_PubRel         = 0x60,
    PacketType_PubComp        = 0x70,
    PacketType_Subscribe      = 0x80,
    PacketType_SubAck         = 0x90,
    PacketType_Unsubscribe    = 0xA0,
    PacketType_UnsubAck       = 0xB0,
    PacketType_PingRec        = 0xC0,
    PacketType_PingResp       = 0xD0,
    PacketType_Disconnect     = 0xE0,
    PacketType_Reserved2Type  = 0xF0
};

// MQTT ���� ������, ������������ ����� ����������.
enum ErrorCode {
    ConnAccepted                     = 0x00,
    ConnRefusedBadProtocolVersion    = 0x01,
    ConnRefusedIDRejected            = 0x02,
    ConnRefusedServerUnavailable     = 0x03,
    ConnRefusedBadUsernameOrPassword = 0x04,
    ConnRefusedNotAuthorized         = 0x05
};

// ������� ��������
enum QoS {
    QoS_0 = 0,
    QoS_1 = 1,
    QoS_2 = 2,
    QoS_3 = 3,
};

#pragma pack(push,1)

// ConnectPacket ���������� ������� ������������
// �������� ��� ����� �� ������ MQTT.
struct ConnectPacket {
    const char* protocolName;
    uint8_t     protocolLevel;
    uint16_t    keepalive = 100;
    const char* clientId = "/:ephor1";
    bool        cleanSession = true;
    bool        willFlag = false;
    const char* willTopic = "";
    const char* willMessage = "";
    uint8_t     willQOS = 0;
    bool        willRetain = true;
    const char* username = "";
    const char* password = "";
};

// ConnAckPacket ������������� ������������
struct ConnAckPacket {
    bool sessionPresent = false;
    uint8_t returnCode;
};

// PublishPacket - ����������� �����, ������������ �������� ��� �������� ��� ��������� ���
// �������� �������� ��������.
struct PublishPacket {
    uint16_t id = 0;
    bool     dup = false;
    QoS      qos = QoS_1;
    bool     retain = false;
    String   topic;
	Buffer   payload = MQTT_PACKET_SIZE;
};

// PubAckPacket - ����������� �����, ������������ �������� ��� �������� � ����� ��
// ������ PUBLISH, ����� QOS ��� ���������� ������ 1.
struct PubAckPacket {
    uint16_t id;
};

// Topic �������� ������
struct Topic {
	String name;
	QoS qos = QoS_0;
};

// SubscribePacket ����������� � ������ ��� ��������� ���������
struct SubscribePacket {
    uint16_t id = 0;
	Topic topics[MQTT_TOPIC_SUBSRIBERS];
};

// SubAckPacket ������������� ��������
struct SubAckPacket {
    uint16_t id = 0;
	uint8_t returnCodes[MQTT_TOPIC_SUBSRIBERS];
};

#pragma pack(pop)


}

#endif
