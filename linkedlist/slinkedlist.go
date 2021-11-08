package linkedlist

import "fmt"

//SLinkedList implement a single linkedlist
type SLinkedList struct {
	head *Node
	tail *Node
	size int
}

//NewSingleLinkedList create a new single list
func NewSingleLinkedList(data int) *SLinkedList {
	var sll *SLinkedList = new(SLinkedList)
	sll.AddFront(data)
	return sll
}

//AddFront add element to front of the linkedlist
func (sll *SLinkedList) AddFront(data int) *Node {
	node := &Node{data: data}
	if sll.head == nil {
		sll.head = node
		sll.tail = node
		sll.size = 1
	} else {
		node.next = sll.head
		sll.head = node
		sll.size = sll.size + 1
	}
	return node
}

//AddEnd adding at the front of linkelist
func (sll *SLinkedList) AddEnd(data int) *Node {
	node := &Node{data: data}
	sll.tail.next = node
	sll.tail = node
	sll.size = sll.size + 1
	return node
}

//AddAt adding at the front of linkelist
func (sll *SLinkedList) AddAt(data int, at int) *Node {
	tn := &Node{data: data}
	var index int = 0
	var node *Node
	for node = sll.head; node != nil; node = node.next {
		index++
		if index == at {
			tn.next = node.next
			node.next = tn
		}
	}
	sll.size = sll.size + 1
	return node
}

//Delete deletes the linkedlist
func (sll *SLinkedList) Delete() {
	sll.head.next = nil
	sll.tail = nil
	sll.size = 0
}

//Size returns the size of linked list
func (sll *SLinkedList) Size() int {
	return sll.size
}

//IterateForward iterates and call func pass as parameter
func (sll *SLinkedList) IterateForward(f func(n *Node)) {
	var node *Node
	for node = sll.head; node != nil; node = node.next {
		f(node)
	}
}

//Nodes returns all nodes of the linkedlist
func (sll *SLinkedList) Nodes() []*Node {
	var node *Node
	var nodes []*Node
	for node = sll.head; node != nil; node = node.next {
		nodes = append(nodes, node)
	}
	return nodes
}

func (sll *SLinkedList) Print() {
	currNode := sll.head
	if currNode == nil {
		fmt.Println("empty list")
	} else {
		fmt.Print("[")
		for i := 0; i < sll.size; i++ {
			if i == sll.size-1 {
				fmt.Printf("%v", currNode.data)
			} else {
				fmt.Printf("%v,", currNode.data)
			}
			currNode = currNode.next
		}
		fmt.Printf("]")
	}
}
