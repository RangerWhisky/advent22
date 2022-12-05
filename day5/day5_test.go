package day5

import "testing"

func TestPartOneSolution(t *testing.T) {
	solution := PartOne("./simplified_example.txt")
	if solution != "CMZ" {
		t.Errorf("Sample solution calculated as %s but should be %s", solution, "CMZ")
	}
}
