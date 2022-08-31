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
	s.Push(2)
	s.Push(3)

	if s.Size() != 3 {
		t.Errorf("got %v and wanted size 1", s.Size())
	}

	e := s.Pop()

	if s.Size() != 2 {
		t.Errorf("got %v and wanted size 1", s.Size())
	}

	if e.value != 3 {
		t.Errorf("got %v and wanted 1", e.value)
	}

	e = s.Pop()

	if s.Size() != 1 {
		t.Errorf("got %v and wanted size 1", s.Size())
	}

	if e.value != 2 {
		t.Errorf("got %v and wanted value 2", e.value)
	}

	e = s.Pop()

	if s.Size() != 0 {
		t.Errorf("got %v and wanted size 1", s.Size())
	}

	if e.value != 1 {
		t.Errorf("got %v and wanted value 2", e.value)
	}

}

func TestPeek(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)

	e := s.Peek()

	if e.value != 2 {
		t.Errorf("got %v and wanted 2", e.value)
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
