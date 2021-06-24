#ifndef __QUEUE_H
#define __QUEUE_H

#ifdef AVR
#include <util/atomic.h>
#elif defined(ARM)
#include "common.h"
#else
#include <stdint.h>
#endif

/*
	����� - ������ ��������� ���������������� ������� ���������� ��������. ������� ������ ���� ������� ����� new. ������ �������, ������ ����.
*/

template <class T>
class Queue
{
	private:
	  T **ptr;
	  uint8_t size;
	  volatile uint8_t readIndex;
	  volatile uint8_t writeIndex; // ������ ������, �� ��������� �� 1 ������, ��� ������ ������. ������� ��� �������� ������ ����� ������� � �������.
	  
	  uint8_t lastWriteIndex() {
		  if (writeIndex==0) return size-1;
		  return writeIndex-1;
	  }
	  
	  uint8_t nextWriteIndex() {
		  if (writeIndex == size-1) return 0;
		  return writeIndex + 1;
	  }
	  
	  uint8_t nextReadIndex() {
		  if (readIndex == size-1) return 0;
		  return readIndex + 1;
	  }
	  
	public:
	  Queue(uint8_t size): size(size+1), readIndex(0), writeIndex(1)  {
		  	// ��������� �� ������ ���������� �������
			if (size < 2) size = 2;
		  	ptr = new T*[size];
	  }
	  
	  ~Queue() {
			clear();
			delete []ptr;
	  }
	  
	  void clear() {
		  T *t = pop();
		  while (t) {
			  delete t;
			  t = pop();
		  }
		  readIndex = 0;
		  writeIndex = 1;
	  }
	  
	  bool isEmpty() {
		 return (readIndex == lastWriteIndex());
	  }
		
		bool isFull() {
			return writeIndex == readIndex;
		}
		 
	  bool push(T *p) {
		  // ���� ������� ������ ������, ������ ����� �����������. ������ ������ �� ���������
		  if (writeIndex == readIndex) return false;

		  ptr[lastWriteIndex()] = p;
		  writeIndex = nextWriteIndex();
		  return true;
	   }
		  
	  T *pop() {
		 // ��������. ��� �������� �� �����
		 T *result = NULL;
		 if (!isEmpty()) 
		 { 
			 result = ptr[readIndex];
			 readIndex = nextReadIndex();			 
		 }
		 return result; 
	  }
};

#endif
