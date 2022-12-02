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

func TestResponseEncode(t *testing.T) {
	if EncodeResponse('X') != 1 {
		t.Error()
	}
	if EncodeResponse('Y') != 2 {
		t.Error()
	}
	if EncodeResponse('Z') != 3 {
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

func TestGetLossRequirements(t *testing.T) {
	if GetLossRequirement(3) != 2 {
		t.Error()
	}
	if GetLossRequirement(1) != 3 {
		t.Error()
	}
	if GetLossRequirement(2) != 1 {
		t.Error()
	}
}
