package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	// Open the file
	file, err := os.Open(fm.InputFilePath)
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

func (fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("Failed to create file.")
	}

	time.Sleep(3 * time.Second) //Slow File Write process
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		return errors.New("Failed to convert data to JSON")
	}

	file.Close()
	return nil
}

func New(inputPath string, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
