package day11

import (
	"strconv"
	"strings"
)

const WorryDivisor int = 3

type Monkey struct {
	items          []int
	op             Operation
	divisor        int
	trueDest       int
	falseDest      int
	itemsInspected int
	calmFactor     int
}

type Pass struct {
	val    int
	target int
}

type Operation struct {
	operand string
	val     int
}

func InitMonkey(input []string, calmFactor int) Monkey {
	var monkey Monkey

	monkey.items = GetStartingItems(input[0])
	monkey.op = GetOperation(input[1])
	monkey.divisor = GetTestDivisor(input[2])
	monkey.trueDest = GetThrowDestination(input[3])
	monkey.falseDest = GetThrowDestination(input[4])
	monkey.calmFactor = calmFactor

	return monkey
}

func TakeTurn(monkey *Monkey) []Pass {
	var passList []Pass
	var emptyList []int
	//TODO - test this
	monkey.itemsInspected += len(monkey.items)

	for i := 0; i < len(monkey.items); i++ {
		Inspect(monkey, i)
		passInstruction := Pass{}
		passInstruction.val = monkey.items[i]
		passInstruction.target = GetThrowTarget(monkey, i)
		passList = append(passList, passInstruction)
	}
	monkey.items = emptyList
	return passList
}

func Inspect(monkey *Monkey, itemIndex int) {
	old := monkey.items[itemIndex]
	old = PerformOperation(monkey.op, old)
	old = old / monkey.calmFactor
	monkey.items[itemIndex] = old
}

func GetThrowTarget(monkey *Monkey, itemIndex int) int {
	remainder := monkey.items[itemIndex] % monkey.divisor

	target := monkey.trueDest
	if remainder != 0 {
		target = monkey.falseDest
	}

	return target
}

func GetStartingItems(input string) []int {
	var startingItems []int
	inputParts := strings.Split(input, ": ")
	itemDescription := strings.Split(inputParts[1], ", ")
	for _, item := range itemDescription {
		val, _ := strconv.Atoi(item)
		startingItems = append(startingItems, val)
	}

	return startingItems
}

func GetTestDivisor(input string) int {
	inputParts := strings.Split(input, " by ")
	val, _ := strconv.Atoi(inputParts[1])
	return val
}

func GetThrowDestination(input string) int {
	inputParts := strings.Split(input, " monkey ")
	val, _ := strconv.Atoi(inputParts[1])
	return val
}

func GetOperation(input string) Operation {
	operation := Operation{"+", 0}

	inputParts := strings.Split(input, "new = old ")

	if inputParts[1] == "* old" {
		operation.operand = "^"
		operation.val = 2
	} else {
		operationParts := strings.Split(inputParts[1], " ")

		operation.operand = operationParts[0]
		operation.val, _ = strconv.Atoi(operationParts[1])
	}

	return operation
}

func PerformOperation(op Operation, val int) int {
	switch op.operand {
	case "+":
		val += op.val
	case "-":
		val -= op.val
	case "*":
		val = val * op.val
	case "^":
		val = val * val
	}

	return val
}
