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

}
