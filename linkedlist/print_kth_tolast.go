package linkedlist

func recurse(head *Node, kth int, index int) *Node {
	if head == nil {
		return nil
	}

	nd := recurse(head.Next, kth, index)
	index++
	if index == kth {
		return head
	}
	return nd
}

func kthToLast(node *Node, kth int) *Node {
	index := 0
	return recurse(node, kth, index)
}
