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
	sll.Insert(2)
	if sll.Length() != 2 {
		t.Errorf("got %v and wanted 2", sll.Length())
	}

	sll.Insert(3)
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

func TestRemoveTail(t *testing.T) {
	sll := NewSingleLinkedList(0)
	sll.Insert(1)
	sll.Insert(2)
	sll.Insert(3)
	sll.IterateForward(andPrint)
	sll.RemoveAt(3)
	if sll.size != 3 {
		t.Errorf("size should be 3 while its %v ", sll.size)
	}
	fmt.Println()
	sll.IterateForward(andPrint)

}

func TestRemoveHead(t *testing.T) {
	sll := NewSingleLinkedList(0)
	sll.Insert(1)
	sll.Insert(2)
	sll.Insert(3)
	sll.IterateForward(andPrint)
	sll.RemoveAt(0)
	if sll.size != 3 {
		t.Errorf("size should be 3 while its %v ", sll.size)
	}
	fmt.Println()
	sll.IterateForward(andPrint)
}

func TestPrint(t *testing.T) {
	sll := NewSingleLinkedList(0)
	sll.Insert(1)
	sll.Insert(2)
	sll.Insert(3)
	sll.Print()
}

func TestIterateForward(t *testing.T) {
	sll := NewSingleLinkedList(4)
	sll.Insert(1)
	sll.Insert(2)
	sll.Insert(3)
	sll.IterateForward(andPrint)
}
