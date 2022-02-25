package hackerrank

func gradeMarks(grades []int32) []int32 {
	for i, v := range grades {
		if v >= 35 {
			r := v % 5
			if r > 0 {
				diff := 5 - r
				if diff < 3 {
					grades[i] = v + diff
				}
			}
		}
	}
	return grades
}
