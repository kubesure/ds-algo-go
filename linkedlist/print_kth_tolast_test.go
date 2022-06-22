package linkedlist

import (
	"testing"
)

func TestPrintKthFromLast(t *testing.T) {
	ll := NewSingleLinkedList(5)
	ll.Insert(4)
	ll.Insert(3)
	ll.Insert(2)
	ll.Insert(1)

	n := kthFromLast(ll.Head, 2)

	if n == nil {
		t.Fatalf("should not be nil")
	}

	if n.Data != 3 {
		t.Errorf("should be 3 but is %v", n.Data)
	}
}
