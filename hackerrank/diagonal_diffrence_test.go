package hackerrank

import "testing"

func TestDiagonalDiff(t *testing.T) {
	ar := [][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}

	diff := diagonalDiff(ar)
	if diff != 2 {
		t.Errorf("diagonal diff is not 2 but %v", diff)
	}
}

func TestDiagonalDiff2(t *testing.T) {
	ar := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}

	diff := diagonalDiff(ar)
	if diff != 15 {
		t.Errorf("diagonal diff is not 15 but %v", diff)
	}
}
