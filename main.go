package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	rope2()
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
