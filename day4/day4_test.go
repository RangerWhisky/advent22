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

func TestConfirmNoOverlapReverse(t *testing.T) {
	sectionList := "6-8,2-4"
	overlap := GetSectionOverlap(sectionList)

	if overlap != false {
		t.Errorf("No overlap on this test but returned True")
	}
}

func TestSectionIdentical(t *testing.T) {
	sectionList := "2-4,2-4"
	overlap := GetSectionOverlap(sectionList)

	if overlap != true {
		t.Errorf("Overlap on this test but returned False")
	}
}

func TestSectionLowerBoundary(t *testing.T) {
	sectionList := "2-4,2-2"
	overlap := GetSectionOverlap(sectionList)

	if overlap != true {
		t.Errorf("Overlap on this test but returned False")
	}
}

func TestSectionUpperBoundary(t *testing.T) {
	sectionList := "2-4,4-4"
	overlap := GetSectionOverlap(sectionList)

	if overlap != true {
		t.Errorf("Overlap on this test but returned False")
	}
}

func TestSectionLowerBoundaryReverse(t *testing.T) {
	sectionList := "2-2,2-4"
	overlap := GetSectionOverlap(sectionList)

	if overlap != true {
		t.Errorf("Overlap on this test but returned False")
	}
}
