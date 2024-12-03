package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type s2403 Puzzle

func (s *s2403) SetDebug(debug bool) error {
	s.Debug = debug
	return nil
}

func (s *s2403) SetInput(input []string) error {
	s.Input = input
	return nil
}

func (s *s2403) SetPart(part int) error {
	s.Part = part
	return nil
}

func (s *s2403) Solve() (string, error) {
	return strconv.Itoa(s.Process()), nil
}

func (s *s2403) Process() int {
	regex := `mul\(\d+\,\d+\)`
	if s.Part == 2 {
		regex = `mul\(\d+\,\d+\)|do\(\)|don\'t\(\)`
	}
	re := regexp.MustCompile(regex)
	sum := 0
	enabled := true
	for _, line := range s.Input {
		matches := re.FindAllString(line, -1)
		if s.Debug {
			fmt.Println(matches)
		}
		for _, match := range matches {
			if match == "do()" {
				enabled = true
			} else if match == "don't()" {
				enabled = false
			} else if enabled {
				ints := ExtractInts(match)
				sum += ints[0] * ints[1]
			}
		}
	}
	return sum
}
