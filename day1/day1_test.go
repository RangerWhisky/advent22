package day1

import (
	file_input "localhost/advent22/utils"
	"testing"
)

func TestElfCalorieTotal(t *testing.T) {
	elfInput := file_input.Read_file("../utils/file_input_sample.txt")
	caloriesList := GetCalorieList(elfInput)

	if caloriesList[0] != 6000 {
		t.Errorf("Calories listed as %d but should be 6000", caloriesList[0])
	}
}

func TestAllElfCalorieTotals(t *testing.T) {

	elfInput := file_input.Read_file("../utils/file_input_sample.txt")
	caloriesList := GetCalorieList(elfInput)
	expectedInputs := []int{6000, 4000, 11000, 24000, 10000}

	for index, expectedValue := range expectedInputs {
		if caloriesList[index] != expectedValue {
			t.Errorf("Calories listed as %d but should be %d", caloriesList[index], expectedValue)
		}
	}
}

func TestGetMaxCalorieLoad(t *testing.T) {
	elfInput := file_input.Read_file("../utils/file_input_sample.txt")
	caloriesList := GetCalorieList(elfInput)
	maxCalorieLoad := GetMaxCalories(caloriesList)

	if maxCalorieLoad != 24000 {
		t.Errorf("Max Calories calculated as %d but should be %d", maxCalorieLoad, 24000)
	}
}

func TestTopOneTotal(t *testing.T) {
	elfInput := file_input.Read_file("../utils/file_input_sample.txt")
	caloriesList := GetCalorieList(elfInput)
	total := GetTopCalorieTotal(caloriesList, 1)

	if total != 24000 {
		t.Errorf("Max Calories calculated as %d but should be %d", total, 24000)
	}
}

func TestEasySolution(t *testing.T) {
	solution := PartOne("../utils/file_input_sample.txt")
	if solution != 24000 {
		t.Errorf("Sample solution Calories calculated as %d but should be %d", solution, 24000)
	}
}

func TestBonusSolution(t *testing.T) {
	solution := PartTwo("../utils/file_input_sample.txt")
	if solution != 45000 {
		t.Errorf("Sample solution Calories calculated as %d but should be %d", solution, 45000)
	}
}
