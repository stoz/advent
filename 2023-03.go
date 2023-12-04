package main

import (
	"strconv"
)

func s2303(filename string, part, debug bool) int {
	lines := ReadGridRune("./data/2023/03/" + filename)
	sum := 0
	// part 2
	var gears [140][140][]int
	for y, line := range lines {
		num := 0
		for x, r := range line {
			if r >= '0' && r <= '9' {
				// is digit
				num = num*10 + int(r-'0')
			}
			// check if the next rune is not a digit
			if x == len(line)-1 || line[x+1] < '0' || line[x+1] > '9' {
				// configure search space
				ystart := 0
				if y-1 > ystart {
					ystart = y - 1
				}
				yend := len(lines) - 1
				if y+1 < yend {
					yend = y + 1
				}
				xstart := 0
				if x-len(strconv.Itoa(num)) > xstart {
					xstart = x - len(strconv.Itoa(num))
				}
				xend := len(line) - 1
				if x+1 < xend {
					xend = x + 1
				}
				foundSymbol := false
				for yy := ystart; yy <= yend; yy++ {
					for xx := xstart; xx <= xend; xx++ {
						if (lines[yy][xx] < '0' || lines[yy][xx] > '9') && lines[yy][xx] != '.' {
							// found a symbol
							foundSymbol = true
							// part 2
							if lines[yy][xx] == '*' && num > 0 { // num > 0 is a weird hack
								gears[xx][yy] = append(gears[xx][yy], num)
							}
						}
					}
				}
				if foundSymbol {
					sum += num
				}
				num = 0
			}
		}
	}
	if !part {
		return sum
	}
	sum = 0
	for _, gear := range gears {
		for _, g := range gear {
			if len(g) == 2 {
				sum += g[0] * g[1]
			}
		}
	}
	return sum
}
