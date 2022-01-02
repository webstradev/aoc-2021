package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var convertTestCases = []struct {
	name       string
	stringList []string
	expected   []int
}{
	{
		"A few numbers",
		[]string{"1", "2", "3"},
		[]int{1, 2, 3},
	},
	{
		"A few numbers and a word",
		[]string{"1", "2", "3", "four"},
		nil,
	},
}

func TestConvertListOfStrings(t *testing.T) {
	for _, test := range convertTestCases {
		t.Run(test.name, func(t *testing.T) {
			got, _ := ConvertListOfStringsToInts(test.stringList)
			assert.Equal(t, test.expected, got)
		})
	}
}
