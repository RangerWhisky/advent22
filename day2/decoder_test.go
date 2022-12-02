package day2

import "testing"

func TestResultDecode(t *testing.T) {
	if Decode('X') != 1 {
		t.Error()
	}

	if Decode('Y') != 2 {
		t.Error()
	}

	if Decode('Z') != 3 {
		t.Error()
	}
}

func TestGetWinRequirements(t *testing.T) {
	if GetWinRequirement(1) != 2 {
		t.Error()
	}
	if GetWinRequirement(2) != 3 {
		t.Error()
	}
	if GetWinRequirement(3) != 1 {
		t.Error()
	}
}
