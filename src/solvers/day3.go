package solvers

import (
	"fmt"
	"helpers"
	"strings"
)

func SolveDay3() {
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

	//solve_1(terrainMap, count, lineLen)
	solve_2(terrainMap, count, lineLen)
}

// go 3 right down 1
func solve_1(terrainMap [][]string, rowCount int, lineLen int) {
	solution1 := countTrees(terrainMap, rowCount, lineLen, 1, 3)
	fmt.Printf("Tree count part 1: %d\n", solution1)
}

func solve_2(terrainMap [][]string, rowCount int, lineLen int) {
	// Right 1, down 1.
	// Right 3, down 1.
	// Right 5, down 1.
	// Right 7, down 1.
	// Right 1, down 2.
	slopes := [5][2]int{{1, 1}, {1, 3}, {1, 5}, {1, 7}, {2, 1}} //{y, x}

	treeCounts := []int{}
	for k := 0; k < len(slopes); k++ {
		slope := slopes[k]
		fmt.Printf("Slope y=%d x=%d\n", slope[0], slope[1])
		treeCounts = append(treeCounts, countTrees(terrainMap, rowCount, lineLen, slope[0], slope[1]))
	}

	fmt.Printf("TreeCounts: %v\n", treeCounts)

	total := treeCounts[0]

	for i := 1; i < len(treeCounts); i++ {
		total *= treeCounts[i]
	}
	fmt.Printf("TreeCount multiplied together: %d\n", total) //4355551200
}

func countTrees(terrainMap [][]string, rowCount int, lineLen int, yIncrement int, xIncrement int) int {
	treeCount := 0
	j := 0
	for i := 0; i < rowCount; i += yIncrement {
		if j >= lineLen {
			j = j % lineLen
		}
		if terrainMap[i][j] == "#" {
			treeCount++
		} else if terrainMap[i][j] != "." {
			fmt.Printf("Something is wrong. Neither tree nor open space found: %s\n", terrainMap[i][j])
			return -1
		}

		j += xIncrement
	}

	fmt.Printf("Tree count: %d\n", treeCount)
	return treeCount
}
