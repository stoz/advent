package main

import (
	"strconv"
	"strings"
)

type s2302 Puzzle

func (s *s2302) SetDebug(debug bool) error {
	s.Debug = debug
	return nil
}

func (s *s2302) SetInput(input []string) error {
	s.Input = input
	return nil
}

func (s *s2302) SetPart(part int) error {
	s.Part = part
	return nil
}

func (s *s2302) Solve() (string, error) {
	var number int
	if s.Part != 2 {
		number = s.processPart1()
	} else {
		number = s.processPart2()
	}
	return strconv.Itoa(number), nil
}

func (s *s2302) processPart1() int {
	score := 0
	for lineNumber, line := range s.Input {
		pass := true
		words := strings.Fields(line)
		for i, word := range words {
			if n, err := strconv.Atoi(word); err == nil {
				switch words[i+1][0] {
				case 'r':
					if n > 12 {
						pass = false
					}
				case 'g':
					if n > 13 {
						pass = false
					}
				case 'b':
					if n > 14 {
						pass = false
					}
				}
			}
		}
		if pass {
			score += lineNumber + 1
		}
	}
	return score
}

func (s *s2302) processPart2() int {
	score := 0
	for _, line := range s.Input {
		var r, g, b int
		words := strings.Fields(line)
		for i, word := range words {
			if n, err := strconv.Atoi(word); err == nil {
				switch words[i+1][0] {
				case 'r':
					if n > r {
						r = n
					}
				case 'g':
					if n > g {
						g = n
					}
				case 'b':
					if n > b {
						b = n
					}
				}
			}
		}
		score += r * g * b
	}
	return score
}
