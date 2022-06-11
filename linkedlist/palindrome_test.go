package linkedlist

import "testing"

func TestIsPalindrome(t *testing.T) {
	ll1 := NewEmptySingleLinkedList()
	ll1.AddFront(1)
	ll1.AddFront(2)
	ll1.AddFront(3)
	ll1.AddFront(2)
	ll1.AddFront(1)
	res := isPalindrome(ll1)
	if !res {
		t.Errorf("Its is a plaindrome")
	}
}

func TestIsNotPalindrome(t *testing.T) {
	ll1 := NewEmptySingleLinkedList()
	ll1.AddFront(1)
	ll1.AddFront(2)
	ll1.AddFront(3)
	ll1.AddFront(2)
	ll1.AddFront(1)
	ll1.AddFront(9)
	res := isPalindrome(ll1)
	if res {
		t.Errorf("Its not a plaindrome")
	}
}
