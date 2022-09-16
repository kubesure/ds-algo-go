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

func minBT() *Node {
	ar := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 14, 18, 20, 23, 34}
	btr := minBTree(ar)
	return btr
}

func TestInOrderBST(t *testing.T) {
	inOrderTraversal(btree4())
}

func btree4() *Node {
	tr := btree{}
	tr.insert(20)
	tr.insert(8)
	tr.insert(4)
	tr.insert(12)
	tr.insert(10)
	tr.insert(14)
	tr.insert(22)
	//inOrderTraversal(tr.root)
	return tr.root
}
