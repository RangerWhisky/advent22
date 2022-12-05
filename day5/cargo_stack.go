package day5

import (
	"fmt"
	"strconv"
	"strings"
)

type CargoStack struct {
	height int
	cargo  []byte
}

const emptySpace byte = ' '

func GetTop(stack CargoStack) byte {
	top := emptySpace
	if stack.height > 0 {
		itemToCheck := stack.height - 1
		top = stack.cargo[itemToCheck]
	}
	return byte(top)
}

func CreateCargoStack(cargoList []byte) CargoStack {
	stack := CargoStack{
		height: (len(cargoList)),
		cargo:  cargoList,
	}
	return stack
}

func GetCargoStackInput(cargoDescription string) []CargoStack {
	var stacks []CargoStack
	cargoLines := strings.Split(cargoDescription, "\n")
	maxHeight := len(cargoLines) - 1

	stackCount := GetStackCountFromDescription(cargoDescription)

	for stack := 1; stack <= stackCount; stack++ {
		stackContents := []byte{}
		for level := 1; level <= maxHeight; level++ {
			lineToRead := maxHeight - level
			cargo := GetCargoValue(cargoLines[lineToRead], stack)
			if cargo != emptySpace {
				stackContents = append(stackContents, cargo)
			}

		}
		stacks = append(stacks, CreateCargoStack(stackContents))
	}

	return stacks
}

func GetCargoValue(cargoLine string, index int) byte {
	// initialise to be empty
	foundCargo := emptySpace
	if index > 0 {
		indexToRead := (index * 4) - 3
		if indexToRead <= len(cargoLine) {
			foundCargo = cargoLine[indexToRead]
		}
	}
	return foundCargo
}

func GetStackCountFromDescription(cargoDescription string) int {
	cargoLines := strings.Split(cargoDescription, "\n")
	stacks := GetStackCount(cargoLines[len(cargoLines)-1])
	return stacks
}

func GetStackCount(cargoLine string) int {
	inputAsBytes := []byte(cargoLine)
	highestStack := cargoLine[len(inputAsBytes)-2]
	stackCount, err := strconv.Atoi(string(highestStack))
	if err != nil {
		stackCount = 0
	}
	return stackCount
}

func BulkMoveCargo(source *CargoStack, dest *CargoStack, quantity int) {
	MoveCargo(source, dest, quantity, true)
}

func PiecemealMoveCargo(source *CargoStack, dest *CargoStack, quantity int) {
	MoveCargo(source, dest, quantity, false)
}

func MoveCargo(source *CargoStack, dest *CargoStack, quantity int, bulk bool) {
	fmt.Printf("Move %d from %q to %q\n", quantity, source, dest)
	if quantity > source.height {
		quantity = source.height
	}

	if bulk {

	} else {
		for i := 0; i < quantity; i++ {
			source.height--
			dest.cargo = append(dest.cargo, source.cargo[source.height])
			dest.height++
		}
		source.cargo = source.cargo[:source.height]
	}
}
