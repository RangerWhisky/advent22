package day5

import (
	"strconv"
	"strings"
)

func GetInstruction(instruction string) (int, int, int) {
	lineParts := strings.Split(instruction, " ")

	quantity, _ := strconv.Atoi(lineParts[1])
	source, _ := strconv.Atoi(lineParts[3])
	dest, _ := strconv.Atoi(lineParts[5])

	return quantity, source, dest
}
