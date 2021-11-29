package strings

func isPlaindromePrem(str string) bool {
	cmap := make(map[rune]int)
	oddFound := false

	for _, r := range str {
		cmap[r] = cmap[r] + 1
	}

	for _, v := range cmap {
		if v%2 == 1 {
			if oddFound {
				return false
			}
			oddFound = true
		}
	}
	return true
}

func isPlaindromePrem2(str string) bool {
	cmap := make(map[rune]int)
	var oddCount int

	for _, r := range str {
		cmap[r] = cmap[r] + 1
		if cmap[r]%2 == 1 {
			oddCount++
		} else {
			oddCount--
		}
	}

	return oddCount <= 1
}
