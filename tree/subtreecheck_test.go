package tree

import "testing"

func TestIsSubTreeCheck(t *testing.T) {
	t1 := minBTree([]int{1, 2, 1, 3, 2, 3, 1})
	t2 := minBTree([]int{2, 3, 1})

	print(t1, 2, 'T')
	print(t2, 2, 'T')

	result := isSubTreeOf(t1, t2)

	if !result {
		t.Errorf("t2 should be a subtree of t1")
	}
}

func TestIsNotSubTreeCheck(t *testing.T) {
	t1 := minBTree([]int{1, 2, 1, 3, 2, 3, 5})
	t2 := minBTree([]int{2, 3, 1})

	print(t1, 2, 'T')
	print(t2, 2, 'T')

	result := isSubTreeOf(t1, t2)

	if result {
		t.Errorf("t2 not a should be a subtree of t1")
	}
}
