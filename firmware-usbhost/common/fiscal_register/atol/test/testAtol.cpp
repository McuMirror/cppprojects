#if 0
#include "fiscal_register/atol/AtolProtocol.h"
#include "logger/include/Logger.h"
#include "test/include/Test.h"

void TestAtol::testAtol() {
/*
������: �������� ���� ������ <1F 00 FF 10 02 03 1A>.
1. ��������� �����, ������ DLE � ETX (10h � 03h): <1F 00 FF 10 10 02 10 03 1A>.
2. ��������� � ����� ETX:                         <1F 00 FF 10 10 02 10 03 1A 03>.
3. ������������ <CRC>: 1F XOR 00 XOR FF XOR 10 XOR 10 XOR 02 XOR 10 XOR 03 XOR 1A XOR 03 = E8.
4. ��������� � ������ STX:                     <02 1F 00 FF 10 10 02 10 03 1A 03>.
5. ��������� � ����� <CRC>:                    <02 1F 00 FF 10 10 02 10 03 1A 03 E8>.
���������� ������� ������������������ ����, ���������� ����� ���� 5.
*/
//todo: ���� ������� ������ ��� ����������� ������ �������
}
#endif
