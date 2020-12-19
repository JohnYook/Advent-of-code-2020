package solvers

import (
	"fmt"
	"helpers"
	"sort"
)

var _XMASdata []int
var _dataLen int

func SolveDay9() {
	lines, err := helpers.ReadInputFile()

	if err != nil {
		fmt.Println("ReadInputFile returned error(s). Exiting.")
		fmt.Println(lines)
		return
	}

	_XMASdata = helpers.ConvertToInts(lines)
	_dataLen = len(_XMASdata)

	if _dataLen < 26 {
		fmt.Printf("Input too short for XMAS code. Only %d numbers.\n", _dataLen)
		return
	}

	invalidNumber := findFirstInvalidXMASCode()

	if invalidNumber != -1 {
		findEncryptionWeakness(invalidNumber)
	} else {
		fmt.Println("Unable to find invalid XMAS code. Skipping encryption weakness search.")
	}
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

func findFirstInvalidXMASCode() int {
	for i := 25; i < _dataLen; i++ {
		if !isValidXMAScode(_XMASdata[i], _XMASdata[i-25:i]) {
			fmt.Printf("Found an invalid XMAS code: %d.\n", _XMASdata[i])
			return _XMASdata[i]
		}
	}
	fmt.Println("All numbers in input seem to be valid XMAS codes.")
	return -1
}

func findEncryptionWeakness(target int) {
	for i := 0; i < _dataLen; i++ {
		if start, end, found := findContiguousSetThatSumsToTarget(target, i); found {
			fmt.Printf("Found contiguous set that sums to target %d. Start index: %d End index: %d\n", target, start, end)
			//fmt.Printf("%v\n", _XMASdata[start:end+1])
			smallest, largest := findSmallestAndLargestInSet(_XMASdata[start : end+1])
			fmt.Printf("Smallest number: %d Largest number: %d Encryption weakness: %d\n", smallest, largest, smallest+largest)
		}
	}
}

func findContiguousSetThatSumsToTarget(target int, start int) (startIndex int, endingIndex int, found bool) {
	runningSum := _XMASdata[start]
	for i := start + 1; i < _dataLen; i++ {
		runningSum += _XMASdata[i]
		if runningSum == target {
			return start, i, true
		} else if runningSum > target {
			return -1, -1, false
		}
	}
	return -1, -1, false
}

func findSmallestAndLargestInSet(nums []int) (smallest int, largest int) {
	numLen := len(nums)
	_nums := make([]int, numLen)
	copy(_nums, nums)
	sort.Ints(_nums)
	return _nums[0], _nums[numLen-1]
}
