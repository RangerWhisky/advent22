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

	if positions[0] != tail {
		t.Error()
	}
}

func TestGetTailPositionsWithRightChange(t *testing.T) {

	tail := Coordinate{1, 1}

	head := Coordinate{1, 4}

	positions := GetTailPositions(tail, head)

	if positions[0] != tail {
		t.Error()
	}

	if len(positions) < 3 {
		t.Errorf("Too few tail positions")
	}

	for i := 1; i < 4 && i < len(positions); i++ {
		coord := Coordinate{1, i}
		if positions[i] != coord {
			t.Error()
		}
	}
}
