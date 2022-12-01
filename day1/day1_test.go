package day1

import (
	file_input "localhost/advent22/utils"
	"testing"
)

func TestHello(t *testing.T) {
	meaning := Meaning()
	if meaning != 42 {
		t.Errorf("Meaning is = %d, need 42", meaning)

	}
}

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
