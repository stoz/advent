package main

func s2224(filename string, part2 bool, debug bool) int {
	// expected results: 264, 789
	g := ReadGridRune("./data/2022/24/" + filename)
	var state [][2]int8
	start := [2]int8{0, 1}
	state = append(state, start)
	var bz [][3]int8 // blizzards
	var wid int
	for y, line := range g {
		wid = len(line)
		for x, r := range line {
			switch r {
			case '>':
				app := [3]int8{int8(y), int8(x), 0}
				bz = append(bz, app)
			case 'v':
				app := [3]int8{int8(y), int8(x), 1}
				bz = append(bz, app)
			case '<':
				app := [3]int8{int8(y), int8(x), 2}
				bz = append(bz, app)
			case '^':
				app := [3]int8{int8(y), int8(x), 3}
				bz = append(bz, app)
			}
		}
	}
	hig := len(g)
	try := [5][2]int8{{0, 0}, {-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	var count, stage int
	for {
		count++
		var new [27][122]bool
		for _, s := range state {
			for _, t := range try {
				if s[0]+t[0] >= 0 && s[0]+t[0] < int8(hig) && g[s[0]+t[0]][s[1]+t[1]] != '#' {
					new[s[0]+t[0]][s[1]+t[1]] = true
				}
			}
		}
		if (!part2 || stage == 2) && new[hig-1][wid-2] {
			return count
		}
		if part2 && stage == 1 && new[0][1] {
			stage = 2
			for yy, row := range new {
				for xx := range row {
					new[yy][xx] = false
				}
			}
			new[0][1] = true
		}
		if part2 && stage == 0 && new[hig-1][wid-2] {
			stage = 1
			for yy, row := range new {
				for xx := range row {
					new[yy][xx] = false
				}
			}
			new[hig-1][wid-2] = true
		}
		for i, b := range bz {
			switch b[2] {
			case 0:
				if g[b[0]][b[1]+1] == '#' {
					new[b[0]][1] = false
					bz[i][1] = 1
				} else {
					new[b[0]][b[1]+1] = false
					bz[i][1] = b[1] + 1
				}
			case 1:
				if g[b[0]+1][b[1]] == '#' {
					new[1][b[1]] = false
					bz[i][0] = 1
				} else {
					new[b[0]+1][b[1]] = false
					bz[i][0] = b[0] + 1
				}
			case 2:
				if g[b[0]][b[1]-1] == '#' {
					new[b[0]][wid-2] = false
					bz[i][1] = int8(wid) - 2
				} else {
					new[b[0]][b[1]-1] = false
					bz[i][1] = b[1] - 1
				}
			case 3:
				if g[b[0]-1][b[1]] == '#' {
					new[hig-2][b[1]] = false
					bz[i][0] = int8(hig) - 2
				} else {
					new[b[0]-1][b[1]] = false
					bz[i][0] = b[0] - 1
				}
			}
		}
		var swap [][2]int8
		for yy, row := range new {
			for xx, b := range row {
				if b {
					app := [2]int8{int8(yy), int8(xx)}
					swap = append(swap, app)
				}
			}
		}
		state = swap
	}
}
