package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_s2401(t *testing.T) {
	var s s2401

	s.SetInput([]string{
		"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3",
	})
	s.SetPart(1)
	result, err := s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "11", result)
	s.SetPart(2)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "31", result)

	s.SetInput(ReadFile("./data/2024/01/input.txt"))
	s.SetPart(1)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "1970720", result)
	s.SetPart(2)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "17191599", result)
}

func Test_s2402(t *testing.T) {
	var s s2402

	// Sample data taken from Advent of Code website as is the same for all users
	s.SetInput([]string{
		"7 6 4 2 1",
		"1 2 7 8 9",
		"9 7 6 2 1",
		"1 3 2 4 5",
		"8 6 4 4 1",
		"1 3 6 7 9",
	})
	s.SetPart(1)
	result, err := s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "2", result)
	s.SetPart(2)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "4", result)

	s.SetInput(ReadFile("./data/2024/02/input.txt"))
	s.SetPart(1)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "486", result)
	s.SetPart(2)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "540", result)
}

func Test_s2403(t *testing.T) {
	var s s2403

	s.SetInput([]string{
		"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
	})
	s.SetPart(1)
	result, err := s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "161", result)

	s.SetInput([]string{
		"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
	})
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "161", result)
	s.SetPart(2)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "48", result)

	s.SetInput(ReadFile("./data/2024/03/input.txt"))
	s.SetPart(1)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "182780583", result)
	s.SetPart(2)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "90772405", result)
}

func Test_s2404(t *testing.T) {
	var s s2404

	s.SetInput([]string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	})
	s.SetPart(1)
	result, err := s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "18", result)
	s.SetPart(2)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "9", result)

	s.SetInput(ReadFile("./data/2024/04/input.txt"))
	s.SetPart(1)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "2297", result)
	s.SetPart(2)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "1745", result)
}

func Test_s2405(t *testing.T) {
	var s s2405

	s.SetInput([]string{
		"47|53",
		"97|13",
		"97|61",
		"97|47",
		"75|29",
		"61|13",
		"75|53",
		"29|13",
		"97|29",
		"53|29",
		"61|53",
		"97|53",
		"61|29",
		"47|13",
		"75|47",
		"97|75",
		"47|61",
		"75|61",
		"47|29",
		"75|13",
		"53|13",
		"",
		"75,47,61,53,29",
		"97,61,53,29,13",
		"75,29,13",
		"75,97,47,61,53",
		"61,13,29",
		"97,13,75,29,47",
	})
	s.SetPart(1)
	result, err := s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "143", result)
	s.SetPart(2)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "123", result)

	s.SetInput(ReadFile("./data/2024/05/input.txt"))
	s.SetPart(1)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "5651", result)
	s.SetPart(2)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "4743", result)
}

func Test_s2406(t *testing.T) {
	var s s2406

	s.SetInput([]string{
		"....#.....",
		".........#",
		"..........",
		"..#.......",
		".......#..",
		"..........",
		".#..^.....",
		"........#.",
		"#.........",
		"......#...",
	})
	s.SetPart(1)
	result, err := s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "41", result)
	s.SetPart(2)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "6", result)

	s.SetInput(ReadFile("./data/2024/06/input.txt"))
	s.SetPart(1)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "5162", result)
	s.SetPart(2)
	result, err = s.Solve()
	assert.NoError(t, err)
	assert.Equal(t, "1909", result)
}
