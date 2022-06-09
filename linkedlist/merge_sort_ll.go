package linkedlist

func merge_sorted(ll1 *SLinkedList, ll2 *SLinkedList) *SLinkedList {

	ml := NewNullSingleLinkedList()

	if ll1.Head == nil {
		return ll2
	} else if ll2.Head == nil {
		return ll1
	}

	if ll1.Head.Data <= ll2.Head.Data {
		ml.AddFront(ll1.Head.Data)
		ml.AddEnd(ll2.Head.Data)
	} else if ll1.Head.Data >= ll2.Head.Data {
		ml.AddFront(ll2.Head.Data)
		ml.AddEnd(ll1.Head.Data)
	}

	ll1.Head = ll1.Head.Next
	ll2.Head = ll2.Head.Next

	for {
		if ll1.Head != nil && ll2.Head != nil {
			if ll1.Head.Data > ll2.Head.Data {
				ml.AddEnd(ll2.Head.Data)
				ml.AddEnd(ll1.Head.Data)
			} else if ll1.Head.Data < ll2.Head.Data {
				ml.AddEnd(ll1.Head.Data)
				ml.AddEnd(ll2.Head.Data)
			}
			ll1.Head = ll1.Head.Next
			ll2.Head = ll2.Head.Next
		} else {
			break
		}
	}

	if ll1.Head != nil {
		//ml.Tail.Next = ll1.Head
		ml.AddEndNode(ll1.Head)
	} else if ll2.Head != nil {
		//ml.Tail.Next = ll2.Head
		ml.AddEndNode(ll2.Head)
	}
	return ml
}
