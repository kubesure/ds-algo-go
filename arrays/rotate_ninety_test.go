package arrays

import (
	"testing"
)

func TestRotate903x3(t *testing.T) {
	cube := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	cube, err := rotate90(cube)

	if err != nil {
		t.Errorf(err.Error())
	}

	printCube(cube)
}

func TestRotate904x4(t *testing.T) {
	cube := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	cube, err := rotate90(cube)

	if err != nil {
		t.Errorf(err.Error())
	}

	printCube(cube)
}
