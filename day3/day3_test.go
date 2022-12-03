package day3

import "testing"

func TestGetDuplicatedLetters(t *testing.T) {
	backpackContents := "vJvr"

	dupeLetter := GetDuplicatedLetters(backpackContents)

	if dupeLetter != 'v' {
		t.Errorf("Expected 'v' but got %q", dupeLetter)
	}
}
