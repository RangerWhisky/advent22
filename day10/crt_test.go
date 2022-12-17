package day10

import "testing"

func TestCrtOff(t *testing.T) {
	cycle := 7
	registerX := 11

	if IsPixelLit(cycle, registerX) {
		t.Error()
	}

	registerX = 5

	if IsPixelLit(cycle, registerX) {
		t.Error()
	}
}

func TestCrtOn(t *testing.T) {
	cycle := 4
	registerX := 3

	if !IsPixelLit(cycle, registerX) {
		t.Error()
	}

	registerX = 4

	if !IsPixelLit(cycle, registerX) {
		t.Error()
	}

	registerX = 5

	if !IsPixelLit(cycle, registerX) {
		t.Error()
	}
}
