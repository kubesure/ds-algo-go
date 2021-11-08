package aws

import "fmt"

func find_sum_of_val(vals [7]int, val int) bool {
	foundVal := make(map[int]int)
	for _, e := range vals {
		if found, ok := foundVal[val-e]; ok {
			fmt.Printf("found pair %v - %v \n", found, e)
			return true
		}
		foundVal[e] = e
	}
	return false
}
