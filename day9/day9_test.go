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

	position := GetSnapPosition(tail, head)
	if position != tail {
		t.Errorf("Tail should not have moved")
	}

}

func TestGetTailPositionsWithRightChange(t *testing.T) {

	tail := Coordinate{1, 1}

	head := Coordinate{1, 3}

	position := GetSnapPosition(tail, head)

	// count back from head to tail to check our positions

	expectedCoordinate := Coordinate{1, 2}
	if position != expectedCoordinate {
		t.Errorf("Expected %q, got %q", expectedCoordinate, position)
	}

}

func TestDiagonalTailPositionsChange(t *testing.T) {

	tail := Coordinate{2, 2}

	head := Coordinate{4, 3}

	position := GetSnapPosition(tail, head)

	// count back from head to tail to check our positions

	expectedCoordinate := Coordinate{3, 3}
	if position != expectedCoordinate {
		t.Errorf("Expected %q, got %q", expectedCoordinate, position)
	}
}

func TestBigNegativeDiagonalChange(t *testing.T) {

	tail := Coordinate{4, 4}

	head := Coordinate{1, 3}

	position := GetSnapPosition(tail, head)

	// count back from head to tail to check our positions

	expectedCoordinate := Coordinate{2, 3}
	if position != expectedCoordinate {
		t.Errorf("Expected %q, got %q", expectedCoordinate, position)
	}
}

func TestUp(t *testing.T) {

	tail := Coordinate{1, 4}

	head := Coordinate{5, 6}

	position := GetSnapPosition(tail, head)

	// count back from head to tail to check our positions

	expectedCoordinate := Coordinate{4, 5}
	if position != expectedCoordinate {
		t.Errorf("Expected %q, got %q", expectedCoordinate, position)
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
