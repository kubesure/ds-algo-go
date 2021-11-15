package linkedlist

import "fmt"

//SLinkedList implement a single linkedlist
type SLinkedList struct {
	Head *Node
	Tail *Node
	size int
}

//NewSingleLinkedList create a new single list
func NewSingleLinkedList(Data int) *SLinkedList {
	var sll *SLinkedList = new(SLinkedList)
	sll.AddFront(Data)
	return sll
}

//NewSingleLinkedList create a new single list
func NewNullSingleLinkedList() *SLinkedList {
	var sll *SLinkedList = new(SLinkedList)
	return sll
}

//AddFront add element to front of the linkedlist
func (sll *SLinkedList) AddFront(Data int) *Node {
	node := &Node{Data: Data}
	if sll.Head == nil {
		sll.Head = node
		sll.Tail = node
		sll.size = 1
	} else {
		node.Next = sll.Head
		sll.Head = node
		sll.size = sll.size + 1
	}
	return node
}

//AddEnd adding at the front of linkelist
func (sll *SLinkedList) AddEnd(Data int) *Node {
	if sll.Head == nil {
		return sll.AddFront(Data)
	} else if sll.size == 1 {
		node := &Node{Data: Data}
		sll.Head.Next = node
		node.Next = nil
		sll.Tail = node
		sll.size = sll.size + 1
		return node
	} else if sll.size >= 2 {
		node := &Node{Data: Data}
		node.Next = nil
		sll.Tail.Next = node
		sll.Tail = node
		sll.size = sll.size + 1
		return node
	}
	return nil
}

//AddEnd adding at the front of linkelist
func (sll *SLinkedList) AddEndNode(n *Node) {
	sll.Tail.Next = n
	n.Next = nil
	sll.size = sll.size + 1
}

func (sll *SLinkedList) AddSLinkedList(ll ...*SLinkedList) {
	for _, list := range ll {
		sll.Tail = list.Head
		sll.size = sll.size + list.Length()
	}
}

//AddAt adding at the front of linkelist
func (sll *SLinkedList) AddAt(Data int, at int) *Node {
	tn := &Node{Data: Data}
	var index int = 0
	var node *Node
	for node = sll.Head; node != nil; node = node.Next {
		index++
		if index == at {
			tn.Next = node.Next
			node.Next = tn
		}
	}
	sll.size = sll.size + 1
	return node
}

//Delete deletes the linkedlist
func (sll *SLinkedList) Delete() {
	sll.Head.Next = nil
	sll.Tail = nil
	sll.size = 0
}

//Size returns the size of linked list
func (sll *SLinkedList) Length() int {
	return sll.size
}

//IterateForward iterates and call func pass as parameter
func (sll *SLinkedList) IterateForward(f func(n *Node) bool) {
	var node *Node
	for node = sll.Head; node != nil; node = node.Next {
		stop := f(node)
		if stop {
			break
		}
	}
}

//Nodes returns all nodes of the linkedlist
func (sll *SLinkedList) Nodes() []*Node {
	var node *Node
	var nodes []*Node
	for node = sll.Head; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	return nodes
}

func (sll *SLinkedList) Print() {
	currNode := sll.Head
	if currNode == nil {
		fmt.Println("empty list")
	} else {
		fmt.Print("[")
		for i := 0; i < sll.size; i++ {
			if i == sll.size-1 {
				fmt.Printf("%v", currNode.Data)
			} else {
				fmt.Printf("%v,", currNode.Data)
			}
			currNode = currNode.Next
		}
		fmt.Printf("]")
	}
}
