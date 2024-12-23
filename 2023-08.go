package main

import (
	"fmt"
	"strconv"
)

type s2308 Puzzle

func (s *s2308) SetDebug(debug bool) error {
	s.Debug = debug
	return nil
}

func (s *s2308) SetInput(input []string) error {
	s.Input = input
	return nil
}

func (s *s2308) SetPart(part int) error {
	s.Part = part
	return nil
}

func (s *s2308) Solve() (string, error) {
	var number int
	if s.Part != 2 {
		number = s.processPart1()
	} else {
		number = s.processPart2()
	}
	return strconv.Itoa(number), nil
}

func (s *s2308) processPart1() int {
	directions := s.Input[0]
	network := make(map[string][2]string)
	pos := "AAA"
	for i, line := range s.Input {
		if i > 1 {
			var net [2]string
			net[0] = line[7:10]
			net[1] = line[12:15]
			network[line[0:3]] = net
		}
	}
	i := 0
	for pos != "ZZZ" {
		lr := 0
		if directions[i%len(directions)] == 'R' {
			lr = 1
		}
		pos = network[pos][lr]
		if s.Debug {
			fmt.Println(pos, lr)
		}
		i++
	}
	return i
}

func (s *s2308) processPart2() int {
	directions := s.Input[0]
	network := make(map[string][2]string)
	var positions []string
	var results []int
	for i, line := range s.Input {
		if i > 1 {
			var net [2]string
			net[0] = line[7:10]
			net[1] = line[12:15]
			network[line[0:3]] = net
			if line[2] == 'A' {
				positions = append(positions, line[0:3])
			}
		}
	}
	for _, pos := range positions {
		i := 0
		for pos[2] != 'Z' {
			lr := 0
			if directions[i%len(directions)] == 'R' {
				lr = 1
			}
			pos = network[pos][lr]
			if s.Debug {
				fmt.Println(pos, lr)
			}
			i++
		}
		results = append(results, i)
	}
	return LCM(results[0], results[1], results[2:]...)
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
