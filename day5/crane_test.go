package day5

import "testing"

func TestParseInstructionLine(t *testing.T) {
	instruction := "move 6 from 2 to 1"
	quantity, from, to := GetInstruction(instruction)

	if quantity != 6 {
		t.Errorf("Quantity wrong - %d", quantity)
	}
	if from != 2 {
		t.Errorf("From wrong - %d", from)
	}

	if to != 1 {
		t.Errorf("To wrong - %d", to)
	}
}
