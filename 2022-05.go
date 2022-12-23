package main

import (
	"strconv"
	"strings"
)

func s22051(filename string) string {
	// expected result: VGBBJCRMN
	lines := ReadFile("./data/2022/05/" + filename)
	var boxInput []string
	boxes := make(map[int]map[int]string)
	for i, line := range lines {
		if line == "" {
			lines = lines[i+1:]
			break
		} else {
			boxInput = append(boxInput, line)
		}
	}
	words := strings.Fields(boxInput[len(boxInput)-2])
	wid, hig := len(words), len(boxInput)-1
	for x := 1; x < wid*4; x += 4 {
		boxes[(x+3)/4] = make(map[int]string)
		for y := hig - 1; y >= 0; y-- {
			if boxInput[y][x] != ' ' {
				boxes[(x+3)/4][hig-1-y] = string(boxInput[y][x])
			}
		}
	}
	for _, line := range lines {
		words := strings.Fields(line)
		qty, _ := strconv.Atoi(words[1])
		src, _ := strconv.Atoi(words[3])
		dst, _ := strconv.Atoi(words[5])
		for i := 1; i <= qty; i++ {
			boxes[dst][len(boxes[dst])] = boxes[src][len(boxes[src])-1]
			delete(boxes[src], len(boxes[src])-1)
		}
	}
	output := ""
	for i := 0; i <= len(boxes); i++ {
		output += boxes[i][len(boxes[i])-1]
	}
	return output
}

func s22052(filename string) string {
	// expected result: LBBVJBRMH
	lines := ReadFile("./data/2022/05/" + filename)
	var boxInput []string
	boxes := make(map[int]map[int]string)
	for i, line := range lines {
		if line == "" {
			lines = lines[i+1:]
			break
		} else {
			boxInput = append(boxInput, line)
		}
	}
	words := strings.Fields(boxInput[len(boxInput)-2])
	wid, hig := len(words), len(boxInput)-1
	for x := 1; x < wid*4; x += 4 {
		boxes[(x+3)/4] = make(map[int]string)
		for y := hig - 1; y >= 0; y-- {
			if boxInput[y][x] != ' ' {
				boxes[(x+3)/4][hig-1-y] = string(boxInput[y][x])
			}
		}
	}
	for _, line := range lines {
		words := strings.Fields(line)
		qty, _ := strconv.Atoi(words[1])
		src, _ := strconv.Atoi(words[3])
		dst, _ := strconv.Atoi(words[5])
		for i := 1; i <= qty; i++ {
			boxes[dst][len(boxes[dst])] = boxes[src][len(boxes[src])-1-qty+i]
		}
		for i := 1; i <= qty; i++ {
			delete(boxes[src], len(boxes[src])-1)
		}
	}
	output := ""
	for i := 0; i <= len(boxes); i++ {
		output += boxes[i][len(boxes[i])-1]
	}
	return output
}
