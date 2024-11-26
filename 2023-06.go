package main

import (
	"fmt"
	"strconv"
	"strings"
)

type s2306 Puzzle

func (s *s2306) SetDebug(debug bool) error {
	s.Debug = debug
	return nil
}

func (s *s2306) SetInput(input []string) error {
	s.Input = input
	return nil
}

func (s *s2306) SetPart(part int) error {
	s.Part = part
	return nil
}

func (s *s2306) Solve() (string, error) {
	return strconv.Itoa(s.Process()), nil
}

func (s *s2306) Process() int {
	if s.Part == 2 {
		s.Input[0] = strings.ReplaceAll(s.Input[0], " ", "")
		s.Input[1] = strings.ReplaceAll(s.Input[1], " ", "")
	}
	times := ExtractInts(s.Input[0])
	distances := ExtractInts(s.Input[1])
	margin := 1
	for i := 0; i < len(times); i++ {
		if s.Debug {
			fmt.Println(times[i], distances[i])
		}
		wins := 0
		for j := 1; j < times[i]; j++ {
			if s.Debug {
				fmt.Println("Speed:", j, "Distance: ", j*(times[i]-j))
			}
			if j*(times[i]-j) > distances[i] {
				wins++
			}
		}
		if s.Debug {
			fmt.Println("Wins:", wins)
		}
		margin = margin * wins
	}
	return margin
}
