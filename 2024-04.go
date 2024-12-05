package main

import (
	"fmt"
	"strconv"
)

type s2404 Puzzle

func (s *s2404) SetDebug(debug bool) error {
	s.Debug = debug
	return nil
}

func (s *s2404) SetInput(input []string) error {
	s.Input = input
	return nil
}

func (s *s2404) SetPart(part int) error {
	s.Part = part
	return nil
}

func (s *s2404) Solve() (string, error) {
	var number int
	if s.Part != 2 {
		number = s.processPart1()
	} else {
		number = s.processPart2()
	}
	return strconv.Itoa(number), nil
}

func (s *s2404) processPart1() int {
	sum := 0
	grid := MakeGridRune(s.Input)
	// we make an assumption here that the input is square
	yMax := len(grid) - 1
	for y := range grid {
		for x, val := range grid[y] {
			if val == 'X' {
				surrounds := s.getSurround(x, y)
				if s.Debug {
					fmt.Print(x, y)
					fmt.Println(surrounds)
				}
				for _, p := range surrounds {
					valid := true
					for _, v := range p {
						if v < 0 || v > yMax {
							valid = false
						}
					}
					if valid && grid[p[1]][p[0]] == 'M' && grid[p[3]][p[2]] == 'A' && grid[p[5]][p[4]] == 'S' {
						if s.Debug {
							fmt.Println(x, y, p)
						}
						sum++
					}
				}
			}
		}
	}
	return sum
}

func (s *s2404) processPart2() int {
	sum := 0
	grid := MakeGridRune(s.Input)
	// we make an assumption here that the input is square
	yMax := len(grid) - 1
	for y := range grid {
		for x, val := range grid[y] {
			if x > 0 && y > 0 && x < yMax && y < yMax && val == 'A' {
				surrounds := s.getCorners(x, y)
				var corners [4]rune
				for pi, p := range surrounds {
					corners[pi] = grid[p[1]][p[0]]
				}
				match := false
				if corners[0] == 'M' && corners[3] == 'S' {
					if (corners[1] == 'M' && corners[2] == 'S') || (corners[1] == 'S' && corners[2] == 'M') {
						match = true
					}
				}
				if corners[1] == 'M' && corners[2] == 'S' {
					if (corners[0] == 'M' && corners[3] == 'S') || (corners[0] == 'S' && corners[3] == 'M') {
						match = true
					}
				}
				if corners[2] == 'M' && corners[1] == 'S' {
					if (corners[0] == 'M' && corners[3] == 'S') || (corners[0] == 'S' && corners[3] == 'M') {
						match = true
					}
				}
				if corners[3] == 'M' && corners[0] == 'S' {
					if (corners[1] == 'M' && corners[2] == 'S') || (corners[1] == 'S' && corners[2] == 'M') {
						match = true
					}
				}
				if match {
					if s.Debug {
						fmt.Println(x, y)
						fmt.Println(string(grid[y-1][x-1]), " ", string(grid[y-1][x]))
						fmt.Println(" ", string(grid[y][x]), " ")
						fmt.Println(string(grid[y+1][x-1]), " ", string(grid[y+1][x+1]))
					}
					sum++
				}
			}
		}
	}
	return sum
}

func (s *s2404) getSurround(x, y int) [][6]int {
	return [][6]int{
		{x - 1, y - 1, x - 2, y - 2, x - 3, y - 3},
		{x, y - 1, x, y - 2, x, y - 3},
		{x + 1, y - 1, x + 2, y - 2, x + 3, y - 3},
		{x - 1, y, x - 2, y, x - 3, y},
		{x + 1, y, x + 2, y, x + 3, y},
		{x - 1, y + 1, x - 2, y + 2, x - 3, y + 3},
		{x, y + 1, x, y + 2, x, y + 3},
		{x + 1, y + 1, x + 2, y + 2, x + 3, y + 3},
	}
}

func (s *s2404) getCorners(x, y int) [][4]int {
	return [][4]int{
		{x - 1, y - 1},
		{x + 1, y - 1},
		{x - 1, y + 1},
		{x + 1, y + 1},
	}
}
