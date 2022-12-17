package day10

import "testing"

func TestPartOne(t *testing.T) {
	solution := PartOne("./long_example.txt")

	if solution != 13140 {
		t.Errorf("Expecting %d, got %d", 13140, solution)
	}
}
