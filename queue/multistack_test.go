package queue

import "testing"

func TestPushMultiStack(t *testing.T) {
	ms := NewMultiStack(3, 5)
	ms.Push(1, 1)
	ms.Push(1, 2)
	ms.Push(1, 3)
	ms.Push(1, 4)
	ms.Push(1, 5)

	if ms.Size(1) != 5 {
		t.Errorf("size of stack 1 should be 5")
	}

	ms.Push(2, 1)
	ms.Push(2, 2)
	ms.Push(2, 3)
	ms.Push(2, 4)
	ms.Push(2, 5)

	if ms.Size(2) != 5 {
		t.Errorf("size of stack 2 should be 5")
	}

	ms.Push(3, 1)
	ms.Push(3, 2)
	ms.Push(3, 3)
	ms.Push(3, 4)
	ms.Push(3, 5)

	if ms.Size(3) != 5 {
		t.Errorf("size of stack 3 should be 5")
	}

	err := ms.Push(1, 6)
	if err == nil {
		t.Errorf("Should not be able to push in stack as 6th element")
	}

	if err1 := ms.Push(-1, 1); err1 == nil {
		t.Errorf("should be error cant push on a less than 1 stack")
	}

}

func TestPopMultiStack(t *testing.T) {
	ms := NewMultiStack(3, 5)
	ms.Push(1, 1)
	ms.Push(1, 2)
	ms.Push(1, 3)
	ms.Push(1, 4)
	ms.Push(1, 5)
	if ms.Size(1) != 5 {
		t.Errorf("size should be 5")
	}

	ms.Pop(1)

	if ms.Size(1) != 4 {
		t.Errorf("size should be 4")
	}

	ms.Push(1, 5)

	if ms.Size(1) != 5 {
		t.Errorf("size should be 5")
	}

	ms.Push(2, 1)
	ms.Push(2, 2)
	ms.Push(2, 3)
	ms.Push(2, 4)
	ms.Push(2, 5)

	if ms.Size(2) != 5 {
		t.Errorf("size should be 5")
	}

	ms.Pop(2)

	if ms.Size(2) != 4 {
		t.Errorf("size should be 4")
	}

	ms.Push(3, 1)
	ms.Push(3, 2)
	ms.Push(3, 3)
	ms.Push(3, 4)
	ms.Push(3, 5)

	if ms.Size(3) != 5 {
		t.Errorf("size should be 5")
	}

	ms.Pop(3)

	if ms.Size(3) != 4 {
		t.Errorf("size should be 4")
	}

	if _, err1 := ms.Pop(-1); err1 == nil {
		t.Errorf("should be error cant pop on a less than 1 stack")
	}
}

func TestPeekMultiStack(t *testing.T) {

	ms := NewMultiStack(3, 5)
	ms.Push(1, 1)
	ms.Push(1, 2)
	ms.Push(1, 3)
	ms.Push(1, 4)
	ms.Push(1, 5)
	if e, _ := ms.Peek(1); e.value != 5 {
		t.Errorf("size should be 5")
	}

	ms.Pop(1)
	ms.Pop(1)
	ms.Pop(1)
	ms.Pop(1)
	ms.Pop(1)

	if _, err := ms.Peek(1); err == nil {
		t.Errorf("should not be able to peek on empty stack")
	}

	if _, err1 := ms.Peek(-1); err1 == nil {
		t.Errorf("should be error cant peek on a less than 1 stack")
	}
}
