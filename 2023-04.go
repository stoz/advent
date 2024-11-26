package main

import (
	"strconv"
)

type s2304 Puzzle

func (s *s2304) SetDebug(debug bool) error {
	s.Debug = debug
	return nil
}

func (s *s2304) SetInput(input []string) error {
	s.Input = input
	return nil
}

func (s *s2304) SetPart(part int) error {
	s.Part = part
	return nil
}

func (s *s2304) Solve() (string, error) {
	return strconv.Itoa(s.Process()), nil
}

func (s *s2304) Process() int {
	div := 0
	total := 0
	var cards [204]int
	for c := range cards {
		cards[c] = 1
	}
	for k, v := range s.Input[0] {
		if v == '|' {
			// subtract 5 for the text and padding at the start of the line,
			// then each number plus padding is 3 characters wide
			div = (k - 5) / 3
		}
	}
	for lineIndex, line := range s.Input {
		score := 0
		a := ExtractInts(line)
		for i := 1; i < div; i++ {
			for j := div; j < len(a); j++ {
				if a[i] == a[j] {
					if s.Part == 2 {
						score++
						if lineIndex+score <= len(cards) {
							cards[lineIndex+score] += cards[lineIndex]
						}
					} else if score == 0 {
						score = 1
					} else {
						score = score * 2
					}
				}
			}
		}
		if s.Part == 2 {
			total += cards[lineIndex]
		} else {
			total += score
		}
	}
	return total
}
