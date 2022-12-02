package day2

import "fmt"

func Decode(character byte) int {
	fmt.Printf("%q", character)
	value := 0
	switch character {
	case 'X', 'A':
		value = 1
	case 'Y', 'B':
		value = 2
	case 'Z', 'C':
		value = 3
	}

	return value
}
