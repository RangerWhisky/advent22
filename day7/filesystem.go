package day7

import (
	"fmt"
	"strings"
)

type Filesystem struct {
	directoryList map[string]Directory
	pwd           []string
}

func InitialiseFilesystem(elfInput []string) Filesystem {

	fs := InitialiseEmptyFilesystem()

	for i := 1; i < len(elfInput); i++ {
		fmt.Println(elfInput[i])
		parts := strings.Split(elfInput[i], " ")

		if parts[1] == "ls" {
			// find the end of the directory list
			start, end := GetDirContentIndices(elfInput, i)

			dir := ParseDirectory(elfInput[start:end])
			SaveDirectory(&fs, dir)
			i = end - 1
		} else if parts[1] == "cd" {
			ChangeDir(&fs, parts[2])
		}
	}

	return fs
}

func GetDirContentIndices(sample []string, ls_line int) (int, int) {
	content_start := ls_line + 1
	content_end := len(sample)
	// find the end of the directory list
	for j := content_start; j < len(sample); j++ {
		peekLineParts := strings.Split(sample[j], " ")
		if IsCmd(peekLineParts[0]) {
			content_end = j
			break
		}
	}

	return content_start, content_end
}

func InitialiseEmptyFilesystem() Filesystem {
	fs_map := make(map[string]Directory)

	fs := Filesystem{
		directoryList: fs_map,
		pwd:           []string{"/"},
	}

	return fs
}

func PrintWorkingDirectory(fs *Filesystem) string {
	return fs.pwd[len(fs.pwd)-1]
}

func ChangeDir(fs *Filesystem, newdir string) {
	switch newdir {
	case "..":
		newDirectoryHead := len(fs.pwd) - 1
		fs.pwd = fs.pwd[:newDirectoryHead]
	default:
		fs.pwd = append(fs.pwd, newdir)
	}
}

func SaveDirectory(fs *Filesystem, dir Directory) {
	currentDir := PrintWorkingDirectory(fs)
	fs.directoryList[currentDir] = dir
}

func GetDirectory(fs *Filesystem, path string) Directory {
	dir := fs.directoryList[path]

	return dir
}

func GetSizeOnDisk(fs *Filesystem, dirname string) int {

	dir := fs.directoryList[dirname]

	runningSize := dir.size

	for i := 0; i < len(dir.subdirs); i++ {
		runningSize += GetSizeOnDisk(fs, dir.subdirs[i])
	}
	return runningSize
}
