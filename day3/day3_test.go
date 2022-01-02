package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var validateBinaryInputTests = []struct {
	name   string
	input  string
	length int
	valid  bool
}{
	{
		"valid binary but incorrect length",
		"101010",
		12,
		false,
	},
	{
		"valid binary and correct length",
		"101010",
		6,
		true,
	},
	{
		"letter in string",
		"10a010",
		6,
		false,
	},
	{
		"number 2 in string",
		"102010",
		6,
		false,
	},
}

func TestValidateBinaryInput(t *testing.T) {
	for _, test := range validateBinaryInputTests {
		t.Run(test.name, func(t *testing.T) {
			valid := validateBinaryInput(test.input, test.length)
			assert.Equal(t, test.valid, valid, "Input Validation returned unexpected boolean value")
		})
	}
}

var calculateGammaTests = []struct {
	name     string
	input    []string
	expected string
}{
	{
		"4digit binary",
		[]string{
			"0101",
			"1010",
			"1010",
			"1010",
		},
		"1010",
	},
}
