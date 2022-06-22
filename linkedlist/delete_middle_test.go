package linkedlist

import (
	"fmt"
	"testing"
)

func TestDeleteFromMiddle(t *testing.T) {
	ll := NewSingleLinkedList(6)
	ll.Insert(5)
	ll.Insert(4)
	ll.Insert(3)
	ll.Insert(2)
	ll.Insert(1)
	ll.IterateForward(andPrint)
	fmt.Println()
	if !deleteFromMiddle(ll.Nodes()[2]) {
		t.Errorf("3 not deleted")
	}
	ll.IterateForward(andPrint)
}
