package tree

import (
	"container/list"
	"fmt"
)

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
