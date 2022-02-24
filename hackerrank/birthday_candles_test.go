package hackerrank

import "testing"

func TestTallCandles(t *testing.T) {
	candles := []int32{5, 4, 5, 5}
	l := tallestBCakeCandles(candles)
	if l != 3 {
		t.Errorf("There should have been two tall candles")
	}
}
