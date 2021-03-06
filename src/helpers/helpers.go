package helpers

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func ReadInputFile() ([]string, error) {
	if len(os.Args) < 2 {
		errorMessage := fmt.Sprintf("Usage: go run %s [inputFile]\n", os.Args[0])
		return nil, errors.New(errorMessage) // fmt.Errorf("Usage: go run %s [inputFile]\n", os.Args[0])
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func ConvertToInts(lines []string) []int {
	var nums []int
	for _, line := range lines {
		intVal, err := strconv.Atoi(line)
		if err != nil {
			panic(fmt.Sprintf("ConvertToInts encountered a number it could not convert: %s\n", line))
		}
		nums = append(nums, intVal)
	}
	return nums
}

type Set struct {
	m map[string]struct{}
}

var exists = struct{}{}

func NewSet() *Set {
	s := &Set{}
	s.m = make(map[string]struct{})
	return s
}

func (s *Set) Add(key string) {
	s.m[key] = exists
}

func (s *Set) Remove(key string) {
	delete(s.m, key)
}

func (s *Set) Contains(key string) bool {
	_, ok := s.m[key]
	return ok
}

func (s *Set) Size() int {
	return len(s.m)
}

func (s *Set) Append(o *Set) {
	for key, _ := range o.m {
		s.Add(key)
	}
}

func (s *Set) Replace(o *Set) {
	for key, _ := range s.m {
		delete(s.m, key)
	}

	for key, _ := range o.m {
		s.Add(key)
	}
}

func (s *Set) Copy(o *Set) {
	for key, _ := range o.m {
		s.Add(key)
	}
}

func (s *Set) Items() []string {
	keys := make([]string, 0, len(s.m))

	for k, _ := range s.m {
		keys = append(keys, k)
	}
	return keys
}

func (s *Set) Intersection(o *Set) *Set {
	i := NewSet()
	for key, _ := range s.m {
		_, ok := o.m[key]
		if ok == true {
			i.Add(key)
		}
	}

	return i
}

func Abs(num int) int {
	if num  < 0 {
		num = -num
	}
	return num
}

func Max(nums []int) int {
	if len(nums) <= 0{
		panic(fmt.Sprintf("Invalid input to Max. Must be non-empty slice of ints.\n"))
	}
	max := nums[0]

	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}