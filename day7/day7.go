package day7

import (
	file_input "localhost/advent22/utils"
	"strings"
)

func PartOne(filepath string) int {
	// get file contents
	// when cd command is found
	//    add new entry to the Filesystem
	//    fill out the directory details
	//

	fs := InitialiseFilesystem()

	elfInput := file_input.Read_file(filepath)

	for i := 0; i < len(elfInput); i++ {
		parts := strings.Split(elfInput[i], " ")

		if parts[1] == "cd" {
			ChangeDir(&fs, parts[2])
			continue
		}

		if parts[1] == "ls" {
			// find the end of the directory list
			for j := 0; j < len(elfInput); j++ {
				peekLineParts := strings.Split(elfInput[j], " ")
				if peekLineParts[0] == "$" {
					dir := ParseDirectory(elfInput[i+1 : j])
					SaveDirectory(&fs, dir)
					i = j
				}
			}
		}
	}
	return 0
}
