#ifndef __SD_H
#define __SD_H

/*
 * ���: ���������
 * ��������: ������, ��� ������ � SD ���������. ����� - 1 ���.
 *
 *
 * �������� �������: http://cxem.net/mc/mc325.php
 */

#include <stdint.h>
#include "common/utils/include/Fifo.h"
#include "defines.h"
#include "common.h"


class SD {
	public:
	  static SD *get();
	  SD();

	  // ������� SD �����.
	  bool hasCard();


	  bool test();

private:

};

#endif
