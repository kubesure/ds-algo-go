package linkedlist

import (
	"fmt"
	"log"
	"testing"
)

func TestNewSingleListList(t *testing.T) {
	sll := NewSingleLinkedList(1)
	if sll.size != 1 {
		t.Errorf("got %v and wanted 1", sll.size)
	}
}

func TestAddFront(t *testing.T) {
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

	for _, v := range sll.Nodes() {
		log.Println(v.data)
	}

}

func TestAddEnd(t *testing.T) {
	sll := NewSingleLinkedList(1)
	if sll.size != 1 {
		t.Errorf("got %v and wanted 1", sll.size)
	}
	sll.AddEnd(3)
	if sll.size != 2 {
		t.Errorf("got %v and wanted 2", sll.size)
	}

	sll.AddEnd(2)
	if sll.size != 3 {
		t.Errorf("got %v and wanted 3", sll.size)
	}

	sll.Iterate(func(n *Node) { fmt.Println(n.data) })
}

func TestAddAt(t *testing.T) {
	sll := NewSingleLinkedList(1)
	if sll.size != 1 {
		t.Errorf("got %v and wanted 1", sll.size)
	}
	sll.AddEnd(2)
	if sll.size != 2 {
		t.Errorf("got %v and wanted 2", sll.size)
	}

	sll.AddEnd(3)
	if sll.size != 3 {
		t.Errorf("got %v and wanted 3", sll.size)
	}

	sll.AddAt(4, 2)
	if sll.size != 4 {
		t.Errorf("got %v and wanted 4", sll.size)
	}

	sll.Iterate(func(n *Node) { fmt.Println(n.data) })
}
