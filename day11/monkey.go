package day11

import (
	"strconv"
	"strings"
)

type Operation struct {
	operand string
	val     int
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
	return 0
}
