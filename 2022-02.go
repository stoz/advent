package main

import (
	"strings"
)

func s22021(filename string) int {
	// expedcted valie: 13052
	lines := ReadFile("./data/2022/02/" + filename)
	count := 0
	for _, line := range lines {
		words := strings.Fields(line)
		if words[1] == "X" {
			if words[0] == "A" {
				count += 4
			} else if words[0] == "B" {
				count += 1
			} else if words[0] == "C" {
				count += 7
			}
		}
		if words[1] == "Y" {
			if words[0] == "A" {
				count += 8
			} else if words[0] == "B" {
				count += 5
			} else if words[0] == "C" {
				count += 2
			}
		}
		if words[1] == "Z" {
			if words[0] == "A" {
				count += 3
			} else if words[0] == "B" {
				count += 9
			} else if words[0] == "C" {
				count += 6
			}
		}
	}
	return count
}

func s22022(filename string) int {
	// expected result: 13693
	lines := ReadFile("./data/2022/02/" + filename)
	key := map[string]map[string]int{
		"A": {"X": 3, "Y": 4, "Z": 8},
		"B": {"X": 1, "Y": 5, "Z": 9},
		"C": {"X": 2, "Y": 6, "Z": 7},
	}
	count := 0
	for _, line := range lines {
		words := strings.Fields(line)
		count += key[words[0]][words[1]]
	}
	return count
}
