package tree

import (
	"fmt"
	"math"
	"testing"

	"golang.org/x/tour/tree"
)

func TestIsBinaryTree(t *testing.T) {
	tr := tree.New(1)
	isBinary := isBST(tr, math.MinInt32, math.MaxInt32, "root")
	fmt.Printf(" is binary tree %v \n", isBinary)
}

func TestLevelOrderTravesal(t *testing.T) {
	tr := tree.New(5)
	level_order_bt(tr)
}

func TestInOrderTraverseBT(t *testing.T) {
	tr := tree.New(5)
	in_order_bt(tr)
}

func TestPostOrderTraverseBT(t *testing.T) {
	tr := tree.New(5)
	post_order_bt(tr)
}

func TestPreOrderTraverseBT(t *testing.T) {
	tr := tree.New(5)
	pre_order_bt(tr)
}
