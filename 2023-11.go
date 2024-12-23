package main

import (
	"fmt"
	"strconv"
)

type s2311 Puzzle

func (s *s2311) SetDebug(debug bool) error {
	s.Debug = debug
	return nil
}

func (s *s2311) SetInput(input []string) error {
	s.Input = input
	return nil
}

func (s *s2311) SetPart(part int) error {
	s.Part = part
	return nil
}

func (s *s2311) Solve() (string, error) {
	var number int
	if s.Part != 2 {
		number = s.processPart1()
	} else {
		number = s.processPart2()
	}
	return strconv.Itoa(number), nil
}

func (s *s2311) processPart1() int {
	lines := MakeGridRune(s.Input)

	universe := make(map[int]map[int]bool)
	universe2 := make(map[int]map[int]bool)

	if s.Debug {
		fmt.Println("Original Universe")
		for _, line := range s.Input {
			fmt.Println(line)
		}
	}

	// first check for columns that need to be expanded
	index := 0
	// iterate through the columns
	for i := 0; i < len(lines[0]); i++ {
		// set empty flag, trip if any cell contains #
		isEmpty := true
		// iterate through the rows
		for j := 0; j < len(lines); j++ {
			if i == 0 {
				universe[j] = make(map[int]bool)
			}
			if lines[j][i] == '#' {
				isEmpty = false
			}
		}
		for j := 0; j < len(lines); j++ {
			universe[j][index] = lines[j][i] == '#'
		}
		index++
		if isEmpty {
			for j := 0; j < len(lines); j++ {
				universe[j][index] = lines[j][i] == '#'
			}
			index++
		}
	}

	if s.Debug {
		fmt.Println("Horizontally Expanded Universe")
		for x := 0; x < len(universe); x++ {
			debugString := ""
			for y := 0; y < len(universe[x]); y++ {
				if universe[x][y] {
					debugString += "#"
				} else {
					debugString += "."
				}
			}
			fmt.Println(debugString)
		}
	}

	// now check for rows that need to be expanded
	index = 0
	// iterate through the rows
	for x := 0; x < len(universe); x++ {
		// set empty flag, trip if any cell contains #
		isEmpty := true
		// iterate through the columns
		for y := 0; y < len(universe[x]); y++ {
			if universe[x][y] {
				isEmpty = false
			}
		}
		for y := 0; y < len(universe[x]); y++ {
			if index > len(universe2)-1 {
				universe2[index] = make(map[int]bool)
			}
			universe2[index][y] = universe[x][y]
		}
		index++
		if isEmpty {
			for y := 0; y < len(universe[x]); y++ {
				if index > len(universe2)-1 {
					universe2[index] = make(map[int]bool)
				}
				universe2[index][y] = universe[x][y]
			}
			index++
		}
	}

	if s.Debug {
		fmt.Println("Vertically Expanded Universe")
		for x := 0; x < len(universe2); x++ {
			debugString := ""
			for y := 0; y < len(universe2[x]); y++ {
				if universe2[x][y] {
					debugString += "#"
				} else {
					debugString += "."
				}
			}
			fmt.Println(debugString)
		}
	}

	// get coordinates of each galaxy
	coordinates := [][2]int{}
	for x := 0; x < len(universe2); x++ {
		for y := 0; y < len(universe2[x]); y++ {
			if universe2[x][y] {
				coordinates = append(coordinates, [2]int{x, y})
			}
		}
	}

	sum := 0
	for c, i := range coordinates {
		for d, j := range coordinates {
			if d > c {
				sum += abs(i[0]-j[0]) + abs(i[1]-j[1])
			}
		}
	}

	return sum
}

func (s *s2311) processPart2() int {
	// this could be used to solve Part 1 by changing multiplier to 2
	lines := MakeGridRune(s.Input)
	multiplier := 1000000
	rows := []bool{}
	cols := []bool{}

	if s.Debug {
		fmt.Println("Original Universe")
		for _, line := range s.Input {
			fmt.Println(line)
		}
	}

	// first check for columns that need to be expanded
	// iterate through the columns
	for i := 0; i < len(lines[0]); i++ {
		// set empty flag, trip if any cell contains #
		isEmpty := true
		// iterate through the rows
		for j := 0; j < len(lines); j++ {
			if lines[j][i] == '#' {
				isEmpty = false
			}
		}
		cols = append(cols, isEmpty)
	}

	// now check for rows that need to be expanded
	// iterate through the rows
	for x := 0; x < len(lines); x++ {
		// set empty flag, trip if any cell contains #
		isEmpty := true
		// iterate through the columns
		for y := 0; y < len(lines[x]); y++ {
			if lines[x][y] == '#' {
				isEmpty = false
			}
		}
		rows = append(rows, isEmpty)
	}

	if s.Debug {
		fmt.Println(cols)
		fmt.Println(rows)
	}

	// get coordinates of each galaxy
	coordinates := [][2]int{}
	for x := 0; x < len(lines); x++ {
		for y := 0; y < len(lines[x]); y++ {
			if lines[x][y] == '#' {
				coordinates = append(coordinates, [2]int{x, y})
			}
		}
	}

	sum := 0
	for c, i := range coordinates {
		for d, j := range coordinates {
			if d > c {
				// walk horizontally and then vertically through the universe, using the multiplier on empty rows/cols
				if i[0] < j[0] {
					for a := i[0]; a < j[0]; a++ {
						if rows[a] {
							sum += multiplier
						} else {
							sum++
						}
					}
				}
				if j[0] < i[0] {
					for a := j[0]; a < i[0]; a++ {
						if rows[a] {
							sum += multiplier
						} else {
							sum++
						}
					}
				}
				if i[1] < j[1] {
					for a := i[1]; a < j[1]; a++ {
						if cols[a] {
							sum += multiplier
						} else {
							sum++
						}
					}
				}
				if j[1] < i[1] {
					for a := j[1]; a < i[1]; a++ {
						if cols[a] {
							sum += multiplier
						} else {
							sum++
						}
					}
				}
			}
		}
	}

	return sum
}
