package main

import (
	"fmt"
	"slices"
)

func s2309(filename string, part, debug bool) int {
	lines := ReadFile("./data/2023/09/" + filename)
	sum := 0
	for lineCount, line := range lines {
		values := ExtractSInts(line)
		if part {
			slices.Reverse(values)
		}
		var steps [][]int
		steps = append(steps, values)
		notAllZero := true
		for notAllZero {
			var diff []int
			for i, value := range steps[len(steps)-1] {
				if i > 0 {
					diff = append(diff, value-steps[len(steps)-1][i-1])
				}
			}
			steps = append(steps, diff)
			diffAllZero := true
			for _, d := range diff {
				if d != 0 {
					diffAllZero = false
				}
			}
			if diffAllZero {
				notAllZero = false
			}
		}
		if debug && lineCount == 0 {
			for _, step := range steps {
				fmt.Println(step)
			}
		}
		calc := 0
		for i := len(steps) - 2; i > -1; i-- {
			calc = steps[i][len(steps[i])-1] + calc
		}
		sum += calc
	}
	return sum
}
