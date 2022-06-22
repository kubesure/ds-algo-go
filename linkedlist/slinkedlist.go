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
	sll.Insert(Data)
	return sll
}

//NewSingleLinkedList create a new single list
func NewNullSingleLinkedList() *SLinkedList {
	var sll *SLinkedList = new(SLinkedList)
	return sll
}

func NewEmptySingleLinkedList() *SLinkedList {
	var sll *SLinkedList = new(SLinkedList)
	return sll
}

//Insert add element to front of the linkedlist
func (sll *SLinkedList) Insert(Data int) *Node {
	node := &Node{Data: Data}
	if sll.Head == nil {
		sll.Head = node
		sll.Tail = node
		sll.size = 1
	} else {
		temp := sll.Head
		node.Next = sll.Head
		sll.Head = node
		temp.Prev = node
		sll.size = sll.size + 1
	}
	return node
}

//AddEnd adding at the front of linkelist
func (sll *SLinkedList) AddEnd(Data int) *Node {
	if sll.Head == nil {
		return sll.Insert(Data)
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

func (sll *SLinkedList) RemoveAt(index int) {
	node := sll.nodeAt(index)
	if sll.Head == node {
		node.Prev = nil
		sll.Head = node.Next
	} else if sll.Tail == node {
		sll.Tail = node.Prev
		sll.Tail.Next = nil
	} else {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
	}
	sll.size--
}

func (sll *SLinkedList) nodeAt(index int) *Node {

	var i int
	for node := sll.Head; node != nil; node = node.Next {
		if i == index {
			return node
		}
		i++
	}
	return nil
}

//Delete deletes the linkedlist
func (sll *SLinkedList) Delete() {
	sll.Head.Next = nil
	sll.Tail = nil
	sll.size = 0
}

func (sll *SLinkedList) Partition(x int) *SLinkedList {
	lower := NewEmptySingleLinkedList()
	higher := SLinkedList{}
	var xnode *Node

	for node := sll.Head; node != nil; node = node.Next {
		if node.Data == x {
			xnode = node
		} else if node.Data < x {
			lower.AddEnd(node.Data)
		} else {
			higher.AddEnd(node.Data)
		}
	}
	if xnode != nil {
		lower.AddEnd(xnode.Data)
	}

	lower.Tail.Next = higher.Head
	higher.Head.Prev = lower.Tail
	return lower
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

func sum(ll1, ll2 *SLinkedList) uint {
	var sum, carry, result uint
	var multiplyfactor uint = 1
	resll := NewEmptySingleLinkedList()
	n1, n2 := ll1.Head, ll2.Head
	for n1 != nil || n2 != nil {
		if n1 == nil {
			sum = uint(n2.Data)
			sum += carry
			n2 = n2.Next
		} else if n2 == nil {
			sum = uint(n1.Data)
			sum += carry
			n1 = n1.Next
		} else {
			sum = uint(n1.Data) + uint(n2.Data)
			sum += carry

			n1 = n1.Next
			n2 = n2.Next
		}
		if sum >= 10 {
			carry = 1
		} else {
			carry = 0
		}
		sum = sum % 10
		resll.AddEnd(int(sum))
	}

	for node := resll.Head; node != nil; node = node.Next {
		result += uint(node.Data) * multiplyfactor
		multiplyfactor *= 10
	}
	return result
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

func andPrint(n *Node) bool {
	fmt.Printf("%v ,", n.Data)
	return false
}
