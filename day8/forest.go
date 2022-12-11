package day8

import (
	utils "localhost/advent22/utils"
)

type Forest struct {
	mapData utils.IntMap
}

func InitialiseForestFromFile(filepath string) Forest {

	var forest Forest

	file := utils.Read_file("./simplified_example.txt")

	forest.mapData = utils.InitialiseMap(file)

	return forest
}

func getTreeCount(forest *Forest) int {
	return utils.GetHeight(&forest.mapData) * utils.GetWidth(&forest.mapData)
}
