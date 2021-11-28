package hackerrank

import (
	"fmt"
	"sort"
)

type sortablearr []int32

func minMax(arr []int32) {
	sort.Sort(sortablearr(arr))
	var min int64
	var max int64

	for _, x := range arr[:len(arr)-1] {
		min = min + int64(x)
	}

	for _, n := range arr[1:] {
		max = max + int64(n)
	}

	fmt.Printf("%v %v", min, max)
}

func (s sortablearr) Len() int {
	return len(s)
}

func (s sortablearr) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortablearr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
