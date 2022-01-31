package linkedlist

func (sll *SLinkedList) remmoveDuplicate() {
	dups := make(map[int]struct{})
	var previous *Node
	var node *Node
	for node = sll.Head; node != nil; node = node.Next {
		if _, ok := dups[node.Data]; ok {
			previous.Next = node.Next
		} else {
			dups[node.Data] = struct{}{}
			previous = node
		}
	}
}
