package tree

import (
	"testing"
)

func TestInOrderBST(t *testing.T) {
	ar := []int{2, 4, 5, 6, 8, 10, 12, 14, 18, 20, 23, 34}
	btr := minBTree(ar)
	inOrder(btr)
}

func TestPreOrderBST(t *testing.T) {
	ar := []int{2, 4, 5, 6, 8, 10, 12, 14, 18, 20, 23, 34}
	btr := minBTree(ar)
	preOrder(btr)
}

func TestPostOrderBST(t *testing.T) {
	ar := []int{2, 4, 5, 6, 8, 10, 12, 14, 18, 20, 23, 34}
	btr := minBTree(ar)
	postOrder(btr)
}

func TestLevelOrderBST(t *testing.T) {
	ar := []int{2, 4, 5, 6, 8, 10, 12, 14, 18, 20, 23, 34}
	btr := minBTree(ar)
	levelOrder(btr)
}
