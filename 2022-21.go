package main

import (
	"strconv"
	"strings"
)

type Yell struct {
	result    int
	set       bool
	targets   [2]string
	operation rune
}

func s22211(filename string) int {
	// expected result: 256997859093114
	lines := ReadFile("./data/2022/21/" + filename)
	monkeys := make(map[string]Yell)
	for _, line := range lines {
		words := strings.Fields(line)
		name := words[0][:4]
		var monkey Yell
		switch len(words) {
		case 2:
			monkey.result, _ = strconv.Atoi(words[1])
			monkey.set = true
			monkeys[name] = monkey
		case 4:
			monkey.targets[0] = words[1]
			monkey.targets[1] = words[3]
			runes := []rune(words[2])
			monkey.operation = runes[0]
			monkeys[name] = monkey
		}
	}
	for {
		for n, m := range monkeys {
			if n == "root" && m.set {
				return m.result
			} else if !m.set && monkeys[m.targets[0]].set && monkeys[m.targets[1]].set {
				switch m.operation {
				case '+':
					m.result = monkeys[m.targets[0]].result + monkeys[m.targets[1]].result
					m.set = true
					monkeys[n] = m
				case '-':
					m.result = monkeys[m.targets[0]].result - monkeys[m.targets[1]].result
					m.set = true
					monkeys[n] = m
				case '*':
					m.result = monkeys[m.targets[0]].result * monkeys[m.targets[1]].result
					m.set = true
					monkeys[n] = m
				case '/':
					m.result = monkeys[m.targets[0]].result / monkeys[m.targets[1]].result
					m.set = true
					monkeys[n] = m
				}
			}
		}
	}
}

func s22212(filename string) int {
	lines := ReadFile("./data/2022/21/" + filename)
	monkeys := make(map[string]Yell)
	for _, line := range lines {
		words := strings.Fields(line)
		name := words[0][:4]
		var monkey Yell
		switch len(words) {
		case 2:
			monkey.result, _ = strconv.Atoi(words[1])
			monkey.set = true
			monkeys[name] = monkey
		case 4:
			monkey.targets[0] = words[1]
			monkey.targets[1] = words[3]
			runes := []rune(words[2])
			if name == "root" {
				monkey.operation = '='
			} else {
				monkey.operation = runes[0]
			}
			monkeys[name] = monkey
		}
	}
	min := 1000000000000
	max := 4000000000000
	for {
		// build array to test
		d := FloorDivision(max-min, 6)
		test := [6]int{min, min + d, min + d*2, min + d*3, min + d*4, min + d*5}
		var results [6]int
		for ti, t := range test {
			for i := t; i < max; i++ {
				diff, check := monkeyCheck(monkeys, i)
				if check {
					return i
				} else if diff != 0 {
					results[ti] = abs(diff)
					break
				}
			}
		}
		var mins [6]int
		for ri, r := range results {
			if mins[0] == 0 {
				mins[0], mins[2], mins[4] = r, r, r
			} else if r < mins[0] {
				mins[4], mins[5] = mins[2], mins[3]
				mins[2], mins[3] = mins[0], mins[1]
				mins[0], mins[1] = r, ri
			} else if r < mins[2] {
				mins[4], mins[5] = mins[2], mins[3]
				mins[2], mins[3] = r, ri
			} else if r < mins[4] {
				mins[4], mins[5] = r, ri
			}
		}
		if mins[1] == len(test)-1 {
			min = test[len(test)-2]
		} else if mins[1] == 0 {
			max = test[1]
		} else if mins[3] < mins[5] {
			min, max = test[mins[3]], test[mins[5]]
		} else {
			min, max = test[mins[5]], test[mins[3]]
		}
		if min == max {
			return -1
		}
	}
}

func monkeyCheck(src map[string]Yell, humn int) (int, bool) {
	monkeys := make(map[string]Yell)
	for k, v := range src {
		monkeys[k] = v
	}
	h := monkeys["humn"]
	h.result = humn
	monkeys["humn"] = h
	for {
		for n, m := range monkeys {
			if !m.set && monkeys[m.targets[0]].set && monkeys[m.targets[1]].set {
				switch m.operation {
				case '+':
					m.result = monkeys[m.targets[0]].result + monkeys[m.targets[1]].result
					m.set = true
					monkeys[n] = m
				case '-':
					m.result = monkeys[m.targets[0]].result - monkeys[m.targets[1]].result
					m.set = true
					monkeys[n] = m
				case '*':
					m.result = monkeys[m.targets[0]].result * monkeys[m.targets[1]].result
					m.set = true
					monkeys[n] = m
				case '/':
					if monkeys[m.targets[0]].result%monkeys[m.targets[1]].result > 0 {
						return 0, false
					}
					m.result = monkeys[m.targets[0]].result / monkeys[m.targets[1]].result
					m.set = true
					monkeys[n] = m
				case '=':
					return monkeys[m.targets[0]].result - monkeys[m.targets[1]].result, monkeys[m.targets[0]].result == monkeys[m.targets[1]].result
				}
			}
		}
	}
}
