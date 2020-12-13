package solvers

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func SolveDay2() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run day2.go [inputFile]")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	count := 0
	scanner := bufio.NewScanner(file)
	delim := "([0-9]+)-([0-9]+)[ ]*([A-Za-z]):[ ]*(.*)"
	re := regexp.MustCompile(delim)
	for scanner.Scan() {
		line := scanner.Text()
		groups := re.FindStringSubmatch(line)
		if len(groups) != 5 {
			fmt.Printf("Unexpected number of groups in regex split. Was expecting 4. Got %d. For line=$s\n", len(groups), line)
			fmt.Printf("The groups matched were:\n")
			for i := range groups {
				fmt.Printf("%d %s\n", i, groups[i])
			}
			log.Fatal("Unexpected number of groups in regex split.")
		}
		firstNum, err := strconv.Atoi(groups[1])
		if err != nil {
			log.Fatal(err)
		}
		secondNum, err := strconv.Atoi(groups[2])
		if err != nil {
			log.Fatal(err)
		}
		char := groups[3]
		passwd := groups[4]

		if IsValidPassword_2(passwd, char, firstNum, secondNum) == true {
			count += 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Number of valid passwords: %d\n", count)
}

func IsValidPassword_1(passwd string, char string, min int, max int) bool {
	occurs := strings.Count(passwd, char)
	return (occurs >= min && occurs <= max)
}

func IsValidPassword_2(passwd string, char string, index1 int, index2 int) bool {
	if index1 < 1 || index1 > len(passwd) || index2 < 0 || index2 > len(passwd) {
		log.Fatal("Invalid index values. Password %s has length %d, but specified indices (1-based) were %d and %d\n", passwd, len(passwd), index1, index2)
	}

	return (passwd[index1-1] == char[0]) != (passwd[index2-1] == char[0])
}
