package main

import (
    "bufio"
    "fmt"
    "log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
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
		min, err := strconv.Atoi(groups[1])
		if err != nil {
			log.Fatal(err)
		}
		max, err := strconv.Atoi(groups[2])
		if err != nil {
			log.Fatal(err)
		}
		char := groups[3]
		passwd := groups[4]

		if IsValidPassword(passwd, char, min, max) == true {
			count += 1
		}
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
	}

	fmt.Printf("Number of valid passwords: %d\n", count)
}

func IsValidPassword(passwd string, char string, min int, max int) bool {
	occurs := strings.Count(passwd, char)
	return (occurs >= min && occurs <= max)
}
