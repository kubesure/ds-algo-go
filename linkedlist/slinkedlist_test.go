package linkedlist

import (
	"testing"
)

func TestNewSingleListList(t *testing.T) {
	sll := NewSingleLinkedList(1)
	if sll.Size() != 1 {
		t.Errorf("got %v and wanted 1", sll.Size())
	}
}

func TestAddFront(t *testing.T) {
	sll := NewSingleLinkedList(1)
	if sll.Size() != 1 {
		t.Errorf("got %v and wanted 1", sll.Size())
	}
	sll.AddFront(2)
	if sll.size != 2 {
		t.Errorf("got %v and wanted 2", sll.Size())
	}

	sll.AddFront(3)
	if sll.size != 3 {
		t.Errorf("got %v and wanted 3", sll.Size())
	}
}

func TestAddEnd(t *testing.T) {
	sll := NewSingleLinkedList(1)
	if sll.Size() != 1 {
		t.Errorf("got %v and wanted 1", sll.Size())
	}
	sll.AddEnd(3)
	if sll.Size() != 2 {
		t.Errorf("got %v and wanted 2", sll.Size())
	}

	sll.AddEnd(2)
	if sll.size != 3 {
		t.Errorf("got %v and wanted 3", sll.Size())
	}
}

func TestAddAt(t *testing.T) {
	sll := NewSingleLinkedList(1)
	if sll.Size() != 1 {
		t.Errorf("got %v and wanted 1", sll.Size())
	}
	sll.AddEnd(2)
	if sll.size != 2 {
		t.Errorf("got %v and wanted 2", sll.Size())
	}

	sll.AddEnd(3)
	if sll.Size() != 3 {
		t.Errorf("got %v and wanted 3", sll.Size())
	}

	sll.AddAt(4, 2)
	if sll.size != 4 {
		t.Errorf("got %v and wanted 4", sll.size)
	}
}

func TestPrint(t *testing.T) {
	sll := NewSingleLinkedList(0)
	sll.AddFront(1)
	sll.AddFront(2)
	sll.AddFront(3)
	sll.Print()
}
