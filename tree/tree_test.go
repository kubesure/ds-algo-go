package tree

import (
	"fmt"
	"testing"
)

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
