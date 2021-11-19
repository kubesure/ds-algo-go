package mergesort

import (
	"testing"

	ll "github.com/kubesure/goalgo/linkedlist"
)

func TestMergeSortedLL1(t *testing.T) {
	ll1 := ll.NewSingleLinkedList(4)
	ll1.AddEnd(8)
	ll1.AddEnd(15)
	ll1.AddEnd(19)
	ll2 := ll.NewSingleLinkedList(7)
	ll2.AddEnd(9)
	ll2.AddEnd(10)
	ll2.AddEnd(16)
	ml := merge_sorted(ll1, ll2)
	ml.Print()
	if ml.Length() != 8 {
		t.Errorf("merge list incomplete")
	}
}

func TestMergeSortedLL2(t *testing.T) {
	ll1 := ll.NewSingleLinkedList(1)
	ll1.AddEnd(3)
	ll1.AddEnd(5)
	ll1.AddEnd(6)
	ll2 := ll.NewSingleLinkedList(2)
	ll2.AddEnd(4)
	ll2.AddEnd(6)
	ll2.AddEnd(20)
	ll2.AddEnd(34)
	ml := merge_sorted(ll1, ll2)
	ml.Print()
	if ml.Length() != 9 {
		t.Errorf("merge list incomplete")
	}
}
