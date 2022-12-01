package file_input

import (
	"bufio"
	"fmt"
	"os"
)

func Read_file(filepath string) []string {
	var fileLines []string
	readFile, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	return fileLines
}
