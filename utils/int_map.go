package utils

import (
	"fmt"
	"strconv"
)

type IntMap struct {
	mapPoints [][]int
}

func InitialiseMap(data []string) IntMap {
	var mapData IntMap
	var mapPoints [][]int
	for _, dataLine := range data {
		var linePoints []int
		charSet := []byte(dataLine)
		for _, character := range charSet {
			value, _ := strconv.Atoi(string(character))
			linePoints = append(linePoints, value)
		}
		mapPoints = append(mapPoints, linePoints)
	}

	mapData.mapPoints = mapPoints

	return mapData
}

func PrintMap(mapData *IntMap) {
	height := GetHeight(mapData)
	width := GetWidth(mapData)
	for line := 0; line < height; line++ {
		for i := 0; i < width; i++ {
			fmt.Printf("%d", GetValue(mapData, line, i))
		}
		fmt.Println()
	}
}

func GetValue(mapData *IntMap, height int, width int) int {
	return mapData.mapPoints[height][width]
}

func GetWidth(mapData *IntMap) int {
	return len(mapData.mapPoints[0])
}

func GetHeight(mapData *IntMap) int {
	return len(mapData.mapPoints)
}
