package strings

func isOnewayEdit(first, second string) bool {

	diff := len(first) - len(second)

	if diff < -1 || diff > 1 {
		return false
	}

	if diff == 0 && isSame(first, second) {
		return false
	} else if diff == 0 {
		return isOneReplace(first, second)
	} else {
		return isOneInsert(first, second)
	}
}

func isSame(first, second string) bool {
	for i := range first {
		if first[i] != second[i] {
			return false
		}
	}
	return true
}

func isOneReplace(first, second string) bool {
	var foundOne bool

	for i := range first {
		if first[i] != second[i] {
			if foundOne {
				return false
			}
			foundOne = true
		}
	}
	return true
}

func isOneInsert(first, second string) bool {
	var foundOne bool
	var index1 int
	var index2 int
	for {
		if index1 < len(first) && index2 < len(second) {
			if first[index1] != second[index2] {
				if foundOne {
					return false
				}
				foundOne = true
				index1++
			}
			index1++
			index2++
		} else {
			break
		}
	}
	return true
}

func getShortLongStr(first, second string) (long string, short string) {

	if len(first) > len(second) {
		return first, second
	} else {
		return second, first
	}

}

/*func isOneInsert1(first, second string) bool {
	var shortStr, longStr = getShortLongStr(first, second)

	return false
}*/
