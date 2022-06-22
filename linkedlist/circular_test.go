package linkedlist

import "testing"

func TestIsCircular(t *testing.T) {
	ll := NewEmptySingleLinkedList()
	ll.Insert(7)
	ll.Insert(6)
	ll.Insert(5)
	ll.Insert(4)
	ll.Insert(3)
	ll.Insert(2)
	ll.Insert(1)
	ll.nodeAt(5).Next = ll.nodeAt(2)
	//ll.IterateForward(andPrint)
	n := isCircular(ll)
	if n.Data != 3 {
		t.Errorf("This is a circular linked list wanted 3 got %v", n.Data)
	}
}
