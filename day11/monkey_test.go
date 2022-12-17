package day11

import "testing"

func TestInitialiseStartingItems(t *testing.T) {
	input := "	Starting items: 79, 98"
	itemList := GetStartingItems(input)

	if len(itemList) != 2 {
		t.Error()
	}

	if itemList[0] != 79 {
		t.Error()
	}
	if itemList[1] != 98 {
		t.Error()
	}
}
