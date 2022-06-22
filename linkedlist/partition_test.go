package linkedlist

import (
	"fmt"
	"testing"
)

func TestPartition(t *testing.T) {
	ll := NewSingleLinkedList(0)
	ll.Insert(8)
	ll.Insert(1)
	ll.Insert(5)
	ll.Insert(2)
	ll.Insert(7)
	ll.Insert(3)
	ll.IterateForward(andPrint)
	fmt.Println()
	nll := ll.Partition(5)
	nll.IterateForward(andPrint)
}
