package util

import (
	"bufio"
	"os"
	"strconv"
)

// Read a whole file into the memory and store it as array of lines
// @source https://www.geeksforgeeks.org/how-to-read-a-file-line-by-line-to-string-in-golang/
func ReadLinesToSlice(path string) ([]string, error) {
	// os.Open() opens specific file in
	// read-only mode and this return
	// a pointer of type os.
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	// The bufio.NewScanner() function is called in which the
	// object os.File passed as its parameter and this returns a
	// object bufio.Scanner which is further used on the
	// bufio.Scanner.Split() method.
	scanner := bufio.NewScanner(file)

	// The bufio.ScanLines is used as an
	// input to the method bufio.Scanner.Split()
	// and then the scanning forwards to each
	// new line using the bufio.Scanner.Scan()
	// method.
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// The method os.File.Close() is called
	// on the os.File object to close the file
	file.Close()

	return lines, nil
}

// Convert an slice of strings to a slice of ints
func ConvertListOfStringsToInts(list []string) ([]int, error) {
	result := []int{}

	for _, string := range list {
		number, err := strconv.Atoi(string)
		if err != nil {
			return nil, err
		}

		result = append(result, number)
	}

	return result, nil
}
