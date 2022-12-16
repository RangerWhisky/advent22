package day9

import "testing"

func TestPartOne(t *testing.T) {
	solution := PartOne("./simplified_example.txt")

	if solution != 13 {
		t.Errorf("Expecting %d, got %d", 13, solution)
	}
}

func TestPartTwo(t *testing.T) {
	solution := PartTwo("long_example.txt")

	if solution != 36 {
		t.Errorf("Expecting %d, got %d", 36, solution)
	}
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

func TestDiagonalTailPositionsChange(t *testing.T) {

	tail := Coordinate{2, 2}

	head := Coordinate{4, 3}

	positions := GetTailPositions(tail, head)

	if len(positions) < 1 {
		t.Errorf("Too few tail positions")
	}

	// count back from head to tail to check our positions

	for i, position := range positions {
		expectedCoordinate := Coordinate{3 - i, 3}
		if position != expectedCoordinate {
			t.Errorf("Expected %q, got %q at %d", expectedCoordinate, position, i)
		}
	}
}

func TestBigNegativeDiagonalChange(t *testing.T) {

	tail := Coordinate{4, 4}

	head := Coordinate{1, 3}

	positions := GetTailPositions(tail, head)

	if len(positions) < 2 {
		t.Errorf("Too few tail positions")
	}

	// count back from head to tail to check our positions

	for i, position := range positions {
		expectedCoordinate := Coordinate{2 + i, 3}
		if position != expectedCoordinate {
			t.Errorf("Expected %q, got %q at %d", expectedCoordinate, position, i)
		}
	}
}

func TestUp(t *testing.T) {

	tail := Coordinate{1, 4}

	head := Coordinate{5, 5}

	positions := GetTailPositions(tail, head)

	if len(positions) < 3 {
		t.Errorf("Too few tail positions")
	}

	// count back from head to tail to check our positions

	for i, position := range positions {
		expectedCoordinate := Coordinate{4 - i, 5}
		if position != expectedCoordinate {
			t.Errorf("Expected %q, got %q at %d", expectedCoordinate, position, i)
		}
	}
}

func TestDiagonal(t *testing.T) {

	tail := Coordinate{1, 1}

	head := Coordinate{3, 3}

	position := GetSnapPosition(tail, head)

	// count back from head to tail to check our positions
	expectedCoordinate := Coordinate{2, 2}
	if position != expectedCoordinate {
		t.Errorf("Diagonal move not calculated correctly - got (%d, %d)", position.height, position.width)
	}
}
