package day1

import (
	file_input "localhost/advent22/utils"
	"strconv"
)

func Meaning() int {
	return 6 * 7
}

func Easy(filepath string) int {
	elfInput := file_input.Read_file(filepath)
	caloriesList := GetCalorieList(elfInput)
	return GetMaxCalories(caloriesList)
}

func GetCalorieList(input []string) []int {
	var calories []int
	running_total := 0
	for i := 0; i < len(input); i++ {
		item_val, err := strconv.Atoi(input[i])
		if err != nil || item_val == 0 {
			// push the value and move on
			calories = append(calories, running_total)
			running_total = 0
		} else {
			running_total += item_val
		}
	}
	calories = append(calories, running_total)
	return calories
}

func GetMaxCalories(input []int) int {
	maxCals := 0
	for _, value := range input {
		if value > maxCals {
			maxCals = value
		}
	}
	return maxCals
}
