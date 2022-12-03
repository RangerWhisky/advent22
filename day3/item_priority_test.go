package day3

import "testing"

func TestGetItemPriority(t *testing.T) {
	itemPriority := GetItemPriority('a')

	if itemPriority != 1 {
		t.Errorf("Item Priority should be 1 for a but is %d", itemPriority)
	}
}

func TestGetItemPriorityEndRange(t *testing.T) {
	itemPriority := GetItemPriority('z')

	if itemPriority != 26 {
		t.Errorf("Item Priority should be 26 for z but is %d", itemPriority)
	}
}

func TestGetCapitalItemPriority(t *testing.T) {
	itemPriority := GetItemPriority('A')

	if itemPriority != 27 {
		t.Errorf("Item Priority should be 27 for A but is %d", itemPriority)
	}
}
