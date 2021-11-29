package strings

func isOnewayEdit(frist, second string) bool {

	if (len(frist) + 2) == len(second) {
		return false
	} else if len(frist) == len(second) {
		return isOneReplace(frist, second)
	} else {
		return isOneInsert(frist, second)
	}

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
	return false
}
