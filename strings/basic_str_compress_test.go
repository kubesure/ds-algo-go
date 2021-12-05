package strings

import "testing"

func TestIsCompressStr(t *testing.T) {
	s := "aabcccccaaa"
	cs := compressStr(s)
	if len(s) <= len(cs) {
		t.Errorf("string should be been smaller")
	}
}
