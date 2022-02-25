package linkedlist

import (
	"fmt"
	"testing"
)

func TestPrintKthToLastRecursion(t *testing.T) {
	ll := NewSingleLinkedList(3)
	ll.AddFront(2)
	ll.AddFront(1)
	//ll.AddFront(4)
	//ll.AddFront(5)
	fmt.Println(kthToLast(ll.Head, 2))
}
