package linkedlist

import (
	"fmt"
	"testing"
)

func TestPartition(t *testing.T) {
	ll := NewSingleLinkedList(0)
	ll.AddFront(8)
	ll.AddFront(1)
	ll.AddFront(5)
	ll.AddFront(2)
	ll.AddFront(7)
	ll.AddFront(3)
	ll.IterateForward(andPrint)
	fmt.Println()
	nll := ll.Partition(5)
	nll.IterateForward(andPrint)
}
