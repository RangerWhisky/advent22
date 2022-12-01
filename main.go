package main

import (
	"fmt"
	"localhost/advent22/day1"
)

func main() {
	fmt.Println("Day 1")

	fmt.Printf("Biggest pack is %d ", day1.Easy("day1/input"))
	fmt.Printf("Total of biggest 3 packs is %d ", day1.Bonus("day1/input"))
}
