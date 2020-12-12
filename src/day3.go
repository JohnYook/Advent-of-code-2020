package main

import (
	"bufio"
	"fmt"
	"helpers"
)

func main() {
    file := helpers.GetInputFile()

	fmt.Printf("Read in inputfile: %s\n", file)

	count := 0
	scanner := bufio.NewScanner(file)
    for scanner.Scan() {

	}
}
