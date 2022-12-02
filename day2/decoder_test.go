package day2

import "testing"

func TestResultDecode(t *testing.T) {
	if Decode('X') != 1 {
		t.Errorf("%d", Decode('X'))
	}

	if Decode('Y') != 2 {
		t.Error()
	}

	if Decode('Z') != 3 {
		t.Error()
	}
}
