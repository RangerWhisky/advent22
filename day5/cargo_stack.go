package day5

import (
	"strconv"
	"strings"
)

type CargoStack struct {
	height int
	cargo  []byte
}

const emptySpace byte = '0'

func GetTop(stack CargoStack) byte {
	top := byte('0')
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
		for level := 1; level < maxHeight; level++ {
			lineToRead := maxHeight - level
			cargo := GetCargoValue(cargoLines[lineToRead], stack)
			if cargo != '0' {
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
	return GetStackCount(cargoLines[len(cargoLines)-1])
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

func MoveCargo(source *CargoStack, dest *CargoStack, quantity int) {
	if quantity > source.height {
		quantity = source.height
	}

	for i := 1; i <= quantity; i++ {
		source.height--
		dest.cargo = append(dest.cargo, source.cargo[source.height])
		dest.height++
	}
}
