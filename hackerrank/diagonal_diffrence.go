package hackerrank

func diagonalDiff(ar [][]int32) int32 {
	var leftd int32
	var rightd int32

	for i := 0; i < len(ar); i++ {
		for j := 0; j < len(ar); j++ {
			if i == j {
				leftd += ar[i][j]
			}

			diff := len(ar) - j - 1
			if i == diff {
				rightd += ar[i][j]
			}
		}
	}

	if leftd < rightd {
		return rightd - leftd
	}

	return leftd - rightd
}
