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

func TestRunGameForPartTwo(t *testing.T) {
	monkeyBusiness := RunGame("example.txt", 1, 20)
	if monkeyBusiness != 99*103 {
		t.Error()
	}

	monkeyBusiness = RunGame("example.txt", 1, 1000)
	if monkeyBusiness != 5204*5192 {
		t.Error()
	}

}
