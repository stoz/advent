package main

import (
	"fmt"
	"strconv"
)

func s2214(filename string, part2, printCave bool) int {
	// expected results: 808, 26625
	lines := ReadFile("./data/2022/14/" + filename)
	cave := make(map[int]map[int]int)

	// get the dimensions
	var dim [3]int
	for _, line := range lines {
		ints := ExtractInts(line)
		for i, j := range ints {
			switch i % 2 {
			case 0: //x
				if j < dim[1] || dim[1] == 0 {
					dim[1] = j
				} else if j > dim[2] {
					dim[2] = j
				}
			case 1: //y
				if j > dim[0] {
					dim[0] = j
				}
			}
		}
	}
	if part2 {
		// add a floor 2 units below the lowest point
		dim[0] += 2
		dim[1] = 499 - dim[0]
		dim[2] = 501 + dim[0]
		lines = append(lines, strconv.Itoa(dim[1])+" "+strconv.Itoa(dim[0])+" "+strconv.Itoa(dim[2])+" "+strconv.Itoa(dim[0]))
	}
	for i := 0; i <= dim[0]; i++ {
		cave[i] = make(map[int]int)
	}
	for _, line := range lines {
		ints := ExtractInts(line)
		var s [2]int // start
		var x int
		for i, j := range ints {
			switch i % 2 {
			case 0:
				x = j
			case 1:
				if s[0] != 0 {
					// do the needful
					if s[0] < j {
						for v := s[0]; v <= j; v++ {
							cave[v][x] = 1
						}
					} else if j < s[0] {
						for v := j; v <= s[0]; v++ {
							cave[v][x] = 1
						}
					} else if s[1] < x {
						for v := s[1]; v <= x; v++ {
							cave[j][v] = 1
						}
					} else if x < s[1] {
						for v := x; v <= s[1]; v++ {
							cave[j][v] = 1
						}
					}
				}
				s[0] = j
				s[1] = x
			}
		}
	}
	do := true
	count := 0
	for do {
		y := 0
		x := 500
		do2 := true
		for do2 {
			if y > dim[0] {
				do2 = false
				do = false
			}
			_, down := cave[y+1][x]
			if down {
				_, left := cave[y+1][x-1]
				if left {
					_, right := cave[y+1][x+1]
					if right {
						count++
						if y > 0 {
							cave[y][x] = 2
						} else {
							do = false
						}
						do2 = false
					} else {
						y++
						x++
					}
				} else {
					y++
					x--
				}
			} else {
				y++
			}
		}
	}
	if printCave {
		PrintCave(cave, dim)
	}
	return count
}

func PrintCave2(cave map[int]map[int]int, dim [4]int) {
	for y := dim[0]; y <= dim[1]; y++ {
		var line string
		for x := dim[2]; x <= dim[3]; x++ {
			c := "?"
			i, ok := cave[y][x]
			if ok {
				switch i {
				case 1:
					c = "#"
				case 2:
					c = "o"
				case 3:
					c = "S"
				case 4:
					c = "B"
				}
			} else {
				c = "."
			}
			line += c
		}
		fmt.Println(line, y)
	}
}

func PrintCave(cave map[int]map[int]int, dim [3]int) {
	for y := 0; y <= dim[0]; y++ {
		var line string
		for x := dim[1]; x <= dim[2]; x++ {
			c := "?"
			i, ok := cave[y][x]
			if ok {
				switch i {
				case 1:
					c = "#"
				case 2:
					c = "o"
				case 3:
					c = "S"
				case 4:
					c = "B"
				}
			} else {
				c = "."
			}
			line += c
		}
		fmt.Println(line)
	}
}
