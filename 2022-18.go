package main

import "fmt"

func s22181(filename string) {
	lines := ReadFile("./data/2022/18/" + filename)
	//rops := make(map[int]map[int]map[int]bool)
	var d [22][22][22]int
	for _, line := range lines {
		i := ExtractInts(line)
		d[i[0]+1][i[1]+1][i[2]+1] = 1
	}
	var sum int
	for x, ys := range d {
		for y, zs := range ys {
			for z, v := range zs {
				if v == 1 {
					if d[x][y][z+1] == 0 {
						sum++
					}
					if d[x][y][z-1] == 0 {
						sum++
					}
					if d[x][y+1][z] == 0 {
						sum++
					}
					if d[x][y-1][z] == 0 {
						sum++
					}
					if d[x+1][y][z] == 0 {
						sum++
					}
					if d[x-1][y][z] == 0 {
						sum++
					}
				}
			}
		}
	}
	fmt.Println(sum)
}

func s22182(filename string) {
	lines := ReadFile("./data/2022/18/" + filename)
	//rops := make(map[int]map[int]map[int]bool)
	// not 2810
	// not 1781
	// answer: 2018
	var d [26][26][26]int
	for _, line := range lines {
		i := ExtractInts(line)
		d[i[0]+3][i[1]+3][i[2]+3] = 1
	}
	var sum int

	//start 0,0,0 as safe air and propogate it around the grid
	d[0][0][0] = 2
	for x, ys := range d {
		for y, zs := range ys {
			for z, v := range zs {
				if v == 0 {
					if z < 25 && d[x][y][z+1] == 2 {
						d[x][y][z] = 2
					}
					if z > 0 && d[x][y][z-1] == 2 {
						d[x][y][z] = 2
					}
					if y < 25 && d[x][y+1][z] == 2 {
						d[x][y][z] = 2
					}
					if y > 0 && d[x][y-1][z] == 2 {
						d[x][y][z] = 2
					}
					if x < 25 && d[x+1][y][z] == 2 {
						d[x][y][z] = 2
					}
					if x > 0 && d[x-1][y][z] == 2 {
						d[x][y][z] = 2
					}
				}
			}
		}
	}
	for x, ys := range d {
		for y, zs := range ys {
			for z, v := range zs {
				if v == 0 {
					if z < 25 && d[x][y][z+1] == 2 {
						d[x][y][z] = 2
					}
					if z > 0 && d[x][y][z-1] == 2 {
						d[x][y][z] = 2
					}
					if y < 25 && d[x][y+1][z] == 2 {
						d[x][y][z] = 2
					}
					if y > 0 && d[x][y-1][z] == 2 {
						d[x][y][z] = 2
					}
					if x < 25 && d[x+1][y][z] == 2 {
						d[x][y][z] = 2
					}
					if x > 0 && d[x-1][y][z] == 2 {
						d[x][y][z] = 2
					}
				}
			}
		}
	}
	for x, ys := range d {
		for y, zs := range ys {
			for z, v := range zs {
				if v == 0 {
					if z < 25 && d[x][y][z+1] == 2 {
						d[x][y][z] = 2
					}
					if z > 0 && d[x][y][z-1] == 2 {
						d[x][y][z] = 2
					}
					if y < 25 && d[x][y+1][z] == 2 {
						d[x][y][z] = 2
					}
					if y > 0 && d[x][y-1][z] == 2 {
						d[x][y][z] = 2
					}
					if x < 25 && d[x+1][y][z] == 2 {
						d[x][y][z] = 2
					}
					if x > 0 && d[x-1][y][z] == 2 {
						d[x][y][z] = 2
					}
				}
			}
		}
	}
	for x, ys := range d {
		for y, zs := range ys {
			for z, v := range zs {
				if v == 0 {
					if z < 25 && d[x][y][z+1] == 2 {
						d[x][y][z] = 2
					}
					if z > 0 && d[x][y][z-1] == 2 {
						d[x][y][z] = 2
					}
					if y < 25 && d[x][y+1][z] == 2 {
						d[x][y][z] = 2
					}
					if y > 0 && d[x][y-1][z] == 2 {
						d[x][y][z] = 2
					}
					if x < 25 && d[x+1][y][z] == 2 {
						d[x][y][z] = 2
					}
					if x > 0 && d[x-1][y][z] == 2 {
						d[x][y][z] = 2
					}
				}
			}
		}
	}

	for x, ys := range d {
		for y, zs := range ys {
			for z, v := range zs {
				if v == 1 {
					if d[x][y][z+1] == 2 {
						sum++
					}
					if d[x][y][z-1] == 2 {
						sum++
					}
					if d[x][y+1][z] == 2 {
						sum++
					}
					if d[x][y-1][z] == 2 {
						sum++
					}
					if d[x+1][y][z] == 2 {
						sum++
					}
					if d[x-1][y][z] == 2 {
						sum++
					}
				}
			}
		}
	}
	fmt.Println(sum)

	//for _, a := range d {
	//fmt.Println(a)
	//}
}
