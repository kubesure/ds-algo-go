package strings

func canSegmentStr(str string, dict map[string]struct{}) bool {
	for i := 1; i <= len(str); i++ {
		fristWord := str[0:i]
		if _, ok := dict[fristWord]; ok {
			remainder := str[i:]
			_, ok := dict[remainder]
			if ok || canSegmentStr(remainder, dict) {
				return true
			}
		}
	}
	return false
}

func segmentedWords(str string, dict map[string]struct{}) []string {
	found := make([]string, 0)
	for i := 1; i <= len(str); i++ {
		fristWord := str[0:i]
		if _, ok := dict[fristWord]; ok {
			found = append(found, fristWord)
			remainder := str[i:]
			_, ok := dict[remainder]
			if !ok {
				found = append(found, segmentedWords(remainder, dict)...)
			}
		}
	}
	return found
}
