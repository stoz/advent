package main

func s2206(filename string, part2 bool) int {
	// expected results: 1275, 3605
	if filename == "sample.txt" {
		// sample data was not provided for day 6
		return -1
	}
	line := ReadLine("./data/2022/06/" + filename)
	i := 4
	if part2 {
		i = 14
	}
	return UniqueWindow(line, i)
}

func UniqueWindow(line string, window int) int {
	var r []rune
	pos := 0
	for k, v := range line {
		r = append(r, v)
		if len(r) > window {
			r = r[1:]
			dup := false
			for a, b := range r {
				for c, d := range r {
					if a != c && b == d {
						dup = true
						break
					}
				}
			}
			if !dup {
				pos = k + 1
				break
			}
		}
	}
	return pos
}
