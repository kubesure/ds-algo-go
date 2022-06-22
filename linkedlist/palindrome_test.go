package linkedlist

import "testing"

func TestIsPalindrome(t *testing.T) {
	ll1 := NewEmptySingleLinkedList()
	ll1.Insert(1)
	ll1.Insert(2)
	ll1.Insert(3)
	ll1.Insert(2)
	ll1.Insert(1)
	res := isPalindrome(ll1)
	if !res {
		t.Errorf("Its is a plaindrome")
	}
}

func TestIsNotPalindrome(t *testing.T) {
	ll1 := NewEmptySingleLinkedList()
	ll1.Insert(1)
	ll1.Insert(2)
	ll1.Insert(3)
	ll1.Insert(2)
	ll1.Insert(1)
	ll1.Insert(9)
	res := isPalindrome(ll1)
	if res {
		t.Errorf("Its not a plaindrome")
	}
}
