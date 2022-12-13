package day9

import "testing"

func TestDayOne(t *testing.T) {
	solution := DayOne("./simplified_example.txt")

	t.Errorf("Expecting %d, got %d", 13, solution)
}
