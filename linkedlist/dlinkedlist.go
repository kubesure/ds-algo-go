package linkedlist

import "fmt"

//DBLinkedList implement a single linkedlist
type DBLinkedList struct {
	head *Node
	tail *Node
	size int
}

//NewDBLinkedList create new double linkedlist
func NewDBLinkedList(data int) *DBLinkedList {
	var dl = &DBLinkedList{}
	dl.AddFront(data)
	return dl
}

//AddFront adds element to front of the linkedlist
func (dll *DBLinkedList) AddFront(data int) *Node {
	node := &Node{data: data}
	if dll.head == nil {
		dll.head = node
		dll.tail = node
		dll.size = 1
	} else {
		node.next = dll.head
		dll.head.prev = node
		node.prev = nil
		dll.head = node
		dll.size = dll.size + 1
	}
	return node
}

//AddEnd adds element to the end the linkedlist
func (dll DBLinkedList) AddEnd(data int) (*Node, error) {
	if dll.head == nil {
		return nil, fmt.Errorf("head not set add to head")
	}
	var node = &Node{data: data}
	node.prev = dll.tail
	dll.tail.next = node
	dll.tail = node
	return node, nil
}

//AddAt add element to front of linkedlist
func (dll *DBLinkedList) AddAt(data int, at int) *Node {
	tn := &Node{data: data}
	var index int = 0
	var atNode *Node
	for atNode = dll.head; atNode != nil; atNode = atNode.next {
		index++
		if index == at {
			tn.prev = atNode
			tn.next = atNode.next
			atNode.next.prev = tn
			atNode.next = tn
		}
	}
	dll.size = dll.size + 1
	return atNode
}

//IterateReverse reverse iteration of linked list
func (dll *DBLinkedList) IterateReverse(f func(n *Node)) {
	var node *Node
	for node = dll.tail; node != nil; node = node.prev {
		f(node)
	}
}

//IterateForward reverse iteration of linked list
func (dll *DBLinkedList) IterateForward(f func(n *Node)) {
	var node *Node
	for node = dll.head; node != nil; node = node.next {
		f(node)
	}
}

//Size return size of double linked list
func (dll *DBLinkedList) Size() int {
	return dll.size
}

//Nodes returns all nodes of the linkedlist
func (dll *DBLinkedList) Nodes() []*Node {
	var node *Node
	var nodes []*Node
	for node = dll.head; node != nil; node = node.next {
		nodes = append(nodes, node)
	}
	return nodes
}

//Delete deletes the linkedlist
func (dll *DBLinkedList) Delete() {
	dll.head.next = nil
	dll.tail = nil
	dll.size = 0
}
