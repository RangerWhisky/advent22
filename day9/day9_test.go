package day9

import "testing"

func TestDayOne(t *testing.T) {
	solution := DayOne("./simplified_example.txt")

	t.Errorf("Expecting %d, got %d", 13, solution)
}

func TestGetTailPositionsWithoutChange(t *testing.T) {

	tail := Coordinate{1, 1}

	head := Coordinate{1, 2}

	positions := GetTailPositions(tail, head)
	if len(positions) != 0 {
		t.Error()
	}

}

func TestGetTailPositionsWithRightChange(t *testing.T) {

	tail := Coordinate{1, 1}

	head := Coordinate{1, 4}

	positions := GetTailPositions(tail, head)

	if len(positions) < 2 {
		t.Errorf("Too few tail positions")
	}

	// count back from head to tail to check our positions

	for i, position := range positions {
		expectedCoordinate := Coordinate{1, 3 - i}
		if position != expectedCoordinate {
			t.Errorf("Expected %q, got %q at %d", expectedCoordinate, position, i)
		}
	}
}
