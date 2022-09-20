package tree

func find(x int, node *Node, path *path) bool {

	if node == nil {
		return false
	}

	path.appendPath(node.value)

	if node.value == x {
		return true
	}

	if find(x, node.left, path) || find(x, node.right, path) {
		return true
	}

	remmoveLast(path)
	return false
}

func (p *path) appendPath(v int) {
	*p = append(*p, v)
}

func remmoveLast(path *path) {
	*path = (*path)[:len(*path)-1]
}
