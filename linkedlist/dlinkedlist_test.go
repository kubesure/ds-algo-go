package linkedlist

import (
	"testing"
)

func TestAddFrontDL(t *testing.T) {
	d := new(DBLinkedList)
	d.AddFront(1)
	if d.Size() != 1 {
		t.Errorf("got %v and wanted 1", d.Size())
	}

	d.AddFront(2)
	if d.Size() != 2 {
		t.Errorf("got %v and wanted 2", d.Size())
	}

	d.AddFront(3)
	if d.Size() != 3 {
		t.Errorf("got %v and wanted 3", d.Size())
	}
}

func TestAddEndDL(t *testing.T) {
	dll := NewDBLinkedList(1)

	_, err := dll.AddEnd(3)

	if err != nil {
		t.Error(err)
	}

	if dll.Size() != 2 {
		t.Errorf("got %v and wanted 2", dll.Size())
	}

	dll.AddEnd(2)

	if dll.Size() != 3 {
		t.Errorf("got %v and wanted 3", dll.Size())
	}

}

func TestAddAtDL(t *testing.T) {
	dll := NewDBLinkedList(1)

	dll.AddFront(2)
	if dll.Size() != 2 {
		t.Errorf("got %v and wanted 2", dll.Size())
	}

	dll.AddFront(3)
	if dll.Size() != 3 {
		t.Errorf("got %v and wanted 3", dll.Size())
	}

	dll.AddAt(4, 1)
	if dll.Size() != 4 {
		t.Errorf("got %v and wanted 4", dll.Size())
	}
}
