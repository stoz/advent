package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_s2401(t *testing.T) {
	var s s2401

	// Sample data taken from Advent of Code website as is the same for all users
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
