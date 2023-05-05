package tree

import (
	"fmt"
	"testing"
)

func TestRandonNode(t *testing.T) {
	bt3 := bsTree3()
	n := randomNode(bt3)

	if n == nil {
		t.Errorf("should not be a nil")
	} else {
		fmt.Printf("randon %v ", n.value)
	}
}
