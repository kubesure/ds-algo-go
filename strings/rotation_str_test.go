package strings

import (
	"testing"
)

func TestIsRotation(t *testing.T) {
	if !isRotation("waterbottle", "erbottlewat") {
		t.Errorf("This should have been a rotation")
	}
}

func TestIsNotARotation(t *testing.T) {
	if isRotation("camera", "macera") {
		t.Errorf("This should have not been a rotation")
	}
}
