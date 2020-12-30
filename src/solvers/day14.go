package solvers

import (
	"fmt"
	"helpers"
	"math"
	"strconv"
	"strings"
)

var _len int64

func SolveDay14() {
	lines, err := helpers.ReadInputFile()

	if err != nil {
		fmt.Println("ReadInputFile returned error(s). Exiting.")
		fmt.Println(lines)
		return
	}

	_len = 36
	runInitializationProgram(lines)
}

func runInitializationProgram(lines []string) {
	memSpace := make(map[int64]int64)
	mask := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	for _, line := range lines {
		toks := strings.Split(line, " = ")
		if toks[0] == "mask" {
			mask = toks[1]
		} else if toks[0][:3] == "mem" {
			address := getAddress(toks[0])
			value := applyMask(toks[1], mask)
			memSpace[address] = value
		} else {
			panic(fmt.Sprintf("Error splitting line into command and value: %s\n", line))
		}
	}

	var sum int64
	sum = 0
	for _, value := range memSpace {
		sum += value
	}

	fmt.Printf("Sum of values in memory: %d\n", sum)
}

func getAddress(token string) int64 {
	if token[:3] != "mem" {
		panic(fmt.Sprintf("Unexpected input to getAddress. Was expecting \"mem[XXX]\". Got %s\n", token))
	}

	startPos := strings.Index(token, "[")
	endPos := strings.Index(token, "]")
	address, ok := strconv.ParseInt(token[startPos+1:endPos], 10, 64)
	if ok != nil {
		panic(fmt.Sprintf("Error converting address to int64: %s\n", token[startPos+1:endPos]))
	}
	return address
}

func applyMask(valTok string, mask string) int64 {
	remainder, ok := strconv.ParseInt(valTok, 10, 64)
	if ok != nil {
		panic(fmt.Sprintf("Error converting value to int64: %s\n", valTok))
	}

	var value int64
	value = 0
	var bitVal int64
	bitVal = int64(math.Exp2(float64(_len - 1)))
	for i := _len - 1; i >= 0 ; i-- {
		if i != _len - 1 {
			bitVal /= 2
		}

		bit := 0
		if remainder >= bitVal {
			bit = 1
			remainder -= bitVal
		}

		if mask[_len - 1 - i] == '0' {
			bit = 0
		} else if mask[_len - 1 - i] == '1' {
			bit = 1
		} else if mask[_len - 1 - i] != 'X' {
			panic(fmt.Sprintf("Unexpected char encountered in mask %s index %d: %v\n", mask, i, mask[i]))
		}

		if bit == 1 {
			value += bitVal
		}
	}
	
	return value
}