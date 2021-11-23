package hackerrank

import (
	"fmt"
	"testing"
)

/*
 * Complete the 'compareTriplets' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

func TestCompareTriplets1(t *testing.T) {
	alice := []int32{1, 2, 3}
	bob := []int32{3, 2, 1}
	results := compareTriplets(alice, bob)
	for _, v := range results {
		fmt.Printf("%v ", v)
	}
}

func TestCompareTriplets2(t *testing.T) {
	alice := []int32{5, 6, 7}
	bob := []int32{3, 6, 10}
	results := compareTriplets(alice, bob)
	for _, v := range results {
		fmt.Printf("%v ", v)
	}
}
