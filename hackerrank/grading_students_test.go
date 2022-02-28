package hackerrank

import "testing"

func TestGrades(t *testing.T) {
	grades := []int32{84, 43, 67, 39}
	gradeMarks(grades)
	if grades[0] != 85 {
		t.Errorf("should have been 85")
	}

	if grades[1] != 45 {
		t.Errorf("should have been 85")
	}

	if grades[2] != 67 {
		t.Errorf("should have been 67")
	}

	if grades[3] != 40 {
		t.Errorf("should have been 40")
	}
}
