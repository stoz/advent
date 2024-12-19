package main

import (
	"fmt"
	"strconv"
)

type s2416 Puzzle

type tile2416 struct {
	c int
	// y, x, direction, cost to get there
	n       [][4]int
	v       bool
	parents [][3]int
}

func (s *s2416) SetDebug(debug bool) error {
	s.Debug = debug
	return nil
}

func (s *s2416) SetInput(input []string) error {
	s.Input = input
	return nil
}

func (s *s2416) SetPart(part int) error {
	s.Part = part
	return nil
}

func (s *s2416) Solve() (string, error) {
	var number int
	if s.Part != 2 {
		number = s.processPart1()
	} else {
		number = s.processPart2()
	}
	return strconv.Itoa(number), nil
}

func (s *s2416) processPart1() int {
	// read the input and transform it into a map keyed off y, x and the direction being faced when entered and the cost to enter
	grid := MakeGridRune(s.Input)
	g := make(map[[3]int]tile2416)
	start := [3]int{}
	target := [2]int{}
	// cost to enter from the 4 directions, then the actual movement offsets
	directions := [][6]int{
		{1, 1001, 2001, 1001, 0, 1},
		{1001, 1, 1001, 2001, 1, 0},
		{2001, 1001, 1, 1001, 0, -1},
		{1001, 2001, 1001, 1, -1, 0},
	}
	yMax := len(grid)
	xMax := len(grid[0])
	for y, line := range grid {
		for x, r := range line {
			switch r {
			case 'S':
				start = [3]int{y, x, 0}
			case 'E':
				target = [2]int{y, x}
			}
			if r == '.' || r == 'S' {
				for di := range directions {
					var nbs [][4]int
					for di2, d2 := range directions {
						directionRune := grid[y+d2[4]][x+d2[5]]
						if directionRune == '.' || directionRune == 'E' {
							nbs = append(nbs, [4]int{y + d2[4], x + d2[5], di2, d2[di]})
						}
					}
					tile := tile2416{c: 999999, n: nbs}
					g[[3]int{y, x, di}] = tile
				}
			}
		}
	}
	// set start pointing east to cost zero
	g[start] = tile2416{c: 0, n: g[start].n}
	if s.Debug {
		fmt.Println(g[start])
		fmt.Println(g[[3]int{13, 2, 0}])
		fmt.Println(g[[3]int{12, 1, 0}])
	}

	result, modGrid := s.dijkstraPart1(g, target)
	if s.Debug {
		sums := [2]int{0, 0}
		for _, m := range modGrid {
			if m.v {
				sums[0]++
			} else {
				sums[1]++
			}
		}
		for y := 0; y < yMax; y++ {
			line := ""
			for x := 0; x < xMax; x++ {
				a := string(grid[y][x])
				visitCounter := 0
				for d := 0; d < 4; d++ {
					if modGrid[[3]int{y, x, d}].v {
						visitCounter++
					}
				}
				if visitCounter > 0 {
					a = strconv.Itoa(visitCounter)
				}
				line += a
			}
			fmt.Println(line)
		}
		fmt.Println(sums)
	}
	return result
}

func (s *s2416) dijkstraPart1(g map[[3]int]tile2416, target [2]int) (int, map[[3]int]tile2416) {
	if s.Debug {
		fmt.Println("target:", target)
	}
	// data structure a map of [y][x] coordinates to points
	// each point has a cost (c) and 0 or more neighbours (n)
	do := true
	result := 999999
	for do {
		// start from the lowest-cost square
		shortest := 999999
		var cur [3]int
		var point tile2416
		for key, tile := range g {
			if !tile.v && tile.c < shortest {
				point = tile
				shortest = tile.c
				cur = key
			}
		}
		if s.Debug {
			fmt.Println(cur, g[cur])
			if cur[0] == 13 && cur[1] == 1 {
				fmt.Println(point)
				fmt.Println(g[cur])
			}
		}
		if shortest == 999999 {
			do = false
			break
		}

		for _, n := range point.n {
			cost := point.c + n[3]
			if n[0] == target[0] && n[1] == target[1] {
				if s.Debug {
					fmt.Println("found it")
				}
				result = cost
				do = false
			}
			tile, ok := g[[3]int{n[0], n[1], n[2]}]
			if ok {
				if cost < g[[3]int{n[0], n[1], n[2]}].c {
					tile.c = cost
					if s.Debug {
						fmt.Println(cost)
					}
				}
				g[[3]int{n[0], n[1], n[2]}] = tile
			}
		}
		tile := g[cur]
		tile.v = true
		g[cur] = tile
	}
	return result, g
}

func (s *s2416) processPart2() int {
	// read the input and transform it into a map keyed off y, x and the direction being faced when entered and the cost to enter
	grid := MakeGridRune(s.Input)
	g := make(map[[3]int]tile2416)
	start := [3]int{}
	target := [2]int{}
	// cost to enter from the 4 directions, then the actual movement offsets
	directions := [][6]int{
		{1, 1001, 2001, 1001, 0, 1},
		{1001, 1, 1001, 2001, 1, 0},
		{2001, 1001, 1, 1001, 0, -1},
		{1001, 2001, 1001, 1, -1, 0},
	}
	yMax := len(grid)
	xMax := len(grid[0])
	for y, line := range grid {
		for x, r := range line {
			switch r {
			case 'S':
				start = [3]int{y, x, 0}
			case 'E':
				target = [2]int{y, x}
			}
			if r == '.' || r == 'S' || r == 'E' {
				for di := range directions {
					var nbs [][4]int
					for di2, d2 := range directions {
						directionRune := grid[y+d2[4]][x+d2[5]]
						if directionRune == '.' || directionRune == 'E' {
							nbs = append(nbs, [4]int{y + d2[4], x + d2[5], di2, d2[di]})
						}
					}
					tile := tile2416{c: 999999, n: nbs}
					g[[3]int{y, x, di}] = tile
				}
			}
		}
	}
	// set start pointing east to cost zero
	g[start] = tile2416{c: 0, n: g[start].n}
	if s.Debug {
		fmt.Println(g[start])
		fmt.Println(g[[3]int{13, 2, 0}])
		fmt.Println(g[[3]int{12, 1, 0}])
	}

	_, modGrid := s.dijkstraPart2(g, target)

	// start at the end and collect all the parents
	seats := make(map[[2]int]bool)
	search := [][3]int{
		{target[0], target[1], 0},
		//{target[0], target[1], 1},
		//{target[0], target[1], 2},
		//{target[0], target[1], 3},
	}
	for len(search) > 0 {
		cur := search[0]
		search = search[1:]
		tile := modGrid[[3]int{cur[0], cur[1], cur[2]}]
		seats[[2]int{cur[0], cur[1]}] = true
		//fmt.Println("parents", tile)
		search = append(search, tile.parents...)
	}

	if s.Debug {
		superSeatCounter := 0
		sums := [3]int{0, 0, 0}
		for _, m := range modGrid {
			if m.v {
				sums[0]++
			} else {
				sums[1]++
			}
		}
		for y := 0; y < yMax; y++ {
			line := ""
			for x := 0; x < xMax; x++ {
				a := string(grid[y][x])
				visitCounter := 0
				for d := 0; d < 4; d++ {
					if modGrid[[3]int{y, x, d}].v {
						visitCounter++
					}
					sums[2] += len(modGrid[[3]int{y, x, d}].parents)
				}
				if visitCounter > 0 {
					//a = strconv.Itoa(visitCounter)
				}
				if _, ok := seats[[2]int{y, x}]; ok {
					a = "O"
					superSeatCounter++
				}
				line += a
			}
			fmt.Println(line)
		}
		fmt.Println(sums, superSeatCounter)
	}
	return len(seats)
}

func (s *s2416) dijkstraPart2(g map[[3]int]tile2416, target [2]int) (int, map[[3]int]tile2416) {
	if s.Debug {
		fmt.Println("target:", target)
	}
	// data structure a map of [y][x] coordinates to points
	// each point has a cost (c) and 0 or more neighbours (n)
	do := true
	result := 999999
	for do {
		// start from the lowest-cost square
		shortest := 999999
		var cur [3]int
		var point tile2416
		for key, tile := range g {
			if !tile.v && tile.c < shortest {
				point = tile
				shortest = tile.c
				cur = key
			}
		}
		if s.Debug {
			fmt.Println(cur, g[cur])
			if cur[0] == 13 && cur[1] == 1 {
				fmt.Println(point)
				fmt.Println(g[cur])
			}
		}
		if shortest == 999999 {
			do = false
			break
		}

		for _, n := range point.n {
			cost := point.c + n[3]
			if n[0] == target[0] && n[1] == target[1] {
				if s.Debug {
					fmt.Println("found it")
				}
				if cost < result {
					result = cost
				}
				// do = false
			}
			tile, ok := g[[3]int{n[0], n[1], n[2]}]
			if ok {
				if cost < g[[3]int{n[0], n[1], n[2]}].c {
					tile.parents = [][3]int{{cur[0], cur[1], cur[2]}}
					tile.c = cost
				} else if cost == g[[3]int{n[0], n[1], n[2]}].c {
					if s.Debug {
						fmt.Println("matched cost", cost)
					}
					tile.parents = append(tile.parents, [3]int{cur[0], cur[1], cur[2]})
				}
				g[[3]int{n[0], n[1], n[2]}] = tile
			}
		}
		tile := g[cur]
		tile.v = true
		g[cur] = tile
	}
	return result, g
}
