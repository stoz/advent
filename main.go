package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/alecthomas/kong"
)

type Context struct {
	Debug bool
}

type SolveCmd struct {
	Year      int  `arg:"" name:"year" help:"Year of the puzzle." type:"int"`
	Day       int  `arg:"" name:"day" help:"Day of the puzzle." type:"int"`
	Part      int  `arg:"" optional:"" name:"part" help:"Part of the puzzle (1 or 2)." type:"int"`
	Benchmark bool `short:"b" help:"Record runtime."`
	Sample    bool `short:"s" help:"Use sample data."`
}

type BenchmarkCmd struct {
	Year int `arg:"" name:"year" help:"Year of the puzzle." type:"int"`
}

var cli struct {
	Debug     bool         `help:"Enable debug mode."`
	Solve     SolveCmd     `cmd:"" help:"Solve a puzzle."`
	Benchmark BenchmarkCmd `cmd:"" help:"Benchmark all puzzles for a given year."`
}

func main() {
	ctx := kong.Parse(&cli)
	err := ctx.Run(&Context{Debug: cli.Debug})
	ctx.FatalIfErrorf(err)
}

func (r *SolveCmd) Run(ctx *Context) error {
	f := "input.txt"
	if r.Sample {
		f = "sample.txt"
	}
	var p bool
	if r.Part == 2 {
		p = true
	}
	start := time.Now()
	switch r.Year {
	case 2022, 22:
		switch r.Day {
		case 1:
			fmt.Println(s2201(f, p))
		case 2:
			switch r.Part {
			case 0, 1:
				fmt.Println(s22021(f))
			case 2:
				fmt.Println(s22022(f))
			}
		case 3:
			switch r.Part {
			case 0, 1:
				fmt.Println(s22031(f))
			case 2:
				fmt.Println(s22032(f))
			}
		case 4:
			fmt.Println(s2204(f, p))
		case 5:
			switch r.Part {
			case 0, 1:
				fmt.Println(s22051(f))
			case 2:
				fmt.Println(s22052(f))
			}
		case 6:
			fmt.Println(s2206(f, p))
		case 7:
			switch r.Part {
			case 0, 1:
				fmt.Println(s22071(f))
			case 2:
				fmt.Println(s22072(f))
			}
		case 8:
			switch r.Part {
			case 0, 1:
				fmt.Println(s22081(f))
			case 2:
				fmt.Println(s22082(f))
			}
		case 9:
			switch r.Part {
			case 0, 1:
				fmt.Println(s22091(f))
			case 2:
				fmt.Println(s22092(f))
			}
		case 10:
			switch r.Part {
			case 0, 1:
				fmt.Println(s22101(f))
			case 2:
				s22102(f, true)
			}
		case 11:
			fmt.Println(s2211(f, p))
		case 12:
			fmt.Println(s2212(f, p))
		case 13:
			switch r.Part {
			case 0, 1:
				fmt.Println(s22131(f))
			case 2:
				fmt.Println(s22132(f))
			}
		case 14:
			fmt.Println(s2214(f, p, cli.Debug))
		case 18:
			switch r.Part {
			case 0, 1:
				fmt.Println(s22181(f))
			case 2:
				fmt.Println(s22182(f))
			}
		case 19:
			fmt.Println(s2219(f, p, cli.Debug))
		case 20:
			fmt.Println(s2220(f, p, cli.Debug))
		case 21:
			switch r.Part {
			case 0, 1:
				fmt.Println(s22211(f))
			case 2:
				fmt.Println(s22212(f))
			}
		case 22:
			switch r.Part {
			case 0, 1:
				fmt.Println(s22221(f, cli.Debug))
			case 2:
				fmt.Println(s22222(f, cli.Debug))
			}
		case 23:
			fmt.Println(s2223(f, p, cli.Debug))
		case 24:
			fmt.Println(s2224(f, p, cli.Debug))
		}
	}
	if r.Benchmark {
		duration := time.Since(start)
		fmt.Println(duration)
	}
	return nil
}

func (r *BenchmarkCmd) Run(ctx *Context) error {
	f := "input.txt"
	switch r.Year {
	case 2022, 22:
		var end [2]time.Duration
		// Day 1
		start := time.Now()
		s2201(f, false)
		end[0] = time.Since(start)
		start = time.Now()
		s2201(f, true)
		end[1] = time.Since(start)
		fmt.Println("Day 01", end[0], end[1])
		// Day 2
		start = time.Now()
		s22021(f)
		end[0] = time.Since(start)
		start = time.Now()
		s22022(f)
		end[1] = time.Since(start)
		fmt.Println("Day 02", end[0], end[1])
		// Day 3
		start = time.Now()
		s22031(f)
		end[0] = time.Since(start)
		start = time.Now()
		s22032(f)
		end[1] = time.Since(start)
		fmt.Println("Day 03", end[0], end[1])
		// Day 4
		start = time.Now()
		s2204(f, false)
		end[0] = time.Since(start)
		start = time.Now()
		s2204(f, true)
		end[1] = time.Since(start)
		fmt.Println("Day 04", end[0], end[1])
		// Day 5
		start = time.Now()
		s22051(f)
		end[0] = time.Since(start)
		start = time.Now()
		s22052(f)
		end[1] = time.Since(start)
		fmt.Println("Day 05", end[0], end[1])
		// Day 6
		start = time.Now()
		s2206(f, false)
		end[0] = time.Since(start)
		start = time.Now()
		s2206(f, true)
		end[1] = time.Since(start)
		fmt.Println("Day 06", end[0], end[1])
		// Day 7
		start = time.Now()
		s22071(f)
		end[0] = time.Since(start)
		start = time.Now()
		s22072(f)
		end[1] = time.Since(start)
		fmt.Println("Day 07", end[0], end[1])
		// Day 8
		start = time.Now()
		s22081(f)
		end[0] = time.Since(start)
		start = time.Now()
		s22082(f)
		end[1] = time.Since(start)
		fmt.Println("Day 08", end[0], end[1])
		// Day 9
		start = time.Now()
		s22091(f)
		end[0] = time.Since(start)
		start = time.Now()
		s22092(f)
		end[1] = time.Since(start)
		fmt.Println("Day 09", end[0], end[1])
		// Day 10
		start = time.Now()
		s22101(f)
		end[0] = time.Since(start)
		start = time.Now()
		s22102(f, false)
		end[1] = time.Since(start)
		fmt.Println("Day 10", end[0], end[1])
		// Day 11
		start = time.Now()
		s2211(f, false)
		end[0] = time.Since(start)
		start = time.Now()
		s2211(f, true)
		end[1] = time.Since(start)
		fmt.Println("Day 11", end[0], end[1])
		// Day 12
		start = time.Now()
		s2212(f, false)
		end[0] = time.Since(start)
		start = time.Now()
		s2212(f, true)
		end[1] = time.Since(start)
		fmt.Println("Day 12", end[0], end[1])
		// Day 13
		start = time.Now()
		s22131(f)
		end[0] = time.Since(start)
		start = time.Now()
		s22132(f)
		end[1] = time.Since(start)
		fmt.Println("Day 13", end[0], end[1])
		// Day 14
		start = time.Now()
		s2214(f, false, false)
		end[0] = time.Since(start)
		start = time.Now()
		s2214(f, true, false)
		end[1] = time.Since(start)
		fmt.Println("Day 14", end[0], end[1])
		// Day 18
		start = time.Now()
		s22181(f)
		end[0] = time.Since(start)
		start = time.Now()
		s22182(f)
		end[1] = time.Since(start)
		fmt.Println("Day 18", end[0], end[1])
		// Day 19
		start = time.Now()
		s2219(f, false, false)
		end[0] = time.Since(start)
		start = time.Now()
		s2219(f, true, false)
		end[1] = time.Since(start)
		fmt.Println("Day 19", end[0], end[1])
		// Day 20
		start = time.Now()
		s2220(f, false, false)
		end[0] = time.Since(start)
		start = time.Now()
		s2220(f, true, false)
		end[1] = time.Since(start)
		fmt.Println("Day 20", end[0], end[1])
		// Day 21
		start = time.Now()
		s22211(f)
		end[0] = time.Since(start)
		start = time.Now()
		s22212(f)
		end[1] = time.Since(start)
		fmt.Println("Day 21", end[0], end[1])
		// Day 22
		start = time.Now()
		s22221(f, false)
		end[0] = time.Since(start)
		start = time.Now()
		s22222(f, false)
		end[1] = time.Since(start)
		fmt.Println("Day 22", end[0], end[1])
		// Day 23
		start = time.Now()
		s2223(f, false, false)
		end[0] = time.Since(start)
		start = time.Now()
		s2223(f, true, false)
		end[1] = time.Since(start)
		fmt.Println("Day 23", end[0], end[1])
		// Day 24
		start = time.Now()
		s2224(f, false, false)
		end[0] = time.Since(start)
		start = time.Now()
		s2224(f, true, false)
		end[1] = time.Since(start)
		fmt.Println("Day 23", end[0], end[1])
	}
	return nil
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func binary() {
	// expected result: 3923414
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
	// expected result: 1567
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
	// expected result: 1529
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
	grid := MakeGridRune(lines)
	return grid
}

func MakeGridRune(lines []string) [][]rune {
	var max int
	for _, line := range lines {
		if len(line) > max {
			max = len(line)
		}
	}
	grid := make([][]rune, len(lines))
	for y, line := range lines {
		grid[y] = make([]rune, max)
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

func ExtractUint8s(s string) []uint8 {
	// https://stackoverflow.com/questions/32987215/find-numbers-in-string-using-golang-regexp
	re := regexp.MustCompile("[0-9]+")
	matches := re.FindAllString(s, -1)
	var ints []uint8
	for _, m := range matches {
		i, err := strconv.Atoi(m)
		if err != nil {
			log.Fatal(err)
		}
		ints = append(ints, uint8(i))
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
}

type StrPointOld struct {
	c int
	n []string
	h []string
}

func DijkstraPaths(g map[string]StrPoint, t string) int {
	// data structure a map of [y][x] coordinates to points
	// each point has a cost (c) and 0 or more neighbours (n)
	do := true
	var result int
	for do {
		// start from the lowest-cost square
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
		if shortest == 999999 {
			do = false
			fmt.Println("Couldn't find path")
			break
		}

		cost := point.c + 1
		for _, n := range point.n {
			nbh, ok := g[n]
			if ok {
				if n == t {
					return cost
				} else if cost < g[n].c {
					nbh.c = cost
					g[n] = nbh
				}
			}
		}
		delete(g, cur)
	}
	return result
}

func DijkstraPathsOld(g map[string]StrPointOld, t string) []string {
	// data structure a map of [y][x] coordinates to points
	// each point has a cost (c) and 0 or more neighbours (n)
	do := true
	var result []string
	for do {
		// start from the lowest-cost square
		shortest := 999999
		var cur string
		var point StrPointOld
		// select shorted
		for s, c := range g {
			if c.c < shortest {
				cur = s
				point = c
				shortest = c.c
			}
		}
		if shortest == 999999 {
			do = false
			fmt.Println("Couldn't find path")
			break
		}

		cost := point.c + 1
		for _, n := range point.n {
			nbh, ok := g[n]
			if ok {
				if n == t {
					result = point.h
					result = append(result, n)
					return result
				} else if cost < g[n].c {
					nbh.c = cost
					nbh.h = point.h
					nbh.h = append(nbh.h, n)
					g[n] = nbh
				}
			}
		}
		delete(g, cur)
	}
	return result
}
