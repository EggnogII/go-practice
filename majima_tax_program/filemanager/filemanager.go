package filemanager

import (
	"bufio"
	"errors"
	"os"
)

func ReadLines(path string) ([]string, error) {
	// Open the file
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("Failed to open file.")
	}

	// Read the contents
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Return an error if there is an error during read
	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("Failed to read file.")
	}

	file.Close()
	return lines, nil
}
