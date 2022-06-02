package linkedlist

func kthFromLast(head *Node, kth int) *Node {
	p1, p2 := head, head
	for i := 0; i < kth+1; i++ {
		p1 = p1.Next
		if p1.Next == nil {
			return nil
		}
	}

	for p1.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p2.Next
}

//    1 2 3 4 5
//p1    _
//p2  _
