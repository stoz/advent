package main

import (
	"fmt"
	"strings"
)

func s2306(filename string, part, debug bool) int {
	lines := ReadFile("./data/2023/06/" + filename)
	if part {
		lines[0] = strings.ReplaceAll(lines[0], " ", "")
		lines[1] = strings.ReplaceAll(lines[1], " ", "")
	}
	times := ExtractInts(lines[0])
	distances := ExtractInts(lines[1])
	margin := 1
	for i := 0; i < len(times); i++ {
		if debug {
			fmt.Println(times[i], distances[i])
		}
		wins := 0
		for j := 1; j < times[i]; j++ {
			if debug {
				fmt.Println("Speed:", j, "Distance: ", j*(times[i]-j))
			}
			if j*(times[i]-j) > distances[i] {
				wins++
			}
		}
		if debug {
			fmt.Println("Wins:", wins)
		}
		margin = margin * wins
	}
	return margin
}
