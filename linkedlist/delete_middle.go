package linkedlist

func deleteFromMiddle(node *Node) bool {
	if node == nil || node.Next == nil {
		return false
	}

	node.Prev.Next = node.Next
	return true
}
