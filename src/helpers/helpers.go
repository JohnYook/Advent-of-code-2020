package helpers

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
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
