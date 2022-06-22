package linkedlist

import "testing"

func TestRemmoveDups(t *testing.T) {
	ll := NewSingleLinkedList(4)
	ll.Insert(3)
	ll.Insert(4)
	ll.Insert(2)
	ll.Insert(1)
	ll.remmoveDuplicate()
	ll.IterateForward(andPrint)
}
