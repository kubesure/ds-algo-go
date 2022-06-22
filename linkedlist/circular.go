package linkedlist

func isCircular(ll *SLinkedList) *Node {
	fast, slow := ll.Head, ll.Head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			break
		}
	}

	if fast == nil || fast.Next == nil {
		return nil
	}

	slow = ll.Head

	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return fast
}
