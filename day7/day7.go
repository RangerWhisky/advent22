package day7

import (
	"fmt"
	file_input "localhost/advent22/utils"
)

func PartOne(filepath string) int {
	// get file contents
	// when cd command is found
	//    add new entry to the Filesystem
	//    fill out the directory details
	//

	fs := GetFsFromFile(filepath)

	PrintFilesystem(&fs, "/", "")

	fmt.Println("***")

	cleanableSize := 0

	for dirName := range fs.directoryList {
		recursiveSize := GetSizeOnDisk(&fs, dirName)
		fmt.Printf("%s (%d)\n", dirName, recursiveSize)
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
