package main

import "fmt"

func s2220(filename string, part2, debug bool) int {
	// expected results: 2827, 7834270093909
	lines := ReadFile("./data/2022/20/" + filename)
	var d [][2]int
	var mixed []int
	factor, iterations := 1, 1
	if part2 {
		factor = 811589153
		iterations = 10
	}
	for i, line := range lines {
		ext := ExtractSInts(line)
		app := [2]int{ext[0] * factor, i}
		d = append(d, app)
	}
	l := len(d)
	for z := 0; z < iterations; z++ {
		for i := 0; i < l; i++ {
			if debug {
				fmt.Println(z, i)
			}
			new := d[i][1] + d[i][0]
			if d[i][0] < 0 {
				new--
			}
			for new < 0 {
				new = l + new%(l-1)
			}
			for new >= l {
				new = new % (l - 1)
			}
			for j := 0; j < l; j++ {
				if new > d[i][1] {
					if d[j][1] > d[i][1] && d[j][1] <= new {
						if d[j][1] == 0 {
							d[j][1] = l - 1
						} else {
							d[j][1]--
						}
					}
				} else if d[i][1] > new {
					if d[j][1] >= new && d[j][1] < d[i][1] {
						if d[j][1] == l-1 {
							d[j][1] = 0
						} else {
							d[j][1]++
						}
					}
				}
			}
			d[i][1] = new
			if debug {
				var dbg []int
				seen := make(map[int]bool)
				for j := 0; j < l; j++ {
					// sanity check
					if _, ok := seen[d[j][1]]; ok {
						fmt.Println("Duplicate", d[j])
						return -1
					}
					seen[d[j][1]] = true
					for _, a := range d {
						if a[1] == j {
							dbg = append(dbg, a[0])
						}
					}
				}
			}
		}
	}
	var zero int
	for i := 0; i < l; i++ {
		if d[i][0] == 0 {
			zero = d[i][1]
		}
		for _, a := range d {
			if a[1] == i {
				mixed = append(mixed, a[0])
			}
		}
	}
	return mixed[(zero+1000)%l] + mixed[(zero+2000)%l] + mixed[(zero+3000)%l]
}
