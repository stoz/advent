package main

import (
	"fmt"
	"strconv"
)

type s2415 Puzzle

func (s *s2415) SetDebug(debug bool) error {
	s.Debug = debug
	return nil
}

func (s *s2415) SetInput(input []string) error {
	s.Input = input
	return nil
}

func (s *s2415) SetPart(part int) error {
	s.Part = part
	return nil
}

func (s *s2415) Solve() (string, error) {
	var number int
	if s.Part != 2 {
		number = s.processPart1()
	} else {
		number = s.processPart2()
	}
	return strconv.Itoa(number), nil
}

func (s *s2415) processPart1() int {
	sum := 0
	grid := make(map[[2]int]rune)
	moves := ""
	first := true
	maxY, maxX, ry, rx := 0, 0, 0, 0
	for y, line := range s.Input {
		if line == "" {
			first = false
		} else if first {
			for x, r := range line {
				if r == '@' {
					ry = y
					rx = x
					r = '.'
				}
				grid[[2]int{y, x}] = r
				if x > maxX {
					maxX = x
				}
			}
			if y > maxY {
				maxY = y
			}
		} else {
			moves += line
		}
	}
	if s.Debug {
		fmt.Println(grid)
		fmt.Println(moves)
		fmt.Println(ry, rx)
	}

	// process all moves
	directions := make(map[rune][2]int)
	directions['>'] = [2]int{0, 1}
	directions['v'] = [2]int{1, 0}
	directions['<'] = [2]int{0, -1}
	directions['^'] = [2]int{-1, 0}
	if s.Debug {
		s.print(grid, maxY, maxX, ry, rx)
	}
	for _, m := range moves {
		target := [2]int{ry + directions[m][0], rx + directions[m][1]}
		switch grid[target] {
		case '.':
			ry = target[0]
			rx = target[1]
		case 'O':
			do := true
			move := false
			var boxes [][2]int
			boxes = append(boxes, target)
			for do {
				newTarget := [2]int{target[0] + directions[m][0], target[1] + directions[m][1]}
				switch grid[newTarget] {
				case '#':
					do = false
				case '.':
					do = false
					move = true
				case 'O':
					target = newTarget
					boxes = append(boxes, target)
				}
			}
			if move {
				for i := len(boxes) - 1; i >= 0; i-- {
					target = [2]int{boxes[i][0] + directions[m][0], boxes[i][1] + directions[m][1]}
					grid[target] = 'O'
				}
				target = [2]int{ry + directions[m][0], rx + directions[m][1]}
				grid[target] = '.'
				ry = target[0]
				rx = target[1]
			}
		}
		if s.Debug {
			s.print(grid, maxY, maxX, ry, rx)
		}
	}

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if grid[[2]int{y, x}] == 'O' {
				sum += y*100 + x
			}
		}
	}
	return sum
}

func (s *s2415) processPart2() int {
	sum := 0
	grid := make(map[[2]int]rune)
	moves := ""
	first := true
	maxY, maxX, ry, rx := 0, 0, 0, 0
	for y, line := range s.Input {
		if line == "" {
			first = false
		} else if first {
			for baseX, r := range line {
				x := baseX * 2
				if r == '@' {
					ry = y
					rx = x
					r = '.'
				}
				r2 := r
				if r == 'O' {
					r = '['
					r2 = ']'
				}
				grid[[2]int{y, x}] = r
				grid[[2]int{y, x + 1}] = r2
				if x > maxX {
					maxX = x + 1
				}
			}
			if y > maxY {
				maxY = y
			}
		} else {
			moves += line
		}
	}
	if s.Debug {
		fmt.Println(grid)
		fmt.Println(moves)
		fmt.Println(ry, rx)
	}

	// process all moves
	directions := make(map[rune][2]int)
	directions['>'] = [2]int{0, 1}
	directions['v'] = [2]int{1, 0}
	directions['<'] = [2]int{0, -1}
	directions['^'] = [2]int{-1, 0}
	if s.Debug {
		s.print(grid, maxY, maxX, ry, rx)
	}
	for mi, m := range moves {
		target := [2]int{ry + directions[m][0], rx + directions[m][1]}
		if grid[target] == '.' {
			ry = target[0]
			rx = target[1]
		} else if grid[target] == '[' || grid[target] == ']' {
			boxes := make(map[[2]int]bool)
			// we use track the left half of the boxes in the map
			if grid[target] == ']' {
				target[1]--
			}
			boxes[target] = true
			oldLength := 0
			// add touching boxes until we don't add any more
			for len(boxes) > oldLength {
				oldLength = len(boxes)
				for t := range boxes {
					newTargetLeft := [2]int{t[0] + directions[m][0], t[1] + directions[m][1]}
					newTargetRight := [2]int{t[0] + directions[m][0], t[1] + directions[m][1] + 1}
					if grid[newTargetLeft] == '[' || grid[newTargetLeft] == ']' {
						if grid[newTargetLeft] == ']' {
							newTargetLeft[1]--
						}
						boxes[newTargetLeft] = true
					}
					if grid[newTargetRight] == '[' || grid[newTargetRight] == ']' {
						if grid[newTargetRight] == ']' {
							newTargetRight[1]--
						}
						boxes[newTargetRight] = true
					}
				}
			}

			// check if any boxes would move into a wall
			move := true
			for t := range boxes {
				newTargetLeft := [2]int{t[0] + directions[m][0], t[1] + directions[m][1]}
				newTargetRight := [2]int{t[0] + directions[m][0], t[1] + directions[m][1] + 1}
				if grid[newTargetLeft] == '#' || grid[newTargetRight] == '#' {
					move = false
				}
			}

			if move {
				// make a copy of the grid
				newGrid := make(map[[2]int]rune)
				for t, r := range grid {
					newGrid[t] = r
				}
				// blank out all the current box coordinates on the new grid
				if s.Debug {
					fmt.Println(boxes)
				}
				for t := range boxes {
					newGrid[t] = '.'
					newGrid[[2]int{t[0], t[1] + 1}] = '.'
				}
				// move all the boxes to their new destination
				for t := range boxes {
					newTargetLeft := [2]int{t[0] + directions[m][0], t[1] + directions[m][1]}
					newTargetRight := [2]int{t[0] + directions[m][0], t[1] + directions[m][1] + 1}
					newGrid[newTargetLeft] = '['
					newGrid[newTargetRight] = ']'
				}
				target = [2]int{ry + directions[m][0], rx + directions[m][1]}
				newGrid[target] = '.'
				ry = target[0]
				rx = target[1]
				grid = newGrid
			}
		}
		if s.Debug && mi < 10 {
			s.print(grid, maxY, maxX, ry, rx)
		}
	}

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if grid[[2]int{y, x}] == '[' {
				sum += y*100 + x
			}
		}
	}
	return sum
}

func (s *s2415) print(grid map[[2]int]rune, maxY, maxX, ry, rx int) {
	fmt.Println("")
	for y := 0; y <= maxY; y++ {
		a := ""
		for x := 0; x <= maxX; x++ {
			if y == ry && x == rx {
				a += "@"
			} else {
				a += string(grid[[2]int{y, x}])
			}
		}
		fmt.Println(a)
	}
}
