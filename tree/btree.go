package tree

import (
	"fmt"
)

func minBTree(arr []int) *Node {
	if len(arr) == 0 {
		return nil
	}

	return convertToBTree(arr)
}

func convertToBTree(arr []int) *Node {

	if len(arr) == 0 {
		return nil
	}

	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid+1:]
	root := &Node{value: arr[mid]}
	root.left = convertToBTree(left)
	root.right = convertToBTree(right)
	return root
}

func (t *btree) insert(value int) *Node {
	if t.root == nil {
		t.root = &Node{value: value}
		return t.root
	} else {
		return t.root.insert(t.root, value)
	}
}

func (n *Node) insert(node *Node, value int) *Node {
	if node == nil {
		return &Node{value: value}
	} else if value <= node.value {
		node.left = n.insert(node.left, value)
	} else if value > node.value {
		node.right = n.insert(node.right, value)
	}
	return node
}

func print(node *Node, indent int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < indent; i++ {
		fmt.Print(" ")
	}

	fmt.Printf("%c:%v\n", ch, node.value)
	print(node.left, indent+2, 'L')
	print(node.right, indent+2, 'R')
}
