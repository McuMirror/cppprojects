#ifndef COMMON_PLATFORM_ARM_PLATFORM_H
#define COMMON_PLATFORM_ARM_PLATFORM_H

#include <platform/arm/_atomic.h>

#include "cmsis_boot/stm32f4xx_conf.h"

#define __ASM  __asm

// ������: PROBE_INIT(D, 1, GPIO_OType_PP); ������������� ���� �����. type = GPIO_OType_PP, GPIO_OType_OD
#define PROBE_INIT(port, pin, type) {\
	RCC_AHB1PeriphClockCmd(RCC_AHB1Periph_GPIO##port, ENABLE);\
	GPIO_InitTypeDef gpio;\
	GPIO_StructInit(&gpio);\
	gpio.GPIO_OType = type;\
	gpio.GPIO_Mode = GPIO_Mode_OUT;\
	gpio.GPIO_Pin = GPIO_Pin_##pin;\
	GPIO_Init(GPIO##port, &gpio);\
}

// ������: PROBE(D, 1); ����������� ������ �� ������
#define PROBE_TOGGLE(port, pin) {\
	GPIO_ToggleBits(GPIO##port, GPIO_Pin_##pin);\
}

#define PROBE_ON(port, pin) {\
	GPIO_SetBits(GPIO##port, GPIO_Pin_##pin);\
}

#define PROBE_OFF(port, pin) {\
	GPIO_ResetBits(GPIO##port, GPIO_Pin_##pin);\
}

#endif
