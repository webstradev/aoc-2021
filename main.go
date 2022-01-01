package main

import (
	"log"

	"github.com/webstradev/advent-of-code/day1"
	"github.com/webstradev/advent-of-code/day2"
)

func main() {
	log.Print("----------")
	log.Println("Day 1:")
	day1puzzle1, day1puzzle2 := day1.SolveDay1()
	log.Printf("day1puzzle1 - amount of increases found : %d", day1puzzle1)
	log.Printf("day1puzzle2 - amount of increases found : %d", day1puzzle2)
	log.Print("----------")
	log.Print("----------")
	log.Println("Day 2:")
	day2puzzle1 := day2.SolveDay2()
	log.Printf("day2puzzle1 - Multiplying the coordinates of the final postion gives: %d", day2puzzle1.Horizontal*day2puzzle1.Depth)
	log.Print("----------")
}
