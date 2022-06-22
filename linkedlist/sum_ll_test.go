package linkedlist

import "testing"

func TestSumLinkedList(t *testing.T) {
	ll1 := NewEmptySingleLinkedList()
	ll1.Insert(6)
	ll1.Insert(1)
	ll1.Insert(7)

	ll2 := NewEmptySingleLinkedList()
	ll2.Insert(2)
	ll2.Insert(9)
	ll2.Insert(5)

	sum := sum(ll1, ll2)

	if sum != uint(912) {
		t.Errorf("sum should be 912 and not %v", sum)
	}
}

func TestSumLinkedList1(t *testing.T) {
	ll1 := NewEmptySingleLinkedList()
	ll1.Insert(6)
	ll1.Insert(1)
	ll1.Insert(7)

	ll2 := NewEmptySingleLinkedList()
	ll2.Insert(2)
	ll2.Insert(9)

	sum := sum(ll1, ll2)

	if sum != uint(646) {
		t.Errorf("sum should be 912 and not %v", sum)
	}
}

func TestSumLinkedList2(t *testing.T) {
	ll1 := NewEmptySingleLinkedList()
	ll1.Insert(7)
	ll1.Insert(8)
	ll1.Insert(6)

	ll2 := NewEmptySingleLinkedList()
	ll2.Insert(8)
	ll2.Insert(6)

	sum := sum(ll1, ll2)

	if sum != uint(872) {
		t.Errorf("sum should be 912 and not %v", sum)
	}
}
