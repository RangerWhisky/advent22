package day7

import (
	file_input "localhost/advent22/utils"
	"math"
)

func PartOne(filepath string) int {
	// get file contents
	// when cd command is found
	//    add new entry to the Filesystem
	//    fill out the directory details
	//

	fs := GetFsFromFile(filepath)

	cleanableSize := 0

	for dirName := range fs.directoryList {
		recursiveSize := GetSizeOnDisk(&fs, dirName)
		if recursiveSize <= 100000 {
			cleanableSize += recursiveSize
		}
	}

	return cleanableSize
}

func PartTwo(filepath string) int {
	fs := GetFsFromFile(filepath)
	usedSpace := GetSizeOnDisk(&fs, "/")
	targetDirectorySize := GetTargetSize(usedSpace)

	bestCandidateSize := usedSpace

	for dirName := range fs.directoryList {
		recursiveSize := GetSizeOnDisk(&fs, dirName)
		if (recursiveSize >= targetDirectorySize) && (recursiveSize < bestCandidateSize) {
			bestCandidateSize = recursiveSize
		}
	}
	return bestCandidateSize
}

func GetFsFromFile(filepath string) Filesystem {
	elfInput := file_input.Read_file(filepath)
	return InitialiseFilesystem(elfInput)
}

func GetTargetSize(usedSpace int) int {
	freeSpace := 70000000 - usedSpace
	required := 30000000 - freeSpace

	return int(math.Max(float64(required), 0))
}
