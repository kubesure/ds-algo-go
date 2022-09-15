package tree

import (
	"testing"

	"golang.org/x/tour/tree"
)

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
