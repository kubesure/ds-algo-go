package arrays

func zero_matrix(ar [][]int) [][]int {
	for rnum, row := range ar {
	out:
		for _, column := range row {
			if column == 0 {
				for i, _ := range ar[rnum] {
					ar[rnum][i] = 0
				}
				break out
			}
		}
	}
	return ar
}
