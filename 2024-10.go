package main

import (
	"fmt"
	"strconv"
)

type s2410 Puzzle

func (s *s2410) SetDebug(debug bool) error {
	s.Debug = debug
	return nil
}

func (s *s2410) SetInput(input []string) error {
	s.Input = input
	return nil
}

func (s *s2410) SetPart(part int) error {
	s.Part = part
	return nil
}

func (s *s2410) Solve() (string, error) {
	return strconv.Itoa(s.Process()), nil
}

func (s *s2410) Process() int {
	sum := 0
	grid := MakeGridRune(s.Input)
	directions := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var searchSpace [][4]int
	uniques := make(map[string]bool)
	maxY := len(grid)
	maxX := len(grid[0])
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == '0' {
				searchSpace = append(searchSpace, [4]int{y, x, y, x})
			}
		}
	}
	for len(searchSpace) > 0 {
		t := searchSpace[len(searchSpace)-1]
		searchSpace = searchSpace[:len(searchSpace)-1]
		for _, d := range directions {
			td := [4]int{t[0] + d[0], t[1] + d[1], t[2], t[3]}
			if td[0] >= 0 && td[0] < maxY && td[1] >= 0 && td[1] < maxX {
				if grid[td[0]][td[1]] == grid[t[0]][t[1]]+1 {
					if grid[td[0]][td[1]] == '9' {
						sum++
						uniques[fmt.Sprintf("%d,%d,%d,%d", td[0], td[1], td[2], td[3])] = true
					} else {
						searchSpace = append(searchSpace, td)
					}
				}
			}
		}
		if s.Debug {
			fmt.Println(string(grid[t[0]][t[1]]))
		}
	}
	if s.Part != 2 {
		return len(uniques)
	} else {
		return sum
	}
}
