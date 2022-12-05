package day5

import (
	file_input "localhost/advent22/utils"
)

func PartOne(filepath string) string {
	crateStarter := []byte{}
	elfInput := file_input.Read_file(filepath)
	instructionStart := 0

	for i := 0; i < len(elfInput); i++ {
		if len(elfInput[i]) != 0 {
			array := []byte(elfInput[i])
			crateStarter = append(crateStarter, array...)
			crateStarter = append(crateStarter, '\n')
		} else {
			instructionStart = i + 1
			break
		}
	}

	stacks := GetCargoStackInput(string(crateStarter))

	for i := instructionStart; i < len(elfInput); i++ {
		quantity, source, dest := GetInstruction(elfInput[i])
		MoveCargo(&stacks[source], &stacks[dest], quantity)
	}

	tops := []byte{}
	for i := 0; i < len(stacks); i++ {
		tops = append(tops, GetTop(stacks[i]))
	}
	return string(tops)
}
