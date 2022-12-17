package day10

type Processor struct {
	cycle     int
	registerX int
}

func CreateProcessor() Processor {
	proc := Processor{0, 1}
	return proc
}
