package aws

import (
	"math/rand"
	"testing"
	"time"
)

func TestFindMissingNum(t *testing.T) {
	nums := make([]int, 0)
	rand.Seed(time.Now().UTC().UnixNano())
	r := randInt(1, 10)

	for i := 1; i <= 10; i++ {
		if r != i {
			nums = append(nums, i)
		}
	}
	//fmt.Printf("num array %v \n", nums)

	n := findMissingNum(nums)

	if n != r {
		t.Errorf("Failed to find missing number")
	}
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
