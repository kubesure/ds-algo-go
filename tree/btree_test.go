package tree

import (
	"testing"
)

func TestCreateBTree(t *testing.T) {
	bt := btree{}
	bt.insert(4)
	bt.insert(5)
	bt.insert(7)
	bt.insert(-1)
	bt.insert(-1)
	bt.insert(5)
	bt.insert(4)
	print(bt.root, 0, 'M')
}

func TestCreateMinBTreeArr(t *testing.T) {
	ar := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	btr := minBTree(ar)
	print(btr, 0, 'M')
}
