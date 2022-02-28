package hackerrank

import "testing"

func TestCountApplesAndOranges(t *testing.T) {
	var start int32 = 7
	var end int32 = 10
	var appleTreeDistance int32 = 4
	var orangeTreeDistance int32 = 12
	var applesdistances []int32 = []int32{2, 3, -4}
	var orangedistances []int32 = []int32{3, -2, -4}
	countApplesAndOranges(start, end, appleTreeDistance, orangeTreeDistance, applesdistances, orangedistances)
}
