package main

import (
	"fmt"
	"strconv"
)

type s2408 Puzzle

func (s *s2408) SetDebug(debug bool) error {
	s.Debug = debug
	return nil
}

func (s *s2408) SetInput(input []string) error {
	s.Input = input
	return nil
}

func (s *s2408) SetPart(part int) error {
	s.Part = part
	return nil
}

func (s *s2408) Solve() (string, error) {
	return strconv.Itoa(s.Process()), nil
}

func (s *s2408) Process() int {
	sum := 0
	//lookup := map[rune]struct{}{'A': {}, '0': {}}
	grid := MakeGridRune(s.Input)

	// put the coordinates of each antenna into a map
	antenna := make(map[rune][][2]int)
	for y, line := range grid {
		for x, g := range line {
			if g != '.' {
				antenna[g] = append(antenna[g], [2]int{y, x})
			}
		}
	}

	if s.Debug {
		fmt.Println(antenna)
	}

	maxY := len(grid)
	maxX := len(grid[0])
	for _, coords := range antenna {
		// for each antenna of the type, calculate the two antinodes it creates with each other antenna of that type
		for i, yx := range coords {
			for i2, yx2 := range coords {
				if i != i2 {
					y := yx2[0] - yx[0]
					x := yx2[1] - yx[1]
					var antinodes [][2]int
					factor := 2
					if s.Part == 2 {
						// for part 2, try up to a factor of 50 in each direction
						factor = 50
					}
					for f := 1; f < factor; f++ {
						antinodes = append(antinodes, [2]int{yx[0] - y*f, yx[1] - x*f})
						antinodes = append(antinodes, [2]int{yx2[0] + y*f, yx2[1] + x*f})
					}
					for _, node := range antinodes {
						if node[0] >= 0 && node[0] < maxY && node[1] >= 0 && node[1] < maxX {
							grid[node[0]][node[1]] = '#'
						}
					}
				}
			}
		}
	}

	for _, line := range grid {
		for _, g := range line {
			if s.Part != 2 {
				// for part 1, count only the antinodes
				if g == '#' {
					sum++
				}
			} else {
				// for part 2, count all antinodes and all antenna
				if g != '.' {
					sum++
				}
			}
		}
	}
	return sum
}
