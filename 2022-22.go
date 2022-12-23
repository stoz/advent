package main

import (
	"fmt"
	"strconv"
)

func s22221(filename string, debug bool) int {
	// expected result: 75254
	lines := ReadFile(("./data/2022/22/" + filename))
	input := lines[len(lines)-1]
	lines = lines[:len(lines)-2]
	grid := MakeGridRune(lines)
	var f, y, x int
	overlay := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	h := len(grid) - 1
	for i, r := range grid[0] {
		if r == '.' {
			x = i
			break
		}
	}
	var buf string
	var moves [][2]int
	for _, r := range input {
		if r == 'L' {
			i, _ := strconv.Atoi(buf)
			app := [2]int{i, 0}
			moves = append(moves, app)
			buf = ""
		} else if r == 'R' {
			i, _ := strconv.Atoi(buf)
			app := [2]int{i, 1}
			moves = append(moves, app)
			buf = ""
		} else {
			buf += string(r)
		}
	}
	if buf != "" {
		i, _ := strconv.Atoi(buf)
		app := [2]int{i, -1}
		moves = append(moves, app)
	}
	if debug {
		fmt.Println(y, x, f)
	}

	for c, m := range moves {
		o := overlay[f]
		for i := 0; i < m[0]; i++ {
			yy := y + o[0]
			if yy < 0 {
				yy = h
			} else if yy > h {
				yy = 0
			}
			w := len(grid[yy]) - 1
			xx := x + o[1]
			if xx < 0 {
				xx = w
			} else if xx > w {
				xx = 0
			}
			for grid[yy][xx] == ' ' || grid[yy][xx] == 0 {
				yy += o[0]
				xx += o[1]
				if yy < 0 {
					yy = h
				} else if yy >= h {
					yy = 0
				}
				if xx < 0 {
					xx = w
				} else if xx >= w {
					xx = 0
				}
			}
			if grid[yy][xx] == '.' {
				y, x = yy, xx
			} else if grid[yy][xx] == '#' {
				break
			}
		}
		if m[1] == 0 {
			// left
			if f == 0 {
				f = 3
			} else {
				f--
			}
		} else if m[1] == 1 {
			if f == 3 {
				f = 0
			} else {
				f++
			}
		}
		if debug {
			fmt.Println("|", y, x, f)
			if c > 20 {
				return -1
			}
		}
	}
	return 1000*(y+1) + 4*(x+1) + f
}

func s22222(filename string, debug bool) int {
	// expected result: 108311
	if filename == "sample.txt" {
		// cube mapping does not support sample data
		return -1
	}
	lines := ReadFile(("./data/2022/22/" + filename))
	input := lines[len(lines)-1]
	lines = lines[:len(lines)-2]
	grid := MakeGridRune(lines)
	var f, y, x int
	overlay := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	h := len(grid) - 1
	for i, r := range grid[0] {
		if r == '.' {
			x = i
			break
		}
	}
	var buf string
	var moves [][2]int
	for _, r := range input {
		if r == 'L' {
			i, _ := strconv.Atoi(buf)
			app := [2]int{i, 0}
			moves = append(moves, app)
			buf = ""
		} else if r == 'R' {
			i, _ := strconv.Atoi(buf)
			app := [2]int{i, 1}
			moves = append(moves, app)
			buf = ""
		} else {
			buf += string(r)
		}
	}
	if buf != "" {
		i, _ := strconv.Atoi(buf)
		app := [2]int{i, -1}
		moves = append(moves, app)
	}
	if debug {
		fmt.Println(y, x, f)
	}

	for c, m := range moves {
		for i := 0; i < m[0]; i++ {
			ff := f
			o := overlay[f]
			yy := y + o[0]
			xx := x + o[1]
			// cube transitons
			// cube 1 up
			if y == 0 && x < 100 && f == 3 {
				yy = x + 100
				xx = 0
				ff = 0
			}
			// cube 2 up
			if y == 0 && x > 99 && f == 3 {
				yy = 199
				xx = x - 100
				ff = 3
			}
			// cube 1 left 0-49 -> 100->149
			if y < 50 && x == 50 && f == 2 {
				yy = 149 - y // 99 or 100?
				xx = 0
				ff = 0
			}
			// cube 2 right 0-49 -> 100->149
			if x == 149 && f == 0 {
				yy = 149 - y // 99 or 100?
				xx = 99
				ff = 2
			}
			// cube 2 down
			if y == 49 && x > 99 && f == 1 {
				yy = x - 50
				xx = 99
				f = 2
			}
			// cube 3 left
			if y > 49 && y < 100 && x == 50 && f == 2 {
				yy = 100
				xx = x - 50
				ff = 1
			}
			// cube 3 right
			if y > 49 && y < 100 && x == 99 && f == 0 {
				yy = 49
				xx = y + 50
				ff = 3
			}
			// cube 4 up
			if y == 100 && x < 50 && f == 3 {
				yy = x + 50
				xx = 50
				f = 0
			}
			// cube 4 left 100-149 -> 0-49
			if y < 150 && x == 0 && f == 2 {
				yy = 149 - y // 149 or 150?
				xx = 50
				ff = 0
			}
			// cube 5 right 100-149 -> 49-0
			if y > 99 && y < 150 && x == 99 && f == 0 {
				yy = 149 - y // 149 or 150
				xx = 149
				ff = 2
			}
			// cube 5 down
			if y == 149 && x > 49 && f == 1 {
				yy = x + 100
				xx = 49
				ff = 2
			}
			// cube 6 left 150-199 -> 50-99
			if y > 149 && x == 0 && f == 2 {
				yy = 0
				xx = y - 100
				ff = 1
			}
			// cube 6 right // 150-199 -> 50-99
			if y > 149 && x == 49 && f == 0 {
				yy = 149
				xx = y - 100
				ff = 3
			}
			// cube 6 down
			if y == 199 && f == 1 {
				yy = 0
				xx = x + 100
				ff = 1
			}

			// checks for incorrect cube mapping. Could be deleted
			if yy < 0 {
				fmt.Println("wrap up")
				return -1
			} else if yy > h {
				fmt.Println("wrap down")
				return -1
			}
			w := len(grid[yy]) - 1
			if xx < 0 {
				fmt.Println("wrap left")
				return -1
			} else if xx > w {
				fmt.Println("wrap right", y, x, f, yy, xx, ff)
				return -1
			}
			for grid[yy][xx] == ' ' || grid[yy][xx] == 0 {
				fmt.Println("zoom", y, x, f, yy, xx, ff)
				return -1
			}
			if grid[yy][xx] == '.' {
				y, x, f = yy, xx, ff
			} else if grid[yy][xx] == '#' {
				break
			}
		}

		if m[1] == 0 {
			// left
			if f == 0 {
				f = 3
			} else {
				f--
			}
		} else if m[1] == 1 {
			if f == 3 {
				f = 0
			} else {
				f++
			}
		}
		if debug {
			fmt.Println("|", y, x, f)
			if c > 20 {
				return -1
			}
		}
	}
	return 1000*(y+1) + 4*(x+1) + f
}
