package arrays

import "testing"

func TestIsPermutations(t *testing.T) {
	str1 := "abdc"
	str2 := "abcz"
	if !isPermutations(str1, str2) {
		t.Errorf("should be a Permutations")
	}
}
