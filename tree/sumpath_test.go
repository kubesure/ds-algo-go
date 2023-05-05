package tree

import "testing"

func TestPathToSum(t *testing.T) {
	bt := PathSumBTree2()
	print(bt, 2, 'T')
}

func TestToArray(t *testing.T) {
	bt := PathSumBTree2()
	arr := toArray(bt)
	if len(arr) != 9 {
		t.Errorf("should be 9 and not %v", len(arr))
	}
}
