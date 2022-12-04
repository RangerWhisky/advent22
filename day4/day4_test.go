package day4

import "testing"

func TestGetSectionRange(t *testing.T) {
	rangeString := "3-4"
	rangeStart, rangeEnd := GetSectionRange(rangeString)
	if rangeStart != 3 {
		t.Errorf("Start of range is %d, should be 3", rangeStart)
	}

	if rangeEnd != 4 {
		t.Errorf("End of range is %d, should be 4", rangeEnd)
	}
}

func TestConfirmNoOverlap(t *testing.T) {
	sectionList := "2-4,6-8"
	overlap := GetSectionOverlap(sectionList)

	if overlap != false {
		t.Errorf("No overlap on this test but returned True")
	}
}
