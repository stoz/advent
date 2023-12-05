package main

import "fmt"

func s23051(filename string, part, debug bool) int {
	lines := ReadFile("./data/2023/05/" + filename)
	var maps [7][][3]int
	// build the maps
	mapIndex := -1
	for lineIndex, line := range lines {
		ints := ExtractInts(line)
		if lineIndex > 0 {
			previousInts := ExtractInts(lines[lineIndex-1])
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
	if debug {
		fmt.Println(maps)
	}

	// iterate through the seeds
	seeds := ExtractInts(lines[0])
	var results []int
	for _, seed := range seeds {
		if debug {
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
			if debug {
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

func s23052(filename string, part, debug bool) int {
	lines := ReadFile("./data/2023/05/" + filename)
	var maps [7][][3]int
	// build the maps
	mapIndex := -1
	for lineIndex, line := range lines {
		ints := ExtractInts(line)
		if lineIndex > 0 {
			previousInts := ExtractInts(lines[lineIndex-1])
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
	if debug {
		fmt.Println(maps)
	}

	// iterate through the seeds
	seeds := ExtractInts(lines[0])
	if debug {
		fmt.Println(seeds)
	}
	min := -1
	for seedIndex, seedRange := range seeds {
		if seedIndex%2 == 0 {
			if debug {
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
