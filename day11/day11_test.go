package day11

import "testing"

func TestPartOne(t *testing.T) {
	monkeyBusiness := PartOne("example.txt")
	if monkeyBusiness != 10605 {
		t.Error()
	}
}

func TestPartTwo(t *testing.T) {
	monkeyBusiness := PartTwo("example.txt")
	if monkeyBusiness != 2713310158 {
		t.Error()
	}
}
