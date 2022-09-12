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
	ar := []int{2, 4, 5, 6, 8, 10, 12, 14, 18, 20, 23, 34}
	btr := minBTree(ar)
	print(btr, 0, 'M')
}
