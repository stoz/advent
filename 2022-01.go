package main

import (
	"strconv"
)

func s2201(filename string, part2 bool) int {
	// expected results: 71780, 212489
	lines := ReadFile("./data/2022/01/" + filename)
	current := 0
	var max [3]int
	for _, line := range lines {
		if line == "" {
			if current > max[0] {
				max[2] = max[1]
				max[1] = max[0]
				max[0] = current
			} else if current > max[1] {
				max[2] = max[1]
				max[1] = current
			} else if current > max[2] {
				max[2] = current
			}
			current = 0
		} else {
			textAsInt, _ := strconv.Atoi(line)
			current += textAsInt
		}
	}
	// no final blank line so repeat
	if current > max[0] {
		max[2] = max[1]
		max[1] = max[0]
		max[0] = current
	} else if current > max[1] {
		max[2] = max[1]
		max[1] = current
	} else if current > max[2] {
		max[2] = current
	}
	if part2 {
		return max[0] + max[1] + max[2]
	}
	return max[0]
}
