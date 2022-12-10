package day7

import (
	file_input "localhost/advent22/utils"
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

func GetFsFromFile(filepath string) Filesystem {
	elfInput := file_input.Read_file(filepath)
	return InitialiseFilesystem(elfInput)
}
