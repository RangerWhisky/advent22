package day7

func IsCmd(cmd []byte) bool {
	return cmd[0] == byte('$')
}

func IsFile(cmd []byte) bool {
	file := !(IsCmd(cmd)) && cmd[0] != byte('d')
	return file
}
