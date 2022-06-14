package linkedlist

import "testing"

func TestIntersect1(t *testing.T) {
	ll1 := NewEmptySingleLinkedList()
	ll1.AddFront(1)
	ll1.AddFront(2)
	ll1.AddFront(7)
	ll1.AddFront(9)
	ll1.AddFront(5)
	ll1.AddFront(1)
	ll1.AddFront(3)

	ll2 := NewEmptySingleLinkedList()
	ll2.AddFront(6)
	ll2.AddFront(4)
	ll2.Tail.Next = ll1.nodeAt(4)
	ll2.Tail = ll1.Tail
	ll2.size += 3

	res := areIntersecting(ll1, ll2)
	if !res {
		t.Errorf("should intersect")
	}
}

func TestIntersect2(t *testing.T) {
	ll1 := NewEmptySingleLinkedList()
	ll1.AddFront(6)
	ll1.AddFront(4)

	ll2 := NewEmptySingleLinkedList()
	ll2.AddFront(1)
	ll2.AddFront(2)
	ll2.AddFront(7)
	ll2.AddFront(9)
	ll2.AddFront(5)
	ll2.AddFront(1)
	ll2.AddFront(3)

	ll1.Tail.Next = ll2.nodeAt(4)
	ll1.Tail = ll2.Tail
	ll1.size += 3

	res := areIntersecting(ll1, ll2)
	if !res {
		t.Errorf("should intersect")
	}
}
