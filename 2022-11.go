package main

import (
	"strings"
)

type Monkey struct {
	items       []int
	operation   int
	modulus     int
	targets     [2]int
	inspections int
}

func s2211(filename string, part2 bool) int {
	// expected results: 55930, 14636993466
	lines := ReadFile("./data/2022/11/" + filename)
	var monkeys []Monkey
	var monkey Monkey
	iterations := 20
	if part2 {
		iterations = 10000
	}
	for i, line := range lines {
		ints := ExtractInts(line)
		switch i % 7 {
		case 1:
			monkey.items = ints
		case 2:
			words := strings.Fields(line)
			if words[5] == "old" {
				monkey.operation = 0
			} else if words[4] == "*" {
				monkey.operation = -ints[0]
			} else {
				monkey.operation = ints[0]
			}
		case 3:
			monkey.modulus = ints[0]
		case 4:
			monkey.targets[0] = ints[0]
		case 5:
			monkey.targets[1] = ints[0]
			monkeys = append(monkeys, monkey)
		}
	}
	modulus := 1
	for _, monkey := range monkeys {
		modulus *= monkey.modulus
	}
	for i := 0; i < iterations; i++ {
		for j, monkey := range monkeys {
			for _, old := range monkey.items {
				monkeys[j].inspections++
				new := 0
				if monkey.operation == 0 {
					new = old * old
				} else if monkey.operation < 0 {
					new = old * -monkey.operation
				} else {
					new = old + monkey.operation
				}
				if part2 {
					new = new % modulus
				} else {
					new = FloorDivision(new, 3)
				}
				target := 0
				if new%monkey.modulus == 0 {
					target = monkey.targets[0]
				} else {
					target = monkey.targets[1]
				}
				monkeys[target].items = append(monkeys[target].items, new)
			}
			monkeys[j].items = []int{}
		}
	}
	var max [2]int
	for _, monkey := range monkeys {
		if monkey.inspections > max[0] {
			max[1] = max[0]
			max[0] = monkey.inspections
		} else if monkey.inspections > max[1] {
			max[1] = monkey.inspections
		}
	}
	return max[0] * max[1]
}
