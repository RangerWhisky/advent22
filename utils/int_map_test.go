package utils

import "testing"

const test_map_file = "./int_map.txt"

func TestInitialiseMap(t *testing.T) {
	file := Read_file(test_map_file)

	mapData := InitialiseMap(file)

	PrintMap(&mapData)

	if GetWidth(&mapData) != 5 {
		t.Errorf("Bad Width")
	}
	if GetHeight(&mapData) != 6 {
		t.Errorf("Bad Height")
	}
}
