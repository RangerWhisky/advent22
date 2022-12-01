package day1

import (
	file_input "localhost/advent22/utils"
	"sort"
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

func Bonus(filepath string) int {
	elfInput := file_input.Read_file(filepath)
	caloriesList := GetCalorieList(elfInput)
	return GetTopCalorieTotal(caloriesList, 3)
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
	// sort before checking
	sort.Ints(input[:])

	return input[len(input)-1]
}

func GetTopCalorieTotal(input []int, count int) int {
	sort.Ints(input[:])
	total := 0
	for i := 1; i <= count; i++ {
		total += input[len(input)-i]
	}
	return 0
}
