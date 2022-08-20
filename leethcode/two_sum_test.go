package leethcode

import "testing"

func TestTwoSum1(t *testing.T) {
	nums := []int{2, 11, 7, 15}
	indices := twoSum(nums, 9)

	if indices[0] != 0 || indices[1] != 2 {
		t.Errorf("indices are not [0] and [2]")
	}
}

func TestTwoSum2(t *testing.T) {
	nums := []int{3, 2, 4}
	indices := twoSum(nums, 6)

	if indices[0] != 1 || indices[1] != 2 {
		t.Errorf("indices are not [1] and [2]")
	}
}
