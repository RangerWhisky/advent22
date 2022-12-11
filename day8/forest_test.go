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
