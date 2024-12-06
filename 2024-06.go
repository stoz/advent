package main

import (
	"fmt"
	"strconv"
)

type s2406 Puzzle

func (s *s2406) SetDebug(debug bool) error {
	s.Debug = debug
	return nil
}

func (s *s2406) SetInput(input []string) error {
	s.Input = input
	return nil
}

func (s *s2406) SetPart(part int) error {
	s.Part = part
	return nil
}

func (s *s2406) Solve() (string, error) {
	var number int
	if s.Part != 2 {
		number = s.processPart1()
	} else {
		number = s.processPart2()
	}
	return strconv.Itoa(number), nil
}

func (s *s2406) processPart1() int {
	sum := 0
	grid := MakeGridRune(s.Input)
	var startY, startX int
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == '^' {
				startY = y
				startX = x
				grid[y][x] = 'X'
			}
		}
	}
	grid = s.patrol(grid, startY, startX)
	// count the Xs
	for y := 0; y < len(grid); y++ {
		if s.Debug {
			fmt.Println(string(grid[y]))
		}
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 'X' {
				sum++
			}
		}
	}
	return sum
}

func (s *s2406) processPart2() int {
	sum := 0
	grid := MakeGridRune(s.Input)
	directions := [4][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	var startY, startX, posY, posX int
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == '^' {
				startY = y
				startX = x
				grid[y][x] = 'X'
			}
		}
	}
	grid = s.patrol(grid, startY, startX)
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 'X' && !(y == startY && x == startX) {
				direction := 0
				newGrid := MakeGridRune(s.Input)
				newGrid[y][x] = '#'
				newGrid[startY][startX] = 'X'
				posY = startY
				posX = startX
				do := true
				counter := 0
				for do {
					counter++
					newPosY := posY + directions[direction][1]
					newPosX := posX + directions[direction][0]
					if counter > 10000 {
						// loop detected
						sum++
						do = false
					} else if newPosY < 0 || newPosY >= len(newGrid) || newPosX < 0 || newPosX >= len(newGrid[newPosY]) {
						// out of bounds
						do = false
					} else if newGrid[newPosY][newPosX] == '.' || newGrid[newPosY][newPosX] == 'X' {
						// move forward
						newGrid[newPosY][newPosX] = 'X'
						posY = newPosY
						posX = newPosX
					} else if newGrid[newPosY][newPosX] == '#' {
						// turn right
						direction = (direction + 1) % 4
					}
				}
			}
		}
	}
	return sum
}

func (s *s2406) patrol(grid [][]rune, posY, posX int) [][]rune {
	directions := [4][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	direction := 0
	do := true
	for do {
		newPosY := posY + directions[direction][1]
		newPosX := posX + directions[direction][0]
		if newPosY < 0 || newPosY >= len(grid) || newPosX < 0 || newPosX >= len(grid[newPosY]) {
			// out of bounds
			do = false
		} else if grid[newPosY][newPosX] == '.' || grid[newPosY][newPosX] == 'X' {
			// move forwards
			grid[newPosY][newPosX] = 'X'
			posY = newPosY
			posX = newPosX
		} else if grid[newPosY][newPosX] == '#' {
			// turn right
			direction = (direction + 1) % 4
		}
	}
	return grid
}
