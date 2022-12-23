package main

func s22081(filename string) int {
	// expected result: 1835
	grid := ReadGridInt("./data/2022/08/" + filename)
	count := 0
	for y, row := range grid {
		for x, a := range row {
			// outside trees
			if y == 0 || y == len(grid)-1 || x == 0 || x == len(row)-1 {
				count++
			} else {
				up := true
				for i := x - 1; i >= 0; i-- {
					if grid[y][i] >= a {
						up = false
					}
				}
				down := true
				for i := x + 1; i < len(row); i++ {
					if grid[y][i] >= a {
						down = false
					}
				}
				left := true
				for i := y - 1; i >= 0; i-- {
					if grid[i][x] >= a {
						left = false
					}
				}
				right := true
				for i := y + 1; i < len(grid); i++ {
					if grid[i][x] >= a {
						right = false
					}
				}
				if up || down || left || right {
					count++
				}
			}
		}
	}
	return count
}

func s22082(filename string) int {
	// expected result: 263670
	grid := ReadGridInt("./data/2022/08/" + filename)
	count := 0
	for y, row := range grid {
		for x, a := range row {
			// skip if any are on the edge of the grid
			if y != 0 && y != len(grid)-1 && x != 0 && x != len(row)-1 {
				up := 0
				for i := x - 1; i >= 0; i-- {
					up++
					if grid[y][i] >= a {
						break
					}
				}
				down := 0
				for i := x + 1; i < len(row); i++ {
					down++
					if grid[y][i] >= a {
						break
					}
				}
				left := 0
				for i := y - 1; i >= 0; i-- {
					left++
					if grid[i][x] >= a {
						break
					}
				}
				right := 0
				for i := y + 1; i < len(grid); i++ {
					right++
					if grid[i][x] >= a {
						break
					}
				}
				score := up * down * left * right
				if score > count {
					count = score
				}
			}
		}
	}
	return count
}
