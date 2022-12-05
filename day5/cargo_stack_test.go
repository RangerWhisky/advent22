package day5

import "testing"

func TestCreateStackFromBytes(t *testing.T) {
	testInput := []byte("MCD")

	stack := CreateCargoStack(testInput)

	if stack.height != 3 {
		t.Errorf("Stack height set to %d but should be 3", stack.height)
	}

	for i := 0; i < len(testInput); i++ {
		if testInput[i] != stack.cargo[i] {
			t.Errorf("Non matching stack descriptions - input %q - saved %q", testInput, stack.cargo)
			t.FailNow()
		}
	}
}

func TestGetStackInputFromDescription(t *testing.T) {
	exampleText := `    [D]    
	[N] [C]    
	[Z] [M] [P]
	 1   2   3 `
	stackInput := GetCargoStackInput(exampleText)

	if len(stackInput) == 0 {
		t.Error()
	}
}

func TestGetStackValueOnLine(t *testing.T) {
	exampleLine := `[Z] [M] [P]`
	expectedValues := []byte{'0', 'Z', 'M', 'P', '0'}
	for i := 0; i <= 4; i++ {
		value := GetCargoValue(exampleLine, i)
		if value != expectedValues[i] {
			t.Errorf("Got %q, expected %q", value, expectedValues[i])
		}
	}
}

func TestGetStackCount(t *testing.T) {
	exampleLine := " 1   2   3   4   5   6   7   8   9 "
	stacks := GetStackCount(exampleLine)
	if stacks != 9 {
		t.Errorf("Expected 9 stacks but got %d", stacks)
	}
}
