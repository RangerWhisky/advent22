package day7

import (
	"testing"
)

func TestPartOneSolution(t *testing.T) {
	solution := PartOne("./simplified_example.txt")
	if solution != 10 {
		t.Errorf("Start of packet marker complete after %d characters but should be %d", solution, 10)
	}
}
