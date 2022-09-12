package tree

import "testing"

func TestIsMinTreeBalanced(t *testing.T) {
	ar := []int{2, 4, 5, 6, 8, 10, 12, 14, 18, 20, 23, 34}
	btr := minBTree(ar)
	if isBalanceTree(btr) != true {
		t.Errorf("Should have been balanced tree")
	}
}

func TestIsBalancedTree(t *testing.T) {
	btr := balancedTree()
	if isBalanceTree(btr) != true {
		t.Errorf("Should have been balanced tree")
	}
}

func TestIsNotBalancedTree(t *testing.T) {
	tr := unBalancedTree()
	if isBalanceTree(tr) != false {
		t.Errorf("This is not a balanced tree")
	}
}

func unBalancedTree() *Node {
	root := &Node{value: 7}
	root.left = &Node{value: 3}
	root.left.left = &Node{value: 1}
	root.left.left = &Node{value: 5}
	root.left.left.left = &Node{value: 4}
	root.left.left.left = &Node{value: 6}
	return root
}

func balancedTree() *Node {
	root := &Node{value: 7}
	//left
	root.left = &Node{value: 3}
	root.left.left = &Node{value: 1}
	root.left.right = &Node{value: 5}
	root.left.right.left = &Node{value: 4}
	root.left.right.right = &Node{value: 6}

	//right
	root.right = &Node{value: 11}
	root.right.left = &Node{value: 9}
	root.right.right = &Node{value: 13}
	root.right.left.left = &Node{value: 8}
	root.right.right.left = &Node{value: 12}
	root.right.right.right = &Node{value: 14}

	return root
}
