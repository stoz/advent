package main

import (
	"strconv"
	"strings"
)

func s22071(filename string) int {
	// expected result: 1325919
	lines := ReadFile("./data/2022/07/" + filename)
	// dir sizes
	var path []string
	sizes := make(map[string]int)
	for _, line := range lines {
		words := strings.Fields(line)
		if line == "$ cd /" {
			path = nil
		} else if words[0] == "$" && words[1] == "cd" {
			if words[2] == ".." {
				path = path[:len(path)-1]
			} else {
				path = append(path, words[2])
			}
		} else if words[0] == "dir" {
			// do nothing
		} else {
			size, _ := strconv.Atoi(words[0])
			pat := ""
			for _, p := range path {
				pat += "/" + p
				sizes[pat] += size
			}
		}
	}
	sum := 0
	for _, size := range sizes {
		if size < 100001 {
			sum += size
		}
	}
	return sum
}

func s22072(filename string) int {
	// expected result: 2050735
	lines := ReadFile("./data/2022/07/" + filename)
	// dir sizes
	var path []string
	sizes := make(map[string]int)
	for _, line := range lines {
		words := strings.Fields(line)
		if line == "$ cd /" {
			path = nil
			path = append(path, "/")
		} else if words[0] == "$" && words[1] == "cd" {
			if words[2] == ".." {
				path = path[:len(path)-1]
			} else {
				path = append(path, words[2])
			}
		} else if words[0] == "dir" {
			// do nothing
		} else {
			size, _ := strconv.Atoi(words[0])
			pat := ""
			for _, p := range path {
				pat += "/" + p
				sizes[pat] += size
			}
		}
	}
	//sum := 0
	space := 70000000 - sizes["//"]
	needed := 30000000
	target := needed - space
	smallest := 70000000
	for _, size := range sizes {
		if size >= target && size < smallest {
			smallest = size
		}
	}
	return smallest
}
