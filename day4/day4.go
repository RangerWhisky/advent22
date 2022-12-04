package day4

import (
	"strconv"
	"strings"
)

func GetSectionRange(rangeString string) (int, int) {
	rangeValues := strings.Split(rangeString, "-")

	start, _ := strconv.Atoi(rangeValues[0])
	end, _ := strconv.Atoi(rangeValues[1])

	return start, end
}

func GetSectionOverlap(sectionDescription string) bool {
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
