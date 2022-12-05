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

const emptySpace byte = '0'

func CreateCargoStack(cargoList []byte) CargoStack {
	stack := CargoStack{
		height: (len(cargoList)),
		cargo:  cargoList,
	}
	return stack
}

func GetCargoStackInput(cargoDescription string) [][]byte {
	var val [][]byte
	cargoLines := strings.Split(cargoDescription, "\n")
	maxHeight := len(cargoLines) - 1

	stackCount := GetStackCountFromDescription(cargoDescription)

	for stack := 1; stack <= stackCount; stack++ {
		stackContents := []byte{}
		for level := 1; level < maxHeight; level++ {
			lineToRead := maxHeight - level
			fmt.Printf("level %d, maxHeight %d, line to check %d\n", level, maxHeight, lineToRead)
			cargo := GetCargoValue(cargoLines[maxHeight-level], stack)
			if cargo != '0' {
				stackContents = append(stackContents, cargo)
			}

		}
		val = append(val, stackContents)
	}

	return val
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
