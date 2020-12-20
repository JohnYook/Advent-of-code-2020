package solvers

import (
	"fmt"
	"helpers"
	"sort"
)

func SolveDay10() {
	lines, err := helpers.ReadInputFile()

	if err != nil {
		fmt.Println("ReadInputFile returned error(s). Exiting.")
		fmt.Println(lines)
		return
	}

	adapters := helpers.ConvertToInts(lines)

	getOneJoltTimesThreeJoltDiffs(adapters)
}

func getOneJoltTimesThreeJoltDiffs(adapters []int) {
	chain := make([]int, len(adapters), len(adapters)+2)
	copy(chain, adapters)
	chain = append(chain, 0)
	sort.Ints(chain)
	chain = append(chain, chain[len(chain)-1]+3)

	ones := 0
	threes := 0
	last := 0
	for _, val := range chain {
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
