package day4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var sumUnmarkedValuesTests = []struct {
	name  string
	board Board
	sum   int
}{
	{
		"sumBoard1",
		Board{
			{-1, 13, 17, 11, -1},
			{8, -1, 23, 4, 24},
			{21, 9, -1, 16, 7},
			{6, -1, 3, -1, 5},
			{1, 12, -1, 15, -1},
		},
		195,
	},
	{
		"sumBoard2",
		Board{
			{-1, 15, 0, -1, 22},
			{-1, 18, 13, -1, 5},
			{-1, 8, -1, 25, 23},
			{-1, 11, -1, 24, 4},
			{-1, 21, 16, 12, -1},
		},
		217,
	},
	{
		"sumBoard3",
		Board{
			{-1, -1, -1, -1, -1},
			{10, 16, 15, -1, 19},
			{18, 8, -1, 26, 20},
			{22, -1, 13, 6, -1},
			{-1, -1, 12, 3, -1},
		},
		188,
	},
}

func TestSumBoard(t *testing.T) {
	for _, test := range sumUnmarkedValuesTests {
		t.Run(test.name, func(t *testing.T) {
			got := sumUnMarkedValues(test.board)
			assert.Equal(t, test.sum, got, "Calculated sum was unexpected")
		})
	}
}
