#ifndef COMMON_BUTTONS_BUTTON_H__
#define COMMON_BUTTONS_BUTTON_H__

#include <stdint.h>

class Button {
public:
	Button(uint16_t buttonId) : buttonId(buttonId), pressed(false) {}
	virtual ~Button() {}
	uint16_t getId() { return buttonId; }
	bool isPressed() { return pressed; }
	/**
	 * ��������� ���������� �� ��������� ������ � ������� ���������� ������ ������� ������.
	 * ������������ ��������:
	 * true - ���������� �� ����������������
	 * false - �� ����������
	 */
	bool isChange();

protected:
	virtual bool checkPressed() = 0;

private:
	uint16_t buttonId;
	bool pressed;
};

#endif
