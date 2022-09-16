package tree

func inOrderSuccessor(root *Node, data int) *Node {
	var successor *Node

	targetNode := findNode(root, data)
	if targetNode == nil {
		return nil
	}

	if targetNode.right != nil {
		return findLeftMost(targetNode.right)
	} else {
		ancestor := root
		for {
			if ancestor != targetNode {
				if targetNode.value < ancestor.value {
					successor = ancestor
					ancestor = ancestor.left
				} else {
					ancestor = ancestor.right
				}
			} else {
				break
			}
		}
	}

	return successor
}

func findNode(n *Node, data int) *Node {
	if n == nil {
		return nil
	}

	if n.value == data {
		return n
	}

	if n.value > data {
		return findNode(n.left, data)
	} else if n.value < data {
		return findNode(n.right, data)
	}
	return nil
}

func findLeftMost(n *Node) *Node {
	if n == nil {
		return nil
	}

	for {
		if n.left != nil {
			n = n.left
		} else {
			break
		}
	}
	return n
}
