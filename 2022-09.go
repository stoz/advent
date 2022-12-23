package main

import (
	"log"
	"strconv"
	"strings"
)

func s22091(filename string) int {
	// expected result: 6376
	lines := ReadFile("./data/2022/09/" + filename)
	hx := 1000
	hy := 1000
	tx := 1000
	ty := 1000
	factor := 10000
	var visited [100000000]bool
	visited[1000*factor+1000] = true
	for _, line := range lines {
		f := strings.Fields(line)
		moves, err := strconv.Atoi(f[1])
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < moves; i++ {
			if f[0] == "U" {
				hx--
			} else if f[0] == "D" {
				hx++
			} else if f[0] == "L" {
				hy--
			} else if f[0] == "R" {
				hy++
			}
			// move if too far away
			if hx-tx > 1 {
				tx++
				if hy-ty > 0 {
					ty++
				} else if ty-hy > 0 {
					ty--
				}
			} else if tx-hx > 1 {
				tx--
				if hy-ty > 0 {
					ty++
				} else if ty-hy > 0 {
					ty--
				}
			} else if hy-ty > 1 {
				ty++
				if hx-tx > 0 {
					tx++
				} else if tx-hx > 0 {
					tx--
				}
			} else if ty-hy > 1 {
				ty--
				if hx-tx > 0 {
					tx++
				} else if tx-hx > 0 {
					tx--
				}
			}
			visited[tx*factor+ty] = true
		}
	}
	count := 0
	for _, v := range visited {
		if v {
			count++
		}
	}
	return count
}

func s22092(filename string) int {
	// expected result: 2607
	lines := ReadFile("./data/2022/09/" + filename)
	x := [10]int{1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000}
	y := [10]int{1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000}
	factor := 10000
	var visited [100000000]bool
	visited[1000*factor+1000] = true
	for _, line := range lines {
		f := strings.Fields(line)
		moves, err := strconv.Atoi(f[1])
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < moves; i++ {
			if f[0] == "U" {
				x[0]--
			} else if f[0] == "D" {
				x[0]++
			} else if f[0] == "L" {
				y[0]--
			} else if f[0] == "R" {
				y[0]++
			}
			for z := 0; z < 9; z++ {
				// move if too far away
				if x[z]-x[z+1] > 1 {
					x[z+1]++
					if y[z]-y[z+1] > 0 {
						y[z+1]++
					} else if y[z+1]-y[z] > 0 {
						y[z+1]--
					}
				} else if x[z+1]-x[z] > 1 {
					x[z+1]--
					if y[z]-y[z+1] > 0 {
						y[z+1]++
					} else if y[z+1]-y[z] > 0 {
						y[z+1]--
					}
				} else if y[z]-y[z+1] > 1 {
					y[z+1]++
					if x[z]-x[z+1] > 0 {
						x[z+1]++
					} else if x[z+1]-x[z] > 0 {
						x[z+1]--
					}
				} else if y[z+1]-y[z] > 1 {
					y[z+1]--
					if x[z]-x[z+1] > 0 {
						x[z+1]++
					} else if x[z+1]-x[z] > 0 {
						x[z+1]--
					}
				}
			}
			visited[x[9]*factor+y[9]] = true
		}
	}
	count := 0
	for _, v := range visited {
		if v {
			count++
		}
	}
	return count
}
