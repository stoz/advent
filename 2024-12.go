package main

import (
	"fmt"
	"strconv"
)

type s2412 Puzzle

func (s *s2412) SetDebug(debug bool) error {
	s.Debug = debug
	return nil
}

func (s *s2412) SetInput(input []string) error {
	s.Input = input
	return nil
}

func (s *s2412) SetPart(part int) error {
	s.Part = part
	return nil
}

func (s *s2412) Solve() (string, error) {
	var number int
	if s.Part != 2 {
		number = s.processPart1()
	} else {
		number = s.processPart2()
	}
	return strconv.Itoa(number), nil
}

func (s *s2412) processPart1() int {
	sum := 0
	grid := MakeGridRune(s.Input)
	directions := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	yMax, xMax := len(grid), len(grid[0])
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] != ' ' {
				target := grid[y][x]
				grid[y][x] = '#'
				if s.Debug {
					fmt.Println(string(target))
				}
				// flood fill with '#'
				sources := [][2]int{{y, x}}
				for len(sources) > 0 {
					t := sources[len(sources)-1]
					sources = sources[:len(sources)-1]
					for _, d := range directions {
						ny, nx := t[0]+d[0], t[1]+d[1]
						if ny >= 0 && ny < yMax && nx >= 0 && nx < xMax && grid[ny][nx] == target {
							grid[ny][nx] = '#'
							sources = append(sources, [2]int{ny, nx})
						}
					}
				}

				// count '#'
				area := 0
				perimeter := 0
				for yy := 0; yy < len(grid); yy++ {
					for xx := 0; xx < len(grid[yy]); xx++ {
						if grid[yy][xx] == '#' {
							area++
							for _, d := range directions {
								if yy+d[0] >= 0 && yy+d[0] < yMax && xx+d[1] >= 0 && xx+d[1] < xMax {
									if grid[yy+d[0]][xx+d[1]] != '#' {
										perimeter++
									}
								} else {
									// off edge of grid, so needs a fence
									perimeter++
								}
							}
						}
					}
				}
				sum += area * perimeter

				// un-fill grid
				for yy := 0; yy < len(grid); yy++ {
					debugString := ""
					for xx := 0; xx < len(grid[yy]); xx++ {
						if grid[yy][xx] == '#' {
							grid[yy][xx] = ' '
						}
						debugString += string(grid[yy][xx])
					}
					if s.Debug {
						fmt.Println(debugString)
					}
				}

			}
		}
	}
	return sum
}

func (s *s2412) processPart2() int {
	sum := 0
	grid := MakeGridRune(s.Input)
	directions := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	yMax, xMax := len(grid), len(grid[0])
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] != ' ' {
				target := grid[y][x]
				grid[y][x] = '#'
				if s.Debug {
					fmt.Println(string(target))
				}
				// flood fill with '#'
				sources := [][2]int{{y, x}}
				for len(sources) > 0 {
					t := sources[len(sources)-1]
					sources = sources[:len(sources)-1]
					for _, d := range directions {
						ny, nx := t[0]+d[0], t[1]+d[1]
						if ny >= 0 && ny < yMax && nx >= 0 && nx < xMax && grid[ny][nx] == target {
							grid[ny][nx] = '#'
							sources = append(sources, [2]int{ny, nx})
						}
					}
				}

				// count '#'
				area := 0
				// setup 4 grids, one for each direction, that we can use to track which side the perimeter is part of
				var sides [4]map[int]map[int]rune
				for d := range directions {
					sides[d] = make(map[int]map[int]rune)
					for yy := -2; yy < len(grid)+2; yy++ {
						sides[d][yy] = make(map[int]rune)
					}
				}
				for yy := 0; yy < len(grid); yy++ {
					for xx := 0; xx < len(grid[yy]); xx++ {
						if grid[yy][xx] == '#' {
							area++
							for di, d := range directions {
								if yy+d[0] >= 0 && yy+d[0] < yMax && xx+d[1] >= 0 && xx+d[1] < xMax {
									if grid[yy+d[0]][xx+d[1]] != '#' {
										sides[di][yy+d[0]][xx+d[1]] = '#'
									}
								} else {
									// off edge of grid, so needs a fence
									sides[di][yy+d[0]][xx+d[1]] = '#'
								}
							}
						}
					}
				}

				// count distinct sides
				perimeter := 0
				if target == 'R' {
					for d := range directions {
						for yy := -2; yy < yMax+2; yy++ {
							debugString := ""
							for xx := -2; xx < xMax+2; xx++ {
								c := " "
								if sides[d][yy][xx] == '#' {
									c = "#"
								}
								debugString += c
							}
							if s.Debug {
								fmt.Println(debugString)
							}
						}
					}
				}
				for _, side := range sides {
					perimeter += s.countSides(side, yMax, xMax)
				}
				if s.Debug {
					fmt.Println(area, perimeter)
				}
				sum += area * perimeter

				// un-fill grid
				for yy := 0; yy < len(grid); yy++ {
					for xx := 0; xx < len(grid[yy]); xx++ {
						if grid[yy][xx] == '#' {
							grid[yy][xx] = ' '
						}
					}
				}

			}
		}
	}
	return sum
}

func (s *s2412) countSides(sides map[int]map[int]rune, yMax int, xMax int) int {
	directions := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	sum := 0
	for y := -1; y <= yMax; y++ {
		for x := -1; x <= xMax; x++ {
			if sides[y][x] == '#' {
				target := sides[y][x]
				sides[y][x] = '.'
				// flood fill with '.'
				sources := [][2]int{{y, x}}
				for len(sources) > 0 {
					t := sources[len(sources)-1]
					sources = sources[:len(sources)-1]
					for _, d := range directions {
						ny, nx := t[0]+d[0], t[1]+d[1]
						if sides[ny][nx] == target {
							sides[ny][nx] = '.'
							sources = append(sources, [2]int{ny, nx})
						}
					}
				}

				// count '#'
				sum++

				// un-fill grid
				for yy := -1; yy <= yMax; yy++ {
					for xx := -1; xx <= xMax; xx++ {
						if sides[yy][xx] == '.' {
							sides[yy][xx] = ' '
						}
					}
				}
			}
		}
	}
	return sum
}
