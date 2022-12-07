package day7

import (
	"strconv"
	"strings"
)

func GetDirectorySize(contents [][]byte) int {
	size := 0
	for i := 0; i < len(contents); i++ {
		if IsDir(contents[i]) {
			size += GetDirectorySize(contents[i:])
		} else if IsFile(contents[i]) {
			parts := strings.Split(string(contents[i]), " ")
			fileSize, _ := strconv.Atoi(parts[0])
			size += fileSize
		}

	}
	return size
}
