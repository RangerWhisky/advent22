package day5

import (
	"fmt"
	file_input "localhost/advent22/utils"
)

func PartOne(filepath string) string {
	return runCrateSimulation(filepath, false)
}

func PartTwo(filepath string) string {
	return runCrateSimulation(filepath, true)
}

func runCrateSimulation(filepath string, is9001 bool) string {
	crateStarterText := ""
	elfInput := file_input.Read_file(filepath)
	instructionStart := 0

	for i := 0; i < len(elfInput); i++ {
		if elfInput[i][1] != '1' {
			array := []byte(elfInput[i])
			crateStarterText = crateStarterText + string(array) + "\n"
			fmt.Printf("line %q \n", string(array))
		} else {
			array := []byte(elfInput[i])
			crateStarterText = crateStarterText + string(array)
			fmt.Printf("numbers %q \n", string(array))
			instructionStart = i + 2
			break
		}
	}

	stacks := GetCargoStackInput(crateStarterText)

	fmt.Printf("whole string %q \n", crateStarterText)

	for i := instructionStart; i < len(elfInput); i++ {
		quantity, source, dest := GetInstruction(elfInput[i])
		if is9001 {
			BulkMoveCargo(&stacks[source-1], &stacks[dest-1], quantity)
		} else {
			PiecemealMoveCargo(&stacks[source-1], &stacks[dest-1], quantity)
		}
	}

	tops := []byte{}
	for i := 0; i < len(stacks); i++ {
		tops = append(tops, GetTop(stacks[i]))
	}
	return string(tops)
}
