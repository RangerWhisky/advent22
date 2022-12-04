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

func TestConfirmNoSuperset(t *testing.T) {
	sectionList := "2-4,6-8"
	superset := IsSectionSuperset(sectionList)

	if superset != false {
		t.Errorf("No superset on this test but returned True")
	}
}

func TestConfirmNoSupersetReverse(t *testing.T) {
	sectionList := "6-8,2-4"
	superset := IsSectionSuperset(sectionList)

	if superset != false {
		t.Errorf("No superset on this test but returned True")
	}
}

func TestConfirmIncompleteSuperset(t *testing.T) {
	sectionList := "6-8,7-9"
	superset := IsSectionSuperset(sectionList)

	if superset != false {
		t.Errorf("No superset on this test but returned True")
	}
}

func TestSectionIdentical(t *testing.T) {
	sectionList := "2-4,2-4"
	superset := IsSectionSuperset(sectionList)

	if superset != true {
		t.Errorf("Superset on this test but returned False")
	}
}

func TestSectionLowerBoundary(t *testing.T) {
	sectionList := "2-4,2-2"
	superset := IsSectionSuperset(sectionList)

	if superset != true {
		t.Errorf("Superset on this test but returned False")
	}
}

func TestSectionUpperBoundary(t *testing.T) {
	sectionList := "2-4,4-4"
	superset := IsSectionSuperset(sectionList)

	if superset != true {
		t.Errorf("Superset on this test but returned False")
	}
}

func TestSectionLowerBoundaryReverse(t *testing.T) {
	sectionList := "2-2,2-4"
	superset := IsSectionSuperset(sectionList)

	if superset != true {
		t.Errorf("Superset on this test but returned False")
	}
}

func TestSectionUpperBoundaryReverse(t *testing.T) {
	sectionList := "4-4,2-4"
	superset := IsSectionSuperset(sectionList)

	if superset != true {
		t.Errorf("Superset on this test but returned False")
	}
}

func TestPartOneSolution(t *testing.T) {
	solution := PartOne("./simplified_example.txt")
	if solution != 2 {
		t.Errorf("Sample solution Score calculated as %d but should be %d", solution, 2)
	}
}

func TestConfirmNoOverlap(t *testing.T) {
	sectionList := "6-8,2-4"
	overlap := IsOverlapping(sectionList)

	if overlap != false {
		t.Errorf("No overlap on this test but returned True")
	}
}

func TestConfirmOverlap(t *testing.T) {
	sectionList := "6-8,7-9"
	overlap := IsOverlapping(sectionList)

	if overlap != true {
		t.Errorf("Overlap on this test but returned False")
	}
}
