package day3

import "strings"

var binaryDigits = "01"

func validateBinaryInput(input string, digits int) bool {
	// Check if string is of correct length.
	if len(input) != digits {
		return false
	}

	// Check if each character is either a 0 or a 1
	for _, char := range input {
		if !strings.ContainsRune(binaryDigits, char) {
			return false
		}
	}

	return true
}
