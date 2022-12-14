package day10

import "testing"

func TestPartOne(t *testing.T) {
	solution := PartOne("./long_example.txt")

	if solution != 13140 {
		t.Errorf("Expecting %d, got %d", 13140, solution)
	}
}

func TestProcessorModelling20(t *testing.T) {
	solution := ModelProcessor("./long_example.txt", 20)

	if solution != 420 {
		t.Errorf("Expecting %d, got %d", 420, solution)
	}
}

func TestProcessorModelling220(t *testing.T) {
	solution := ModelProcessor("./long_example.txt", 220)

	if solution != 420+1140+1800+2940+2880+3960 {
		t.Errorf("Expecting %d, got %d", 420+1140+1800+2940+2880+3960, solution)
	}
}

func TestCrtModelling(t *testing.T) {
	expectedOutput := "##..##..##..##..##..#"

	solution := ModelCrt("./long_example.txt", 21)

	if solution != expectedOutput {
		t.Errorf("Expected %s, got %s", expectedOutput, solution)
	}
}
func TestFullCrtModelling(t *testing.T) {
	expectedOutput := "##..##..##..##..##..##..##..##..##..##..\n###...###...###...###...###...###...###.\n####....####....####....####....####....\n#####.....#####.....#####.....#####.....\n######......######......######......####\n#######.......#######.......#######.....\n"

	solution := ModelCrt("./long_example.txt", 240)

	if solution != expectedOutput {
		t.Errorf("Expected\n%s \n got\n%s", expectedOutput, solution)
	}
}
