package day10

import (
	utils "localhost/advent22/utils"
	"strconv"
	"strings"
)

const MaxCycle int = 220

func PartOne(filepath string) int {
	instructionList := utils.Read_file(filepath)
	instructionPointer := 0

	signalStrengthSum := 0

	proc := CreateProcessor()

	for clock := 1; clock <= MaxCycle; clock++ {
		if IsIdle(&proc) {
			SendInstruction(&proc, instructionList[instructionPointer])
			instructionPointer++
		}
		ClockTick(&proc)
		if NeedSignalStrengthReading(clock) {
			signalStrengthSum += GetSignalStrength(&proc)
		}
	}
	return signalStrengthSum
}

func SendInstruction(proc *Processor, ins string) {
	stringParts := strings.Split(ins, " ")
	switch stringParts[0] {
	case "noop":
		QueueNoop(proc)
	case "addx":
		val, _ := strconv.Atoi(stringParts[1])
		QueueAddx(proc, val)
	}
}

func NeedSignalStrengthReading(clock int) bool {
	required := false
	switch clock {
	case 20, 60, 100, 140, 180, 220:
		required = true
	}
	return required
}
