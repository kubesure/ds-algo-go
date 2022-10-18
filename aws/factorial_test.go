package aws

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	f := factorial(4)
	if f != 24 {
		t.Errorf("Factorial of 4 should be 24 and not %v", f)
	}
}
