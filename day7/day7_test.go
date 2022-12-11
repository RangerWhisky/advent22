package day7

import (
	"testing"
)

func TestPartOneSolution(t *testing.T) {
	solution := PartOne("./simplified_example.txt")
	if solution != 95437 {
		t.Errorf("totalSize of cleanable directories is %d but should be %d", solution, 95437)
	}
}

func TestPartTwoSolution(t *testing.T) {

	solution := PartTwo("./simplified_example.txt")
	if solution != 24933642 {
		t.Errorf("Directory to delete is size %d, should be 24933642", solution)
	}
}

func TestGetFsFromFile(t *testing.T) {
	fs := GetFsFromFile("./simplified_example.txt")

	size := GetSizeOnDisk(&fs, "/a/e")

	if size != 584 {
		t.Errorf("Size of simple folder e not correct - %d", size)
	}
}

func TestGetTargetSize(t *testing.T) {
	targetDeletionSize := GetTargetSize(48381165)
	if targetDeletionSize != 8381165 {
		t.Errorf("Size of target deletion not correct - %d", targetDeletionSize)
	}
}
