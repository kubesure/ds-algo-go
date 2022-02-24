package hackerrank

func tallestBCakeCandles(candles []int32) int32 {
	var count, tallest int32
	for _, v := range candles {
		if v > tallest {
			count = 1
			tallest = v
		} else if v == tallest {
			count++
		}
	}
	return count
}
