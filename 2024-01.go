package main

import (
	"sort"
	"strconv"
)

type s2401 Puzzle

func (s *s2401) SetDebug(debug bool) error {
	s.Debug = debug
	return nil
}

func (s *s2401) SetInput(input []string) error {
	s.Input = input
	return nil
}

func (s *s2401) SetPart(part int) error {
	s.Part = part
	return nil
}

func (s *s2401) Solve() (string, error) {
	return strconv.Itoa(s.Process()), nil
}

func (s *s2401) Process() int {
	var left, right []int
	for _, line := range s.Input {
		ints := ExtractInts(line)
		left = append(left, ints[0])
		right = append(right, ints[1])
	}
	sum := 0
	if s.Part != 2 {
		sort.Ints(left)
		sort.Ints(right)
		for i, l := range left {
			sum += abs(right[i] - l)
		}
	} else {
		for _, l := range left {
			for _, r := range right {
				if l == r {
					sum += l
				}
			}
		}
	}
	return sum
}
