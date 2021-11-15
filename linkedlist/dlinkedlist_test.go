package linkedlist

import (
	"testing"
)

func TestAddFrontDL(t *testing.T) {
	d := new(DBLinkedList)
	d.AddFront(1)
	if d.Length() != 1 {
		t.Errorf("got %v and wanted 1", d.Length())
	}

	d.AddFront(2)
	if d.Length() != 2 {
		t.Errorf("got %v and wanted 2", d.Length())
	}

	d.AddFront(3)
	if d.Length() != 3 {
		t.Errorf("got %v and wanted 3", d.Length())
	}
}

func TestAddEndDL(t *testing.T) {
	dll := NewDBLinkedList(1)

	_, err := dll.AddEnd(3)

	if err != nil {
		t.Error(err)
	}

	if dll.Length() != 2 {
		t.Errorf("got %v and wanted 2", dll.Length())
	}

	dll.AddEnd(2)

	if dll.Length() != 3 {
		t.Errorf("got %v and wanted 3", dll.Length())
	}

}

func TestAddAtDL(t *testing.T) {
	dll := NewDBLinkedList(1)

	dll.AddFront(2)
	if dll.Length() != 2 {
		t.Errorf("got %v and wanted 2", dll.Length())
	}

	dll.AddFront(3)
	if dll.Length() != 3 {
		t.Errorf("got %v and wanted 3", dll.Length())
	}

	dll.AddAt(4, 1)
	if dll.Length() != 4 {
		t.Errorf("got %v and wanted 4", dll.Length())
	}
}
