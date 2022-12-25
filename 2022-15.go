package main

import (
	"fmt"
)

func s22151(filename string, debug bool) int {
	// expected result: 4907780
	lines := ReadFile("./data/2022/15/" + filename)
	var dim [4]int
	cave := make(map[int]map[int]int)
	sens := make(map[int]map[int]int)
	for _, line := range lines {
		ints := ExtractSInts(line)
		for i, j := range ints {
			switch i % 2 {
			case 0: //x
				if j < dim[2] || dim[2] == 0 {
					dim[2] = j - 10000000
				} else if j > dim[3] {
					dim[3] = j + 10000000
				}
			case 1: //y
				if j < dim[0] || dim[0] == 0 {
					dim[0] = j - 10
				} else if j > dim[1] {
					dim[1] = j + 10
				}
			}
		}
	}
	for i := dim[0]; i <= dim[1]; i++ {
		cave[i] = make(map[int]int)
	}
	for _, line := range lines {
		ints := ExtractInts(line)
		var x int
		var sensor [2]int
		for i, y := range ints {
			switch i % 4 {
			case 0, 2:
				x = y
			case 1:
				cave[y][x] = 3
				sensor[0] = y
				sensor[1] = x
			case 3:
				cave[y][x] = 4
				_, ok := sens[sensor[0]]
				if !ok {
					sens[sensor[0]] = make(map[int]int)
				}
				sens[sensor[0]][sensor[1]] = abs(sensor[0]-y) + abs(sensor[1]-x)
			}
		}
	}
	if debug {
		fmt.Println(dim)
		PrintCave2(cave, dim)
	}
	var sum int
	y := 2000000
	for x := dim[2]; x <= dim[3]; x++ {
		value, ok := cave[y][x]
		safe := 0
		if !ok {
			// square is empty
			// check if any sensors are in range
			safe = InRangeOfSensor(y, x, sens)
		} else {
			if value == 3 {
				safe = 1
			}
		}
		sum += safe
	}
	return sum
}

func s22152(filename string, debug bool) int {
	// expected result: 13639962836448
	lines := ReadFile("./data/2022/15/" + filename)
	var dim [4]int
	cave := make(map[int]map[int]int)
	sens := make(map[int]map[int]int)
	for _, line := range lines {
		ints := ExtractSInts(line)
		for i, j := range ints {
			switch i % 2 {
			case 0: //x
				if j < dim[2] || dim[2] == 0 {
					dim[2] = j - 10000000
				} else if j > dim[3] {
					dim[3] = j + 10000000
				}
			case 1: //y
				if j < dim[0] || dim[0] == 0 {
					dim[0] = j - 10
				} else if j > dim[1] {
					dim[1] = j + 10
				}
			}
		}
	}
	for i := dim[0]; i <= dim[1]; i++ {
		cave[i] = make(map[int]int)
	}
	for _, line := range lines {
		ints := ExtractSInts(line)
		var x int
		var sensor [2]int
		for i, y := range ints {
			switch i % 4 {
			case 0, 2:
				x = y
			case 1:
				cave[y][x] = 3
				sensor[0] = y
				sensor[1] = x
			case 3:
				cave[y][x] = 4
				_, ok := sens[sensor[0]]
				if !ok {
					sens[sensor[0]] = make(map[int]int)
				}
				var diff int
				if sensor[0] < y {
					diff += y - sensor[0]
				} else {
					diff += sensor[0] - y
				}
				if sensor[1] < x {
					diff += x - sensor[1]
				} else {
					diff += sensor[1] - x
				}
				sens[sensor[0]][sensor[1]] = diff
			}
		}
	}
	if debug {
		fmt.Println(dim)
	}
	// 4000000
	var lim = 4000000
	var fin []int
	cache := make(map[int]bool)

	//13639963409990 too high
	for sy, line := range sens {
		for sx, d := range line {
			// check around the edge
			e := d + 1
			for y := sy - e; y <= sy+e; y++ {
				if y >= 0 && y <= lim {
					var test [2]int
					if y < sy {
						test[0] = sx - e + sy - y
						test[1] = sx + e - sy - y
					} else {
						test[0] = sx - e + y - sy
						test[1] = sx + e - y - sy
					}
					for _, x := range test {
						if x >= 0 && x <= lim {
							_, ok := sens[x*lim+y]
							if !ok {
								var diff int
								if sy < y {
									diff += y - sy
								} else {
									diff += sy - y
								}
								if sx < x {
									diff += x - sx
								} else {
									diff += sx - x
								}
								if diff == e {
									safe := InRangeOfSensor(y, x, sens)
									cache[x*lim+y] = true
									if safe == 0 {
										return x*lim + y
									}
								}
							}
						}
					}
				}
			}
		}
		if len(fin) > 2 {
			fmt.Println("too many", fin)
			fmt.Println("---", fin[1]*lim+fin[0])
			return -1
		}
	}
	if debug {
		fmt.Println(fin)
		if len(fin) == 2 {
			fmt.Print(fin[1]*lim + fin[0])
		}
	}
	return -1
}

func InRangeOfSensor(y, x int, sens map[int]map[int]int) int {
	for sy, line := range sens {
		for sx, d := range line {
			if abs(sy-y)+abs(sx-x) <= d {
				return 1
			}
		}
	}
	return 0
}
