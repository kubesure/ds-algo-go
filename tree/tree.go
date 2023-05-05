package tree

import (
	"container/list"
	"fmt"
)

func print(node *Node, indent int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < indent; i++ {
		fmt.Print(" ")
	}

	fmt.Printf("%c:%v  %v \n", ch, node.value, node.index)
	print(node.left, indent+2, 'L')
	print(node.right, indent+2, 'R')
}

func inOrder(root *Node) {
	if root == nil {
		return
	}

	inOrder(root.left)
	fmt.Printf("%v>", root.value)
	inOrder(root.right)
}

func preOrder(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("%v>", root.value)
	preOrder(root.left)
	preOrder(root.right)
}

func postOrder(root *Node) {
	if root == nil {
		return
	}
	postOrder(root.left)
	postOrder(root.right)
	fmt.Printf("%v>", root.value)
}

func levelOrder(root *Node) {
	queue := list.New()
	queue.PushBack(root)
	for {
		if queue.Len() > 0 {
			f := queue.Front()
			currNode := f.Value.(*Node)
			fmt.Printf("%v ", currNode.value)
			queue.Remove(f)

			if currNode.left != nil {
				queue.PushBack(currNode.left)
			}

			if currNode.right != nil {
				queue.PushBack(currNode.right)
			}
		} else {
			break
		}
	}
}
