package day7

import (
	"strconv"
	"strings"
)

type Directory struct {
	subdirs []string
	size    int
}

func GetDirectorySize(contents [][]byte) int {
	dir := ParseDirectory(contents)
	return dir.size
}

func ParseDirectory(contents [][]byte) Directory {
	var dir = new(Directory)

	for i := 0; i < len(contents); i++ {
		if IsDir(contents[i]) {
			parts := strings.Split(string(contents[i]), " ")
			dir.subdirs = append(dir.subdirs, parts[1])
		} else if IsFile(contents[i]) {
			parts := strings.Split(string(contents[i]), " ")
			fileSize, _ := strconv.Atoi(parts[0])
			dir.size += fileSize
		}

	}
	return *dir
}
