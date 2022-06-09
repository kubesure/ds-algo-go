package hackerrank

import (
	"fmt"
	"strconv"
)

func main() {
	//var h int

	//fmt.Scan(&h)
	h := 4

	if h < 1 || h > 100 {
		return
	}

	str := ""
	for i := h; i > 0; i-- {
		str += "#"
		s := strconv.Itoa(h)
		fmt.Printf("%"+s+"s\n", str)
	}
}
