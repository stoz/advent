package main

import (
	"strconv"
	"strings"
)

func s2204(filename string, part2 bool) int {
	// expected result: 542
	lines := ReadFile("./data/2022/04/" + filename)
	count := 0
	f := func(c rune) bool {
		return c == '-' || c == ','
	}
	for _, line := range lines {
		words := strings.FieldsFunc(line, f)
		a, _ := strconv.Atoi(words[0])
		b, _ := strconv.Atoi(words[1])
		c, _ := strconv.Atoi(words[2])
		d, _ := strconv.Atoi(words[3])
		if part2 {
			if (a <= d && b >= c) || (d <= a && c >= b) {
				count++
			}
		} else {
			if (a >= c && b <= d) || (c >= a && d <= b) {
				count++
			}
		}
	}
	return count
}
