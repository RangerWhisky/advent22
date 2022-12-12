package day8

import "testing"

func TestPartOneExample(t *testing.T) {
	visibleTreeCount := PartOne("./simplified_example.txt")

	if visibleTreeCount != 21 {
		t.Errorf("Visible tree count is %d, should be 21", visibleTreeCount)
	}
}
