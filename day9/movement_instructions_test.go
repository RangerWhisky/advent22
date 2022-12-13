package day9

import (
	"testing"
)

func TestDecodeRight(t *testing.T) {
	heightDiff, widthDiff := Decode("R 4")
	if heightDiff != 0 {
		t.Error()
	}
	if widthDiff != 4 {
		t.Error()
	}
}

func TestDecodeLeft(t *testing.T) {
	heightDiff, widthDiff := Decode("L 4")
	if heightDiff != 0 {
		t.Error()
	}
	if widthDiff != -4 {
		t.Error()
	}
}

func TestDecodeUp(t *testing.T) {
	heightDiff, widthDiff := Decode("U 3")
	if heightDiff != 3 {
		t.Error()
	}
	if widthDiff != 0 {
		t.Error()
	}
}

func TestDecodeDown(t *testing.T) {
	heightDiff, widthDiff := Decode("D 3")
	if heightDiff != -3 {
		t.Error()
	}
	if widthDiff != 0 {
		t.Error()
	}
}

func TestGetMapRequirements(t *testing.T) {
	filepath := "./simplified_example.txt"

	height, width := GetMapRequirements(filepath)
	if height != 5 {
		t.Errorf("Need height %d but got %d", 5, height)
	}
	if width != 6 {
		t.Errorf("Need width %d but got %d", 6, width)
	}
}
