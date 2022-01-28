package strings

import "strings"

func isRotation(s1 string, s2 string) bool {

	if len(s1) > 0 && len(s1) == len(s2) {
		s1s1 := s1 + s1
		result := strings.Contains(s1s1, s2)
		return result
	}
	return false
}
