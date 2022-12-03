package day3

import "testing"

func TestGetItemPriority(t *testing.T) {
	itemPriority := GetItemPriority('a')

	if itemPriority != 1 {
		t.Errorf("Item Priority should be 1 for a but is %d", itemPriority)
	}
}
