package linkedlist

import (
	"testing"
)

func TestPrintKthFromLast(t *testing.T) {
	ll := NewSingleLinkedList(5)
	ll.AddFront(4)
	ll.AddFront(3)
	ll.AddFront(2)
	ll.AddFront(1)

	n := kthFromLast(ll.Head, 2)

	if n == nil {
		t.Fatalf("should not be nil")
	}

	if n.Data != 3 {
		t.Errorf("should be 3 but is %v", n.Data)
	}
}
