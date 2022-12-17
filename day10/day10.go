package day10

import (
	utils "localhost/advent22/utils"
	"strconv"
	"strings"
)

const MaxCycle int = 220
const MaxCrtCycle int = 240

func PartOne(filepath string) int {
	return ModelProcessor(filepath, MaxCycle)
}

func PartTwo(filepath string) string {
	return ModelCrt(filepath, MaxCrtCycle)
}

func ModelProcessor(filepath string, cycleCount int) int {
	instructionList := utils.Read_file(filepath)
	instructionPointer := 0

	signalStrengthSum := 0

	proc := CreateProcessor()

	for clock := 1; clock <= cycleCount; clock++ {
		if IsIdle(&proc) {
			SendInstruction(&proc, instructionList[instructionPointer])
			instructionPointer++
		}
		signalStrength := ClockTick(&proc)
		if NeedSignalStrengthReading(clock) {
			signalStrengthSum += signalStrength
		}
	}
	return signalStrengthSum
}

func ModelCrt(filepath string, cycleCount int) string {
	instructionList := utils.Read_file(filepath)
	instructionPointer := 0

	CrtString := ""

	proc := CreateProcessor()

	for clock := 1; clock <= cycleCount; clock++ {
		if IsIdle(&proc) {
			SendInstruction(&proc, instructionList[instructionPointer])
			instructionPointer++
		}
		register := proc.registerX
		_ = ClockTick(&proc)
		CrtString = CrtString + GetCrtOutput(clock, register)
	}
	return CrtString
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
