package hackerrank

import (
	"fmt"
	"strconv"
)

func plusMinus(arr []int32) {
	var countPositive float64
	var countNegative float64
	var countZero float64

	for _, v := range arr {
		if v < 0 {
			countNegative++
		} else if v > 0 {
			countPositive++
		} else {
			countZero++
		}
	}

	length := float64(len(arr))
	fmt.Println(strconv.FormatFloat((countPositive / length), 'f', 6, 64))
	fmt.Println(strconv.FormatFloat((countNegative / length), 'f', 6, 64))
	fmt.Println(strconv.FormatFloat((countZero / length), 'f', 6, 64))
}
