package day11

import "testing"

func TestPartOne(t *testing.T) {
	monkeyBusiness := PartOne("example.txt")
	if monkeyBusiness != 10605 {
		t.Error()
	}
}
