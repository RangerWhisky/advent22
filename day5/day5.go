package day5

import (
	"fmt"
	file_input "localhost/advent22/utils"
)

func PartOne(filepath string) string {
	crateStarter := []byte{}
	elfInput := file_input.Read_file(filepath)
	stackCountLine := 0
	instructionStart := 0

	for i := 0; i < len(elfInput); i++ {
		if elfInput[i][1] != '1' {
			array := []byte(elfInput[i])
			crateStarter = append(crateStarter, array...)
			crateStarter = append(crateStarter, '\n')
		} else {
			stackCountLine = i
			instructionStart = i + 2
			break
		}
	}

	stacks := GetCargoStackInput(string(elfInput[stackCountLine]))

	tops := []byte{}
	for i := 0; i < len(stacks); i++ {
		tops = append(tops, GetTop(stacks[i]))
	}
	fmt.Printf("%q\n", tops)

	tops = []byte{}

	for i := instructionStart; i < len(elfInput); i++ {
		quantity, source, dest := GetInstruction(elfInput[i])
		MoveCargo(&stacks[source-1], &stacks[dest-1], quantity)
	}

	for i := 0; i < len(stacks); i++ {
		tops = append(tops, GetTop(stacks[i]))
	}
	return string(tops)
}
