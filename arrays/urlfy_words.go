package arrays

func replacesSpaces(str string) string {

	endIndex := findEndOfStr(str)
	r := []rune(str)

	endcodeArr := make([]rune, 0)
	for i := 0; i <= endIndex; i++ {
		if r[i] == ' ' {
			endcodeArr = append(endcodeArr, '%')
			endcodeArr = append(endcodeArr, '2')
			endcodeArr = append(endcodeArr, '0')
		} else {
			endcodeArr = append(endcodeArr, r[i])
		}
	}
	return string(endcodeArr)
}

func findEndOfStr(str string) int {
	r := []rune(str)
	for i := len(r) - 1; i >= 0; i-- {
		if r[i] != ' ' {
			return i
		}
	}
	return -1
}
