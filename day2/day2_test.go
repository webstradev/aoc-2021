package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var commandCreationTests = []struct {
	input    string
	expected *Command
}{
	{
		"forward 1",
		&Command{Direction: "forward", Step: 1},
	},
	{
		"forward 2",
		&Command{Direction: "forward", Step: 2},
	},
	{
		"forward 5",
		&Command{Direction: "forward", Step: 5},
	},
	{
		"down 3",
		&Command{Direction: "down", Step: 3},
	},
	{
		"down 1",
		&Command{Direction: "down", Step: 1},
	},
	{
		"up 1",
		&Command{Direction: "up", Step: 1},
	},
	{
		"up 3",
		&Command{Direction: "up", Step: 3},
	},
	{
		"x 4",
		nil,
	},
	{
		"testblah",
		nil,
	},
	{
		"test test",
		nil,
	},
}

var validateDirectionTests = []struct {
	input    string
	expected bool
}{
	{
		"Forward",
		false,
	},
	{
		"forward",
		true,
	},
	{
		"down",
		true,
	},
	{
		"up",
		true,
	},
	{
		"bal",
		false,
	},
	{
		"asdf asdf",
		false,
	},
}

var calculateCommandResultTests = []struct {
	name          string
	commands      []string
	expectedPos   int
	expectedDepth int
}{
	{
		"forward 1 and down 1",
		[]string{"forward 1", "down 1"},
		1,
		1,
	},
	{
		"forward 2 and down 1",
		[]string{"forward 2", "down 1"},
		2,
		1,
	},
	{
		"forward 3 and down 1 but split",
		[]string{"forward 1", "down 1", "forward 1", "forward 1"},
		3,
		1,
	},
	{
		"forward 2 and down 3 and back up 1",
		[]string{"forward 1", "down 1", "forward 1", "down 1", "down 1", "up 1"},
		2,
		2,
	},
	{
		"forward 1 and invalid",
		[]string{"forward 1", "test"},
		1,
		0,
	},
}

func TestValidateDirectionString(t *testing.T) {
	for _, test := range validateDirectionTests {
		t.Run(test.input, func(t *testing.T) {
			got := validateDirectionString(test.input)
			assert.Equal(t, test.expected, got, "Direction Validation Returned unexpected boolean value")
		})
	}
}

func TestCreateCommandFromString(t *testing.T) {
	for _, test := range commandCreationTests {
		t.Run(test.input, func(t *testing.T) {
			got := createCommandFromString(test.input)
			assert.Equal(t, test.expected, got, "Creating Command did not return expected")
		})
	}
}

func TestCalculateCommandResult(t *testing.T) {
	for _, test := range calculateCommandResultTests {
		t.Run(test.name, func(t *testing.T) {
			gotPos, gotDepth := calculateCommandResults(test.commands)
			assert.Equal(t, test.expectedPos, gotPos, "Calculated X value did not match expected")
			assert.Equal(t, test.expectedDepth, gotDepth, "Calculated Y value did not match expected")
		})
	}
}
