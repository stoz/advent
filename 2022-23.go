package main

import "fmt"

func s22231(filename string, debug bool) int {
	lines := ReadGridRune("./data/2022/23/" + filename)
	// padding
	var og [][]rune
	pad := 100
	wid := len(lines[0]) + pad*2
	var g [][]rune
	for i := 0; i < pad; i++ {
		var buf []rune
		for j := 0; j < wid; j++ {
			buf = append(buf, '.')
		}
		g = append(g, buf)
	}
	for _, line := range lines {
		var buf []rune
		for i := 0; i < pad; i++ {
			buf = append(buf, '.')
		}
		for _, r := range line {
			buf = append(buf, r)
		}
		for i := 0; i < pad; i++ {
			buf = append(buf, '.')
		}
		g = append(g, buf)
	}
	for i := 0; i < pad; i++ {
		var buf []rune
		for j := 0; j < wid; j++ {
			buf = append(buf, '.')
		}
		g = append(g, buf)
	}
	dir := [4]int{3, 1, 2, 0}
	for s := 0; s < 10; s++ {
		var pro [][5]int
		for y, line := range g {
			for x, r := range line {
				if r == '#' {
					// check for any other elves
					if g[y-1][x-1] == '#' || g[y-1][x] == '#' || g[y-1][x+1] == '#' || g[y][x-1] == '#' || g[y][x+1] == '#' || g[y+1][x-1] == '#' || g[y+1][x] == '#' || g[y+1][x+1] == '#' {
						// propose a move
						var proposed bool
						for _, d := range dir {
							switch d {
							case 0:
								if !proposed && g[y-1][x+1] != '#' && g[y][x+1] != '#' && g[y+1][x+1] != '#' {
									p := [5]int{y, x, y, x + 1}
									pro = append(pro, p)
									proposed = true
								}
							case 1:
								if !proposed && g[y+1][x-1] != '#' && g[y+1][x] != '#' && g[y+1][x+1] != '#' {
									p := [5]int{y, x, y + 1, x}
									pro = append(pro, p)
									proposed = true
								}
							case 2:
								if !proposed && g[y-1][x-1] != '#' && g[y][x-1] != '#' && g[y+1][x-1] != '#' {
									p := [5]int{y, x, y, x - 1}
									pro = append(pro, p)
									proposed = true
								}
							case 3:
								if !proposed && g[y-1][x-1] != '#' && g[y-1][x] != '#' && g[y-1][x+1] != '#' {
									p := [5]int{y, x, y - 1, x}
									pro = append(pro, p)
									proposed = true
								}
							}
						}
					}
				}
			}
		}
		// iterate through proposals
		for i, p := range pro {
			safe := true
			for j, c := range pro {
				if i != j && p[2] == c[2] && p[3] == c[3] {
					safe = false
				}
			}
			if safe {
				pro[i][4] = 1
			}
		}
		//fmt.Println(pro)
		for _, p := range pro {
			if p[4] == 1 {
				g[p[0]][p[1]] = '.'
			}
		}
		for _, p := range pro {
			if p[4] == 1 {
				g[p[2]][p[3]] = '#'
			}
		}
		// increment dir
		// int{3, 1, 2, 0}
		switch dir[0] {
		case 0:
			dir[0], dir[1], dir[2], dir[3] = 3, 1, 2, 0
		case 1:
			dir[0], dir[1], dir[2], dir[3] = 2, 0, 3, 1
		case 2:
			dir[0], dir[1], dir[2], dir[3] = 0, 3, 1, 2
		case 3:
			dir[0], dir[1], dir[2], dir[3] = 1, 2, 0, 3
		}
	}
	// print
	for _, y := range g {
		var buf string
		for _, r := range y {
			buf += string(r)
		}
		fmt.Println(buf)
	}
	var sum int
	for _, line := range g {
		for _, r := range line {
			if r == '.' {
				sum++
			}
		}
	}
	return sum
}

func s22232(filename string, debug bool) int {
	return 2
}
