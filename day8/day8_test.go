package day8

import "testing"

func TestPartOneExample(t *testing.T) {
	visibleTreeCount := PartOne("./simplified_example.txt")

	if visibleTreeCount != 21 {
		t.Errorf("Visible tree count is %d, should be 21", visibleTreeCount)
	}
}

func TestPartTwoExample(t *testing.T) {
	maxScenicScore := PartTwo("./simplified_example.txt")

	if maxScenicScore != 8 {
		t.Errorf("Max scenic score is %d, should be 8", maxScenicScore)
	}
}
