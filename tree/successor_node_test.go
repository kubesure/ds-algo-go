package tree

import "testing"

func TestInOrderSuccessorNodeLeftTreeData1(t *testing.T) {
	bst := bsTree2()
	s := inOrderSuccessor(bst, 2)

	if s == nil {
		t.Errorf("inorder sucessor of 2 should not be nil")
	} else if s.value != 3 {
		t.Errorf("inorder sucessor of 2 should be 3 and not %v", s.value)
	}
}

func TestInOrderSuccessorNodeLeftTreeData2(t *testing.T) {
	bst := bsTree2()
	s := inOrderSuccessor(bst, 0)

	if s == nil {
		t.Errorf("inorder sucessor of 1 should not be nil")
	} else if s.value != 1 {
		t.Errorf("inorder sucessor of 0 should be 1 and not %v", s.value)
	}
}

func TestInOrderSucessorNodeLeftMost(t *testing.T) {
	bst := bsTree2()
	s := inOrderSuccessor(bst, 3)

	if s == nil {
		t.Errorf("inorder sucessor of 3 should not be nil")
	} else if s.value != 4 {
		t.Errorf("inorder sucessor of 3 should be 4 and not %v", s.value)
	}
}

func TestInOrderSucessorNodeRightTreeDate1(t *testing.T) {
	bst := bsTree2()
	s := inOrderSuccessor(bst, 4)

	if s == nil {
		t.Errorf("inorder sucessor of 4 should not be nil")
	} else if s.value != 5 {
		t.Errorf("inorder sucessor of 4 should be 5 and not %v", s.value)
	}
}

func TestInOrderSucessorNodeRightTreeDate2(t *testing.T) {
	bst := bsTree2()
	s := inOrderSuccessor(bst, 5)

	if s == nil {
		t.Errorf("inorder sucessor of 5 should not be nil")
	} else if s.value != 6 {
		t.Errorf("inorder sucessor of 5 should be 6 and not %v", s.value)
	}
}

func TestInOrderSucessorNodeRightTreeDate6(t *testing.T) {
	bst := bsTree2()
	s := inOrderSuccessor(bst, 6)

	if s != nil {
		t.Errorf("inorder sucessor of 5 should be nil")
	}
}

func TestInOrderMinBST(t *testing.T) {
	tr := minBT()
	s := inOrderSuccessor(tr, 12)

	if s == nil {
		t.Errorf("inorder sucessor of 12 should not be nil")
	} else if s.value != 14 {
		t.Errorf("inorder sucessor of 12 should be 14 and not %v", s.value)
	}
}

func TestInOrderTree4(t *testing.T) {
	tr := btree4()
	s := inOrderSuccessor(tr, 4)

	if s == nil {
		t.Errorf("inorder sucessor of 20 should not be nil")
	} else if s.value != 8 {
		t.Errorf("inorder sucessor of 20 should be 22 and not %v", s.value)
	}
}

func TestInOrderTree5(t *testing.T) {
	tr := btree4()
	s := inOrderSuccessor(tr, 8)

	if s == nil {
		t.Errorf("inorder sucessor of 8 should not be nil")
	} else if s.value != 10 {
		t.Errorf("inorder sucessor of 8 should be 10 and not %v", s.value)
	}
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
