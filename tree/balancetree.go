package tree

func isBalanceTree(root *Node) bool {
	return dsfTree(root) != -1
}

func dsfTree(root *Node) int {
	if root == nil {
		return 0
	}

	var dl int
	if dl = dsfTree(root.left); dl == -1 {
		return -1
	}

	var dr int
	if dr = dsfTree(root.right); dr == -1 {
		return -1
	}

	if abs(dl-dr) > 1 {
		return -1
	}

	return max(dl, dr) + 1

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
