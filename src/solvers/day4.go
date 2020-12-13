package solvers

import (
	"fmt"
	"helpers"
	"strings"
)

type Passport struct {
	birthYear, issueYear, expirationYear, height, passportId, countryId int
	hairColor, eyeColor                                                 string
}

func SolveDay4() {
	lines, err := helpers.ReadInputFile()

	if err != nil {
		fmt.Println("ReadInputFile returned error(s). Exiting.")
		return
	}

	part_1(lines)
}

func part_1(lines []string) {
	count := len(lines)

	validCount := 0
	currentPassport := make(map[string]string)
	for i := 0; i < count; i++ {
		line := lines[i]
		if line == "" {
			// blank line encountered. Validate currentPassport and start new passport
			validCount += countIncrement(isValidPassport(currentPassport))
			currentPassport = make(map[string]string)
		} else {
			readIntoPassport(currentPassport, line)
		}
	}

	// check if last currentPassport has been processed
	if len(currentPassport) != 0 {
		validCount += countIncrement(isValidPassport(currentPassport))
	}

	fmt.Printf("Found %d valid passports.", validCount)
}

func countIncrement(meetsCondition bool) int {
	if meetsCondition {
		return 1
	} else {
		return 0
	}
}

func isValidPassport(passport map[string]string) bool {
	// Required fields: [byr iyr eyr hgt hcl ecl pid].  Optional fields: [cid]
	requiredKeys := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, key := range requiredKeys {
		_, ok := passport[key]
		if ok == false {
			return false
		}
	}
	return true
}

func readIntoPassport(passport map[string]string, line string) {
	// tokenize line into list of "key:value"
	kvPairs := strings.Split(line, " ")
	for _, token := range kvPairs {
		keyValue := strings.Split(token, ":")
		key := keyValue[0]
		value := keyValue[1]
		if key == "" || value == "" {
			fmt.Printf("Unexpected error. Blank key or value for token %s, line %s\n Key-Value pairs: ", token, line)
			fmt.Println(kvPairs)
		}
		passport[key] = value
	}
}
