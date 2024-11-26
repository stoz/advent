package main

import (
	"fmt"
	"slices"
	"strconv"
)

type s2309 Puzzle

func (s *s2309) SetDebug(debug bool) error {
	s.Debug = debug
	return nil
}

func (s *s2309) SetInput(input []string) error {
	s.Input = input
	return nil
}

func (s *s2309) SetPart(part int) error {
	s.Part = part
	return nil
}

func (s *s2309) Solve() (string, error) {
	return strconv.Itoa(s.Process()), nil
}

func (s *s2309) Process() int {
	sum := 0
	for lineCount, line := range s.Input {
		values := ExtractSInts(line)
		if s.Part == 2 {
			slices.Reverse(values)
		}
		var steps [][]int
		steps = append(steps, values)
		notAllZero := true
		for notAllZero {
			var diff []int
			for i, value := range steps[len(steps)-1] {
				if i > 0 {
					diff = append(diff, value-steps[len(steps)-1][i-1])
				}
			}
			steps = append(steps, diff)
			diffAllZero := true
			for _, d := range diff {
				if d != 0 {
					diffAllZero = false
				}
			}
			if diffAllZero {
				notAllZero = false
			}
		}
		if s.Debug && lineCount == 0 {
			for _, step := range steps {
				fmt.Println(step)
			}
		}
		calc := 0
		for i := len(steps) - 2; i > -1; i-- {
			calc = steps[i][len(steps[i])-1] + calc
		}
		sum += calc
	}
	return sum
}
