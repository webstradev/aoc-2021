package day4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumBoard(t *testing.T) {
	tests := []struct {
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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := sumUnMarkedValues(test.board)
			assert.Equal(t, test.sum, got, "Calculated sum was unexpected")
		})
	}
}

func TestIsBoardWinning(t *testing.T) {
	tests := []struct {
		name    string
		board   Board
		winning bool
	}{
		{
			"Winning board - horizontal win",
			Board{
				{-1, -1, -1, -1, -1},
				{10, 16, 15, -1, 19},
				{18, 8, -1, 26, 20},
				{22, -1, 13, 6, -1},
				{-1, -1, 12, 3, -1},
			},
			true,
		},
		{
			"Winning board - vertical win",
			Board{
				{-1, 15, 0, -1, 22},
				{-1, 18, 13, -1, 5},
				{-1, 8, -1, 25, 23},
				{-1, 11, -1, 24, 4},
				{-1, 21, 16, 12, -1},
			},
			true,
		}, {
			"Losing board",
			Board{
				{-1, 15, 0, -1, 22},
				{15, 18, 13, -1, 5},
				{-1, 8, -1, 25, 23},
				{-1, 11, -1, 24, 4},
				{-1, 21, 16, 12, -1},
			},
			false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := isBoardWinning(test.board)
			assert.Equal(t, test.winning, got, "The function did not return an expected winning bool")
		})
	}
}

func TestParseGameFromInput(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected *Game
		success  bool
	}{
		{
			"Invalid Path",
			"./test.xyz",
			nil,
			false,
		},
		{
			"Valid Input",
			"./input_test.txt",
			&Game{
				[]int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1},
				[]Board{
					{
						{22, 13, 17, 11, 0},
						{8, 2, 23, 4, 24},
						{21, 9, 14, 16, 7},
						{6, 10, 3, 18, 5},
						{1, 12, 20, 15, 19},
					},
					{
						{3, 15, 0, 2, 22},
						{9, 18, 13, 17, 5},
						{19, 8, 7, 25, 23},
						{20, 11, 10, 24, 4},
						{14, 21, 16, 12, 6},
					},
					{
						{14, 21, 17, 24, 4},
						{10, 16, 15, 9, 19},
						{18, 8, 23, 26, 20},
						{22, 11, 13, 6, 5},
						{2, 0, 12, 3, 7},
					},
				},
			},
			true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := parseGameFromInput(test.path)
			assert.Equal(t, test.success, err == nil, "Expected success does not match error state")
			assert.Equal(t, test.expected, got, "The function did not return the expected game")
		})
	}
}
