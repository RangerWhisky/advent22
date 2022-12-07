package day7

func IsCmd(cmd []byte) bool {
	return cmd[0] == byte('$')
}
