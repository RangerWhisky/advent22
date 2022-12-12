package day8

import (
	"testing"
)

func TestInitialiseForest(t *testing.T) {
	forest := InitialiseForestFromFile("./simplified_example.txt")

	treeCount := getTreeCount(&forest)

	if treeCount != 25 {
		t.Errorf("Got %d expected %d", treeCount, 25)
	}
}

func TestAssessVisibility(t *testing.T) {
	forest := InitialiseForestFromFile("./simplified_example.txt")

	AssessVisibility(&forest)
	output := GetVisibleTreeCount(&forest)
	if output != 21 {
		PrintVisibilityMap(&forest)
		t.Errorf("Should be 21 visible trees, instead get %d", output)
	}
}

func TestVisibleFromNorth(t *testing.T) {
	forest := InitialiseForestFromFile("./simplified_example.txt")

	if !IsVisibleFromNorth(&forest, 1, 1, 5) {
		t.Errorf("False negative")
	}

	if IsVisibleFromNorth(&forest, 2, 2, 3) {
		t.Errorf("False positive")
	}

}

func TestVisibleFromSouth(t *testing.T) {
	forest := InitialiseForestFromFile("./simplified_example.txt")

	if !IsVisibleFromSouth(&forest, 3, 2, 5) {
		t.Errorf("False negative")
	}

	if IsVisibleFromSouth(&forest, 3, 1, 3) {
		t.Errorf("False positive")
	}

}

func TestVisibleFromWest(t *testing.T) {
	forest := InitialiseForestFromFile("./simplified_example.txt")

	if !IsVisibleFromWest(&forest, 1, 1, 5) {
		t.Errorf("False negative")
	}

	if IsVisibleFromWest(&forest, 2, 1, 5) {
		t.Errorf("False positive")
	}
}
func TestVisibleFromEast(t *testing.T) {
	forest := InitialiseForestFromFile("./simplified_example.txt")

	if !IsVisibleFromEast(&forest, 0, 3, 7) {
		t.Errorf("False negative")
	}

	if IsVisibleFromEast(&forest, 1, 3, 1) {
		t.Errorf("False positive")
	}
}

func TestViewToNorth(t *testing.T) {
	forest := InitialiseForestFromFile("./simplified_example.txt")

	if GetViewToNorth(&forest, 0, 1, 5) != 0 {
		t.Error("Bad error handling for top line")
	}

	viewDistance := GetViewToNorth(&forest, 1, 1, 5)

	if viewDistance != 1 {
		t.Errorf("Incorrect view distance north - %d", viewDistance)
	}

	viewDistance = GetViewToNorth(&forest, 3, 2, 5)
	if viewDistance != 2 {
		t.Errorf("Incorrect view distance north - %d", viewDistance)
	}

}

func TestViewToSouth(t *testing.T) {
	forest := InitialiseForestFromFile("./simplified_example.txt")

	if GetViewToSouth(&forest, 4, 1, 5) != 0 {
		t.Error("Bad error handling for bottom line")
	}

	viewDistance := GetViewToSouth(&forest, 2, 1, 5)

	if viewDistance != 2 {
		t.Errorf("Incorrect view distance south - %d", viewDistance)
	}
}
