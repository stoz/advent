package main

func s2212(filename string, pt2 bool) int {
	// expected results: 534, 525
	lines := ReadGridRune("./data/2022/12/" + filename)
	grid := make(map[int]map[int]GridPoint)
	max := [2]int{len(lines), len(lines[0])}
	var start [2]int
	var target [2]int

	// pre-process to replace S and E with a and z
	for y, line := range lines {
		for x, a := range line {
			switch a {
			case 'S':
				lines[y][x] = 'a'
				start[0] = y
				start[1] = x
			case 'E':
				lines[y][x] = 'z'
				target[0] = y
				target[1] = x
			}
		}
	}
	for y, line := range lines {
		grid[y] = make(map[int]GridPoint)
		for x, a := range line {
			var point GridPoint
			if y == start[0] && x == start[1] {
				point.c = 0
			} else if pt2 && a == 'a' {
				point.c = 0
			} else {
				point.c = 999999
			}
			var nbh [4][2]int
			for i := 0; i < len(nbh); i++ {
				nbh[i][0] = y
				nbh[i][1] = x
			}
			nbh[0][0]--
			nbh[1][0]++
			nbh[2][1]--
			nbh[3][1]++
			for _, n := range nbh {
				if n[0] > -1 && n[1] > -1 && n[0] < max[0] && n[1] < max[1] && lines[n[0]][n[1]] < a+2 {
					point.n = append(point.n, n)
				}
			}
			grid[y][x] = point
		}
	}
	result := Dijkstra(grid, target)
	return result
}
