package main

import (
	"fmt"
	"helpers"
	"strings"
)

func main() {
	lines, err := helpers.ReadInputFile()

	if err != nil {
		fmt.Println("Looks like ReadInputFile failed and returned an error.")
		return
	}

	count := len(lines)
	fmt.Printf("Total lines: %d\n", count)

	if count <= 0 {
		fmt.Println("Empty input file. No lines were read in.")
		return
	}

	lineLen := len(lines[0])
	fmt.Printf("Line length: %d\n", lineLen)

	terrainMap := make([][]string, count)
	for i := 0; i < count; i++ {
		line := lines[i]
		if len(line) != lineLen {
			fmt.Printf("Error. Lines of different length detected. First line was len=%d, this line is len=%d: %s\n", lineLen, len(line), line)
		}
		lineChars := strings.Split(line, "")
		terrainMap[i] = make([]string, lineLen)
		for j, c := range lineChars {
			terrainMap[i][j] = c
		}
	}

	// go 3 right down 1
	treeCount := 0
	j := 0
	for i := 0; i < count; i++ {
		if j >= lineLen {
			j = j % lineLen
		}
		if terrainMap[i][j] == "#" {
			treeCount++
		} else if terrainMap[i][j] != "." {
			fmt.Printf("Something is wrong. Neither tree nor open space found: %s\n", terrainMap[i][j])
		}

		j += 3
	}

	fmt.Printf("Tree count: %d\n", treeCount)
}
