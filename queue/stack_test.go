package queue

import (
	"testing"
)

func TestPush(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)

	if s.Size() != 2 {
		t.Errorf("got %v and wanted 2", s.Size())
	}

	s.Push(3)
	if s.Size() != 3 {
		t.Errorf("got %v and wanted 2", s.Size())
	}

}

func TestPop(t *testing.T) {
	s := NewStack()
	s.Push(1)

	if s.Size() != 1 {
		t.Errorf("got %v and wanted 1", s.Size())
	}
	s.Pop()

	if s.Size() != 0 {
		t.Errorf("got %v and wanted 0", s.Size())
	}

	s.Push(1)
	s.Pop()

	if s.Pop() != nil {
		t.Errorf("got %v and wanted nil", s.Size())
	}
}

func TestPop2(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Pop()
	s.Print()
}

func TestPeek(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)

	e := s.Peek()

	if e.value != 2 {
		t.Errorf("got %v and wanted 2", s.Size())
	}

	if s.size != 2 {
		t.Errorf("got %v and wanted 2", s.Size())
	}
}

func TestIsEmpty(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Pop()

	empty := s.IsEmpty()

	if !empty {
		t.Errorf("got %v and wanted 0", s.Size())
	}
}

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

func TestMinStack(t *testing.T) {
	ms := NewMinStack()
	ms.Push(50)
	ms.Push(25)
	ms.Push(10)
	if ms.Min().value != 10 {
		t.Errorf("Should have been 10 after push")
	}

	ms.Pop()
	if ms.Peek().value != 25 {
		t.Errorf("Should have been 25 after 1 pop")
	}

	ms.Pop()

	if ms.Min().value != 50 {
		t.Errorf("Should have been 50 after 2 pops")
	}

}
