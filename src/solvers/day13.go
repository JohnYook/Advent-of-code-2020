package solvers

import (
	"fmt"
	"helpers"
	"sort"
	"strconv"
	"strings"
)

func SolveDay13() {
	lines, err := helpers.ReadInputFile()

	if err != nil {
		fmt.Println("ReadInputFile returned error(s). Exiting.")
		fmt.Println(lines)
		return
	}

	findEarliestBus(lines)
}

func findEarliestBus(lines []string) {
	earliestTime, ok := strconv.Atoi(lines[0])
	if ok != nil {
		panic(fmt.Sprintf("Error during processing line: %s\n", lines[0]))
	}

	busIds := getBusIds(lines[1])

	smallestWait := earliestTime
	bestBus := 0
	for _, bus := range busIds {
		floor := earliestTime / bus	// int division.
		waitTime := (floor + 1) * bus - earliestTime
		if waitTime < smallestWait {
			smallestWait = waitTime
			bestBus = bus
		}
	}

	fmt.Printf("Solution part 1: %d\n", bestBus * smallestWait)
}

func getBusIds(line string) []int {
	idToks := strings.Split(line, ",")

	var busIds []int
	for _, idTok := range idToks {
		if idTok == "x" {
			continue
		}
		busId, ok := strconv.Atoi(idTok)
		if ok != nil {
			panic(fmt.Sprintf("Error converting bus lin %s to int.\n", idTok))
		}
		busIds = append(busIds, busId)
	}

	sort.Ints(busIds)
	return busIds
}
