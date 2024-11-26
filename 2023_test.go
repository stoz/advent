package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_s2301(t *testing.T) {
	var s s2301

	// Sample data taken from Advent of Code website as is the same for all users
	s.SetInput([]string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	})
	s.SetPart(1)
	result, err := s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "142", result)

	s.SetInput([]string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	})
	s.SetPart(2)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "281", result)

	s.SetInput(ReadFile("./data/2023/01/input.txt"))
	s.SetPart(1)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "55834", result)
	s.SetPart(2)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "53221", result)
}

func Test_s2302(t *testing.T) {
	var s s2302

	s.SetInput([]string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	})
	s.SetPart(1)
	result, err := s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "8", result)
	s.SetPart(2)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "2286", result)

	s.SetInput(ReadFile("./data/2023/02/input.txt"))
	s.SetPart(1)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "2278", result)
	s.SetPart(2)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "67953", result)
}

func Test_s2303(t *testing.T) {
	var s s2303

	s.SetInput([]string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	})
	s.SetPart(1)
	result, err := s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "4361", result)
	s.SetPart(2)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "467835", result)

	s.SetInput(ReadFile("./data/2023/03/input.txt"))
	s.SetPart(1)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "532331", result)
	s.SetPart(2)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "82301120", result)
}

func Test_s2304(t *testing.T) {
	var s s2304

	s.SetInput([]string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	})
	s.SetPart(1)
	result, err := s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "13", result)
	s.SetPart(2)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "30", result)

	s.SetInput(ReadFile("./data/2023/04/input.txt"))
	s.SetPart(1)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "19855", result)
	s.SetPart(2)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "10378710", result)
}

func Test_s2305(t *testing.T) {
	var s s2305

	s.SetInput([]string{
		"seeds: 79 14 55 13",
		"",
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
		"",
		"soil-to-fertilizer map:",
		"0 15 37",
		"37 52 2",
		"39 0 15",
		"",
		"fertilizer-to-water map:",
		"49 53 8",
		"0 11 42",
		"42 0 7",
		"57 7 4",
		"",
		"water-to-light map:",
		"88 18 7",
		"18 25 70",
		"",
		"light-to-temperature map:",
		"45 77 23",
		"81 45 19",
		"68 64 13",
		"",
		"temperature-to-humidity map:",
		"0 69 1",
		"1 0 69",
		"",
		"humidity-to-location map:",
		"60 56 37",
		"56 93 4",
	})
	s.SetPart(1)
	result, err := s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "35", result)
	s.SetPart(2)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "46", result)

	s.SetInput(ReadFile("./data/2023/05/input.txt"))
	s.SetPart(1)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "107430936", result)
	s.SetPart(2)
	// TODO: Fix part 2 to be much faster
	// result, err = s.Solve()
	// assert.NoError(t, err)
	// assert.Equal(t, "23738616")
}

func Test_s2306(t *testing.T) {
	var s s2306

	s.SetInput([]string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	})
	s.SetPart(1)
	result, err := s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "288", result)
	s.SetPart(2)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "71503", result)

	s.SetInput(ReadFile("./data/2023/06/input.txt"))
	s.SetPart(1)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "588588", result)
	s.SetPart(2)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "34655848", result)
}

func Test_s2307(t *testing.T) {
	var s s2307

	s.SetInput([]string{
		"32T3K 765",
		"T55J5 684",
		"KK677 28",
		"KTJJT 220",
		"QQQJA 483",
	})
	s.SetPart(1)
	result, err := s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "6440", result)
	s.SetPart(2)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "5905", result)

	s.SetInput(ReadFile("./data/2023/07/input.txt"))
	s.SetPart(1)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "248217452", result)
	s.SetPart(2)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "245576185", result)
}

func Test_s2308(t *testing.T) {
	var s s2308

	s.SetInput([]string{
		"RL",
		"",
		"AAA = (BBB, CCC)",
		"BBB = (DDD, EEE)",
		"CCC = (ZZZ, GGG)",
		"DDD = (DDD, DDD)",
		"EEE = (EEE, EEE)",
		"GGG = (GGG, GGG)",
		"ZZZ = (ZZZ, ZZZ)",
	})
	s.SetPart(1)
	result, err := s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "2", result)

	// Two different samples are provided for Day 8
	s.SetInput([]string{
		"LLR",
		"",
		"AAA = (BBB, BBB)",
		"BBB = (AAA, ZZZ)",
		"ZZZ = (ZZZ, ZZZ)",
	})
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "6", result)

	s.SetInput([]string{
		"LR",
		"",
		"11A = (11B, XXX)",
		"11B = (XXX, 11Z)",
		"11Z = (11B, XXX)",
		"22A = (22B, XXX)",
		"22B = (22C, 22C)",
		"22C = (22Z, 22Z)",
		"22Z = (22B, 22B)",
		"XXX = (XXX, XXX)",
	})
	s.SetPart(2)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "6", result)

	s.SetInput(ReadFile("./data/2023/08/input.txt"))
	s.SetPart(1)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "14257", result)
	s.SetPart(2)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "16187743689077", result)
}

func Test_s2309(t *testing.T) {
	var s s2309

	s.SetInput([]string{
		"0 3 6 9 12 15",
		"1 3 6 10 15 21",
		"10 13 16 21 30 45",
	})
	s.SetPart(1)
	result, err := s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "114", result)
	s.SetPart(2)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "2", result)

	s.SetInput(ReadFile("./data/2023/09/input.txt"))
	s.SetPart(1)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "2005352194", result)
	s.SetPart(2)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "1077", result)
}

func Test_s2310(t *testing.T) {
	var s s2310

	s.SetInput([]string{
		".....",
		".S-7.",
		".|.|.",
		".L-J.",
		".....",
	})
	s.SetPart(1)
	result, err := s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "4", result)

	s.SetInput([]string{
		"..F7.",
		".FJ|.",
		"SJ.L7",
		"|F--J",
		"LJ...",
	})
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "8", result)

	s.SetInput([]string{
		"...........",
		".S-------7.",
		".|F-----7|.",
		".||.....||.",
		".||.....||.",
		".|L-7.F-J|.",
		".|..|.|..|.",
		".L--J.L--J.",
		"...........",
	})
	s.SetPart(2)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "4", result)

	s.SetInput([]string{
		".F----7F7F7F7F-7....",
		".|F--7||||||||FJ....",
		".||.FJ||||||||L7....",
		"FJL7L7LJLJ||LJ.L-7..",
		"L--J.L7...LJS7F-7L7.",
		"....F-J..F7FJ|L7L7L7",
		"....L7.F7||L7|.L7L7|",
		".....|FJLJ|FJ|F7|.LJ",
		"....FJL-7.||.||||...",
		"....L---J.LJ.LJLJ...",
	})
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "8", result)

	s.SetInput([]string{
		"FF7FSF7F7F7F7F7F---7",
		"L|LJ||||||||||||F--J",
		"FL-7LJLJ||||||LJL-77",
		"F--JF--7||LJLJIF7FJ-",
		"L---JF-JLJIIIIFJLJJ7",
		"|F|F-JF---7IIIL7L|7|",
		"|FFJF7L7F-JF7IIL---7",
		"7-L-JL7||F7|L7F-7F7|",
		"L.L7LFJ|||||FJL7||LJ",
		"L7JLJL-JLJLJL--JLJ.L",
	})
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "10", result)

	s.SetInput(ReadFile("./data/2023/10/input.txt"))
	s.SetPart(1)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "6979", result)
	s.SetPart(2)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "443", result)
}
