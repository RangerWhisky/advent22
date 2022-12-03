package day3

import "testing"

func TestGetCompartmentDuplicate(t *testing.T) {
	backpackContents := "vJvr"

	dupeLetter := GetCompartmentDuplicate(backpackContents)

	if dupeLetter != 'v' {
		t.Errorf("Expected 'v' but got %q", dupeLetter)
	}
}

func TestGetIncorrectItemPriority(t *testing.T) {
	backpackContents := "vJrv"

	priority := GetIncorrectItemPriority(backpackContents)

	if priority != 22 {
		t.Errorf("Expected 22 for 'v' but got %d", priority)
	}
}

func TestGetSampleResult(t *testing.T) {
	solution := PartOne("./simplified_example.txt")

	if solution != 157 {
		t.Errorf("Should be 157 but is %d", solution)
	}
}

func TestGetBadgeLetter(t *testing.T) {
	var backpacks = []string{"vrvF", "frgf", "prgFrp"}
	badge := GetBadgeLetter(backpacks)

	if badge != byte('r') {
		t.Errorf("Expected r but got %q", badge)
	}
}

func TestPartTwoExample(t *testing.T) {
	solution := PartTwo("./simplified_example.txt")

	if solution != 70 {
		t.Errorf("Should be 70 but is %d", solution)
	}
}
