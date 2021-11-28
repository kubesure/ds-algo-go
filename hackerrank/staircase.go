package hackerrank

import "fmt"

func staircase(n int32) {
	var i int32
	var j int32
	for i = 0; i < n; i++ {
		spaces := n - i - 1
		fmt.Printf("%*s", spaces, "")
		for j = 0; j <= i; j++ {
			fmt.Printf("%v", "#")
		}
		fmt.Printf("\n")
	}
}
