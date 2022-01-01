package main

import (
	"log"

	"github.com/webstradev/advent-of-code/day1"
)

func main() {
	day1puzzle1, day1puzzle2 := day1.SolveDay1()
	log.Printf("day1puzzle1 - amount of increases found : %d", day1puzzle1)
	log.Printf("day1puzzle2 - amount of increases found : %d", day1puzzle2)
}
