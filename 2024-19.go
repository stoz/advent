package main

import (
	"fmt"
	"strconv"
)

type s2419 Puzzle

func (s *s2419) SetDebug(debug bool) error {
	s.Debug = debug
	return nil
}

func (s *s2419) SetInput(input []string) error {
	s.Input = input
	return nil
}

func (s *s2419) SetPart(part int) error {
	s.Part = part
	return nil
}

func (s *s2419) Solve() (string, error) {
	var number int
	if s.Part != 2 {
		number = s.processPart1()
	} else {
		number = s.processPart2()
	}
	return strconv.Itoa(number), nil
}

func (s *s2419) processPart1() int {
	sum := 0
	towels := []string{}
	t := ""
	for _, r := range s.Input[0] {
		if r == ',' {
			towels = append(towels, t)
			t = ""
		} else if r != ' ' {
			t += string(r)
		}
	}
	towels = append(towels, t)
	if s.Debug {
		fmt.Println(towels)
	}

	designs := []string{}
	for i, line := range s.Input {
		if i > 1 {
			designs = append(designs, line)
		}
	}
	if s.Debug {
		fmt.Println(designs)
	}

	for _, design := range designs {
		partials := map[string]bool{}
		// populate the initial set of partial matches
		for _, t := range towels {
			if t == design[:len(t)] {
				partials[t] = false
			}
		}

		count := len(partials)
		do := false
		if count > 0 {
			do = true
		}
		for do {
			for p, b := range partials {
				if !b && do {
					for _, t := range towels {
						if design == p+t && do {
							do = false
							sum++
						} else if len(p)+len(t) <= len(design) && t == design[len(p):len(p)+len(t)] {
							partials[p+t] = false
						}
					}
					partials[p] = true
				}
			}
			if len(partials) > count {
				count = len(partials)
			} else {
				do = false
			}
		}
	}
	return sum
}

func (s *s2419) processPart2() int {
	sum := 0
	towels := []string{}
	t := ""
	for _, r := range s.Input[0] {
		if r == ',' {
			towels = append(towels, t)
			t = ""
		} else if r != ' ' {
			t += string(r)
		}
	}
	towels = append(towels, t)
	if s.Debug {
		fmt.Println(towels)
	}

	designs := []string{}
	for i, line := range s.Input {
		if i > 1 {
			designs = append(designs, line)
		}
	}
	if s.Debug {
		fmt.Println(designs)
	}

	for _, design := range designs {
		partials := map[string][2]int{} // new, count
		// populate the initial set of partial matches
		for _, t := range towels {
			if t == design[:len(t)] {
				partials[t] = [2]int{0, 1}
			}
		}

		count := len(partials)
		do := false
		if count > 0 {
			do = true
		}
		minLength := 1
		for do {
			for p, b := range partials {
				if b[0] < 1 && len(p)+len(t) < minLength {
					for _, t := range towels {
						if design == p+t {
							// do the final comparison later to make sure we have all the parent scores
						} else if len(p)+len(t) <= len(design) && t == design[len(p):len(p)+len(t)] {
							update := partials[p+t]
							update[1] += b[1]
							partials[p+t] = update
						}
					}
					fullUpdate := partials[p]
					if fullUpdate[0] == 0 {
						fullUpdate[0] = 1
						partials[p] = fullUpdate
					}
				}
			}
			if len(partials) > count || minLength <= len(design) {
				count = len(partials)
				minLength++
			} else {
				for p, b := range partials {
					for _, t := range towels {
						if design == p+t {
							sum += b[1]
						}
					}
				}
				do = false
			}
		}

		//fmt.Println(design, partials)
	}
	return sum
}
