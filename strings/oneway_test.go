package strings

import "testing"

func TestIsOnewayEdit(t *testing.T) {
	if !isOnewayEdit("apple", "aple") {
		t.Errorf("should have been an one way edit")
	}

}
