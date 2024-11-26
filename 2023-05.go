package main

import (
	"fmt"
	"strconv"
)

type s2305 Puzzle

func (s *s2305) SetDebug(debug bool) error {
	s.Debug = debug
	return nil
}

func (s *s2305) SetInput(input []string) error {
	s.Input = input
	return nil
}

func (s *s2305) SetPart(part int) error {
	s.Part = part
	return nil
}

func (s *s2305) Solve() (string, error) {
	var number int
	if s.Part != 2 {
		number = s.processPart1()
	} else {
		number = s.processPart2()
	}
	return strconv.Itoa(number), nil
}

func (s *s2305) processPart1() int {
	var maps [7][][3]int
	// build the maps
	mapIndex := -1
	for lineIndex, line := range s.Input {
		ints := ExtractInts(line)
		if lineIndex > 0 {
			previousInts := ExtractInts(s.Input[lineIndex-1])
			if len(ints) == 3 {
				if len(previousInts) != 3 {
					mapIndex++
				}
				var mapRange [3]int
				mapRange[0] = ints[0]
				mapRange[1] = ints[1]
				mapRange[2] = ints[2]
				maps[mapIndex] = append(maps[mapIndex], mapRange)
			}
		}
	}
	if s.Debug {
		fmt.Println(maps)
	}

	// iterate through the seeds
	seeds := ExtractInts(s.Input[0])
	var results []int
	for _, seed := range seeds {
		if s.Debug {
			fmt.Println(seed)
		}
		for _, step := range maps {
			newSeed := seed
			for _, m := range step {
				if seed >= m[1] && seed < m[1]+m[2] {
					newSeed = m[0] + seed - m[1]
				}
			}
			seed = newSeed
			if s.Debug {
				fmt.Println(newSeed)
			}
		}
		results = append(results, seed)
	}
	min := results[0]
	for _, r := range results {
		if r < min {
			min = r
		}
	}
	return min
}

func (s *s2305) processPart2() int {
	var maps [7][][3]int
	// build the maps
	mapIndex := -1
	for lineIndex, line := range s.Input {
		ints := ExtractInts(line)
		if lineIndex > 0 {
			previousInts := ExtractInts(s.Input[lineIndex-1])
			if len(ints) == 3 {
				if len(previousInts) != 3 {
					mapIndex++
				}
				var mapRange [3]int
				mapRange[0] = ints[0]
				mapRange[1] = ints[1]
				mapRange[2] = ints[2]
				maps[mapIndex] = append(maps[mapIndex], mapRange)
			}
		}
	}
	if s.Debug {
		fmt.Println(maps)
	}

	// iterate through the seeds
	seeds := ExtractInts(s.Input[0])
	if s.Debug {
		fmt.Println(seeds)
	}
	min := -1
	for seedIndex, seedRange := range seeds {
		if seedIndex%2 == 0 {
			if s.Debug {
				fmt.Println(seedIndex)
			}
			for xseed := seedRange; xseed < seedRange+seeds[seedIndex+1]; xseed++ {
				seed := xseed
				for _, step := range maps {
					newSeed := seed
					for _, m := range step {
						if seed >= m[1] && seed < m[1]+m[2] {
							newSeed = m[0] + seed - m[1]
						}
					}
					seed = newSeed
				}
				if min == -1 || seed < min {
					min = seed
				}
			}
		}
	}
	return min
}
