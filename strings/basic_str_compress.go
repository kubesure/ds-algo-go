package strings

import (
	"strconv"
)

func compressStr(s string) string {
	var fcs string
	var countConsecutive int

	for i, v := range s {
		countConsecutive++
		if (i+1) >= len(s) || s[i] != s[i+1] {
			fcs += string(v) + strconv.Itoa(countConsecutive)
			countConsecutive = 0
		}
	}
	if len(s) > len(fcs) {
		return fcs
	}
	return s
}
