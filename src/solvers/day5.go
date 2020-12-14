package solvers

import (
	"fmt"
	"helpers"
	"regexp"
	"sort"
	"strings"
)

func SolveDay5() {
	lines, err := helpers.ReadInputFile()

	if err != nil {
		fmt.Println("ReadInputFile returned error(s). Exiting.")
		fmt.Println(lines)
		return
	}

	getSeatId_1(lines)
}

func getSeatId_1(lines []string) {

	var seatIds []int
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if validateLine(line) == false {
			fmt.Printf("Error. Line fails validation: %s\n", line)
			continue
		}

		seatId := getSeatId(line)
		seatIds = append(seatIds, seatId)
	}

	sort.Ints(seatIds)
	fmt.Printf("Highest seatId is: %d\n", seatIds[len(seatIds)-1])
}

func getSeatId(line string) int {
	rowCode := line[:7]
	seatCode := line[7:]

	row := computeRow(rowCode)
	seat := computeSeat(seatCode)
	return (row * 8) + seat
}

func validateLine(line string) bool {
	// FFBFFFBLLL
	pattern := "[FB]{7}[LR]{3}"
	reg, _ := regexp.Compile(pattern)
	return reg.MatchString(line)
}

func computeRow(rowCode string) int {
	rowChars := strings.Split(rowCode, "")
	var rowBits []int
	for i := len(rowCode) - 1; i >= 0; i-- {
		switch rowChars[i] {
		case "F":
			rowBits = append(rowBits, 0)
		case "B":
			rowBits = append(rowBits, 1)
		}
	}

	row := 0
	multiplier := 1
	for i := 0; i < len(rowBits); i++ {
		row += rowBits[i] * multiplier
		multiplier *= 2
	}

	return row
}

func computeSeat(seatCode string) int {
	seatChars := strings.Split(seatCode, "") // L R
	var seatBits []int
	for i := len(seatCode) - 1; i >= 0; i-- {
		switch seatChars[i] {
		case "L":
			seatBits = append(seatBits, 0)
		case "R":
			seatBits = append(seatBits, 1)
		}
	}

	seat := 0
	multiplier := 1
	for i := 0; i < len(seatBits); i++ {
		seat += seatBits[i] * multiplier
		multiplier *= 2
	}

	return seat
}
