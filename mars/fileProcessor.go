package mars

import (
	"bufio"
	"log"
	"os"
)

// ReadFile reads a string filename, splits the contents into a string slice and returns it
func ReadFile(filename string) ([]string, error) {

	// Open file
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("failed to open file: %v error: %v\n", filename, err)
		return nil, err
	}

	// Scan file and split it into lines
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// Add file to []string object
	var fileContent []string
	for scanner.Scan() {
		fileContent = append(fileContent, scanner.Text())
	}

	// Close the file and return the output object
	file.Close()
	return fileContent, nil
}
