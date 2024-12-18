package main

import (
	"fmt"
	"strconv"
)

type s2418 Puzzle

func (s *s2418) SetDebug(debug bool) error {
	s.Debug = debug
	return nil
}

func (s *s2418) SetInput(input []string) error {
	s.Input = input
	return nil
}

func (s *s2418) SetPart(part int) error {
	s.Part = part
	return nil
}

func (s *s2418) Solve() (string, error) {
	if s.Part != 2 {
		return strconv.Itoa(s.processPart1()), nil
	} else {
		ints := s.processPart2()
		return strconv.Itoa(ints[0]) + "," + strconv.Itoa(ints[1]), nil
	}
}

func (s *s2418) processPart1() int {
	grid := make(map[[2]int]uint8)
	yMax := 70
	xMax := 70
	limit := 1024
	// sample data has less lines
	if len(s.Input) < 100 {
		yMax = 6
		xMax = 6
		limit = 12
	}

	for y := 0; y < yMax+1; y++ {
		for x := 0; x < xMax+1; x++ {
			grid[[2]int{y, x}] = 1
		}
	}
	for i, line := range s.Input {
		if i < limit { // 12 for sample
			ints := ExtractInts(line)
			grid[[2]int{ints[0], ints[1]}] = 2
		}
	}

	// print grid
	if s.Debug {
		s.print(grid, yMax, xMax)
	}

	// find the shortest path to the exit
	grid[[2]int{0, 0}] = 3
	do := true
	directions := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	i := 0
	for do {
		i++
		for y := 0; y < yMax+1; y++ {
			for x := 0; x < xMax+1; x++ {
				if grid[[2]int{y, x}] == 3 {
					for _, d := range directions {
						if grid[[2]int{y + d[0], x + d[1]}] == 1 {
							grid[[2]int{y + d[0], x + d[1]}] = 4
						}
					}
				}
			}
		}
		for y := 0; y < yMax+1; y++ {
			for x := 0; x < xMax+1; x++ {
				if grid[[2]int{y, x}] == 4 {
					grid[[2]int{y, x}] = 3
				}
			}
		}
		if grid[[2]int{yMax, xMax}] == 3 {
			return i
		}
		if i > 500 {
			do = false
		}
	}
	s.print(grid, yMax, xMax)

	return 0
}
func (s *s2418) processPart2() [2]int {
	grid := make(map[[2]int]uint8)

	// pre-extract the input
	bytes := [][2]int{}
	for _, line := range s.Input {
		b := ExtractInts(line)
		bytes = append(bytes, [2]int{b[0], b[1]})
	}

	yMax := 70
	xMax := 70
	bounds := [2]int{1025, len(bytes)}
	// sample data has less lines
	if len(s.Input) < 100 {
		yMax = 6
		xMax = 6
		bounds[0] = 1
	}

	do := true
	for do {
		// set initial state for grid
		if s.Debug {
			fmt.Println(bounds)
		}
		for y := 0; y < yMax+1; y++ {
			for x := 0; x < xMax+1; x++ {
				grid[[2]int{y, x}] = 1
			}
		}
		test := FloorDivision(bounds[0]+bounds[1], 2)
		for i := 0; i < test; i++ {
			grid[[2]int{bytes[i][0], bytes[i][1]}] = 2
		}
		grid[[2]int{0, 0}] = 3
		if s.checkExit(grid, yMax, xMax) {
			bounds[0] = test
		} else {
			bounds[1] = test
		}
		if bounds[1]-bounds[0] == 1 {
			return bytes[bounds[0]]
		}
	}
	return [2]int{0, 0}
}

func (s *s2418) checkExit(grid map[[2]int]uint8, yMax, xMax int) bool {
	directions := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for i := 0; i < 1000; i++ {
		for y := 0; y <= yMax; y++ {
			for x := 0; x <= xMax; x++ {
				if grid[[2]int{y, x}] == 3 {
					for _, d := range directions {
						if grid[[2]int{y + d[0], x + d[1]}] == 1 {
							grid[[2]int{y + d[0], x + d[1]}] = 4
						}
					}
				}
			}
		}
		for y := 0; y <= yMax; y++ {
			for x := 0; x <= xMax; x++ {
				if grid[[2]int{y, x}] == 4 {
					grid[[2]int{y, x}] = 3
				}
			}
		}
		if grid[[2]int{yMax, xMax}] == 3 {
			return true
		}
	}
	return false
}

func (s *s2418) print(grid map[[2]int]uint8, yMax, xMax int) {
	for y := 0; y < yMax+1; y++ {
		for x := 0; x < xMax+1; x++ {
			switch grid[[2]int{x, y}] {
			case 4:
				print("X")
			case 3:
				print("O")
			case 2:
				print("#")
			case 1:
				print(".")
			}
		}
		println()
	}
}
