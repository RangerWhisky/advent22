package day3

func GetItemPriority(character byte) int {
	priority := int(character) - 96

	if priority < 0 {
		priority += 58
	}

	return priority
}
