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
	success  bool // should be false if err != nil
}{
	{
		"4 * 4digit binary",
		[]string{
			"0101",
			"1010",
			"1010",
			"0010",
		},
		"1010",
		true,
	},
	{
		"5 * 4digit binary",
		[]string{
			"0101",
			"1010",
			"1010",
			"1010",
			"1101",
		},
		"1010",
		true,
	},
	{
		"5 * 4digit binary with invalid input",
		[]string{
			"0101",
			"1010",
			"1010",
			"1010",
			"110a",
		},
		"",
		false,
	},
}

func TestCalculateGammaResult(t *testing.T) {
	for _, test := range calculateGammaTests {
		t.Run(test.name, func(t *testing.T) {
			got, err := calculateGammaValue(test.input)
			assert.Equal(t, test.success, err == nil, "Expected success does not match error state")
			assert.Equal(t, test.expected, got, "Calculated gamma result was unexpected")
		})
	}
}

var calculateEpsilonFromGammaTests = []struct {
	gamma   string
	epsilon string
	success bool
}{
	{
		"1010",
		"0101",
		true,
	},
	{
		"101a",
		"",
		false,
	},
	{
		"1010001101",
		"0101110010",
		true,
	},
	{
		"10100011",
		"01011100",
		true,
	},
}

func TestCalculateEpsilonFromGamma(t *testing.T) {
	for _, test := range calculateEpsilonFromGammaTests {
		t.Run(test.gamma, func(t *testing.T) {
			got, err := calculateEpsilonFromGamma(test.gamma)
			assert.Equal(t, test.success, err == nil, "Expected success does not match error state")
			assert.Equal(t, test.epsilon, got, "Calculated epsilon result was unexpected")
		})
	}
}

var calculatePowerConsumptionTests = []struct {
	name     string
	input    []string
	expected int
	success  bool
}{
	{
		"normal valid input",
		[]string{
			"0101",
			"1010",
			"1010",
			"0010",
		},
		50,
		true,
	},
	{
		"ivalid input in one of the binaries",
		[]string{
			"0101",
			"1b10",
			"1010",
			"0010",
		},
		-1,
		false,
	},
	{
		"incorrect length in one of the binaries",
		[]string{
			"0101",
			"11110",
			"1010",
			"0010",
		},
		-1,
		false,
	},
}

func TestCalculatePowerConsumption(t *testing.T) {
	for _, test := range calculatePowerConsumptionTests {
		t.Run(test.name, func(t *testing.T) {
			got, err := calculatePowerConsumption(test.input)
			assert.Equal(t, test.success, err == nil, "Expected success does not match error state")
			assert.Equal(t, test.expected, got, "Calculated power consumption rate was unexpected")
		})
	}
}

var calculateRatingTests = []struct {
	name       string
	input      []string
	ratingType string
	expected   string
	success    bool
}{
	{
		"invalid rating type - normal valid input",
		[]string{
			"00100",
			"11110",
			"10110",
			"10111",
			"10101",
			"01111",
			"00111",
			"11100",
			"10000",
			"11001",
			"00010",
			"01010",
		},
		"oxygen",
		"",
		false,
	},
	{
		"o2 - normal valid input",
		[]string{
			"00100",
			"11110",
			"10110",
			"10111",
			"10101",
			"01111",
			"00111",
			"11100",
			"10000",
			"11001",
			"00010",
			"01010",
		},
		"o2",
		"10111",
		true,
	},
	{
		"o2 - ivalid input in one of the binaries",
		[]string{
			"00100",
			"1111a",
			"10110",
			"10111",
			"10101",
			"01111",
			"00111",
			"11100",
			"10000",
			"11001",
			"00010",
			"01010",
		},
		"o2",
		"",
		false,
	},
	{
		"o2 - incorrect length in one of the binaries",
		[]string{
			"00100",
			"1111a",
			"10110",
			"10111",
			"10101",
			"01111",
			"00111",
			"11100",
			"10000",
			"11001",
			"00010",
			"01010",
		},
		"o2",
		"",
		false,
	},
	{
		"co2 - normal valid input",
		[]string{
			"00100",
			"11110",
			"10110",
			"10111",
			"10101",
			"01111",
			"00111",
			"11100",
			"10000",
			"11001",
			"00010",
			"01010",
		},
		"co2",
		"01010",
		true,
	},
	{
		"co2 - ivalid input in one of the binaries",
		[]string{
			"00100",
			"1111a",
			"10110",
			"10111",
			"10101",
			"01111",
			"00111",
			"11100",
			"10000",
			"11001",
			"00010",
			"01010",
		},
		"co2",
		"",
		false,
	},
	{
		"co2 - incorrect length in one of the binaries",
		[]string{
			"00100",
			"1111a",
			"10110",
			"10111",
			"10101",
			"01111",
			"00111",
			"11100",
			"10000",
			"11001",
			"00010",
			"01010",
		},
		"co2",
		"",
		false,
	},
}

func TestCalculateRating(t *testing.T) {
	for _, test := range calculateRatingTests {
		t.Run(test.name, func(t *testing.T) {
			got, err := calculateRating(test.input, test.ratingType, 0)
			assert.Equal(t, test.success, err == nil, "Expected success does not match error state")
			assert.Equal(t, test.expected, got, "Calculated Rating was unexpected")
		})
	}
}

var calculateLifeSupportRatingTests = []struct {
	name     string
	input    []string
	expected int
	success  bool
}{
	{
		"normal valid input",
		[]string{
			"00100",
			"11110",
			"10110",
			"10111",
			"10101",
			"01111",
			"00111",
			"11100",
			"10000",
			"11001",
			"00010",
			"01010",
		},
		230,
		true,
	},
	{
		"ivalid input in one of the binaries",
		[]string{
			"00100",
			"1111a",
			"10110",
			"10111",
			"10101",
			"01111",
			"00111",
			"11100",
			"10000",
			"11001",
			"00010",
			"01010",
		},
		-1,
		false,
	},
	{
		"incorrect length in one of the binaries",
		[]string{
			"00100",
			"1111a",
			"10110",
			"10111",
			"10101",
			"01111",
			"00111",
			"11100",
			"10000",
			"11001",
			"00010",
			"01010",
		},
		-1,
		false,
	},
}

func TestCalculateLifeSupportRating(t *testing.T) {
	for _, test := range calculateLifeSupportRatingTests {
		t.Run(test.name, func(t *testing.T) {
			got, err := calculateLifeSupportRating(test.input)
			assert.Equal(t, test.success, err == nil, "Expected success does not match error state")
			assert.Equal(t, test.expected, got, "Calculated Life Support Rating was unexpected")
		})
	}
}
