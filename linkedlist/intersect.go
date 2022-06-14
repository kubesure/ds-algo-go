package linkedlist

func areIntersecting(ll1, ll2 *SLinkedList) bool {
	if ll1.Tail != ll2.Tail {
		return false
	}

	n1 := ll1.Head
	n2 := ll2.Head

	if ll1.size > ll2.size {
		for i := 0; i < ll1.size-ll2.size; i++ {
			n1 = n1.Next
		}
	}

	if ll2.size > ll1.size {
		for i := 0; i < ll2.size-ll1.size; i++ {
			n2 = n2.Next
		}
	}

	for n1 != nil {
		if n1 == n2 {
			return true
		}
		n1, n2 = n1.Next, n2.Next
	}

	return false
}
