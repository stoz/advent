package main

func s22171(filename string) int {
	// expected result: 3159
	line := ReadLine("./data/2022/17/" + filename)
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
		}

		// iteraite through air
		i++
		if i >= len(line) {
			i = 0
		}
	}
	return len(grid)
}

func s22172(filename string) int {
	// expected result: 1566272189352
	line := ReadLine("./data/2022/17/" + filename)
	patternRocks, patternOffset := s22172find(filename)
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
	var rocks, y, x, i, offset int
	shape := -1
	target := 1000000000000
	for rocks < 281 {
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
		}

		// iteraite through air
		i++
		if i >= len(line) {
			i = 0
		}
	}
	// pattern: height 2647, rocks 1690
	for rocks+patternRocks < target {
		rocks += patternRocks
		offset += patternOffset
	}

	for rocks < target {
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
		}

		// iteraite through air
		i++
		if i >= len(line) {
			i = 0
		}
	}
	return len(grid) + offset
}

// return rocks and offset
func s22172find(filename string) (int, int) {
	// expected result: 3159
	line := ReadLine("./data/2022/17/" + filename)
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
	var rocks, y, x, i, offset int
	var states [][4]int
	shape := -1
	for rocks < 2022 {
		if shape == -1 {
			shape = rocks % 5
			y = len(grid) + 3
			x = 2

			// check for full lines
			var ng [][7]int
			var fullLine bool
			for yy, gg := range grid {
				if fullLine {
					ng = append(ng, gg)
				}
				if gg[0] == 1 && gg[1] == 1 && gg[2] == 1 && gg[3] == 1 && gg[4] == 1 && gg[5] == 1 && gg[6] == 1 {
					fullLine = true
					offset += yy + 1
				}
			}
			if fullLine {
				app := [4]int{shape, i, offset, rocks}
				states = append(states, app)
				grid = ng
				y = len(grid) + 3
			}
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
		}

		// iteraite through air
		i++
		if i >= len(line) {
			i = 0
		}
	}
	for _, s := range states {
		for _, t := range states {
			if s[0] == t[0] && s[1] == t[1] && s[2] != t[2] {
				return t[3] - s[3], t[2] - s[2]
			}
		}
	}
	return 0, 0
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
