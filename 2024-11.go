package main

import (
	"fmt"
	"strconv"
)

type s2411 Puzzle

func (s *s2411) SetDebug(debug bool) error {
	s.Debug = debug
	return nil
}

func (s *s2411) SetInput(input []string) error {
	s.Input = input
	return nil
}

func (s *s2411) SetPart(part int) error {
	s.Part = part
	return nil
}

func (s *s2411) Solve() (string, error) {
	return strconv.Itoa(s.Process()), nil
}

func (s *s2411) Process() int {
	input := ExtractInts(s.Input[0])
	ints := make(map[int]int)
	for _, inp := range input {
		ints[inp]++
	}
	if s.Debug {
		fmt.Println(ints)
	}
	length := 25
	if s.Part == 2 {
		length = 75
	}
	for i := 0; i < length; i++ {
		ints = s.blink(ints)
		if s.Debug && i < 5 {
			fmt.Println(ints)
		}
	}
	sum := 0
	for _, v := range ints {
		if v > 0 {
			sum += v
		}
	}
	return sum
}

func (s *s2411) blink(a map[int]int) map[int]int {
	b := make(map[int]int)
	for val, count := range a {
		if count > 0 {
			y := strconv.Itoa(val)
			if val == 0 {
				b[1] += count
			} else if len(y)%2 == 0 {
				first, _ := strconv.Atoi(y[:(len(y))/2])
				second, _ := strconv.Atoi(y[len(y)/2:])
				b[first] += count
				b[second] += count
			} else {
				b[val*2024] += count
			}
		}
	}
	return b
}
