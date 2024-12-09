package main

import (
	"fmt"
	"strconv"
)

type s2407 Puzzle

func (s *s2407) SetDebug(debug bool) error {
	s.Debug = debug
	return nil
}

func (s *s2407) SetInput(input []string) error {
	s.Input = input
	return nil
}

func (s *s2407) SetPart(part int) error {
	s.Part = part
	return nil
}

func (s *s2407) Solve() (string, error) {
	var number int
	if s.Part != 2 {
		number = s.processPart1()
	} else {
		number = s.processPart2()
	}
	return strconv.Itoa(number), nil
}

func (s *s2407) processPart1() int {
	sum := 0
	for _, line := range s.Input {
		equation := ExtractInts(line)
		do := true
		counter := 0
		operators := [12]bool{}
		result := equation[0]
		for do {
			if result == 292 && counter < 20 {
				fmt.Println(operators)
			}
			calculation := 0
			for i, v := range equation {
				if i == 1 {
					calculation = v
				}
				if i > 1 {
					if operators[i-2] {
						// muliply
						calculation *= v
					} else {
						// addition
						calculation += v
					}
				}
			}
			if operators[11] {
				do = false
			} else if calculation == result {
				sum += result
				fmt.Println(result)
				do = false
			} else {
				counter++
				// do some dumb binary stuff
				x := counter
				if result == 292 && counter < 20 {
					fmt.Println(x)
				}
				lookup := [12]int{2048, 1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}
				for lk, lv := range lookup {
					if x >= lv {
						operators[11-lk] = true
						x -= lv
					} else {
						operators[11-lk] = false
					}
				}
			}
		}
	}
	return sum
}

func (s *s2407) processPart2() int {
	sum := 0
	for _, line := range s.Input {
		equation := ExtractInts(line)
		do := true
		counter := 0
		operators := [12]int{}
		result := equation[0]
		for do {
			if result == 292 && counter < 20 {
				//fmt.Println(operators)
			}
			calculation := 0
			for i, v := range equation {
				if i == 1 {
					calculation = v
				}
				if i > 1 {
					if operators[i-2] == 1 {
						// muliply
						calculation *= v
					} else if operators[i-2] == 2 {
						// concat
						concat, _ := strconv.Atoi(strconv.Itoa(calculation) + strconv.Itoa(v))
						if result == 156 && counter < 6 {
							fmt.Println(operators, calculation, v, concat)
						}
						calculation = concat
					} else {
						// addition
						calculation += v
					}
				}
			}
			if operators[11] == 2 {
				do = false
			} else if calculation == result {
				sum += result
				fmt.Println(result)
				do = false
			} else {
				counter++
				for i, o := range operators {
					if o < 2 {
						operators[i]++
						// reset everything to the left
						if i > 0 {
							for r := 0; r < i; r++ {
								operators[r] = 0
							}
						}
						break
					}
				}
			}
		}
	}
	return sum
}
