package main

import (
	"fmt"
	"strconv"
)

type s2311 Puzzle

func (s *s2311) SetDebug(debug bool) error {
	s.Debug = debug
	return nil
}

func (s *s2311) SetInput(input []string) error {
	s.Input = input
	return nil
}

func (s *s2311) SetPart(part int) error {
	s.Part = part
	return nil
}

func (s *s2311) Solve() (string, error) {
	return strconv.Itoa(s.Process()), nil
}

func (s *s2311) Process() int {
	lines := MakeGridRune(s.Input)

	universe := make(map[int]map[int]bool)

	if s.Debug {
		for _, line := range lines {
			fmt.Printf("%q\n", line)
		}
	}

	// first check for columns that need to be expanded
	index := 0
	for i := 0; i < len(lines[0]); i++ {
		isEmpty := true
		for y, line := range lines {
			if i == 0 {
				universe[y] = make(map[int]bool)
			}
			if line[i] == '#' {
				isEmpty = false
			}
		}
		loops := 1
		if isEmpty {
			loops = 2
		}
		for j := 0; j < loops; j++ {
			for y, line := range lines {
				universe[y][index] = line[i] == '#'
			}
			index++
		}
	}

	if s.Debug {
		for _, line := range universe {
			fmt.Println(line)
		}
	}

	i := 0
	for _, line := range lines {
		isEmpty := true
		for _, r := range line {
			if r == '#' {
				isEmpty = false
			}
		}
		if isEmpty {
			for j := len(universe); j > i; j++ {
				universe[j] = universe[j-1]
			}
			i++
		}
		i++
	}

	if s.Debug {
		for _, line := range universe {
			fmt.Println(line)
		}
	}

	return len(lines)
}
