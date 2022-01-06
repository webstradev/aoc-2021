package day4

import (
	"bufio"
	"log"
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
func (b *Board) sumUnMarkedValues() int {
	count := 0

	// Loop over each row of the board
	for _, x := range b {
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
func (b *Board) isBoardWinning() bool {
	// Check if there is a row with all -1's
	for _, x := range b {
		if x[0] == -1 && x[1] == -1 && x[2] == -1 && x[3] == -1 && x[4] == -1 {
			return true
		}
	}

	// Check if there is a column with all -1's
	for i := 0; i < 5; i++ {
		if b[0][i] == -1 && b[1][i] == -1 && b[2][i] == -1 && b[3][i] == -1 && b[4][i] == -1 {
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

func (b *Board) playTurn(number int) bool {
	// Loop over the rows
	for x, row := range b {
		// Loop over the numbers on that row
		for y, value := range row {
			// If we match the games next number then mark the number as -1
			if value == number {
				b[x][y] = -1
			}
		}
	}
	// Return whether or not the board is now winning
	return b.isBoardWinning()
}

// handleTurns loops over the input and marks of all boards until there is a winning board and returns this board's remaining sum
func (g *Game) playBingo() int {
	for _, nextNumber := range g.numbers {
		for index, board := range g.boards {
			// Play a turn on the board and check if it is now a winning
			boardIsNowWinning := board.playTurn(nextNumber)
			// update the board on the actual game
			g.boards[index] = board
			// If this is now a winning board then f
			if boardIsNowWinning {
				return board.sumUnMarkedValues() * nextNumber
			}
		}
	}

	return -1
}

func (g *Game) DeleteBoard(board Board) {
	for i, b := range g.boards {
		if b == board {
			g.boards = append(g.boards[:i], g.boards[i+1:]...)
			break
		}
	}
}

// handleTurns loops over the input and marks of all boards if there is a winning board it is removed and with one board remaining
func (g *Game) playReverseBingo() int {

	var lastWinningBoard *Board
	var lastInput int
	for _, nextNumber := range g.numbers {
		// Mark off the boards
		for index, board := range g.boards {
			// Play a turn on the board
			_ = board.playTurn(nextNumber)

			// update the board on the actual game
			g.boards[index] = board
		}

		// Check for winning boards and add them to the winning boards array
		winningBoards := []Board{}
		for _, board := range g.boards {
			if board.isBoardWinning() {
				winningBoards = append(winningBoards, board)
			}
		}

		if len(winningBoards) > 0 {
			for _, board := range winningBoards {
				lastInput = nextNumber
				lastWinningBoard = &board
				g.DeleteBoard(board)
			}
		}
	}

	return lastWinningBoard.sumUnMarkedValues() * lastInput
}

func playBingoFromInput(fileName string, findFirst bool) (int, error) {
	// Get the game input and parse it into a game struct
	game, err := parseGameFromInput(fileName)
	if err != nil {
		return -1, err
	}

	result := -1
	if findFirst {
		result = game.playBingo()
	} else {
		result = game.playReverseBingo()
	}

	if result == -1 {
		return -1, err
	}

	return result, nil
}

func Solve() {
	result1, err := playBingoFromInput("./day4/input.txt", true)
	if err != nil {
		log.Printf("[ERROR] Unable to find a winning board for this input")
	}
	result2, err := playBingoFromInput("./day4/input.txt", false)
	if err != nil {
		log.Printf("[ERROR] Unable to find a winning board for this input")
	}

	log.Println("----------")
	log.Println("Day 4:")
	log.Printf("day4puzzle1 - The score of the first winning board is %d", result1)
	log.Printf("day4puzzle2 - The score of the last winning board is %d", result2)
	log.Print("----------")
}
