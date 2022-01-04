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
