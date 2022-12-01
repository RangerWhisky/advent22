package day1

import (
	"fmt"
	"strconv"
)

func Meaning() int {
	return 6 * 7
}

func GetCalorieList(input []string) []int {
	var calories []int
	running_total := 0
	for i := 0; i < len(input); i++ {
		item_val, err := strconv.Atoi(input[i])
		fmt.Printf("value is %d", item_val)
		if err != nil || item_val == 0 {
			// push the value and move on
			calories = append(calories, running_total)
		} else {
			running_total += item_val
		}
	}
	return calories
}
