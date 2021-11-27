package arrays

func isUniqueStrInArr(arr []string) bool {
	charsfound := make(map[string]struct{})

	for _, v := range arr {
		if _, ok := charsfound[v]; ok {
			return false
		}
		charsfound[v] = struct{}{}
	}

	return true
}
