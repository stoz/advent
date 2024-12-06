package main

import (
	"fmt"
	"strconv"
)

type s2405 Puzzle

func (s *s2405) SetDebug(debug bool) error {
	s.Debug = debug
	return nil
}

func (s *s2405) SetInput(input []string) error {
	s.Input = input
	return nil
}

func (s *s2405) SetPart(part int) error {
	s.Part = part
	return nil
}

func (s *s2405) Solve() (string, error) {
	var number int
	if s.Part != 2 {
		number = s.processPart1()
	} else {
		number = s.processPart2()
	}
	return strconv.Itoa(number), nil
}

func (s *s2405) processPart1() int {
	sum := 0
	var rules, updates [][]uint8
	firstPart := true
	for _, line := range s.Input {
		if line == "" {
			firstPart = false
			continue
		}
		if firstPart {
			rules = append(rules, ExtractUint8s(line))
		} else {
			updates = append(updates, ExtractUint8s(line))
		}
	}
	for _, update := range updates {
		good := s.validate(update, rules)
		if good {
			sum += int(update[(len(update)-1)/2])
		}
	}
	return sum
}

func (s *s2405) processPart2() int {
	sum := 0
	var rules, updates [][]uint8
	firstPart := true
	for _, line := range s.Input {
		if line == "" {
			firstPart = false
			continue
		}
		if firstPart {
			rules = append(rules, ExtractUint8s(line))
		} else {
			updates = append(updates, ExtractUint8s(line))
		}
	}
	for updateIndex, update := range updates {
		good := s.validate(update, rules)
		if !good {
			for !good {
				for _, rule := range rules {
					firstPos := -1
					secondPos := -1
					for i, u := range update {
						if u == rule[0] {
							firstPos = i
						}
						if u == rule[1] {
							secondPos = i
						}
					}
					if firstPos > -1 && secondPos > -1 && firstPos > secondPos {
						// swap the positions of any pair that doesn't match a rule
						update[firstPos], update[secondPos] = update[secondPos], update[firstPos]
					}
				}
				good = s.validate(update, rules)
			}
			sum += int(update[(len(update)-1)/2])
			if s.Debug {
				fmt.Println("update: ", updateIndex)
			}
		}
	}
	return sum
}

func (s *s2405) validate(update []uint8, rules [][]uint8) bool {
	good := true
	for _, rule := range rules {
		firstPos := -1
		secondPos := -1
		for i, u := range update {
			if u == rule[0] {
				firstPos = i
			}
			if u == rule[1] {
				secondPos = i
			}
		}
		if firstPos > -1 && secondPos > -1 && firstPos > secondPos {
			good = false
		}
	}
	return good
}
