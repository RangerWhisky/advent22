package day7

func IsCmd(cmd string) bool {
	return cmd[0] == byte('$')
}

func IsFile(cmd string) bool {
	file := !(IsCmd(cmd)) && cmd[0] != byte('d')
	return file
}

func IsDir(cmd string) bool {
	file := !(IsCmd(cmd)) && cmd[0] == byte('d')
	return file
}
