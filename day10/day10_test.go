package day10

import "testing"

func TestPartOne(t *testing.T) {
	solution := PartOne("./long_example.txt")

	if solution != 13140 {
		t.Errorf("Expecting %d, got %d", 13140, solution)
	}
}

func TestProcessorModelling20(t *testing.T) {
	solution := ModelProcessor("./long_example.txt", 20)

	if solution != 420 {
		t.Errorf("Expecting %d, got %d", 420, solution)
	}
}
