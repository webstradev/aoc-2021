package day4

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
