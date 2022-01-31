package linkedlist

import "testing"

func TestRemmoveDups(t *testing.T) {
	ll := NewSingleLinkedList(4)
	ll.AddFront(3)
	ll.AddFront(4)
	ll.AddFront(2)
	ll.AddFront(1)
	ll.remmoveDuplicate()
	ll.IterateForward(andPrint)
}
