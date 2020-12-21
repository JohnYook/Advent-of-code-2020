package solvers

import (
	"fmt"
	"helpers"
	"sort"
)

var _chain []int
var _chainLen int
var _numPaths []int

func SolveDay10() {
	lines, err := helpers.ReadInputFile()

	if err != nil {
		fmt.Println("ReadInputFile returned error(s). Exiting.")
		fmt.Println(lines)
		return
	}

	adapters := helpers.ConvertToInts(lines)

	_chain = getOutletToDeviceChain(adapters)
	_chainLen = len(_chain)

	getOneJoltTimesThreeJoltDiffs()
	fmt.Printf("Number of ways to connect: %d\n", countWaysToConnect())
}

func getOutletToDeviceChain(adapters []int) []int {
	chain := make([]int, len(adapters), len(adapters)+2)
	copy(chain, adapters)
	chain = append(chain, 0)
	sort.Ints(chain)
	return append(chain, chain[len(chain)-1]+3)
}

func getOneJoltTimesThreeJoltDiffs() {
	ones := 0
	threes := 0
	last := 0
	for _, val := range _chain {
		if val == 0 {
			continue
		}
		if val == last+1 {
			ones += 1
		} else if val == last+3 {
			threes += 1
		}
		last = val
	}

	fmt.Printf("Jumps of one: %d Jumps of three: %d Product: %d\n", ones, threes, ones*threes)
}

func countWaysToConnect() int {
	_numPaths = make([]int, _chainLen)
	_numPaths[0] = 1

	for i := 1; i < _chainLen; i++ {
		_numPaths[i] = getPossiblePathsHere(i)
	}

	return _numPaths[_chainLen-1]
}

func getPossiblePathsHere(currentIndex int) int {
	// for i, check i - 1, i - 2, i -3
	start := currentIndex - 3
	if start < 0 {
		start = 0
	}
	possiblePaths := 0
	for i := start; i < currentIndex; i++ {
		if _chain[currentIndex]-_chain[i] <= 3 {
			possiblePaths += _numPaths[i]
		}
	}

	return possiblePaths
}
