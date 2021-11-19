package aws

import (
	"testing"
)

func TestSumOfVals(t *testing.T) {
	vals := [7]int{6, 7, 4, 2, 8, 4, 3}
	if find_sum_of_val(vals, 10) != true {
		t.Errorf("should have found true for val %v", 10)
	}

	if find_sum_of_val(vals, 20) != false {
		t.Errorf("should have not found true for val %v", 20)
	}
}
