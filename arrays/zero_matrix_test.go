package arrays

import "testing"

func TestZeroMatrix(t *testing.T) {
	ar := [][]int{
		{1, 2, 3, 4},
		{5, 6, 0, 8},
		{9, 0, 9, 10},
	}
	ar = zero_matrix(ar)
	printCube(ar)
}
