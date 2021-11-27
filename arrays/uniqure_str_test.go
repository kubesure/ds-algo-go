package arrays

import "testing"

func TestUniqueStrInArr(t *testing.T) {
	chars := []string{"a", "b", "c"}
	isUnique := isUniqueStrInArr(chars)
	if !isUnique {
		t.Errorf("duplicates should have not been found")
	}
}

func TestDupStrInArr(t *testing.T) {
	chars := []string{"a", "b", "a"}
	isUnique := isUniqueStrInArr(chars)
	if isUnique {
		t.Errorf("duplicates should have been found")
	}
}
