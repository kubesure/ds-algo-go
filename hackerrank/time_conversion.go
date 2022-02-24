package hackerrank

import (
	"strconv"
)

func timeConversion(s string) string {
	hour, _ := strconv.Atoi(s[:2])
	if s[8:] == "PM" {
		if hour <= 11 {
			hour = hour + 12
			return strconv.Itoa(hour) + s[2:8]
		} else {
			return s[0:8]
		}
	} else {
		if hour == 12 {
			return "00" + s[2:8]
		}
	}
	return s[0:8]
}
