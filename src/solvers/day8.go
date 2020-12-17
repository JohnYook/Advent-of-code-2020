package solvers

import (
	"fmt"
	"helpers"
	"strconv"
	"strings"
)

func SolveDay8() {
	lines, err := helpers.ReadInputFile()

	if err != nil {
		fmt.Println("ReadInputFile returned error(s). Exiting.")
		fmt.Println(lines)
		return
	}

	fmt.Printf("Accumulator value at first repeated instruction: %d\n", getAccumulatorValueAtLoop(lines))
}

func getAccumulatorValueAtLoop(program []string) int {
	programLen := len(program)

	executedInstructions := helpers.NewSet()

	current := 0
	accumulator := 0
	for {
		if executedInstructions.Contains(fmt.Sprintf("%d", current)) {
			break
		} else {
			executedInstructions.Add(fmt.Sprintf("%d", current))
		}

		instruction := program[current]
		offset, value := processInstruction(instruction)

		current += offset
		accumulator += value

		if current < 0 || current >= programLen {
			panic(fmt.Sprintf("Somehow got out of bounds instruction line: %d. Instruction set length is: %d\n", current, programLen))
		}
	}
	return accumulator
}

func processInstruction(line string) (int, int) {
	tokens := strings.Split(line, " ")
	instruction := tokens[0]
	value, ok := strconv.Atoi(tokens[1])
	if ok != nil {
		panic(fmt.Sprintf("Something went wrong coverting %s to integer.", tokens[1]))
	}
	offset := 1
	accumulator := 0
	switch instruction {
	case "nop":
		break
	case "acc":
		accumulator = value
	case "jmp":
		offset = value
	}

	return offset, accumulator
}
