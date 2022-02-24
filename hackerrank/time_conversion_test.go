package hackerrank

import "testing"

func TestTimeConversion(t *testing.T) {
	time := timeConversion("07:05:45PM")
	if time != "19:05:45" {
		t.Errorf("Invalid conversion expected 19:05:45")
	}

	time = timeConversion("12:01:00PM")
	if time != "12:01:00" {
		t.Errorf("Invalid conversion expected 12:01:00")
	}

	time = timeConversion("12:01:00AM")
	if time != "00:01:00" {
		t.Errorf("Invalid conversion expected 00:01:00")
	}

	time = timeConversion("10:01:00AM")
	if time != "10:01:00" {
		t.Errorf("Invalid conversion expected 10:01:00")
	}
}
