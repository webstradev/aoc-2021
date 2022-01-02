package day1

import (
	"log"

	"github.com/webstradev/advent-of-code/util"
)

// countIncreasedMeasurements counts how many measurements are an increase compared to the previous measurement
func countIncreasedMeasurements(measurements []int) int {
	count := 0

	for index, measurement := range measurements {
		// Skip the first element
		if index == 0 {
			continue
		}

		// Calculate if the measurement is greater than previous measurement
		if measurement > measurements[index-1] {
			count += 1
		}
	}

	return count
}

// createSlidingWindowMeasurements creates a sliding window of n days for the measurements passed
func createSlidingWindowMeasurements(measurements []int, windowSize int) []int {
	result := []int{}

	// Loop over all measurements
	for index, measurement := range measurements {
		// If there aren't enough measurements to create a n-day window
		if index > len(measurements)-windowSize {
			break
		}

		// Start with the value of the current measurement we are on
		windowSum := measurement
		// Grab the next n elements and add them to the count
		for i := index + 1; i < index+windowSize; i++ {
			windowSum += measurements[i]
		}

		// Append this result to the slice
		result = append(result, windowSum)
	}

	return result
}

func Solve() {
	// Read input file
	inputStrings, err := util.ReadLinesToSlice("./day1/input.txt")
	if err != nil {
		log.Println("[ERROR] Failed to read input file for day 1")
		return
	}

	// Parse integers from input result
	input, err := util.ConvertListOfStringsToInts(inputStrings)
	if err != nil {
		log.Println("[ERROR] Failed to convert strings to ints")
		return
	}

	// Solve puzzles
	day1puzzle1, day1puzzle2 := countIncreasedMeasurements(input), countIncreasedMeasurements(createSlidingWindowMeasurements(input, 3))

	log.Println("----------")
	log.Println("Day 1:")
	log.Printf("day1puzzle1 - amount of increases found : %d", day1puzzle1)
	log.Printf("day1puzzle2 - amount of increases found : %d", day1puzzle2)
	log.Print("----------")
}
