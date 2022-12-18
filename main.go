package main

import (
	"fmt"
	"localhost/advent22/day1"
	"localhost/advent22/day10"
	"localhost/advent22/day11"
	"localhost/advent22/day2"
	"localhost/advent22/day3"
	"localhost/advent22/day4"
	"localhost/advent22/day5"
	"localhost/advent22/day6"
	"localhost/advent22/day7"
	"localhost/advent22/day8"
	"localhost/advent22/day9"
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

	fmt.Println("\nDay 4")
	fmt.Printf("Count of superset sections is %d\n", day4.PartOne("day4/input"))
	fmt.Printf("Count of overlapping sections is %d\n", day4.PartTwo("day4/input"))

	fmt.Println("\nDay 5")
	fmt.Printf("Top Level Containers are %s\n", day5.PartOne("day5/input"))
	fmt.Printf("Top Level Containers for 9001 version are %s\n", day5.PartTwo("day5/input"))

	fmt.Println("\nDay 6")
	fmt.Printf("Signal starts at %d\n", day6.PartOne("day6/input"))
	fmt.Printf("Message starts at %d\n", day6.PartTwo("day6/input"))

	fmt.Println("\nDay 7")
	fmt.Printf("Cleanable file total size is %d\n", day7.PartOne("day7/input"))
	fmt.Printf("Dir size to remove is %d\n", day7.PartTwo("day7/input"))

	fmt.Println("\nDay 8")
	fmt.Printf("Number of visible trees is %d\n", day8.PartOne("day8/input"))
	fmt.Printf("Best Scenic Score is %d\n", day8.PartTwo("day8/input"))

	fmt.Println("\nDay 9")
	fmt.Printf("Number of visited squares is %d\n", day9.PartOne("day9/input"))
	fmt.Printf("Number of visited squares with 10 long rope is %d\n", day9.PartTwo("day9/input"))

	fmt.Println("\nDay 10")
	fmt.Printf("Signal strength at 220 is %d\n", day10.PartOne("day10/input"))
	fmt.Printf("Cipher is\n%s\n", day10.PartTwo("day10/input"))

	fmt.Println("\nDay 11")
	fmt.Printf("Monkey Business score is %d\n", day11.PartOne("day11/input"))
	fmt.Printf("Monkey Business without calming measures is %d\n", day11.PartTwo("day11/input"))

	fmt.Println("\nDone!")
}
