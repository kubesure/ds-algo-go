package tree

import (
	"container/list"
)

func find(x int, node *Node, path *list.List) bool {

	if node == nil {
		return false
	}

	path.PushBack(node.value)

	if node.value == x {
		return true
	}

	if find(x, node.left, path) || find(x, node.right, path) {
		return true
	}
	remmovePaths(path)
	return false
}

func remmovePaths(path *list.List) {
	for {
		if path.Len() > 0 {
			path.Remove(path.Front())
		} else {
			break
		}
	}
}
