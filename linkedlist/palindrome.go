package linkedlist

func isPalindrome(ll *SLinkedList) bool {
	if ll.Head == nil && ll.Head == ll.Tail {
		return false
	}

	h := ll.Head
	t := ll.Tail

	for {
		if h != nil || t != nil {
			if h.Data != t.Data {
				return false
			}
			h = h.Next
			t = t.Prev
		}
		if h == nil || t == nil {
			break
		}
	}

	return true
}
