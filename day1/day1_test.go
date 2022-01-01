package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	name           string
	measurements   []int
	expectedAmount int
}{
	{
		"empty input",
		[]int{},
		0,
	},
	{
		"5 numbers all increasing",
		[]int{1, 2, 3, 4, 5, 6},
		5,
	},
	{
		"5 numbers 5 increases, but interrupted by one",
		[]int{1, 2, 3, 2, 4, 5, 6},
		5,
	},
	{
		"5 numbers 0 increases, but interrupted by one",
		[]int{6, 5, 4, 3, 2, 1, 0},
		0,
	},
}

var slidingTestCases = []struct {
	name                string
	measurements        []int
	windowSize          int
	expectedMasurements []int
}{
	{
		"6 numbers - 2 day window",
		[]int{1, 2, 3, 4, 5, 6},
		2,
		[]int{3, 5, 7, 9, 11},
	},
	{
		"7 numbers - 2 day window",
		[]int{1, 2, 3, 4, 5, 6, 7},
		2,
		[]int{3, 5, 7, 9, 11, 13},
	},
	{
		"6 numbers - 3 day window",
		[]int{1, 2, 3, 4, 5, 6},
		3,
		[]int{6, 9, 12, 15},
	},
	{
		"7 numbers - 3 day window",
		[]int{1, 2, 3, 4, 5, 6, 7},
		3,
		[]int{6, 9, 12, 15, 18},
	},
}

func TestCountIncreasedMeasurements(t *testing.T) {
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got := countIncreasedMeasurements(test.measurements)
			assert.Equal(t, test.expectedAmount, got, "Did not calculate expected amount correctly")
		})
	}
}

func TestCreateSlidingWindow(t *testing.T) {
	for _, test := range slidingTestCases {
		t.Run(test.name, func(t *testing.T) {
			got := createSlidingWindowMeasurements(test.measurements, test.windowSize)
			assert.ElementsMatch(t, got, test.expectedMasurements, "Dit not caluclate sliding window sums correctly")
		})
	}
}
