package tree

import (
	"container/list"
	"fmt"

	"golang.org/x/tour/tree"
)

func pre_order_bt(t *tree.Tree) {
	if t == nil || t.Value == 0 {
		return
	} else {
		fmt.Printf("%v ", t.Value)
		pre_order_bt(t.Left)
		pre_order_bt(t.Right)
	}
}

func in_order_bt(t *tree.Tree) {
	if t == nil || t.Value == 0 {
		return
	} else {
		in_order_bt(t.Left)
		fmt.Printf("%v ", t.Value)
		in_order_bt(t.Right)
	}
}

func post_order_bt(t *tree.Tree) {
	if t == nil || t.Value == 0 {
		return
	} else {
		post_order_bt(t.Left)
		post_order_bt(t.Right)
		fmt.Printf("%v ", t.Value)
	}
}

func level_order_bt(t *tree.Tree) {
	queue := list.New()
	queue.PushBack(t)
	for {
		if queue.Len() > 0 {
			f := queue.Front()
			currNode := f.Value.(*tree.Tree)
			fmt.Printf("%v ", currNode.Value)
			queue.Remove(f)

			if currNode.Left != nil {
				queue.PushBack(currNode.Left)
			}

			if currNode.Right != nil {
				queue.PushBack(currNode.Right)
			}
		} else {
			break
		}
	}
}
