package queue

import (
	"testing"
)

func TestPush(t *testing.T) {
	s := NewStack(1)
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
	s := NewStack(1)
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

func TestPeek(t *testing.T) {
	s := NewStack(1)
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
	s := NewStack(1)
	s.Push(1)
	s.Pop()

	empty := s.IsEmpty()

	if !empty {
		t.Errorf("got %v and wanted 0", s.Size())
	}
}
