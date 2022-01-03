package main

import (
	"log"
	"os"

	"github.com/webstradev/advent-of-code/day1"
	"github.com/webstradev/advent-of-code/day2"
	"github.com/webstradev/advent-of-code/day3"
)

// Maps passed day variable to the solver functions in the various packages
var dayToFunctionMapper = map[string]func(){
	"1": day1.Solve,
	"2": day2.Solve,
	"3": day3.Solve,
}

func main() {
	day := os.Getenv("day")

	// No day specified run all
	if day == "" {
		for _, dayFunction := range dayToFunctionMapper {
			dayFunction()
		}
		return
	}

	// Incorrect day specified return error
	if dayToFunctionMapper[day] == nil {
		log.Printf("[ERROR] Invalid Day parameter passed")
		return
	}

	// Run solver function for a specific day
	dayToFunctionMapper[day]()

}
