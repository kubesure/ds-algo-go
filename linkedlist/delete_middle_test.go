package linkedlist

import (
	"fmt"
	"testing"
)

func TestDeleteFromMiddle(t *testing.T) {
	ll := NewSingleLinkedList(6)
	ll.AddFront(5)
	ll.AddFront(4)
	ll.AddFront(3)
	ll.AddFront(2)
	ll.AddFront(1)
	ll.IterateForward(andPrint)
	fmt.Println()
	if !deleteFromMiddle(ll.Nodes()[2]) {
		t.Errorf("3 not deleted")
	}
	ll.IterateForward(andPrint)
}
