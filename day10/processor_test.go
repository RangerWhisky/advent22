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

func TestAddx(t *testing.T) {
	proc := CreateProcessor()

	QueueAddx(&proc, 5)
	if GetSignalStrength(&proc) != 0 {
		t.Errorf("Addx impacted state before first tick")
	}
	ClockTick(&proc)
	if GetSignalStrength(&proc) != 1 {
		t.Errorf("Addx impacted register before second tick")
	}
	ClockTick(&proc)
	if proc.registerX != 6 {
		t.Errorf("Addx didn't execute after second tick")
	}
	ClockTick(&proc)
	if proc.registerX != 6 {
		t.Errorf("instruction didn't clear after execution")
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

func TestIsIdle(t *testing.T) {
	proc := CreateProcessor()

	if !IsIdle(&proc) {
		t.Error()
	}
	QueueAddx(&proc, 5)

	if IsIdle(&proc) {
		t.Error()
	}
}
