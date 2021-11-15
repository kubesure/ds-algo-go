package aws

import "testing"

func TestStrSegmentation(t *testing.T) {
	str := "applepiepearpieapple"
	dict := make(map[string]struct{})
	words := []string{"apple", "pear", "pie", "pier"}
	for _, v := range words {
		dict[v] = struct{}{}
	}
	found := segmentedWords(str, dict)
	if len(found) != 4 {
		t.Errorf("should have found 4 found %v", len(found))
	}

	/*if found[0] != "apple" {
		t.Errorf("apple not found.. found%v", found[0])
	}

	if found[1] != "pie" {
		t.Errorf("pie not found.. found%v", found[0])
	}*/
}

func TestCanSegmentation(t *testing.T) {

	str := "applepiepearpieapple"
	dict := make(map[string]struct{})
	words := []string{"apple", "pear", "pie", "pier"}
	for _, v := range words {
		dict[v] = struct{}{}
	}

	ok := canSegmentStr(str, dict)

	if !ok {
		t.Errorf("should have segmented")
	}
}
