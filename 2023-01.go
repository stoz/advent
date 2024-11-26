package main

import (
	"strconv"
)

type s2301 Puzzle

func (s *s2301) SetDebug(debug bool) error {
	s.Debug = debug
	return nil
}

func (s *s2301) SetInput(input []string) error {
	s.Input = input
	return nil
}

func (s *s2301) SetPart(part int) error {
	s.Part = part
	return nil
}

func (s *s2301) Solve() (string, error) {
	return strconv.Itoa(s.Process()), nil
}

func (s *s2301) Process() int {
	// expected results: 55834, 53221
	part2 := s.Part == 2
	sum := 0
	for _, line := range s.Input {
		first := -1
		last := -1
		for i, c := range line {
			n := -1
			switch c {
			case '0':
				n = 0
			case '1':
				n = 1
			case '2':
				n = 2
			case '3':
				n = 3
			case '4':
				n = 4
			case '5':
				n = 5
			case '6':
				n = 6
			case '7':
				n = 7
			case '8':
				n = 8
			case '9':
				n = 9
			case 'e': // one, three, five, nine
				if part2 && i > 1 && line[i-2] == 'o' && line[i-1] == 'n' {
					n = 1
				}
				if part2 && i > 3 && line[i-4] == 't' && line[i-3] == 'h' && line[i-2] == 'r' && line[i-1] == 'e' {
					n = 3
				}
				if part2 && i > 2 && line[i-3] == 'f' && line[i-2] == 'i' && line[i-1] == 'v' {
					n = 5
				}
				if part2 && i > 2 && line[i-3] == 'n' && line[i-2] == 'i' && line[i-1] == 'n' {
					n = 9
				}
			case 'o': // two
				if part2 && i > 1 && line[i-2] == 't' && line[i-1] == 'w' {
					n = 2
				}
			case 'r': // four
				if part2 && i > 2 && line[i-3] == 'f' && line[i-2] == 'o' && line[i-1] == 'u' {
					n = 4
				}
			case 'x': // six
				if part2 && i > 1 && line[i-2] == 's' && line[i-1] == 'i' {
					n = 6
				}
			case 'n': // seven
				if part2 && i > 3 && line[i-4] == 's' && line[i-3] == 'e' && line[i-2] == 'v' && line[i-1] == 'e' {
					n = 7
				}
			case 't': // eight
				if part2 && i > 3 && line[i-4] == 'e' && line[i-3] == 'i' && line[i-2] == 'g' && line[i-1] == 'h' {
					n = 8
				}
			}
			if n > -1 && first < 0 {
				first = n
			}
			if n > -1 {
				last = n
			}
		}
		sum += 10*first + last
	}
	return sum
}
