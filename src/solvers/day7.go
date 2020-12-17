package solvers

import (
	"fmt"
	"helpers"
	"regexp"
	"strconv"
	"strings"
)

func SolveDay7() {
	lines, err := helpers.ReadInputFile()

	if err != nil {
		fmt.Println("ReadInputFile returned error(s). Exiting.")
		fmt.Println(lines)
		return
	}

	countBags(lines)
}

func countBags(lines []string) {
	bagRules := make(map[string]map[string]int)
	containedByLUT := make(map[string]*helpers.Set)

	parseLine := func(line string) {
		bagAndRulesPattern := "([a-z ]+) bags contain ([0-9a-z, ]+)."
		bagAndRulesRegexp := regexp.MustCompile(bagAndRulesPattern)
		numberOfBagsPattern := "([0-9]+) ([a-z ]+) bag[s]?"
		numberOfBagsRegexp := regexp.MustCompile(numberOfBagsPattern)
		noOtherBagsPattern := "no other bags"
		noOtherBagsRegexp := regexp.MustCompile(noOtherBagsPattern)

		subjectContentgroups := bagAndRulesRegexp.FindStringSubmatch(line)
		if len(subjectContentgroups) != 3 {
			panic(fmt.Sprintf("Unexpected number (%d) of matching groups for line %s.\nGroups: %v\n ", len(subjectContentgroups), line, subjectContentgroups))
		}
		subjectBag := subjectContentgroups[1]
		contents := subjectContentgroups[2]

		if _, ok := bagRules[subjectBag]; ok {
			panic(fmt.Sprintf("Encountered rule for %s more than once\n.", subjectBag))
		}

		bagRules[subjectBag] = make(map[string]int)

		if noOtherBagsRegexp.MatchString(contents) {
			return
		}

		currentRules := bagRules[subjectBag] // this is supposed to be a reference
		containedBags := strings.Split(contents, ",")
		for _, bagItem := range containedBags {
			contentBagGroups := numberOfBagsRegexp.FindStringSubmatch(bagItem)
			if len(contentBagGroups) != 3 {
				panic(fmt.Sprintf("Unexpected number (%d) of matching groups for content bag group %s.\nGroups: %v\n ", len(contentBagGroups), line, contentBagGroups))
			}
			containedBagNumber, _ := strconv.Atoi(contentBagGroups[1])
			containedBag := contentBagGroups[2]
			if _, ok := currentRules[containedBag]; ok {
				panic(fmt.Sprintf("Encountered contained bag %s more than once\n.", containedBag))
			}
			currentRules[containedBag] = containedBagNumber
			if _, ok := containedByLUT[containedBag]; ok {
				if !containedByLUT[containedBag].Contains(subjectBag) {
					containedByLUT[containedBag].Add(subjectBag)
				} else {
					panic(fmt.Sprintf("Adding bag %s to containing bag %s again.\n.", subjectBag, containedBag))
				}
			} else {
				containedByLUT[containedBag] = helpers.NewSet()
				containedByLUT[containedBag].Add(subjectBag)
			}
		}

		return
	}

	countNumberOfBagsThatContain := func(bag string) int {
		listOfBags := helpers.NewSet()

		addBagsThatContainToList(listOfBags, containedByLUT, bag)

		return listOfBags.Size()
	}

	countNumberOfBagsContainedIn := func(bag string) int {
		countOfBags := countBagsContainedIn(bagRules, bag)

		return countOfBags
	}

	for _, line := range lines {
		parseLine(line)
	}

	fmt.Printf("Number of bags that can contain a shiny gold bag: %d\n", countNumberOfBagsThatContain("shiny gold"))
	fmt.Printf("Number of bags contained in a shiny gold bag: %d\n", countNumberOfBagsContainedIn("shiny gold"))
}

func addBagsThatContainToList(listOfBags *helpers.Set, lookup map[string]*helpers.Set, bagColor string) {
	if bagsThatContain, ok := lookup[bagColor]; ok {
		for _, key := range bagsThatContain.Items() {
			if !listOfBags.Contains(key) {
				listOfBags.Add(key)
			}
			addBagsThatContainToList(listOfBags, lookup, key)
		}
	}
}

func countBagsContainedIn(lookup map[string]map[string]int, bagColor string) int {
	nestedBagsCount := 0
	if containedBags, ok := lookup[bagColor]; ok && len(containedBags) > 0 {
		for key, value := range containedBags {
			// e.g. if [brown: 3], then it's 3 brown bags + 3 * what each brown bag contains
			nestedBagsCount += value + value*countBagsContainedIn(lookup, key)
		}
	}

	return nestedBagsCount
}
