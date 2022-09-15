package tree

import "testing"

func TestInOrderSuccessorNodeRightTreeData(t *testing.T) {
	bst := bsTree2()
	s := inOrderSuccessor(bst, 2)

	if s == nil {
		t.Errorf("inorder sucessor of 2 should not be nil")
	} else if s.value != 3 {
		t.Errorf("inorder sucessor of 2 should be 3 and not %v", s.value)
	}
}

func TestInOrderSucessorNodeLeftTreeData(t *testing.T) {
	bst := bsTree2()
	s := inOrderSuccessor(bst, 0)

	if s == nil {
		t.Errorf("inorder sucessor of 1 should not be nil")
	} else if s.value != 1 {
		t.Errorf("inorder sucessor of 0 should be 1 and not %v", s.value)
	}
}
