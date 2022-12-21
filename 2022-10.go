package main

import (
	"fmt"
	"strconv"
	"strings"
)

func s22101(filename string) int {
	// expected result: 14620
	lines := ReadFile("./data/2022/10/" + filename)
	x := 1
	c := 0
	interesting := map[int]struct{}{20: {}, 60: {}, 100: {}, 140: {}, 180: {}, 220: {}}
	sum := 0
	for _, line := range lines {
		f := strings.Fields(line)
		cycles := 1
		add := 0
		if f[0] == "addx" {
			cycles = 2
			add, _ = strconv.Atoi(f[1])
		}
		for i := 0; i < cycles; i++ {
			c++
			if _, ok := interesting[c]; ok {
				sum += x * c
			}
		}
		x += add
	}
	return sum
}

func s22102(filename string, output bool) {
	// expected result: BJFRHRFU
	lines := ReadFile("./data/2022/10/" + filename)
	x := 1
	width := 40
	buffer := ""
	for _, line := range lines {
		f := strings.Fields(line)
		cycles := 1
		add := 0
		if f[0] == "addx" {
			cycles = 2
			add, _ = strconv.Atoi(f[1])
		}
		for i := 0; i < cycles; i++ {
			c := len([]rune(buffer)) // get character count rather than byte count
			if c == x-1 || c == x || c == x+1 {
				buffer += "█"
			} else {
				buffer += " "
			}
			if c == width-1 {
				if output {
					fmt.Println(buffer)
				}
				buffer = ""
			}
		}
		x += add
	}
}
