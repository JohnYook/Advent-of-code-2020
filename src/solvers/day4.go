package solvers

import (
	"fmt"
	"helpers"
	"regexp"
	"strconv"
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

	//countValidPassports_1(lines)
	countValidPassports_2(lines)
}

func countValidPassports_1(lines []string) {
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

func countValidPassports_2(lines []string) {
	count := len(lines)

	validCount := 0
	currentPassport := make(map[string]string)
	for i := 0; i < count; i++ {
		line := lines[i]
		if line == "" {
			// blank line encountered. Validate currentPassport and start new passport
			validCount += countIncrement(isValidPassport2(currentPassport))
			currentPassport = make(map[string]string)
		} else {
			readIntoPassport(currentPassport, line)
		}
	}

	// check if last currentPassport has been processed
	if len(currentPassport) != 0 {
		validCount += countIncrement(isValidPassport2(currentPassport))
	}

	fmt.Printf("Found %d valid passports.", validCount)
}

func isValidPassport2(passport map[string]string) bool {
	// Required fields: [byr iyr eyr hgt hcl ecl pid].  Optional fields: [cid]
	requiredKeys := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, key := range requiredKeys {
		value, ok := passport[key]
		if ok == false {
			return false
		}

		switch key {
		case "byr":
			if validateBirthYear(value) == false {
				return false
			}
		case "iyr":
			if validateIssueYear(value) == false {
				return false
			}
		case "eyr":
			if validateExpirationYear(value) == false {
				return false
			}
		case "hgt":
			if validateHeight(value) == false {
				return false
			}
		case "hcl":
			if validateHairColor(value) == false {
				return false
			}
		case "ecl":
			if validateEyeColor(value) == false {
				return false
			}
		case "pid":
			if validatePassportId(value) == false {
				return false
			}
		}
	}
	return true
}

func validateBirthYear(yearString string) bool {
	return validateYear(yearString, 1920, 2002)
}

func validateIssueYear(yearString string) bool {
	return validateYear(yearString, 2010, 2020)
}

func validateExpirationYear(yearString string) bool {
	return validateYear(yearString, 2020, 2030)
}

func validateYear(yearString string, min int, max int) bool {
	yearVal, _ := strconv.Atoi(yearString)

	return yearVal >= min && yearVal <= max
}

func validateHeight(heightString string) bool {
	units := heightString[len(heightString)-2:]
	heightVal, _ := strconv.Atoi(heightString[:len(heightString)-2])

	if units == "cm" {
		return heightVal >= 150 && heightVal <= 193
	} else if units == "in" {
		return heightVal >= 59 && heightVal <= 76
	} else {
		return false
	}
}

func validateHairColor(hairColor string) bool {
	pattern := "#[0-9a-f]{6}"
	reg, _ := regexp.Compile(pattern)
	return reg.MatchString(hairColor)
}

func validateEyeColor(eyeColor string) bool {
	switch eyeColor {
	case
		"amb",
		"blu",
		"brn",
		"gry",
		"grn",
		"hzl",
		"oth":
		return true
	default:
		return false
	}
}

func validatePassportId(passportId string) bool {
	digits := len(passportId)
	if digits != 9 {
		return false
	}
	pattern := "[0-9]{9}"
	reg, _ := regexp.Compile(pattern)
	return reg.MatchString(passportId)
}
