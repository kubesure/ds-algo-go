package tree

import (
	"testing"
)

func TestIsBST1(t *testing.T) {
	bst := isBinarySearchTree(bsTree1())
	if !bst {
		t.Errorf("Should be a BST")
	}
}

func TestIsBST2(t *testing.T) {
	bst := isBinarySearchTree(bsTree2())
	if !bst {
		t.Errorf("Should be a BST")
	}
}

func bsTree1() *Node {
	root := &Node{value: 4}

	root.left = &Node{value: 2}
	root.right = &Node{value: 5}

	root.left.left = &Node{value: 1}
	root.left.right = &Node{value: 3}

	return root
}

func bsTree3() *Node {
	root := &Node{value: 4}
	root.left = root.insertBSTNode(2)
	root.left.insertBSTNode(1)
	root.left.insertBSTNode(3)
	root.right = root.insertBSTNode(5)
	return root
}

func PathSumBTree2() *Node {
	root := &Node{value: 10}

	rr := root.insertBTNode(-3, 'R')
	rl := root.insertBTNode(5, 'L')

	rr.insertBTNode(11, 'R')
	rlr := rl.insertBTNode(1, 'R')
	rll := rl.insertBTNode(3, 'L')

	rlr.insertBTNode(2, 'R')

	rll.insertBTNode(3, 'L')
	rll.insertBTNode(-2, 'R')

	return root

}

func PathSumBTree() *Node {
	root := &Node{value: 10}
	root.left = &Node{value: 5}

	root.left.left = &Node{value: 3}
	root.left.right = &Node{value: 2}
	root.left.right.right = &Node{value: 1}

	root.left.left.left = &Node{value: 3}
	root.left.left.right = &Node{value: -2}

	root.right = &Node{value: -3}
	root.right.right = &Node{value: 11}

	return root

}

func bsTree2() *Node {
	root := &Node{value: 3}

	root.left = &Node{value: 1}
	root.left.left = &Node{value: 0}
	root.left.right = &Node{value: 2}

	root.right = &Node{value: 5}
	root.right.left = &Node{value: 4}
	root.right.right = &Node{value: 6}

	return root
}
