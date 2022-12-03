package main

import (
	"fmt"
	"localhost/advent22/day1"
	"localhost/advent22/day2"
	"localhost/advent22/day3"
)

func main() {
	fmt.Println("Day 1")

	fmt.Printf("Biggest pack is %d \n", day1.PartOne("day1/input"))
	fmt.Printf("Total of biggest 3 packs is %d \n", day1.PartTwo("day1/input"))

	fmt.Println("\nDay 2")
	fmt.Printf("Score for strategy card is %d\n", day2.PartOne("day2/input"))
	fmt.Printf("Score for strategy card with target result %d\n", day2.PartTwo("day2/input"))

	fmt.Println("\nDay 3")
	fmt.Printf("Total priority for incorrect packing is %d\n", day3.PartOne("day3/input"))
	fmt.Printf("Total badge priority for groups is %d\n", day3.PartTwo("day3/input"))

	fmt.Println("\nDone!")
}
