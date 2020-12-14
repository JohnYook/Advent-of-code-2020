package solvers

import (
	"fmt"
	"helpers"
)

func SolveDay6() {
	lines, err := helpers.ReadInputFile()

	if err != nil {
		fmt.Println("ReadInputFile returned error(s). Exiting.")
		fmt.Println(lines)
		return
	}

	countAnswers_1(lines)
}

func countAnswers_1(lines []string) {
	answerCount := 0
	currentGroup := helpers.NewSet()
	for _, line := range lines {
		if line == "" {
			answerCount += countGroupAnswers(currentGroup)
			currentGroup = helpers.NewSet()
		} else {
			addPersonAnswers(currentGroup, line)
		}
	}

	if currentGroup.Size() != 0 {
		answerCount += countGroupAnswers(currentGroup)
	}

	fmt.Printf("Total answers counted: %d\n", answerCount)
}

func countGroupAnswers(groupAnswers *helpers.Set) int {
	return groupAnswers.Size()
}

func addPersonAnswers(groupAnswers *helpers.Set, line string) {
	for _, r := range line {
		char := string(r)
		if groupAnswers.Contains(char) == false {
			groupAnswers.Add(char)
		}
	}
}
