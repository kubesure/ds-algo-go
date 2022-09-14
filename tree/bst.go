package tree

import "math"

func isBinarySearchTree(node *Node) bool {
	return checkBST(node, math.MinInt, math.MaxInt)
}

func checkBST(node *Node, min, max int) bool {
	if node == nil {
		return true
	}

	if (node.value < min) || (node.value > max) {
		return false
	}

	if !checkBST(node.left, min, node.value) || !checkBST(node.right, node.value, max) {
		return false
	}

	return true
}
