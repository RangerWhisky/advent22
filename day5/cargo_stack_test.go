package day5

import (
	"fmt"
	"testing"
)

func TestGetTop(t *testing.T) {
	stack := CreateCargoStack([]byte("MCD"))
	if GetTop(stack) != 'D' {
		t.Error()
	}
}

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
	if GetTop(stackInput[0]) != byte('N') {
		t.Errorf("Expected ZN, got %q", stackInput[0])
	}
}

func TestGetStackValueOnLine(t *testing.T) {
	exampleLine := `[Z] [M] [P]`
	expectedValues := []byte{' ', 'Z', 'M', 'P', ' '}
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

	exampleLine = " 1   2   3 "
	stacks = GetStackCount(exampleLine)
	if stacks != 3 {
		t.Errorf("Expected 3 stacks but got %d", stacks)
	}
}

func TestGetStackCountFromDescription(t *testing.T) {
	exampleText := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 `
	stackCount := GetStackCountFromDescription(exampleText)
	if stackCount != 3 {
		t.Errorf("Stack count from full description is %d, should be 3", stackCount)
	}
}

func TestSimpleMove(t *testing.T) {
	sourceStack := CreateCargoStack([]byte("MCD"))
	destStack := CreateCargoStack([]byte("P"))

	PiecemealMoveCargo(&sourceStack, &destStack, 1)
	if sourceStack.height != 2 {
		t.Error()
	}
}

func TestBigMove(t *testing.T) {
	sourceStack := CreateCargoStack([]byte("MCD"))
	destStack := CreateCargoStack([]byte("P"))

	PiecemealMoveCargo(&sourceStack, &destStack, 2)
	if destStack.height != 3 {
		t.Error()
	}

	fmt.Printf("source %d, dest %d", sourceStack.height, destStack.height)

	if GetTop(sourceStack) != 'M' {
		t.Error()
	}

	if GetTop(destStack) != 'C' {
		t.Error()
	}
}

func TestSampleMovesStepwise(t *testing.T) {
	exampleText := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 `
	stackInput := GetCargoStackInput(exampleText)

	step := []byte{GetTop(stackInput[0]), GetTop(stackInput[1]), GetTop(stackInput[2])}
	if string(step) != "NDP" {
		t.Errorf("Step zero failed %q", step)
	}

	PiecemealMoveCargo(&stackInput[1], &stackInput[0], 1)
	step = []byte{GetTop(stackInput[0]), GetTop(stackInput[1]), GetTop(stackInput[2])}
	if string(step) != "DCP" {
		t.Errorf("Step one failed %q", step)
	}

	PiecemealMoveCargo(&stackInput[0], &stackInput[2], 3)
	step = []byte{GetTop(stackInput[0]), GetTop(stackInput[1]), GetTop(stackInput[2])}
	if string(step) != " CZ" {
		t.Errorf("Step two failed %q", step)
	}

	PiecemealMoveCargo(&stackInput[1], &stackInput[0], 2)
	step = []byte{GetTop(stackInput[0]), GetTop(stackInput[1]), GetTop(stackInput[2])}
	if string(step) != "M Z" {
		t.Errorf("Step three failed %q", step)
	}

	PiecemealMoveCargo(&stackInput[0], &stackInput[1], 1)
	step = []byte{GetTop(stackInput[0]), GetTop(stackInput[1]), GetTop(stackInput[2])}
	if string(step) != "CMZ" {
		t.Errorf("Step four failed %q", step)
	}
}
