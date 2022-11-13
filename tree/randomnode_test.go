package tree

import (
	"fmt"
	"testing"
)

func TestFindNode(t *testing.T) {
	bt3 := bsTree3()
	n := randomNode(bt3)

	if n == nil {
		t.Errorf("should not be a nil")
	} else {
		fmt.Printf("randon %v ", n.value)
	}
}

func TestFindDepth(t *testing.T) {
	s := sizeOf(bsTree1())
	if s != 5 {
		t.Errorf("Size should be 5 and not %v", s)
	}
}

func TestFindDepthSubTree(t *testing.T) {
	s := sizeOf(bsTree1().left)
	if s != 3 {
		t.Errorf("Size should be 3 and not %v", s)
	}
}
