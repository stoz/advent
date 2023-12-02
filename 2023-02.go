package main

import (
	"strconv"
	"strings"
)

func s23021(filename string, part, debug bool) int {
	lines := ReadFile("./data/2023/02/" + filename)
	score := 0
	for lineNumber, line := range lines {
		pass := true
		words := strings.Fields(line)
		for i, word := range words {
			if n, err := strconv.Atoi(word); err == nil {
				switch words[i+1][0] {
				case 'r':
					if n > 12 {
						pass = false
					}
				case 'g':
					if n > 13 {
						pass = false
					}
				case 'b':
					if n > 14 {
						pass = false
					}
				}
			}
		}
		if pass {
			score += lineNumber + 1
		}
	}
	return score
}

func s23022(filename string, part, debug bool) int {
	lines := ReadFile("./data/2023/02/" + filename)
	score := 0
	for _, line := range lines {
		var r, g, b int
		words := strings.Fields(line)
		for i, word := range words {
			if n, err := strconv.Atoi(word); err == nil {
				switch words[i+1][0] {
				case 'r':
					if n > r {
						r = n
					}
				case 'g':
					if n > g {
						g = n
					}
				case 'b':
					if n > b {
						b = n
					}
				}
			}
		}
		score += r * g * b
	}
	return score
}
