package day3

import "testing"

func TestGetDuplicatedLetters(t *testing.T) {
	backpackContents := "vJvr"

	dupeLetter := GetDuplicatedLetters(backpackContents)

	if dupeLetter != 'v' {
		t.Errorf("Expected 'v' but got %q", dupeLetter)
	}
}

func TestGetIncorrectItemPriority(t *testing.T) {
	backpackContents := "vJrv"

	priority := GetIncorrectItemPriority(backpackContents)

	if priority != 22 {
		t.Errorf("Expected 22 for 'v' but got %d", priority)
	}
}
