package strings

import "testing"

func TestIsPlaindromePrem(t *testing.T) {
	if !isPlaindromePrem("tactcoapapa") {
		t.Errorf("should be a plaindrome permunation")
	}
}

func TestIsPlaindromePrem2(t *testing.T) {
	if !isPlaindromePrem2("tactcozapapa") {
		t.Errorf("should be a plaindrome permunation")
	}
}
