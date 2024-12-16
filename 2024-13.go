package main

import (
	"fmt"
	"strconv"
)

type s2413 Puzzle

type claw struct {
	a [2]int
	b [2]int
	p [2]int
}

func (s *s2413) SetDebug(debug bool) error {
	s.Debug = debug
	return nil
}

func (s *s2413) SetInput(input []string) error {
	s.Input = input
	return nil
}

func (s *s2413) SetPart(part int) error {
	s.Part = part
	return nil
}

func (s *s2413) Solve() (string, error) {
	var number int
	if s.Part != 2 {
		number = s.processPart1()
	} else {
		number = s.processPart2()
	}
	return strconv.Itoa(number), nil
}

func (s *s2413) processPart1() int {
	sum := 0
	var claws []claw
	var c claw
	for i, line := range s.Input {
		ints := ExtractInts(line)
		switch i % 4 {
		case 0:
			c.a = [2]int{ints[0], ints[1]}
		case 1:
			c.b = [2]int{ints[0], ints[1]}
		case 2:
			c.p = [2]int{ints[0], ints[1]}
			claws = append(claws, c)
		}
	}
	for _, c := range claws {
		lowScore := 0
		for a := 0; a < 101; a++ {
			for b := 0; b < 101; b++ {
				if c.a[0]*a+c.b[0]*b == c.p[0] && c.a[1]*a+c.b[1]*b == c.p[1] {
					score := a*3 + b
					if lowScore == 0 || score < lowScore {
						lowScore = score
					}
				}
			}
		}
		sum += lowScore
	}
	return sum
}

func (s *s2413) processPart2() int {
	sum := 0
	var claws []claw
	var c claw
	for i, line := range s.Input {
		ints := ExtractInts(line)
		switch i % 4 {
		case 0:
			c.a = [2]int{ints[0], ints[1]}
		case 1:
			c.b = [2]int{ints[0], ints[1]}
		case 2:
			c.p = [2]int{ints[0] + 10000000000000, ints[1] + 10000000000000}
			claws = append(claws, c)
		}
	}
	for _, c := range claws {
		requiredBPresses := (c.p[1]*c.a[0] - c.p[0]*c.a[1]) / (c.b[1]*c.a[0] - c.b[0]*c.a[1])
		requiredAPresses := (c.p[0] - requiredBPresses*c.b[0]) / c.a[0]
		if s.Debug {
			fmt.Println(requiredBPresses, c.a[0]*requiredAPresses+c.b[0]*requiredBPresses, c.p[0], requiredAPresses)
		}
		if c.a[0]*requiredAPresses+c.b[0]*requiredBPresses == c.p[0] && c.a[1]*requiredAPresses+c.b[1]*requiredBPresses == c.p[1] {
			sum += requiredAPresses*3 + requiredBPresses
		}
	}
	return sum
}
