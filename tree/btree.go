package tree

import (
	"math"
	"math/rand"
	"time"
)

func (t *btree) insert(value int) *Node {
	if t.root == nil {
		t.root = &Node{value: value}
		t.size++
		return t.root
	} else {
		t.size++
		return t.root.insert(t.root, value)
	}
}

func (n *Node) insert(node *Node, value int) *Node {
	if node == nil {
		return &Node{value: value, parent: n, size: 1}
	} else if value <= node.value {
		node.left = n.insert(node.left, value)
	} else if value > node.value {
		node.right = n.insert(node.right, value)
	}
	return node
}

func (n *Node) insertBSTNode(value int) *Node {
	var newnode *Node
	if value <= n.value {
		n.left = n.insert(n.left, value)
		newnode = n.left
	} else if value > n.value {
		n.right = n.insert(n.right, value)
		newnode = n.right
	}
	n.size++
	newnode.index = n.size + n.index
	return newnode
}

func (n *Node) insertBTNode(value int, side rune) *Node {
	var newnode *Node
	if side == 'L' {
		n.left = n.insert(n.left, value)
		newnode = n.left
	} else if side == 'R' {
		n.right = n.insert(n.right, value)
		newnode = n.right
	}
	n.size++
	newnode.index = n.size + n.index
	return newnode
}

func findIthNode(n *Node, index int) *Node {
	if n == nil {
		return nil
	}

	c := n.index + 1
	if c == index {
		return n
	}
	rn := findIthNode(n.left, index)
	if rn != nil {
		return rn
	}
	rn = findIthNode(n.right, index)
	if rn != nil {
		return rn
	}
	return nil
}

func sizeOf(n *Node) int {
	if n == nil {
		return 0
	}

	sl := sizeOf(n.left)
	sr := sizeOf(n.right)
	size := sl + sr + 1
	return size
}

func toArray(n *Node) []int {
	e := make(elements, 0)
	convertArray(n, &e)
	return e
}

func convertArray(n *Node, e *elements) {
	if n == nil {
		return
	}
	e.appendElements(n.value)
	convertArray(n.left, e)
	convertArray(n.right, e)
}

func (e *elements) appendElements(v int) {
	*e = append(*e, v)
}

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

func minBTree(arr []int) *Node {
	if len(arr) == 0 {
		return nil
	}
	return convertToBTree(arr)
}

func convertToBTree(arr []int) *Node {

	if len(arr) == 0 {
		return nil
	}

	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid+1:]
	root := &Node{value: arr[mid]}
	root.left = convertToBTree(left)
	root.right = convertToBTree(right)
	return root
}

func randomNode(n *Node) *Node {
	s := sizeOf(n)
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(s)
	return findIthNode(n, r)
}

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

//t2 a sub tree of t1
func isSubTreeOf(t1, t2 *Node) bool {
	if t1 == nil || t2 == nil {
		return false
	}
	return checkSubTree(t1, t2)
}

func checkSubTree(t1, t2 *Node) bool {
	if t1 == nil {
		return false
	}

	if t1.value == t2.value && matchAll(t1, t2) {
		return true
	}

	cl := checkSubTree(t1.left, t2)
	cr := checkSubTree(t1.right, t2)

	return cl || cr
}

func matchAll(t1, t2 *Node) bool {
	if t1 == nil && t2 == nil {
		return true
	} else if t1 == nil || t2 == nil {
		return false
	} else if t1.value != t2.value {
		return false
	} else {
		ml := matchAll(t1.left, t2.left)
		mr := matchAll(t1.right, t2.right)
		return ml && mr
	}
}
