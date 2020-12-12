package linkedlist

import (
	"testing"
)

func TestNewSingleListList(t *testing.T) {
	sll := NewSingleLinkedList(1)
	if sll.size != 1 {
		t.Errorf("got %v and wanted 1", sll.size)
	}
}

func TestAddSingleListList(t *testing.T) {
	sll := NewSingleLinkedList(1)
	if sll.size != 1 {
		t.Errorf("got %v and wanted 1", sll.size)
	}
	sll.AddFront(2)
	if sll.size != 2 {
		t.Errorf("got %v and wanted 2", sll.size)
	}

	sll.AddFront(3)
	if sll.size != 3 {
		t.Errorf("got %v and wanted 3", sll.size)
	}

	sll.IteratePrint()

}
