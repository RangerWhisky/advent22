package day5

import "testing"

func TestPartOneSolution(t *testing.T) {
	solution := PartOne("./simplified_example.txt")
	if solution != "CMZ" {
		t.Errorf("Sample solution calculated as %s but should be %s", solution, "CMZ")
	}
}

func TestPartTwo(t *testing.T) {
	step := PartTwo("simplified_example.txt")
	if string(step) != "MCD" {
		t.Errorf("Step four failed %q", step)
	}
}
