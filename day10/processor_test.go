package day10

import "testing"

func TestCreate(t *testing.T) {
	proc := CreateProcessor()

	if proc.cycle != 0 {
		t.Error()
	}

	if proc.registerX != 1 {
		t.Error()
	}

}

func TestNoop(t *testing.T) {
	proc := CreateProcessor()

	QueueNoop(&proc)
	if GetSignalStrength(&proc) != 0 {
		t.Errorf("Noop impacted state when it shouldn't have")
	}
	ClockTick(&proc)
	if GetSignalStrength(&proc) != 1 {
		t.Errorf("Noop impacted register when it shouldn't have")
	}
}

func TestSignalStrength(t *testing.T) {
	proc := CreateProcessor()

	proc.cycle = 7
	proc.registerX = 11
	signal := GetSignalStrength(&proc)
	if signal != 77 {
		t.Errorf("Expected %d, got %d", 77, signal)
	}
}
