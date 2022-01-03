package day3

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/webstradev/advent-of-code/util"
)

var binaryDigits = "01"

// validateBinaryInput checks if a string only contains 0's and 1's and optionally checks the length as well
func validateBinaryInput(input string, digits int) bool {
	// Check if string is of correct length, only do this if digits param is greater than 0.
	if digits > 0 && len(input) != digits {
		return false
	}

	// Check if each character is either a 0 or a 1
	for _, char := range input {
		if !strings.ContainsRune(binaryDigits, char) {
			return false
		}
	}

	return true
}

// caclulateGammaValue calculates the gamma value by returning the most common bit for each position
func calculateGammaValue(input []string) (string, error) {
	// Number of digits will be defined by the length of the first input and all the other inputs will be checkd to match the same length
	digits := len(input[0])

	// This will track how many ones are found for each position in the list of inputs
	ones := make([]int, digits)

	// Loop over each of the binary inputs
	for _, binary := range input {
		// Check if input is valid
		valid := validateBinaryInput(binary, digits)
		if !valid {
			return "", fmt.Errorf("[ERROR] Invalid Binary Input Found: '%s'", binary)
		}
		// Loop over each digit of this binary
		for index, digit := range binary {
			// If this digit is 1 then increase the count for that position
			if digit == '1' {
				ones[index] += 1
			}
		}
	}

	result := ""

	// Loop over each of the counts
	for _, count := range ones {
		// If the count is larger or equal than half the length of the enitre input, then one is the most common digit, otherwise it will be 0
		if float64(count) >= float64(len(input))/2 {
			result += "1"
		} else {
			result += "0"
		}
	}

	return result, nil
}

// caclulateEpsilonFromGamma reverses a binary string (1's become 0's and v.v.)
func calculateEpsilonFromGamma(gammaInput string) (string, error) {
	// Validate epsilon value (digits 0) ignores the length check
	if valid := validateBinaryInput(gammaInput, 0); !valid {
		return "", fmt.Errorf("[ERROR] Invalid Binary Input Found: '%s'", gammaInput)
	}

	result := ""

	// Flip 1's and 0's
	for _, digit := range gammaInput {
		if digit == '0' {
			result += "1"
		}
		if digit == '1' {
			result += "0"
		}
	}

	return result, nil
}

// calculatePowerConsumption Calculate the power consumption rate given a list of binary inputs
func calculatePowerConsumption(input []string) (int, error) {
	// Calculate gamma rate
	gammaBin, err := calculateGammaValue(input)
	if err != nil {
		return -1, err
	}

	// Convert binary to decimal
	gammaDec, err := strconv.ParseInt(gammaBin, 2, 64)
	if err != nil {
		return -1, err
	}

	// Calculate epsilon rate
	epsilonBin, err := calculateEpsilonFromGamma(gammaBin)
	if err != nil {
		return -1, err
	}

	// Convert binary to decimal
	epsilonDec, err := strconv.ParseInt(epsilonBin, 2, 64)
	if err != nil {
		return -1, err
	}

	// return power consumption
	return int(gammaDec * epsilonDec), nil
}

// calculateRating Calculate the o2 generator or co2 scrubber rating for an input
func calculateRating(input []string, ratingType string, position int) (string, error) {

	// Validate Rating Type
	if !(ratingType == "o2" || ratingType == "co2") {
		return "", errors.New("invalid rating type passed")
	}

	// If the input only has one binary then return this as the result
	if len(input) == 1 {
		return input[0], nil
	}

	// Calculate gamma rate and set it to the bit condition
	bitCondition, err := calculateGammaValue(input)
	if err != nil {
		return "", err
	}

	// If the rating we want is the co2 scrubber rating then we use the epsilon rate instead of the gamma rate
	if ratingType == "co2" {
		bitCondition, err = calculateEpsilonFromGamma(bitCondition)
		if err != nil {
			return "", err
		}
	}

	// Empty result slice
	result := []string{}

	// Find only the binary codes that have the matching position
	for _, binary := range input {
		if binary[position] == bitCondition[position] {
			result = append(result, binary)
		}
	}

	return calculateRating(result, ratingType, position+1)
}

// calculateLifeSupportRating calculates the o2 generator rating and the co2 scrubber rating and multiplies them to return the life support rating
func calculateLifeSupportRating(input []string) (int, error) {
	// calulate o2 generator rating recursively
	o2RatingBin, err := calculateRating(input, "o2", 0)
	if err != nil {
		return -1, err
	}

	// Convert binary to decimal
	o2RatingDec, err := strconv.ParseInt(o2RatingBin, 2, 64)
	if err != nil {
		return -1, err
	}

	// calculate c02 scrubber rating recursively
	co2RatingBin, err := calculateRating(input, "co2", 0)
	if err != nil {
		return -1, err
	}

	co2RatingDec, err := strconv.ParseInt(co2RatingBin, 2, 64)
	if err != nil {
		return -1, err
	}

	return int(o2RatingDec * co2RatingDec), nil
}

func Solve() {
	// Read input file
	input, err := util.ReadLinesToSlice("./day3/input.txt")
	if err != nil {
		log.Println("[ERROR] Failed to read input file for day 3")
		return
	}

	// Solve Puzzle 1
	powerConsumtion, err := calculatePowerConsumption(input)
	if err != nil {
		log.Printf("[ERROR] An error occurred while calulating power consumption: %v", err)
		return
	}

	// Solve Puzzle 2
	lifeSupportRating, err := calculateLifeSupportRating(input)
	if err != nil {
		log.Printf("[ERROR] An error occurred while calulating Life Support rating: %v", err)
		return
	}

	log.Println("----------")
	log.Println("Day 3:")
	log.Printf("day3puzzle1 - Power Consumption Rate: %d", powerConsumtion)
	log.Printf("day3puzzle2 - Life Support Rating: %d", lifeSupportRating)
	log.Print("----------")
}
