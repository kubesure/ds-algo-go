package arrays

import (
	"fmt"
	"testing"
)

func TestReplacesSpaces(t *testing.T) {
	str := "Mr John Smith  abc  "
	str2 := replacesSpaces(str)
	fmt.Println(str2)
}
