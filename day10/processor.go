package day10

type Processor struct {
	cycle         int
	registerX     int
	nextFreeCycle int
	commandQueue  Instruction
}

type Instruction struct {
	ins InstructionType
	val int
}

type InstructionType int

const (
	none InstructionType = iota
	noop
	addx
)

func CreateProcessor() Processor {
	proc := Processor{0, 1, 0, getBlankInstruction()}
	return proc
}

func QueueNoop(proc *Processor) {
	noop := Instruction{InstructionType(noop), 0}
	proc.commandQueue = noop
	proc.nextFreeCycle = proc.cycle + 1
}

func QueueAddx(proc *Processor, val int) {
	addx := Instruction{InstructionType(addx), val}
	proc.commandQueue = addx
	proc.nextFreeCycle = proc.cycle + 2
}

func SetIdle(proc *Processor) {
	proc.commandQueue = getBlankInstruction()
}

func GetSignalStrength(proc *Processor) int {
	return proc.cycle * proc.registerX
}

func ClockTick(proc *Processor) int {
	proc.cycle++
	currentSignal := GetSignalStrength(proc)
	if proc.nextFreeCycle == proc.cycle {
		processInstruction(proc)
		SetIdle(proc)
	}
	return currentSignal
}
func IsIdle(proc *Processor) bool {
	return proc.commandQueue.ins == InstructionType(none)
}

func processInstruction(proc *Processor) {
	switch proc.commandQueue.ins {
	case addx:
		proc.registerX += proc.commandQueue.val
	}
}

func getBlankInstruction() Instruction {
	return Instruction{InstructionType(none), 0}
}
