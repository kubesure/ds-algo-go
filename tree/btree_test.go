package tree

import (
	"fmt"
	"testing"
)

func TestCreateBTree(t *testing.T) {
	bt := aBtree()
	print(bt.root, 0, 'M')
}

func TestCreateMinBTreeArr(t *testing.T) {
	ar := []int{2, 4, 5, 6, 8, 10, 12, 14, 18, 20, 23, 34}
	btr := minBTree(ar)
	print(btr, 0, 'M')
}

func minBT() *Node {
	ar := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 14, 18, 20, 23, 34}
	btr := minBTree(ar)
	return btr
}

func TestFindPath(t *testing.T) {
	root := aTree()

	path := &path{}

	result := find(9, root, path)

	if result != true {
		t.Errorf("should have found path to 9")
	}

	if len(*path) != 3 {
		t.Errorf("depth should have been 3")
	}

	for k, v := range *path {
		fmt.Print(v)
		if k != len(*path)-1 {
			fmt.Printf("->")
		}
	}
}

func aTree() *Node {
	root := &Node{value: 5}
	root.left = &Node{value: 3}
	root.right = &Node{value: 7}

	root.left.left = &Node{value: 6}
	root.left.right = &Node{value: 8}

	root.right.left = &Node{value: 10}
	root.right.right = &Node{value: 9}
	return root
}

func aBtree() *btree {

	bt := btree{}
	bt.insert(4)
	bt.insert(5)
	bt.insert(7)
	bt.insert(-1)
	bt.insert(-1)
	bt.insert(5)
	bt.insert(4)
	return &bt
}
