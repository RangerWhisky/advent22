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
