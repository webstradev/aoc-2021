package day2

import (
	"log"
	"strconv"
	"strings"

	"github.com/webstradev/advent-of-code/util"
)

// Command is the type for a parsed and validated command
type Command struct {
	Direction string
	Step      int
}

// validateDirectionString check if string part is valid
func validateDirectionString(input string) bool {
	switch input {
	case "forward", "down", "up":
		return true
	default:
		return false
	}
}

// Converts a string command into a struct
func createCommandFromString(input string) *Command {
	// Split input on the space seperator
	s := strings.Split(input, " ")
	if len(s) != 2 {
		return nil
	}

	// Check if first part is a valid direction
	if !validateDirectionString(s[0]) {
		return nil
	}

	// Parse second string to int
	stepSize, err := strconv.Atoi(s[1])
	if err != nil {
		return nil
	}

	return &Command{Direction: s[0], Step: stepSize}
}

// calculateCommandResults caluclates the postion and depth of the submarine given a list of commands
func calculateCommandResults(commands []string) (int, int) {
	position, depth := 0, 0

	for _, commandString := range commands {
		command := createCommandFromString(commandString)
		if command == nil {
			log.Printf("[ERROR] Failed to parse '%s' into a valid command. command ignored", commandString)
		} else {
			switch command.Direction {
			case "forward":
				position += command.Step
			case "up":
				if depth-command.Step < 0 {
					depth = 0
				} else {
					depth -= command.Step
				}
			case "down":
				depth += command.Step
			}
		}

	}

	return position, depth
}

// calculateNewCommandResults calculates the postion and depth of the submarine using aim method for a given list of commands
func calculateNewCommandResults(commands []string) (int, int) {
	position, depth, aim := 0, 0, 0

	for _, commandString := range commands {
		command := createCommandFromString(commandString)
		if command == nil {
			log.Printf("[ERROR] Failed to parse '%s' into a valid command. command ignored", commandString)
		} else {
			switch command.Direction {
			case "forward":
				position += command.Step
				depth += command.Step * aim
			case "up":
				aim -= command.Step
			case "down":
				aim += command.Step
			}
		}

	}

	return position, depth
}

type Position struct {
	Horizontal int
	Depth      int
}

func Solve() {
	// Read input file
	input, err := util.ReadLinesToSlice("./day2/input.txt")
	if err != nil {
		log.Println("[ERROR] Failed to read input file for day 2")
		return
	}

	// Solve Puzzle 1
	x, y := calculateCommandResults(input)

	// Solve Puzzle 2
	xNew, yNew := calculateNewCommandResults(input)

	log.Println("----------")
	log.Println("Day 2:")
	log.Printf("day2puzzle1 - Multiplying the coordinates of the final postion gives: %d", x*y)
	log.Printf("day2puzzle2 - Multiplying the coordinates of the final postion gives: %d", xNew*yNew)
	log.Print("----------")
}
