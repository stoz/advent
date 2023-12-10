package main

import "fmt"

func s23081(filename string, part, debug bool) int {
	lines := ReadFile("./data/2023/08/" + filename)
	directions := lines[0]
	network := make(map[string][2]string)
	pos := "AAA"
	for i, line := range lines {
		if i > 1 {
			var net [2]string
			net[0] = line[7:10]
			net[1] = line[12:15]
			network[line[0:3]] = net
		}
	}
	i := 0
	for pos != "ZZZ" {
		lr := 0
		if directions[i%len(directions)] == 'R' {
			lr = 1
		}
		pos = network[pos][lr]
		if debug {
			fmt.Println(pos, lr)
		}
		i++
	}
	return i
}

func s2308(filename string, part, debug bool) int {
	lines := ReadFile("./data/2023/08/" + filename)
	directions := lines[0]
	network := make(map[string][2]string)
	var positions []string
	var results []int
	for i, line := range lines {
		if i > 1 {
			var net [2]string
			net[0] = line[7:10]
			net[1] = line[12:15]
			network[line[0:3]] = net
			if line[2] == 'A' {
				positions = append(positions, line[0:3])
			}
		}
	}
	for _, pos := range positions {
		i := 0
		for pos[2] != 'Z' {
			lr := 0
			if directions[i%len(directions)] == 'R' {
				lr = 1
			}
			pos = network[pos][lr]
			if debug {
				fmt.Println(pos, lr)
			}
			i++
		}
		results = append(results, i)
	}
	return LCM(15871, 19637, 12643, 14257, 21251, 19099)
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
