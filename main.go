package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//start := time.Now()
	xyz4()
	//duration := time.Since(start)
	//fmt.Println(duration.Nanoseconds())
}

type Valve struct {
	p int
	n []string
}

func xyz4() {
	lines := ReadFile("sample.txt")
	valves := make(map[string]Valve)
	grid := make(map[string]StrPoint)
	var visit []string
	for _, line := range lines {
		var valve Valve
		var point StrPoint
		words := strings.Fields(line)
		ints := ExtractInts(line)
		v := words[1]
		valve.p = ints[0]
		point.c = 999999
		for i, w := range words {
			if i > 8 {
				valve.n = append(valve.n, w)
				point.n = append(point.n, w)
			}
		}
		if valve.p > 0 {
			visit = append(visit, v)
		}
		valves[v] = valve
		grid[v] = point
	}
	//max := 0
	//pos := "AA"

	fmt.Println(valves)
	fmt.Println(visit)
	quickest := make(map[string]map[string][]string)
	for _, v := range visit {
		quickest[v] = make(map[string][]string)
		fmt.Println("Find routes for: " + v)
		for _, t := range visit {
			if v != t {
				gridCopy := make(map[string]StrPoint)
				for a, b := range grid {
					gridCopy[a] = b
				}
				start := gridCopy[v]
				start.c = 0
				gridCopy[v] = start
				fmt.Println(gridCopy)
				quickest[v][t] = DijkstraPaths(gridCopy, t)
				fmt.Println("Find routes for: " + v + " to " + t)
				fmt.Println(quickest[v][t])
			}
		}
	}
	//fmt.Println(quickest)
	for s, q := range quickest {
		fmt.Println(s, q)
	}
}

func xyz3() {
	// 4350538 too low
	// 4352518 too low
	// 4370518 too low
	lines := ReadFile("input.txt")
	var dim [4]int
	cave := make(map[int]map[int]int)
	sens := make(map[int]map[int]int)
	for _, line := range lines {
		ints := ExtractSInts(line)
		for i, j := range ints {
			switch i % 2 {
			case 0: //x
				if j < dim[2] || dim[2] == 0 {
					dim[2] = j - 10000000
				} else if j > dim[3] {
					dim[3] = j + 10000000
				}
			case 1: //y
				if j < dim[0] || dim[0] == 0 {
					dim[0] = j - 10
				} else if j > dim[1] {
					dim[1] = j + 10
				}
			}
		}
	}
	for i := dim[0]; i <= dim[1]; i++ {
		cave[i] = make(map[int]int)
	}
	for _, line := range lines {
		ints := ExtractSInts(line)
		var x int
		var sensor [2]int
		for i, y := range ints {
			switch i % 4 {
			case 0, 2:
				x = y
			case 1:
				cave[y][x] = 3
				sensor[0] = y
				sensor[1] = x
			case 3:
				cave[y][x] = 4
				_, ok := sens[sensor[0]]
				if !ok {
					sens[sensor[0]] = make(map[int]int)
				}
				var diff int
				if sensor[0] < y {
					diff += y - sensor[0]
				} else {
					diff += sensor[0] - y
				}
				if sensor[1] < x {
					diff += x - sensor[1]
				} else {
					diff += sensor[1] - x
				}
				sens[sensor[0]][sensor[1]] = diff
			}
		}
	}
	fmt.Println(dim)
	// 4000000
	var lim = 4000000
	var fin []int
	cache := make(map[int]bool)

	//13639963409990 too high
	for sy, line := range sens {
		for sx, d := range line {
			// check around the edge
			e := d + 1
			for y := sy - e; y <= sy+e; y++ {
				if y >= 0 && y <= lim {
					var test [2]int
					if y < sy {
						test[0] = sx - e + sy - y
						test[1] = sx + e - sy - y
					} else {
						test[0] = sx - e + y - sy
						test[1] = sx + e - y - sy
					}
					for _, x := range test {
						if x >= 0 && x <= lim {
							_, ok := sens[x*lim+y]
							if !ok {
								var diff int
								if sy < y {
									diff += y - sy
								} else {
									diff += sy - y
								}
								if sx < x {
									diff += x - sx
								} else {
									diff += sx - x
								}
								if diff == e {
									safe := InRangeOfSensor(y, x, sens)
									cache[x*lim+y] = true
									if safe == 0 {
										fin = append(fin, y, x)
										fmt.Println("!!!", x*lim+y)
									}
								}
							}
						}
					}
				}
			}
		}
		if len(fin) > 2 {
			fmt.Println("too many", fin)
			fmt.Println("---", fin[1]*lim+fin[0])
			os.Exit(1)
		}
	}
	fmt.Println(fin)
	if len(fin) == 2 {
		fmt.Print(fin[1]*lim + fin[0])
	}
}

func InRangeOfSensor(y, x int, sens map[int]map[int]int) int {
	for sy, line := range sens {
		for sx, d := range line {
			if abs(sy-y)+abs(sx-x) <= d {
				return 1
			}
		}
	}
	return 0
}

func xyz2() {
	// 4350538 too low
	// 4352518 too low
	// 4370518 too low
	lines := ReadFile("input.txt")
	var dim [4]int
	cave := make(map[int]map[int]int)
	sens := make(map[int]map[int]int)
	for _, line := range lines {
		ints := ExtractInts(line)
		for i, j := range ints {
			switch i % 2 {
			case 0: //x
				if j < dim[2] || dim[2] == 0 {
					dim[2] = j - 10000000
				} else if j > dim[3] {
					dim[3] = j + 10000000
				}
			case 1: //y
				if j < dim[0] || dim[0] == 0 {
					dim[0] = j - 10
				} else if j > dim[1] {
					dim[1] = j + 10
				}
			}
		}
	}
	for i := dim[0]; i <= dim[1]; i++ {
		cave[i] = make(map[int]int)
	}
	for _, line := range lines {
		ints := ExtractInts(line)
		var x int
		var sensor [2]int
		for i, y := range ints {
			switch i % 4 {
			case 0, 2:
				x = y
			case 1:
				cave[y][x] = 3
				sensor[0] = y
				sensor[1] = x
			case 3:
				cave[y][x] = 4
				_, ok := sens[sensor[0]]
				if !ok {
					sens[sensor[0]] = make(map[int]int)
				}
				sens[sensor[0]][sensor[1]] = abs(sensor[0]-y) + abs(sensor[1]-x)
			}
		}
	}
	fmt.Println(dim)
	// PrintCave2(cave, dim)
	var sum int
	y := 2000000
	for x := dim[2]; x <= dim[3]; x++ {
		value, ok := cave[y][x]
		safe := 0
		if !ok {
			// square is empty
			// check if any sensors are in range
			safe = InRangeOfSensor(y, x, sens)
		} else {
			if value == 3 {
				safe = 1
			}
		}
		sum += safe
	}
	fmt.Println(sum)
}

func xyz() {
	lines := ReadFile("input.txt")
	var dim [4]int
	cave := make(map[int]map[int]int)
	for _, line := range lines {
		ints := ExtractInts(line)
		for i, j := range ints {
			switch i % 2 {
			case 0: //x
				if j < dim[2] || dim[2] == 0 {
					dim[2] = j - 10
				} else if j > dim[3] {
					dim[3] = j + 10
				}
			case 1: //y
				if j < dim[0] || dim[0] == 0 {
					dim[0] = j - 10
				} else if j > dim[1] {
					dim[1] = j + 10
				}
			}
		}
	}
	for i := dim[0]; i <= dim[1]; i++ {
		cave[i] = make(map[int]int)
	}
	for _, line := range lines {
		ints := ExtractInts(line)
		var x int
		var sensor [2]int
		for i, y := range ints {
			fmt.Println(i)
			switch i % 4 {
			case 0, 2:
				x = y
			case 1:
				cave[y][x] = 3
				sensor[0] = y
				sensor[1] = x
			case 3:
				cave[y][x] = 4
				//measure distance
				d := abs(sensor[0]-y) + abs(sensor[1]-x)
				for dy := sensor[0] - d; dy <= sensor[0]+d; dy++ {
					// 1, 3, 5 etc
					width := d - abs(sensor[1]-dy)
					for dx := sensor[1] - width; dx <= sensor[1]+width; dx++ {
						_, ok := cave[dy]
						if ok {
							_, ok2 := cave[dy][dx]
							if !ok2 {
								cave[dy][dx] = 1
							}
						}
					}
				}
			}
		}
	}
	fmt.Println(dim)
	PrintCave2(cave, dim)
	var sum int
	for _, c := range cave[2000000] {
		if c == 1 {
			sum++
		}
	}
	//fmt.Println(sum)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func sand(part2, printCave bool) {
	// expected results: 808, 26625
	lines := ReadFile("sand.txt")
	cave := make(map[int]map[int]int)

	// get the dimensions
	var dim [3]int
	for _, line := range lines {
		ints := ExtractInts(line)
		for i, j := range ints {
			switch i % 2 {
			case 0: //x
				if j < dim[1] || dim[1] == 0 {
					dim[1] = j
				} else if j > dim[2] {
					dim[2] = j
				}
			case 1: //y
				if j > dim[0] {
					dim[0] = j
				}
			}
		}
	}
	if part2 {
		// add a floor 2 units below the lowest point
		dim[0] += 2
		dim[1] = 499 - dim[0]
		dim[2] = 501 + dim[0]
		lines = append(lines, strconv.Itoa(dim[1])+" "+strconv.Itoa(dim[0])+" "+strconv.Itoa(dim[2])+" "+strconv.Itoa(dim[0]))
	}
	for i := 0; i <= dim[0]; i++ {
		cave[i] = make(map[int]int)
	}
	for _, line := range lines {
		ints := ExtractInts(line)
		var s [2]int // start
		var x int
		for i, j := range ints {
			switch i % 2 {
			case 0:
				x = j
			case 1:
				if s[0] != 0 {
					// do the needful
					if s[0] < j {
						for v := s[0]; v <= j; v++ {
							cave[v][x] = 1
						}
					} else if j < s[0] {
						for v := j; v <= s[0]; v++ {
							cave[v][x] = 1
						}
					} else if s[1] < x {
						for v := s[1]; v <= x; v++ {
							cave[j][v] = 1
						}
					} else if x < s[1] {
						for v := x; v <= s[1]; v++ {
							cave[j][v] = 1
						}
					}
				}
				s[0] = j
				s[1] = x
			}
		}
	}
	do := true
	count := 0
	for do {
		y := 0
		x := 500
		do2 := true
		for do2 {
			if y > dim[0] {
				do2 = false
				do = false
			}
			_, down := cave[y+1][x]
			if down {
				_, left := cave[y+1][x-1]
				if left {
					_, right := cave[y+1][x+1]
					if right {
						count++
						if y > 0 {
							cave[y][x] = 2
						} else {
							do = false
						}
						do2 = false
					} else {
						y++
						x++
					}
				} else {
					y++
					x--
				}
			} else {
				y++
			}
		}
	}
	if printCave {
		PrintCave(cave, dim)
	}
	fmt.Println(count)
}

func PrintCave2(cave map[int]map[int]int, dim [4]int) {
	for y := dim[0]; y <= dim[1]; y++ {
		var line string
		for x := dim[2]; x <= dim[3]; x++ {
			c := "?"
			i, ok := cave[y][x]
			if ok {
				switch i {
				case 1:
					c = "#"
				case 2:
					c = "o"
				case 3:
					c = "S"
				case 4:
					c = "B"
				}
			} else {
				c = "."
			}
			line += c
		}
		fmt.Println(line, y)
	}
}

func PrintCave(cave map[int]map[int]int, dim [3]int) {
	for y := 0; y <= dim[0]; y++ {
		var line string
		for x := dim[1]; x <= dim[2]; x++ {
			c := "?"
			i, ok := cave[y][x]
			if ok {
				switch i {
				case 1:
					c = "#"
				case 2:
					c = "o"
				case 3:
					c = "S"
				case 4:
					c = "B"
				}
			} else {
				c = "."
			}
			line += c
		}
		fmt.Println(line)
	}
}

func distress2() {
	lines := ReadFile("distress.txt")
	var linesInterface []any
	add := []string{"2", "6"}
	for _, a := range add {
		lines = append(lines, "[["+a+"]]")
	}
	for _, line := range lines {
		if line != "" {
			var data any
			err := json.Unmarshal([]byte(line), &data)
			if err != nil {
				log.Fatal("Cannot unmarshal the json ", err)
			}
			linesInterface = append(linesInterface, data)
		}
	}

	// sort using the CompareDistress function
	sort.Slice(linesInterface, func(i, j int) bool {
		return CompareDistress(linesInterface[i], linesInterface[j]) < 0
	})

	fmt.Println(linesInterface)

	sum := 1
	for i, line := range linesInterface {
		data, _ := json.Marshal(line)
		for _, a := range add {
			if string(data) == "[["+a+"]]" {
				fmt.Println(sum, i+1)
				sum *= i + 1
			}
		}
	}
	fmt.Println(sum)
}

func distress() {
	lines := ReadFile("distress.txt")
	var lline, rline string
	var count int

	for i, line := range lines {
		switch i % 3 {
		case 0:
			lline = line
		case 1:
			rline = line
			// comparison
			var ldata, rdata any
			err := json.Unmarshal([]byte(lline), &ldata)
			if err != nil {
				log.Fatal("Cannot unmarshal the json ", err)
			}
			err = json.Unmarshal([]byte(rline), &rdata)
			if err != nil {
				log.Fatal("Cannot unmarshal the json ", err)
			}
			if CompareDistress(ldata, rdata) <= 0 {
				count += ((i - i%3) / 3) + 1
			}
		}
	}
	fmt.Println(count)
}

func CompareDistress(left, right any) int {
	// float64 type assertion
	l, lok := left.(float64)
	r, rok := right.(float64)
	// if both are float64, return the difference
	if lok && rok {
		return int(l) - int(r)
	}

	var llist, rlist []any
	switch left.(type) {
	case []any, []float64:
		llist = left.([]any)
	case float64:
		llist = []any{left}
	}
	switch right.(type) {
	case []any, []float64:
		rlist = right.([]any)
	case float64:
		rlist = []any{right}
	}

	for i := range llist {
		if len(rlist) <= i {
			return 1
		}
		if ret := CompareDistress(llist[i], rlist[i]); ret != 0 {
			return ret
		}
	}
	if len(llist) == len(rlist) {
		return 0
	}
	return -1
}

func hill(pt2 bool) {
	// expected results: 534, 525
	lines := ReadGridRune("hill.txt")
	grid := make(map[int]map[int]GridPoint)
	max := [2]int{len(lines), len(lines[0])}
	var start [2]int
	var target [2]int

	// pre-process to replace S and E with a and z
	for y, line := range lines {
		for x, a := range line {
			switch a {
			case 'S':
				lines[y][x] = 'a'
				start[0] = y
				start[1] = x
			case 'E':
				lines[y][x] = 'z'
				target[0] = y
				target[1] = x
			}
		}
	}
	for y, line := range lines {
		grid[y] = make(map[int]GridPoint)
		for x, a := range line {
			var point GridPoint
			if y == start[0] && x == start[1] {
				point.c = 0
			} else if pt2 && a == 'a' {
				point.c = 0
			} else {
				point.c = 999999
			}
			var nbh [4][2]int
			for i := 0; i < len(nbh); i++ {
				nbh[i][0] = y
				nbh[i][1] = x
			}
			nbh[0][0]--
			nbh[1][0]++
			nbh[2][1]--
			nbh[3][1]++
			for _, n := range nbh {
				if n[0] > -1 && n[1] > -1 && n[0] < max[0] && n[1] < max[1] && lines[n[0]][n[1]] < a+2 {
					point.n = append(point.n, n)
				}
			}
			grid[y][x] = point
		}
	}
	result := Dijkstra(grid, target)
	fmt.Println(result)
}

type Monkey struct {
	items       []int
	operation   int
	modulus     int
	targets     [2]int
	inspections int
}

func monkey2() {
	// not 14636993466
	lines := ReadFile("monkey.txt")
	monkeys := new([8]Monkey)
	index := 0
	for i, line := range lines {
		ints := ExtractInts(line)
		switch i % 7 {
		case 0:
			index = ints[0]
		case 1:
			monkeys[index].items = ints
		case 2:
			words := strings.Fields(line)
			if words[5] == "old" {
				monkeys[index].operation = 0
			} else if words[4] == "*" {
				monkeys[index].operation = -ints[0]
			} else {
				monkeys[index].operation = ints[0]
			}
		case 3:
			monkeys[index].modulus = ints[0]
		case 4:
			monkeys[index].targets[0] = ints[0]
		case 5:
			monkeys[index].targets[1] = ints[0]
		}
	}
	modulus := 1
	for _, monkey := range monkeys {
		modulus *= monkey.modulus
	}
	for i := 0; i < 10000; i++ {
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
				new = new % modulus
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
	fmt.Println(max[0] * max[1])
}

func monkey() {
	// expected result: 55930
	lines := ReadFile("monkey.txt")
	monkeys := new([8]Monkey)
	index := 0
	for i, line := range lines {
		ints := ExtractInts(line)
		switch i % 7 {
		case 0:
			index = ints[0]
		case 1:
			monkeys[index].items = ints
		case 2:
			words := strings.Fields(line)
			if words[5] == "old" {
				monkeys[index].operation = 0
			} else if words[4] == "*" {
				monkeys[index].operation = -ints[0]
			} else {
				monkeys[index].operation = ints[0]
			}
		case 3:
			monkeys[index].modulus = ints[0]
		case 4:
			monkeys[index].targets[0] = ints[0]
		case 5:
			monkeys[index].targets[1] = ints[0]
		}
	}
	for i := 0; i < 20; i++ {
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
				new = FloorDivision(new, 3)
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
	fmt.Println(max[0] * max[1])
}

func crt2() {
	// expected result: BJFRHRFU
	lines := ReadFile("crt.txt")
	x := 1
	width := 40
	output := ""
	for _, line := range lines {
		f := strings.Fields(line)
		cycles := 1
		add := 0
		if f[0] == "addx" {
			cycles = 2
			add, _ = strconv.Atoi(f[1])
		}
		for i := 0; i < cycles; i++ {
			c := len([]rune(output))
			if c == x-1 || c == x || c == x+1 {
				output += "█"
			} else {
				output += " "
			}
			if c == width-1 {
				fmt.Println(output)
				output = ""
			}
		}
		x += add
	}
}

func crt() {
	// expected result: 14620
	lines := ReadFile("crt.txt")
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
	fmt.Println(sum)
}

func rope2() {
	// expected result: 2607
	lines := ReadFile("rope.txt")
	x := [10]int{1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000}
	y := [10]int{1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000}
	factor := 10000
	var visited [100000000]bool
	visited[1000*factor+1000] = true
	for _, line := range lines {
		f := strings.Fields(line)
		moves, err := strconv.Atoi(f[1])
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < moves; i++ {
			if f[0] == "U" {
				x[0]--
			} else if f[0] == "D" {
				x[0]++
			} else if f[0] == "L" {
				y[0]--
			} else if f[0] == "R" {
				y[0]++
			}
			for z := 0; z < 9; z++ {
				// move if too far away
				if x[z]-x[z+1] > 1 {
					x[z+1]++
					if y[z]-y[z+1] > 0 {
						y[z+1]++
					} else if y[z+1]-y[z] > 0 {
						y[z+1]--
					}
				} else if x[z+1]-x[z] > 1 {
					x[z+1]--
					if y[z]-y[z+1] > 0 {
						y[z+1]++
					} else if y[z+1]-y[z] > 0 {
						y[z+1]--
					}
				} else if y[z]-y[z+1] > 1 {
					y[z+1]++
					if x[z]-x[z+1] > 0 {
						x[z+1]++
					} else if x[z+1]-x[z] > 0 {
						x[z+1]--
					}
				} else if y[z+1]-y[z] > 1 {
					y[z+1]--
					if x[z]-x[z+1] > 0 {
						x[z+1]++
					} else if x[z+1]-x[z] > 0 {
						x[z+1]--
					}
				}
			}
			//fmt.Println(tx, ty)
			visited[x[9]*factor+y[9]] = true
		}
	}
	count := 0
	for _, v := range visited {
		if v {
			count++
		}
	}
	fmt.Println(count)
}

func rope() {
	// expected result: 6376
	lines := ReadFile("rope.txt")
	hx := 1000
	hy := 1000
	tx := 1000
	ty := 1000
	factor := 10000
	var visited [100000000]bool
	visited[1000*factor+1000] = true
	for _, line := range lines {
		f := strings.Fields(line)
		moves, err := strconv.Atoi(f[1])
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < moves; i++ {
			if f[0] == "U" {
				hx--
				fmt.Println("U")
			} else if f[0] == "D" {
				fmt.Println("D")
				hx++
			} else if f[0] == "L" {
				fmt.Println("L")
				hy--
			} else if f[0] == "R" {
				fmt.Println("R")
				hy++
			}
			// move if too far away
			if hx-tx > 1 {
				tx++
				if hy-ty > 0 {
					ty++
				} else if ty-hy > 0 {
					ty--
				}
			} else if tx-hx > 1 {
				tx--
				if hy-ty > 0 {
					ty++
				} else if ty-hy > 0 {
					ty--
				}
			} else if hy-ty > 1 {
				ty++
				if hx-tx > 0 {
					tx++
				} else if tx-hx > 0 {
					tx--
				}
			} else if ty-hy > 1 {
				ty--
				if hx-tx > 0 {
					tx++
				} else if tx-hx > 0 {
					tx--
				}
			}
			//fmt.Println(tx, ty)
			visited[tx*factor+ty] = true
		}
	}
	count := 0
	for _, v := range visited {
		if v {
			count++
		}
	}
	fmt.Println(count)
}

func trees2() {
	// expected result: 263670
	grid := ReadGridInt("trees.txt")
	count := 0
	for y, row := range grid {
		for x, a := range row {
			// skip if any are on the edge of the grid
			if y != 0 && y != len(grid)-1 && x != 0 && x != len(row)-1 {
				up := 0
				for i := x - 1; i >= 0; i-- {
					up++
					if grid[y][i] >= a {
						break
					}
				}
				down := 0
				for i := x + 1; i < len(row); i++ {
					down++
					if grid[y][i] >= a {
						break
					}
				}
				left := 0
				for i := y - 1; i >= 0; i-- {
					left++
					if grid[i][x] >= a {
						break
					}
				}
				right := 0
				for i := y + 1; i < len(grid); i++ {
					right++
					if grid[i][x] >= a {
						break
					}
				}
				score := up * down * left * right
				if score > count {
					count = score
				}
			}
		}
	}
	fmt.Println(count)
}

func trees() {
	// expected result: 1835
	grid := ReadGridInt("trees.txt")
	count := 0
	for y, row := range grid {
		for x, a := range row {
			// outside trees
			if y == 0 || y == len(grid)-1 || x == 0 || x == len(row)-1 {
				count++
			} else {
				up := true
				for i := x - 1; i >= 0; i-- {
					if grid[y][i] >= a {
						up = false
					}
				}
				down := true
				for i := x + 1; i < len(row); i++ {
					if grid[y][i] >= a {
						down = false
					}
				}
				left := true
				for i := y - 1; i >= 0; i-- {
					if grid[i][x] >= a {
						left = false
					}
				}
				right := true
				for i := y + 1; i < len(grid); i++ {
					if grid[i][x] >= a {
						right = false
					}
				}
				if up || down || left || right {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}

func space2() {
	// expected result: 2050735
	lines := ReadFile("space.txt")
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
	fmt.Println(smallest)
}

func space() {
	lines := ReadFile("space.txt")
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
	fmt.Println(sum)
}

func tuning2() {
	// expeced result: 3605
	line := ReadLine("tuning.txt")
	fmt.Println(UniqueWindow(line, 14))
}

func tuning() {
	// expected result: 1275
	line := ReadLine("tuning.txt")
	fmt.Println(UniqueWindow(line, 4))
}

func stacks2() {
	// expected result: LBBVJBRMH
	lines := ReadFile("stacks-mod.txt")
	var boxes = map[int]map[int]string{
		1: {0: "F", 1: "T", 2: "C", 3: "L", 4: "R", 5: "P", 6: "G", 7: "Q"},
		2: {0: "N", 1: "Q", 2: "H", 3: "W", 4: "R", 5: "F", 6: "S", 7: "J"},
		3: {0: "F", 1: "B", 2: "H", 3: "W", 4: "P", 5: "M", 6: "Q"},
		4: {0: "V", 1: "S", 2: "T", 3: "D", 4: "F"},
		5: {0: "Q", 1: "L", 2: "D", 3: "W", 4: "V", 5: "F", 6: "Z"},
		6: {0: "Z", 1: "C", 2: "L", 3: "S"},
		7: {0: "Z", 1: "B", 2: "M", 3: "V", 4: "D", 5: "F"},
		8: {0: "T", 1: "J", 2: "B"},
		9: {0: "Q", 1: "N", 2: "B", 3: "G", 4: "L", 5: "S", 6: "P", 7: "H"},
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
	fmt.Println(output)
}

func stacks() {
	// expected result: VGBBJCRMN
	lines := ReadFile("stacks-mod.txt")
	var boxes = map[int]map[int]string{
		1: {0: "F", 1: "T", 2: "C", 3: "L", 4: "R", 5: "P", 6: "G", 7: "Q"},
		2: {0: "N", 1: "Q", 2: "H", 3: "W", 4: "R", 5: "F", 6: "S", 7: "J"},
		3: {0: "F", 1: "B", 2: "H", 3: "W", 4: "P", 5: "M", 6: "Q"},
		4: {0: "V", 1: "S", 2: "T", 3: "D", 4: "F"},
		5: {0: "Q", 1: "L", 2: "D", 3: "W", 4: "V", 5: "F", 6: "Z"},
		6: {0: "Z", 1: "C", 2: "L", 3: "S"},
		7: {0: "Z", 1: "B", 2: "M", 3: "V", 4: "D", 5: "F"},
		8: {0: "T", 1: "J", 2: "B"},
		9: {0: "Q", 1: "N", 2: "B", 3: "G", 4: "L", 5: "S", 6: "P", 7: "H"},
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
	fmt.Println(output)
}

func cleanup2() {
	// expected result: 900
	lines := ReadFile("cleanup-mod.txt")
	count := 0
	for _, line := range lines {
		words := strings.Fields(line)
		a, _ := strconv.Atoi(words[0])
		b, _ := strconv.Atoi(words[1])
		c, _ := strconv.Atoi(words[2])
		d, _ := strconv.Atoi(words[3])
		if (a <= d && b >= c) || (d <= a && c >= b) {
			count++
		}
	}
	fmt.Println(count)
}

func cleanup() {
	// expected result: 542
	lines := ReadFile("cleanup-mod.txt")
	count := 0
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for _, line := range lines {
		words := strings.Fields(line)
		a, _ := strconv.Atoi(words[0])
		b, _ := strconv.Atoi(words[1])
		c, _ := strconv.Atoi(words[2])
		d, _ := strconv.Atoi(words[3])
		if (a >= c && b <= d) || (c >= a && d <= b) {
			count++
		}
	}
	fmt.Println(count)
}

func rucksack2() {
	// expected result: 2607
	lines := ReadFile("rucksack.txt")
	count := 0
	i := 0
	var elves [3]string
	for _, line := range lines {
		elves[i] = line
		if i > 1 {
			matched := false
			for _, a := range elves[0] {
				for _, b := range elves[1] {
					for _, c := range elves[2] {
						if a == b && a == c && !matched {
							matched = true
							// Go runes, A-Z = 65-90, a-z=97-122
							if a < 91 {
								// uppercase letter
								count += (int(a) - 38)
							} else {
								// lowercase letter
								count += (int(a) - 96)
							}
						}
					}
				}
			}
			i = 0
		} else {
			i++
		}
	}
	fmt.Println(count)
}

func rucksack() {
	// expected value: 7597
	lines := ReadFile("rucksack.txt")
	count := 0
	var chunks []string
	for _, line := range lines {
		chunks = Chunks(line, len(line)/2)
		matched := false
		for _, a := range chunks[0] {
			for _, b := range chunks[1] {
				if a == b && !matched {
					matched = true
					// Go runes, A-Z = 65-90, a-z=97-122
					if a < 91 {
						// uppercase letter
						count += (int(a) - 38)
					} else {
						// lowercase letter
						count += (int(a) - 96)
					}
				}
			}
		}
	}
	fmt.Println(count)
}

func rock2() {
	// expected value: 13693
	lines := ReadFile("rock.txt")
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
	fmt.Println(count)
}

func rock() {
	// expedcted valie: 13052
	lines := ReadFile("rock.txt")
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
	fmt.Println(count)
}

func calories() {
	// expected value: 212489
	lines := ReadFile("calories.txt")
	current := 0
	var max [3]int
	//a := []int{}
	for _, line := range lines {
		if line == "" {
			if current > max[0] {
				max[2] = max[1]
				max[1] = max[0]
				max[0] = current
			} else if current > max[1] {
				max[2] = max[1]
				max[1] = current
			} else if current > max[2] {
				max[2] = current
			}
			current = 0
		} else {
			textAsInt, _ := strconv.Atoi(line)
			current += textAsInt
		}
	}
	fmt.Println(max[0] + max[1] + max[2])
}

func binary() {
	// expected value: 3923414
	lines := ReadFile("binary.txt")
	var a [12]int
	gamma := 0
	epsilon := 0
	for _, line := range lines {
		for pos, char := range line {
			if char == rune('1') {
				a[pos]++
			} else {
				a[pos]--
			}
		}
	}
	for i, s := range a {
		factor := IntPow(2, 11-i)
		if s > 0 {
			gamma += factor
		} else {
			epsilon += factor
		}
	}
	fmt.Println(gamma * epsilon)
}

func sonar2() {
	// expected value: 1567
	lines := ReadFile("sonar.txt")
	w1 := 999999
	w2 := 999999
	w3 := 999999
	count := 0
	for _, line := range lines {
		textAsInt, _ := strconv.Atoi(line)
		if textAsInt > w1 {
			count++
		}
		w1 = w2
		w2 = w3
		w3 = textAsInt
	}
	fmt.Println(count)
}

func sonar() {
	// expected value: 1529
	lines := ReadFile("sonar.txt")
	buffer := 999999
	count := 0
	for _, line := range lines {
		textAsInt, _ := strconv.Atoi(line)
		if textAsInt > buffer {
			count++
		}
		buffer = textAsInt
	}
	fmt.Println(count)
}

func ReadFile(filename string) []string {
	// https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func ReadLine(filename string) string {
	lines := ReadFile(filename)
	return lines[0]
}

func IntPow(n, m int) int {
	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func ReadGrid(filename string) [][]string {
	lines := ReadFile(filename)
	grid := make([][]string, len(lines))
	for y, line := range lines {
		grid[y] = make([]string, len(line))
		for x, a := range line {
			grid[y][x] = string(a)
		}
	}
	return grid
}

func ReadGridInt(filename string) [][]int {
	lines := ReadFile(filename)
	grid := make([][]int, len(lines))
	for y, line := range lines {
		grid[y] = make([]int, len(line))
		for x, a := range line {
			var err error
			grid[y][x], err = strconv.Atoi(string(a))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return grid
}

func ReadGridRune(filename string) [][]rune {
	lines := ReadFile(filename)
	grid := make([][]rune, len(lines))
	for y, line := range lines {
		grid[y] = make([]rune, len(line))
		for x, a := range line {
			grid[y][x] = a
		}
	}
	return grid
}

func Chunks(s string, chunkSize int) []string {
	// https://stackoverflow.com/questions/25686109/split-string-by-length-in-golang
	if len(s) == 0 {
		return nil
	}
	if chunkSize >= len(s) {
		return []string{s}
	}
	var chunks []string = make([]string, 0, (len(s)-1)/chunkSize+1)
	currentLen := 0
	currentStart := 0
	for i := range s {
		if currentLen == chunkSize {
			chunks = append(chunks, s[currentStart:i])
			currentLen = 0
			currentStart = i
		}
		currentLen++
	}
	chunks = append(chunks, s[currentStart:])
	return chunks
}

func UniqueWindow(line string, window int) int {
	var r []rune
	pos := 0
	for k, v := range line {
		r = append(r, v)
		if len(r) > window {
			r = r[1:]
			dup := false
			for a, b := range r {
				for c, d := range r {
					if a != c && b == d {
						dup = true
						break
					}
				}
			}
			if !dup {
				pos = k + 1
				break
			}
		}
	}
	return pos
}

func ExtractSInts(s string) []int {
	// https://stackoverflow.com/questions/32987215/find-numbers-in-string-using-golang-regexp
	re := regexp.MustCompile(`[-]?\d[\d]*[\.]?[\d{2}]*`)
	matches := re.FindAllString(s, -1)
	var ints []int
	for _, m := range matches {
		i, err := strconv.Atoi(m)
		if err != nil {
			log.Fatal(err)
		}
		ints = append(ints, i)
	}
	return ints
}

func ExtractInts(s string) []int {
	// https://stackoverflow.com/questions/32987215/find-numbers-in-string-using-golang-regexp
	re := regexp.MustCompile("[0-9]+")
	matches := re.FindAllString(s, -1)
	var ints []int
	for _, m := range matches {
		i, err := strconv.Atoi(m)
		if err != nil {
			log.Fatal(err)
		}
		ints = append(ints, i)
	}
	return ints
}

func FloorDivision(i int, d int) int {
	i = i - (i % d)
	i = i / d
	return i
}

type GridPoint struct {
	c int
	n [][2]int
}

func Dijkstra(g map[int]map[int]GridPoint, t [2]int) int {
	// data structure a map of [y][x] coordinates to points
	// each point has a cost (c) and 0 or more neighbours (n)
	do := true
	result := 999999
	for do {
		// start from the lowes-cost square
		shortest := 999999
		var cur [2]int
		var point GridPoint
		for y, line := range g {
			for x, c := range line {
				if c.c < shortest {
					cur[0] = y
					cur[1] = x
					point = c
					shortest = c.c
				}
			}
		}
		if shortest == 999999 {
			do = false
			break
		}

		cost := point.c + 1
		for _, n := range point.n {
			nbh, ok := g[n[0]][n[1]]
			if ok {
				if n[0] == t[0] && n[1] == t[1] {
					result = cost
					do = false
				} else if cost < g[n[0]][n[1]].c {
					nbh.c = cost
					g[n[0]][n[1]] = nbh
				}
			}
		}
		delete(g[cur[0]], cur[1])
	}
	return result
}

type StrPoint struct {
	c int
	n []string
	h []string
}

func DijkstraPaths(g map[string]StrPoint, t string) []string {
	// data structure a map of [y][x] coordinates to points
	// each point has a cost (c) and 0 or more neighbours (n)
	do := true
	var result []string
	for do {
		// start from the lowes-cost square
		shortest := 999999
		var cur string
		var point StrPoint
		// select shorted
		for s, c := range g {
			if c.c < shortest {
				cur = s
				point = c
				shortest = c.c
			}
		}
		fmt.Println(t, cur, shortest, point, g)
		if shortest == 999999 {
			do = false
			fmt.Println("Couldn't find path")
			break
		}

		cost := point.c + 1
		fmt.Println("!!!", point.n)
		for _, n := range point.n {
			fmt.Println("---", n, t)
			nbh, ok := g[n]
			if ok {
				if n == t {
					result = point.h
					result = append(result, n)
					//fmt.Println(result)
					return result
				} else if cost < g[n].c {
					fmt.Println(t, cur, n)
					nbh.c = cost
					nbh.h = point.h
					nbh.h = append(nbh.h, n)
					g[n] = nbh
					//fmt.Println(g)
				}
			}
		}
		delete(g, cur)
	}
	return result
}
