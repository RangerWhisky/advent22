package day5

import "testing"

func TestCreateStackFromBytes(t *testing.T) {
	testInput := []byte("MCD")

	stack := CreateCargoStack(testInput)

	if stack.height != 3 {
		t.Errorf("Stack height set to %d but should be 3", stack.height)
	}
}
