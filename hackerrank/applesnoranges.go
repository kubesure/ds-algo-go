package hackerrank

import "fmt"

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	var totalapples, totalOranges int32

	for _, v := range apples {
		d := v + a
		if d >= s && d <= t {
			totalapples++
		}
	}

	for _, v := range oranges {
		d := v + b
		if d >= s && d <= t {
			totalOranges++
		}
	}

	fmt.Printf("%v\n%v\n", totalapples, totalOranges)

}
