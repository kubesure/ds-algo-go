package arrays

import (
	"sort"
	"strings"
)

func isPermutations(str1 string, str2 string) bool {

	isEqualLen := len(str1) == len(str2)
	if !isEqualLen {
		return false
	}
	return sortStr(str1) == sortStr(str2)
}

func sortStr(str1 string) string {
	ss1 := strings.Split(str1, "")
	sort.Strings(ss1)
	return strings.Join(ss1, "")
}
