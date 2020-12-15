package linkedlist

//Node of a linkedlist
type Node struct {
	data int
	prev *Node
	next *Node
}

//LinkedList intrface for single and double linked list
// type LinkedList interface {
// 	AddFront(data int) *Node
// 	AddEnd(data int) *Node
// 	AddAt(data int, at int) *Node
// 	Nodes() []*Node
// 	IterateForward(f func(n *Node))
// 	IterateReverse(f func(n *Node))
// 	Delete()
// 	Size() int
// }
