package tree

import (
	"container/list"
	"strconv"
	"testing"
)

func TestFindPath(t *testing.T) {
	root := &Node{value: 5}
	root.left = &Node{value: 3}
	root.right = &Node{value: 7}

	root.left.left = &Node{value: 6}
	root.left.right = &Node{value: 8}

	root.right.left = &Node{value: 10}
	root.right.right = &Node{value: 9}

	path := &list.List{}

	result := find(9, root, path)

	if result != true {
		t.Errorf("should have found path to 9")
	}

	var pathstr string

	for {
		if path.Len() > 0 {
			e := path.Front()
			pathstr = pathstr + strconv.Itoa(e.Value.(int))
			path.Remove(e)
		} else {
			break
		}
	}

}
