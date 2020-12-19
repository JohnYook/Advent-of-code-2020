package solvers

import (
	"fmt"
	"helpers"
)

func SolveDay9() {
	lines, err := helpers.ReadInputFile()

	if err != nil {
		fmt.Println("ReadInputFile returned error(s). Exiting.")
		fmt.Println(lines)
		return
	}

	nums := helpers.ConvertToInts(lines)

	findFirstInvalidXMASCode(nums)
}

func isValidXMAScode(num int, preceding []int) bool {
	if len(preceding) < 25 {
		fmt.Printf("%d is INVALID code. Insufficient preamble length: %d\n", num, len(preceding))
		return false
	}
	preamble := preceding[len(preceding)-25:]

	numTable := make(map[int]int)
	for _, val := range preamble {
		if _, ok := numTable[val]; ok {
			numTable[val] += 1
		} else {
			numTable[val] = 1
		}
	}

	for _, firstAddend := range preamble {
		target := num - firstAddend

		if val, ok := numTable[target]; ok {
			if target == firstAddend { // have to check the number occurs twice
				if val > 1 {
					return true
				}
			} else { // enough if target is present once
				return true
			}
		}
	}
	return false
}

func findFirstInvalidXMASCode(nums []int) {
	inputLen := len(nums)
	if inputLen < 26 {
		fmt.Printf("Input too short for XMAS code. Only %d numbers.\n", inputLen)
		return
	}

	for i := 25; i < inputLen; i++ {
		if !isValidXMAScode(nums[i], nums[i-25:i]) {
			fmt.Printf("Found an invalid XMAS code: %d.\n", nums[i])
			return
		}
	}
	fmt.Println("All numbers in input seem to be valid XMAS codes.")
}
