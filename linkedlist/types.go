package linkedlist

//Node of a linkedlist
type Node struct {
	Data int
	Prev *Node
	Next *Node
}

//LinkedList intrface for single and double linked list
// type LinkedList interface {
// 	AddFront(Data int) *Node
// 	AddEnd(Data int) *Node
// 	AddAt(Data int, at int) *Node
// 	Nodes() []*Node
// 	IterateForward(f func(n *Node))
// 	IterateReverse(f func(n *Node))
// 	Delete()
// 	Size() int
// }
