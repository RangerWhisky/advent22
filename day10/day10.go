package day10

import (
	utils "localhost/advent22/utils"
)

const MaxCycle int = 220

func PartOne(filepath string) int {
	instructionList := utils.Read_file(filepath)
	instructionPointer := 0

	proc := CreateProcessor()

	for clock := 0; clock <= MaxCycle; clock++ {
		if IsIdle(&proc) {
			SendInstruction(&proc, instructionList[instructionPointer])
		}
	}
	return 0
}

func SendInstruction(proc *Processor, ins string) {

}
