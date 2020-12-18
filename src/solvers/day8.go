package solvers

import (
	"fmt"
	"helpers"
	"strconv"
	"strings"
)

var program []string
var programLen int

func SolveDay8() {
	lines, err := helpers.ReadInputFile()

	if err != nil {
		fmt.Println("ReadInputFile returned error(s). Exiting.")
		fmt.Println(lines)
		return
	}

	program = lines
	programLen = len(program)

	fmt.Printf("Accumulator value at first repeated instruction: %d\n", getAccumulatorValueAtLoop())
	getAccumulatorValueAtProgramEnd()
}

func getAccumulatorValueAtLoop() int {
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

func getAccumulatorValueAtProgramEnd() {
	executedInstructions := helpers.NewSet()

	accumulator, finished := executeProgramPaths(executedInstructions, 0, 0, false)

	if finished {
		fmt.Printf("Program successfully ran to completion. Accumulator value at end: %d\n", accumulator)
	} else {
		fmt.Printf("Program did not run to completion. All paths encountered loops.")
	}
}

func executeProgramPaths(pastInstructions *helpers.Set, next int, startingValue int, changed bool) (int, bool) {
	current := next
	accumulator := startingValue
	for {
		if pastInstructions.Contains(fmt.Sprintf("%d", current)) {
			break
		} else {
			pastInstructions.Add(fmt.Sprintf("%d", current))
		}

		instruction := program[current]
		if changed == false && isJumpOrNoop(instruction) {
			alternatePastInstructions := helpers.NewSet()
			alternatePastInstructions.Copy(pastInstructions)
			alternateInstruction := getAlternateInstruction(instruction)
			alternateOffset, alternateValue := processInstruction(alternateInstruction)
			alternateCurrent := current + alternateOffset
			alternateAccumulator := accumulator + alternateValue
			alternateAccumulatorValue, didAlternateFinish := executeProgramPaths(alternatePastInstructions, alternateCurrent, alternateAccumulator, true)
			if didAlternateFinish == true {
				return alternateAccumulatorValue, true
			}
			// Otherwise continue executing w original instructions and changed = false. So we can try branching off at next fork, etc
		}
		offset, value := processInstruction(instruction)
		current += offset
		accumulator += value

		if current == programLen {
			// program finished execution successfully
			return accumulator, true
		} else if current < 0 || current > programLen {
			panic(fmt.Sprintf("Somehow got out of bounds instruction line: %d. Instruction set length is: %d\n", current, programLen))
		}
	}
	return accumulator, false
}

func isJumpOrNoop(instruction string) bool {
	return instruction[:3] == "nop" || instruction[:3] == "jmp"
}

func getAlternateInstruction(instruction string) string {
	if instruction[:3] == "nop" {
		return strings.Replace(instruction, "nop", "jmp", 1)
	} else if instruction[:3] == "jmp" {
		return strings.Replace(instruction, "jmp", "nop", 1)
	}
	return instruction
}
