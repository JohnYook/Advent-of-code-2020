package helpers

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

func ReadInputFile() ([]string, error) {
	if len(os.Args) < 2 {
		errorMessage := fmt.Sprintf("Usage: go run %s [inputFile]\n", os.Args[0])
		return nil, errors.New(errorMessage) // fmt.Errorf("Usage: go run %s [inputFile]\n", os.Args[0])
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
