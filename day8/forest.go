package day8

import (
	utils "localhost/advent22/utils"
)

type Forest struct {
	mapData    utils.IntMap
	visibility [][]bool
}

func InitialiseForestFromFile(filepath string) Forest {

	var forest Forest

	file := utils.Read_file("./simplified_example.txt")

	forest.mapData = utils.InitialiseMap(file)

	return forest
}

func AssessVisibility(forest *Forest) {
	// trade off of memory usage vs performance
	// for performance, store the top tree height observed from each direction, then iterate all
	// tree points and compare against neighbours only

	// for memory, for each point compare against neighbours until failure

	height := utils.GetHeight(&forest.mapData)
	width := utils.GetHeight(&forest.mapData)

	var visibilityMap [][]bool

	for i := 0; i < height; i++ {
		var boolArray []bool
		for j := 0; j < width; j++ {
			boolArray = append(boolArray, IsTreeVisible(forest, i, j))
		}
		visibilityMap = append(visibilityMap, boolArray)
	}

	forest.visibility = visibilityMap
}

func IsTreeVisible(forest *Forest, height int, width int) bool {
	treeHeight := utils.GetValue(&forest.mapData, width, height)

	// from north
	visibleFromDirection := IsVisibleFromNorth(forest, width, height, treeHeight) ||
		IsVisibleFromSouth(forest, width, height, treeHeight) ||
		IsVisibleFromWest(forest, width, height, treeHeight) ||
		IsVisibleFromEast(forest, width, height, treeHeight)

	return visibleFromDirection
}

func IsVisibleFromNorth(forest *Forest, height int, width int, treeHeight int) bool {
	visibleFromDirection := true
	for i := 0; i < height && visibleFromDirection; i++ {
		if utils.GetValue(&forest.mapData, i, width) >= treeHeight {
			visibleFromDirection = false
		}
	}
	return visibleFromDirection
}

func IsVisibleFromSouth(forest *Forest, height int, width int, treeHeight int) bool {
	visibleFromDirection := true
	maxHeight := utils.GetHeight(&forest.mapData)
	for i := height + 1; i < maxHeight && visibleFromDirection; i++ {
		if utils.GetValue(&forest.mapData, i, width) >= treeHeight {
			visibleFromDirection = false
		}
	}
	return visibleFromDirection
}

func IsVisibleFromWest(forest *Forest, height int, width int, treeHeight int) bool {
	visibleFromDirection := true
	for i := 0; i < width && visibleFromDirection; i++ {
		if utils.GetValue(&forest.mapData, height, i) >= treeHeight {
			visibleFromDirection = false
		}
	}
	return visibleFromDirection
}

func IsVisibleFromEast(forest *Forest, height int, width int, treeHeight int) bool {
	visibleFromDirection := true
	maxWidth := utils.GetWidth(&forest.mapData)
	for i := width + 1; i < maxWidth && visibleFromDirection; i++ {
		if utils.GetValue(&forest.mapData, height, i) >= treeHeight {
			visibleFromDirection = false
		}
	}
	return visibleFromDirection
}

func getTreeCount(forest *Forest) int {
	return utils.GetHeight(&forest.mapData) * utils.GetWidth(&forest.mapData)
}
