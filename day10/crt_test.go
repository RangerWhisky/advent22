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

func TestCrtLineTwo(t *testing.T) {
	cycle := 44
	registerX := 3
	if !IsPixelLit(cycle, registerX) {
		t.Error()
	}

	registerX = 5
	if !IsPixelLit(cycle, registerX) {
		t.Error()
	}
}

func TestCrtOutput(t *testing.T) {
	cycle := 44
	registerX := 3
	if GetCrtOutput(cycle, registerX) != "#" {
		t.Error()
	}

	registerX = 6
	if GetCrtOutput(cycle, registerX) != "." {
		t.Error()
	}
}

func TestCrtOutputEol(t *testing.T) {
	cycle := 39
	registerX := 39
	if GetCrtOutput(cycle, registerX) != "#\n" {
		t.Error()
	}

	registerX = 6
	if GetCrtOutput(cycle, registerX) != ".\n" {
		t.Error()
	}
}

func TestGetWidthFromPixel(t *testing.T) {
	pixel := 40

	if getWidthFromPixel(pixel) != 0 {
		t.Error()
	}

	if getWidthFromPixel(pixel-1) != 39 {
		t.Error()
	}

	if getWidthFromPixel(pixel+1) != 1 {
		t.Error()
	}
}
