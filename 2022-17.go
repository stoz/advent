package main

import "fmt"

func z2() {
	// expected result: 3159
	// repeating pattern input 2647 = 4148 lines
	// repeating pattern sample 106 = 166 lines
	// 1000000000000
	// 1566037735850
	// 1514285714288 // real answer
	line := ReadLine("input.txt")
	var grid [][7]int
	var pixels [5][5][2]int
	// horizontal
	pixels[0][1][1] = 1
	pixels[0][2][1] = 2
	pixels[0][3][1] = 3
	// plus
	pixels[1][0][0] = 0
	pixels[1][0][1] = 1
	pixels[1][1][0] = 1
	pixels[1][2][0] = 1
	pixels[1][2][1] = 1
	pixels[1][3][0] = 1
	pixels[1][3][1] = 2
	pixels[1][4][0] = 2
	pixels[1][4][1] = 1
	// l
	pixels[2][1][1] = 1
	pixels[2][2][1] = 2
	pixels[2][3][0] = 1
	pixels[2][3][1] = 2
	pixels[2][4][0] = 2
	pixels[2][4][1] = 2
	// vertical
	pixels[3][1][0] = 1
	pixels[3][2][0] = 2
	pixels[3][3][0] = 3
	// square
	pixels[4][1][0] = 1
	pixels[4][2][1] = 1
	pixels[4][3][0] = 1
	pixels[4][3][1] = 1
	var rocks, y, x, i int
	shape := -1
	var log []int
	sum := (1000000000000 - 1000000000000%245) / 245 * 378
	for rocks < 1000000000000%245 {
		if shape == -1 {
			shape = rocks % 5
			y = len(grid) + 3
			x = 2
		}

		// push
		air := -1
		if line[i] == '>' {
			air = 1
		}
		move := true
		for _, p := range pixels[shape] {
			if isCellTaken(y+p[0], x+p[1]+air, grid) {
				move = false
			}
		}
		if move {
			x += air
		}

		// fall
		move = true
		for _, p := range pixels[shape] {
			if isCellTaken(y+p[0]-1, x+p[1], grid) {
				move = false
			}
		}
		if move && y > 0 {
			y--
		} else {
			// stop
			for _, p := range pixels[shape] {
				if y+p[0] == len(grid) {
					grid = append(grid, [7]int{})
				}
				grid[y+p[0]][x+p[1]] = 1
			}
			log = append(log, y)
			rocks++
			shape = -1
			//fmt.Println(rocks)
		}

		// iteraite through air
		i++
		if i >= len(line) {
			i = 0
		}
	}

	// look for a repeating pattern

	fmt.Println(len(grid) + sum)
}

func z1() {
	// expected result: 3159
	// repeating pattern input 2647 = 4148 lines
	// repeating pattern sample 106 = 166 lines
	// sample pattern might be 7 * 40 (air) * 5 (rocks) = 1400
	line := ReadLine("kelvin.txt")
	var grid [][7]int
	var pixels [5][5][2]int
	// horizontal
	pixels[0][1][1] = 1
	pixels[0][2][1] = 2
	pixels[0][3][1] = 3
	// plus
	pixels[1][0][0] = 0
	pixels[1][0][1] = 1
	pixels[1][1][0] = 1
	pixels[1][2][0] = 1
	pixels[1][2][1] = 1
	pixels[1][3][0] = 1
	pixels[1][3][1] = 2
	pixels[1][4][0] = 2
	pixels[1][4][1] = 1
	// l
	pixels[2][1][1] = 1
	pixels[2][2][1] = 2
	pixels[2][3][0] = 1
	pixels[2][3][1] = 2
	pixels[2][4][0] = 2
	pixels[2][4][1] = 2
	// vertical
	pixels[3][1][0] = 1
	pixels[3][2][0] = 2
	pixels[3][3][0] = 3
	// square
	pixels[4][1][0] = 1
	pixels[4][2][1] = 1
	pixels[4][3][0] = 1
	pixels[4][3][1] = 1
	var rocks, y, x, i int
	shape := -1
	var log []int
	var solid, drops []int
	var cycles [][9]int
	for rocks < 1000000 {
		if shape == -1 {
			shape = rocks % 5
			y = len(grid) + 3
			x = 2
			// store cycle as shape, wind index and depth of all 7 columns relative to top
			var cycle [9]int
			for cyclei := 0; cyclei < 7; cyclei++ {
				var tmpcycle = len(grid)
				for cyclej := 1; cyclej < tmpcycle; cyclej++ {
					if grid[len(grid)-cyclej][cyclei] == 1 {
						tmpcycle = cyclej
						break
					}
				}
				cycle[cyclei] = tmpcycle
				cycles = append(cycles, cycle)
			}
			cycle[7] = shape
			cycle[8] = i
		}

		// push
		air := -1
		if line[i] == '>' {
			air = 1
		}
		move := true
		for _, p := range pixels[shape] {
			if isCellTaken(y+p[0], x+p[1]+air, grid) {
				move = false
			}
		}
		if move {
			x += air
		}

		// fall
		move = true
		for _, p := range pixels[shape] {
			if isCellTaken(y+p[0]-1, x+p[1], grid) {
				move = false
			}
		}
		if move && y > 0 {
			y--
		} else {
			// stop
			for _, p := range pixels[shape] {
				if y+p[0] == len(grid) {
					grid = append(grid, [7]int{})
				}
				grid[y+p[0]][x+p[1]] = 1
			}
			log = append(log, y)
			// represent each "drop" as the shape * final x position * wind index
			drops = append(drops, i*35+x*5+shape)
			rocks++
			shape = -1
			//fmt.Println(rocks)
		}

		// iteraite through air
		i++
		if i >= len(line) {
			i = 0
		}
	}

	for z, g := range grid {
		if g[0] == 1 && g[1] == 1 && g[2] == 1 && g[3] == 1 && g[4] == 1 && g[5] == 1 && g[6] == 1 {
			solid = append(solid, z)
		}
	}

	// look for a repeating pattern
	fmt.Println("Looking for repeating pattern")
	for i := 100; i < 10000; i++ {
		for o := 200; o < i+200; o++ {
			match := true
			for c := 0; c < i; c++ {
				for m := 1; m < 9; m++ {
					if cycles[i+o+c][m] != cycles[i*2+o+c][m] {
						match = false
						break
					}
				}
				if !match {
					break
				}
			}
			if match {
				fmt.Println("Possible match:", i, o, log[i-o+200])
				break
			}
		}
	}

	fmt.Println(len(grid))
	fmt.Println("solid", solid)
	for z := 2500; z > 2490; z-- {
		var out string
		for zz := 0; zz < 7; zz++ {
			if grid[z][zz] == 1 {
				out += "#"
			} else {
				out += "."
			}
		}
		fmt.Println(z, out)
	}
}

func z() {
	// expected result: 3159
	line := ReadLine("sample.txt")
	var grid [][7]int
	var pixels [5][5][2]int
	// horizontal
	pixels[0][1][1] = 1
	pixels[0][2][1] = 2
	pixels[0][3][1] = 3
	// plus
	pixels[1][0][0] = 0
	pixels[1][0][1] = 1
	pixels[1][1][0] = 1
	pixels[1][2][0] = 1
	pixels[1][2][1] = 1
	pixels[1][3][0] = 1
	pixels[1][3][1] = 2
	pixels[1][4][0] = 2
	pixels[1][4][1] = 1
	// l
	pixels[2][1][1] = 1
	pixels[2][2][1] = 2
	pixels[2][3][0] = 1
	pixels[2][3][1] = 2
	pixels[2][4][0] = 2
	pixels[2][4][1] = 2
	// vertical
	pixels[3][1][0] = 1
	pixels[3][2][0] = 2
	pixels[3][3][0] = 3
	// square
	pixels[4][1][0] = 1
	pixels[4][2][1] = 1
	pixels[4][3][0] = 1
	pixels[4][3][1] = 1
	var rocks, y, x, i int
	shape := -1
	for rocks < 2022 {
		if shape == -1 {
			shape = rocks % 5
			y = len(grid) + 3
			x = 2
		}

		// push
		air := -1
		if line[i] == '>' {
			air = 1
		}
		move := true
		for _, p := range pixels[shape] {
			if isCellTaken(y+p[0], x+p[1]+air, grid) {
				move = false
			}
		}
		if move {
			x += air
		}

		// fall
		move = true
		for _, p := range pixels[shape] {
			if isCellTaken(y+p[0]-1, x+p[1], grid) {
				move = false
			}
		}
		if move && y > 0 {
			y--
		} else {
			// stop
			for _, p := range pixels[shape] {
				if y+p[0] == len(grid) {
					grid = append(grid, [7]int{})
				}
				grid[y+p[0]][x+p[1]] = 1
			}
			rocks++
			shape = -1
			//fmt.Println(rocks)
		}

		// iteraite through air
		i++
		if i >= len(line) {
			i = 0
		}
		if len(grid)%1000000 == 0 {
			fmt.Println(len(grid))
		}
	}
	fmt.Println(len(grid))
}

func isCellTaken(y, x int, g [][7]int) bool {
	if x < 0 || x > 6 || y < 0 {
		return true
	}
	if y < len(g) {
		return g[y][x] == 1
	}
	return false
}
