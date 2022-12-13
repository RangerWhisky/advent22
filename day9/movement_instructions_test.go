package day9

import "testing"

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
