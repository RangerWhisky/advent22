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
	caloriesList := GetCalories(elfInput)
	if caloriesList[0] != 6000 {
		t.Errorf("Calories listed as %d but should be 6000", caloriesList[0])
	}
}
