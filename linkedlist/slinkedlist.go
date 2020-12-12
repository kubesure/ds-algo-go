package linkedlist

import "log"

//Node of a linkedlist
type Node struct {
	data int
	next *Node
}

//SLinkedList implement a single linkedlist
type SLinkedList struct {
	head *Node
	tail *Node
	size int
}

//NewSingleLinkedList create a new single list
func NewSingleLinkedList(data int) *SLinkedList {
	var ll *SLinkedList
	if data == 0 {
		ll = &SLinkedList{}
		ll.size = 0
	} else {
		ll = &SLinkedList{}
		ll.AddFront(data)
	}
	return ll
}

//AddFront adding at the front of linkelist
func (ll *SLinkedList) AddFront(data int) *Node {
	node := &Node{data: data}
	if ll.head == nil {
		ll.head = node
		ll.tail = node
		ll.size = 1
	} else {
		node.next = ll.head
		ll.head = node
		ll.size = ll.size + 1
	}

	return node
}

//AddEnd adding at the front of linkelist
func (ll *SLinkedList) AddEnd(data int) *Node {
	node := &Node{data: data}
	node.next = nil
	ll.tail = node
	return node
}

//AddAt adding at the front of linkelist
func (ll *SLinkedList) AddAt(data int, at int) *Node {
	tn := &Node{data: data}
	var index int = 0
	var node *Node
	for node = ll.head; node != nil; node = node.next {
		index++
		if index == at {
			tn.next = node.next
			node.next = tn
		}
	}
	return node
}

//IteratePrint prints the linkedlist
func (ll *SLinkedList) IteratePrint() {
	var node *Node
	for node = ll.head; node != nil; node = node.next {
		log.Println(node.data)
	}
}
