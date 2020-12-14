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

	countAnswers_2(lines)
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

func countAnswers_2(lines []string) {
	answerCount := 0
	currentGroup := helpers.NewSet()
	startOfGroup := true
	for _, line := range lines {
		if line == "" {
			answerCount += countGroupAnswers(currentGroup)
			currentGroup = helpers.NewSet()
			startOfGroup = true
		} else {
			incorporpatePersonAnswers(currentGroup, line, startOfGroup)
			startOfGroup = false
		}
	}

	if currentGroup.Size() != 0 {
		answerCount += countGroupAnswers(currentGroup)
	}

	fmt.Printf("Total answers counted: %d\n", answerCount)
}

func incorporpatePersonAnswers(groupAnswers *helpers.Set, line string, startOfGroup bool) {
	currentPerson := helpers.NewSet()

	for _, r := range line {
		char := string(r)
		currentPerson.Add(char)
	}

	if startOfGroup {
		groupAnswers.Append(currentPerson)
	} else {
		groupAnswers.Replace(groupAnswers.Intersection(currentPerson))
	}
}
