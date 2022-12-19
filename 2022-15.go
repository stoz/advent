package main

import (
	"fmt"
	"os"
)

func xyz3() {
	// 4350538 too low
	// 4352518 too low
	// 4370518 too low
	lines := ReadFile("input.txt")
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
	fmt.Println(dim)
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
										fin = append(fin, y, x)
										fmt.Println("!!!", x*lim+y)
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
			os.Exit(1)
		}
	}
	fmt.Println(fin)
	if len(fin) == 2 {
		fmt.Print(fin[1]*lim + fin[0])
	}
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

func xyz2() {
	// 4350538 too low
	// 4352518 too low
	// 4370518 too low
	lines := ReadFile("input.txt")
	var dim [4]int
	cave := make(map[int]map[int]int)
	sens := make(map[int]map[int]int)
	for _, line := range lines {
		ints := ExtractInts(line)
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
	fmt.Println(dim)
	// PrintCave2(cave, dim)
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
	fmt.Println(sum)
}

func xyz() {
	lines := ReadFile("input.txt")
	var dim [4]int
	cave := make(map[int]map[int]int)
	for _, line := range lines {
		ints := ExtractInts(line)
		for i, j := range ints {
			switch i % 2 {
			case 0: //x
				if j < dim[2] || dim[2] == 0 {
					dim[2] = j - 10
				} else if j > dim[3] {
					dim[3] = j + 10
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
			fmt.Println(i)
			switch i % 4 {
			case 0, 2:
				x = y
			case 1:
				cave[y][x] = 3
				sensor[0] = y
				sensor[1] = x
			case 3:
				cave[y][x] = 4
				//measure distance
				d := abs(sensor[0]-y) + abs(sensor[1]-x)
				for dy := sensor[0] - d; dy <= sensor[0]+d; dy++ {
					// 1, 3, 5 etc
					width := d - abs(sensor[1]-dy)
					for dx := sensor[1] - width; dx <= sensor[1]+width; dx++ {
						_, ok := cave[dy]
						if ok {
							_, ok2 := cave[dy][dx]
							if !ok2 {
								cave[dy][dx] = 1
							}
						}
					}
				}
			}
		}
	}
	fmt.Println(dim)
	PrintCave2(cave, dim)
	var sum int
	for _, c := range cave[2000000] {
		if c == 1 {
			sum++
		}
	}
	//fmt.Println(sum)
}
