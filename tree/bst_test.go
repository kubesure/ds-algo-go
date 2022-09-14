package tree

import (
	"testing"
)

func TestIsBST(t *testing.T) {
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
