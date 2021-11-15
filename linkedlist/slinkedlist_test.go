package linkedlist

import (
	"fmt"
	"testing"
)

func TestNewSingleListList(t *testing.T) {
	sll := NewSingleLinkedList(1)
	if sll.Length() != 1 {
		t.Errorf("got %v and wanted 1", sll.Length())
	}
}

func TestAddFront(t *testing.T) {
	sll := NewSingleLinkedList(1)
	if sll.Length() != 1 {
		t.Errorf("got %v and wanted 1", sll.Length())
	}
	sll.AddFront(2)
	if sll.Length() != 2 {
		t.Errorf("got %v and wanted 2", sll.Length())
	}

	sll.AddFront(3)
	if sll.Length() != 3 {
		t.Errorf("got %v and wanted 3", sll.Length())
	}
}

func TestAddEnd(t *testing.T) {
	sll := NewSingleLinkedList(1)

	if sll.Length() != 1 {
		t.Errorf("got %v and wanted 1", sll.Length())
	}
	sll.AddEnd(3)

	if sll.Length() != 2 {
		t.Errorf("got %v and wanted 2", sll.Length())
	}

	sll.AddEnd(2)
	if sll.Length() != 3 {
		t.Errorf("got %v and wanted 3", sll.Length())
	}
}

func TestAddAt(t *testing.T) {
	sll := NewSingleLinkedList(1)
	if sll.Length() != 1 {
		t.Errorf("got %v and wanted 1", sll.Length())
	}
	sll.AddEnd(2)
	if sll.Length() != 2 {
		t.Errorf("got %v and wanted 2", sll.Length())
	}

	sll.AddEnd(3)
	if sll.Length() != 3 {
		t.Errorf("got %v and wanted 3", sll.Length())
	}

	sll.AddAt(4, 2)
	if sll.Length() != 4 {
		t.Errorf("got %v and wanted 4", sll.Length())
	}
}

func TestPrint(t *testing.T) {
	sll := NewSingleLinkedList(0)
	sll.AddFront(1)
	sll.AddFront(2)
	sll.AddFront(3)
	sll.Print()
}

func TestIterateForward(t *testing.T) {
	sll := NewSingleLinkedList(0)
	sll.AddFront(1)
	sll.AddFront(2)
	sll.AddFront(3)
	sll.IterateForward(print)
}

func print(n *Node) bool {
	fmt.Printf("%v ,", n.Data)
	return true
}
