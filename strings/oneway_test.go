package strings

import "testing"

func TestIsOnewayEditInsert(t *testing.T) {
	if !isOnewayEdit("apple", "appl") {
		t.Errorf("should have been an one edit combination")
	}
}

func TestIsOnewayEditReplace(t *testing.T) {
	if !isOnewayEdit("apple", "applz") {
		t.Errorf("should have been an one edit combination")
	}
}

func TestNotEditable(t *testing.T) {
	result := isOnewayEdit("apple", "apple")
	if result {
		t.Errorf("should have been no edit combination")
	}
}
