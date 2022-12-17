package day11

import (
	"strconv"
	"strings"
)

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
