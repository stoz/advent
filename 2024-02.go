package main

import (
	"fmt"
	"strconv"
)

type s2402 Puzzle

func (s *s2402) SetDebug(debug bool) error {
	s.Debug = debug
	return nil
}

func (s *s2402) SetInput(input []string) error {
	s.Input = input
	return nil
}

func (s *s2402) SetPart(part int) error {
	s.Part = part
	return nil
}

func (s *s2402) Solve() (string, error) {
	var number int
	if s.Part != 2 {
		number = s.processPart1()
	} else {
		number = s.processPart2()
	}
	return strconv.Itoa(number), nil
}

func (s *s2402) processPart1() int {
	sum := 0
	for _, report := range s.Input {
		levels := ExtractInts(report)
		valid := s.validateLevels(levels)
		if valid {
			sum++
		}
		fmt.Println(valid, levels)
	}
	return sum
}

func (s *s2402) processPart2() int {
	sum := 0
	for _, report := range s.Input {
		levels := ExtractInts(report)
		// first check the full report
		valid := s.validateLevels(levels)
		if !valid {
			// iterate through each each level in the report, and make a new report
			// with that level removed
			for i := range levels {
				var newLevels []int
				for j, level := range levels {
					if i != j {
						newLevels = append(newLevels, level)
					}
				}
				newValid := s.validateLevels(newLevels)
				if newValid {
					valid = true
					break
				}
			}
		}
		if valid {
			sum++
		}
		if s.Debug {
			fmt.Println(valid, levels)
		}
	}
	return sum
}

func (s *s2402) validateLevels(levels []int) bool {
	previousLevel := 0
	goingDown := false
	for i, level := range levels {
		if i > 0 {
			if i == 1 {
				if level < previousLevel {
					goingDown = true
				}
			}
			if goingDown {
				if level >= previousLevel || previousLevel-level > 3 {
					return false
				}
			} else {
				if level <= previousLevel || level-previousLevel > 3 {
					return false
				}
			}
		}
		previousLevel = level
	}
	return true
}
