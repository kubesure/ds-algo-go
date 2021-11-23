package hackerrank

import "testing"

func TestSimpleArraySum(t *testing.T) {
	ar := []int32{1, 3, 4, 5}
	if simpleArraySum(ar) != 13 {
		t.Errorf("Sum should be 13")
	}
}
