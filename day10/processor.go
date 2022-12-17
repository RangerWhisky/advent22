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
}

func GetSignalStrength(proc *Processor) int {
	return proc.cycle * proc.registerX
}

func getBlankInstruction() Instruction {
	return Instruction{InstructionType(none), 0}
}
