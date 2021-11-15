package linkedlist

import "fmt"

//DBLinkedList implement a single linkedlist
type DBLinkedList struct {
	Head *Node
	Tail *Node
	size int
}

//NewDBLinkedList create new double linkedlist
func NewDBLinkedList(Data int) *DBLinkedList {
	var dl = &DBLinkedList{}
	dl.AddFront(Data)
	return dl
}

//AddFront adds element to front of the linkedlist
func (dll *DBLinkedList) AddFront(Data int) *Node {
	node := &Node{Data: Data}
	if dll.Head == nil {
		dll.Head = node
		dll.Tail = node
		dll.size = 1
	} else {
		node.Next = dll.Head
		dll.Head.Prev = node
		node.Prev = nil
		dll.Head = node
		dll.size = dll.size + 1
	}
	return node
}

//AddEnd adds element to the end the linkedlist
func (dll *DBLinkedList) AddEnd(Data int) (*Node, error) {
	if dll.Head == nil {
		return nil, fmt.Errorf("head not set add to head")
	}
	var node = &Node{Data: Data}
	node.Prev = dll.Tail
	dll.Tail.Next = node
	dll.Tail = node
	dll.size = dll.size + 1
	return node, nil
}

//AddAt add element to front of linkedlist
func (dll *DBLinkedList) AddAt(Data int, at int) *Node {
	tn := &Node{Data: Data}
	var index int = 0
	var atNode *Node
	for atNode = dll.Head; atNode != nil; atNode = atNode.Next {
		index++
		if index == at {
			tn.Prev = atNode
			tn.Next = atNode.Next
			atNode.Next.Prev = tn
			atNode.Next = tn
		}
	}
	dll.size = dll.size + 1
	return atNode
}

//IterateReverse reverse iteration of linked list
func (dll *DBLinkedList) IterateReverse(f func(n *Node)) {
	var node *Node
	for node = dll.Tail; node != nil; node = node.Prev {
		f(node)
	}
}

//IterateForward reverse iteration of linked list
func (dll *DBLinkedList) IterateForward(f func(n *Node)) {
	var node *Node
	for node = dll.Head; node != nil; node = node.Next {
		f(node)
	}
}

//Size return size of double linked list
func (dll *DBLinkedList) Length() int {
	return dll.size
}

//Nodes returns all nodes of the linkedlist
func (dll *DBLinkedList) Nodes() []*Node {
	var node *Node
	var nodes []*Node
	for node = dll.Head; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	return nodes
}

//Delete deletes the linkedlist
func (dll *DBLinkedList) Delete() {
	dll.Head.Next = nil
	dll.Tail = nil
	dll.size = 0
}
