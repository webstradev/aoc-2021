package day4

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// A game contains the numbers to draw and a list of bingo boards
type Game struct {
	numbers []int
	boards  []Board
}

// A board is a 5x5 matrix of ints
type Board [5][5]int

// sumBoard sums all the unmarked values on a board
func sumUnMarkedValues(board Board) int {
	count := 0

	// Loop over each row of the board
	for _, x := range board {
		// Loop over each column of that row
		for _, y := range x {
			// marked values are representend as -1 so we don't want to count these
			if y > -1 {
				count += y
			}

		}
	}

	return count
}

// isBoardWinning checks if a board is a winning board
func isBoardWinning(board Board) bool {
	// Check if there is a row with all -1's
	for _, x := range board {
		if x[0] == -1 && x[1] == -1 && x[2] == -1 && x[3] == -1 && x[4] == -1 {
			return true
		}
	}

	// Check if there is a column with all -1's
	for i := 0; i < 5; i++ {
		if board[0][i] == -1 && board[1][i] == -1 && board[2][i] == -1 && board[3][i] == -1 && board[4][i] == -1 {
			return true
		}
	}

	// No winning rows or columns found:
	return false
}

// splitStringProperly splits a string and ignores empty strings and new line symbols
func splitStringProperly(s, delim string) []string {
	ss := strings.Split(s, delim)
	var ns []string
	// Loop over each value after the split and only include it in the result if it isn't empty
	for _, v := range ss {
		if v == "" || v == "\n" {
			continue
		}
		ns = append(ns, v)
	}
	return ns
}

// parseGameFromInput parse the game numbers and board from an input file
func parseGameFromInput(path string) (*Game, error) {
	var game Game

	// Read the file
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	// Create a new scanner to loop over lines
	fs := bufio.NewScanner(f)

	// Select the first line (this will contain the game input)
	fs.Scan()

	// Loop over the numbers on the line and parse them to ints, then add them to the game
	for _, v := range strings.Split(fs.Text(), ",") {
		parsed, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		game.numbers = append(game.numbers, parsed)
	}

	// Scan over the remaining lines
	for fs.Scan() {
		var board Board

		// Ignore whitespace/ empty lines (these are used to seperate boards)
		if fs.Text() == " " || fs.Text() == "\n" || fs.Text() == "" {
			continue
		}

		// If we reach here we have reached the start of a board

		// Loop over the 5 rows of the board
		for x := 0; x < 5; x++ {
			// Loop over the items on the row (uses splitStringProperly to ignore extra spaces and whitespaces)
			for y, value := range splitStringProperly(fs.Text(), " ") {
				parsed, err := strconv.Atoi(string(value))
				if err != nil {
					return nil, err
				}
				// Add value to the board at the correct position
				board[x][y] = parsed
			}
			// Move the pointer to the next line so it doesn't get run again when we finish the board loop
			fs.Scan()
		}
		// Add the board to the game
		game.boards = append(game.boards, board)
	}
	return &game, nil
}
