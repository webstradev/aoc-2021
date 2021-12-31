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

func TestCountIncreasedMeasurements(t *testing.T) {
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got := countIncreasedMeasurements(test.measurements)
			assert.Equal(t, test.expectedAmount, got, "Did not calculate expected amount correctly")
		})
	}
}
