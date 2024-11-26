package main

import (
	"fmt"
	"strconv"
)

type s2310 Puzzle

func (s *s2310) SetDebug(debug bool) error {
	s.Debug = debug
	return nil
}

func (s *s2310) SetInput(input []string) error {
	s.Input = input
	return nil
}

func (s *s2310) SetPart(part int) error {
	s.Part = part
	return nil
}

func (s *s2310) Solve() (string, error) {
	var number int
	if s.Part != 2 {
		number = s.processPart1()
	} else {
		number = s.processPart2()
	}
	return strconv.Itoa(number), nil
}

func (s *s2310) processPart1() int {
	// exists for each pipe type
	exits := make(map[rune][2][4]int)
	exits['|'] = [2][4]int{{-1, 0}, {1, 0}}
	exits['-'] = [2][4]int{{0, -1}, {0, 1}}
	exits['L'] = [2][4]int{{-1, 0}, {0, 1}}
	exits['J'] = [2][4]int{{-1, 0}, {0, -1}}
	exits['7'] = [2][4]int{{0, -1}, {1, 0}}
	exits['F'] = [2][4]int{{1, 0}, {0, 1}}
	grid := MakeGridRune(s.Input)
	// find start
	var x, y, prevX, prevY, nextX, nextY int
	for yval, yy := range grid {
		for xval, xx := range yy {
			if xx == 'S' {
				x = xval
				y = yval
			}
		}
	}
	if s.Debug {
		fmt.Println(y, x)
	}
	// sample = F input = 7
	pipe := '7'
	prevY = -1
	prevX = -1
	counter := 0
	for grid[nextY][nextX] != 'S' {
		counter++
		for _, ex := range exits[pipe] {
			var proX, proY int
			proY = y + ex[0]
			proX = x + ex[1]
			if proY > -1 && proY < len(grid) && proX > -1 && proX < len(grid[0]) && (proY != prevY || proX != prevX) {
				nextY = proY
				nextX = proX
			}
		}
		prevY = y
		prevX = x
		y = nextY
		x = nextX
		pipe = grid[y][x]
		if s.Debug {
			fmt.Println(prevY, prevX, y, x)
		}
	}

	return counter / 2
}

// too high: 660
func (s *s2310) processPart2() int {
	// exists for each pipe type
	exits := make(map[rune][2][4]int)
	bigGrid := make(map[int]map[int]int)
	exits['|'] = [2][4]int{{-1, 0}, {1, 0}}
	exits['-'] = [2][4]int{{0, -1}, {0, 1}}
	exits['L'] = [2][4]int{{-1, 0}, {0, 1}}
	exits['J'] = [2][4]int{{-1, 0}, {0, -1}}
	exits['7'] = [2][4]int{{0, -1}, {1, 0}}
	exits['F'] = [2][4]int{{1, 0}, {0, 1}}
	grid := MakeGridRune(s.Input)
	for i := -1; i < len(grid)*2; i++ {
		bigGrid[i] = make(map[int]int)
	}
	// find start
	var x, y, prevX, prevY, nextX, nextY int
	for yval, yy := range grid {
		for xval, xx := range yy {
			if xx == 'S' {
				x = xval
				y = yval
			}
		}
	}
	if s.Debug {
		fmt.Println(y, x)
	}
	// sample = F input = 7
	pipe := '7'
	prevY = -1
	prevX = -1
	counter := 0
	for grid[nextY][nextX] != 'S' {
		counter++
		modY := 0
		modX := 0
		for _, ex := range exits[pipe] {
			var proX, proY int
			proY = y + ex[0]
			proX = x + ex[1]
			if proY > -1 && proY < len(grid) && proX > -1 && proX < len(grid[0]) && (proY != prevY || proX != prevX) {
				nextY = proY
				nextX = proX
				modY = ex[0]
				modX = ex[1]
			}
		}
		bigGrid[y*2+modY][x*2+modX] = 1
		bigGrid[nextY*2][nextX*2] = 1
		prevY = y
		prevX = x
		y = nextY
		x = nextX
		pipe = grid[y][x]
		//fmt.Println(prevY, prevX, y, x)
	}
	// flood fill
	bigGrid[-1][-1] = 2
	modified := true
	for modified {
		modified = false
		for yy := -1; yy < len(grid)*2; yy++ {
			for xx := -1; xx < len(grid[0])*2; xx++ {
				if _, ok1 := bigGrid[yy][xx]; !ok1 {
					if val, ok := bigGrid[yy-1][xx]; ok && val == 2 {
						bigGrid[yy][xx] = 2
						modified = true
					}
					if val, ok := bigGrid[yy+1][xx]; ok && val == 2 {
						bigGrid[yy][xx] = 2
						modified = true
					}
					if val, ok := bigGrid[yy][xx-1]; ok && val == 2 {
						bigGrid[yy][xx] = 2
						modified = true
					}
					if val, ok := bigGrid[yy][xx+1]; ok && val == 2 {
						bigGrid[yy][xx] = 2
						modified = true
					}
				}
			}
		}
	}
	counter = 0
	for yy := 0; yy < len(grid)*2; yy += 2 {
		var buf string
		for xx := 0; xx < len(grid[0])*2; xx += 2 {
			if val, ok := bigGrid[yy][xx]; ok {
				if val == 2 {
					buf += "."
				} else {
					buf += "X"
				}
			} else {
				counter++
				buf += ","
			}
		}
		if s.Debug {
			fmt.Println(buf)
		}
	}

	return counter
}
