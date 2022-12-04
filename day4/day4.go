package day4

import (
	file_input "localhost/advent22/utils"
	"strconv"
	"strings"
)

func PartOne(filepath string) int {

	runningTotal := 0
	elfInput := file_input.Read_file(filepath)

	for i := 0; i < len(elfInput); i++ {
		if IsSectionSuperset(elfInput[i]) {
			runningTotal++
		}
	}

	return runningTotal
}

func GetSectionRange(rangeString string) (int, int) {
	rangeValues := strings.Split(rangeString, "-")

	start, _ := strconv.Atoi(rangeValues[0])
	end, _ := strconv.Atoi(rangeValues[1])

	return start, end
}

func IsSectionSuperset(sectionDescription string) bool {
	sections := strings.Split(sectionDescription, ",")
	leftStart, leftEnd := GetSectionRange(sections[0])
	rightStart, rightEnd := GetSectionRange(sections[1])

	fullOverlap := false

	if leftStart == rightStart {
		fullOverlap = true
	} else if leftStart <= rightStart {
		fullOverlap = rightEnd <= leftEnd
	} else {
		fullOverlap = leftEnd <= rightEnd
	}

	return fullOverlap
}

func IsOverlapping(sectionDescription string) bool {
	return false
}
