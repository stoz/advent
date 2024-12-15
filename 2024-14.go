package main

import (
	"fmt"
	"strconv"
)

type robot struct {
	x  int
	y  int
	vx int
	vy int
}

type s2414 Puzzle

func (s *s2414) SetDebug(debug bool) error {
	s.Debug = debug
	return nil
}

func (s *s2414) SetInput(input []string) error {
	s.Input = input
	return nil
}

func (s *s2414) SetPart(part int) error {
	s.Part = part
	return nil
}

func (s *s2414) Solve() (string, error) {
	var number int
	if s.Part != 2 {
		number = s.processPart1()
	} else {
		number = s.processPart2()
	}
	return strconv.Itoa(number), nil
}

func (s *s2414) processPart1() int {
	sum := 0
	var robots []robot
	maxX := 11
	maxY := 7
	// real input has 500 lines and a different grid size
	if len(s.Input) > 100 {
		maxX = 101
		maxY = 103
	}
	for _, line := range s.Input {
		ints := ExtractSInts(line)
		var r robot
		r.x = ints[0]
		r.y = ints[1]
		r.vx = ints[2]
		r.vy = ints[3]
		robots = append(robots, r)
	}
	for i := 0; i < 100; i++ {
		if s.Debug && i < 6 {
			fmt.Println(i, robots[0])
			s.debugRobots(robots, maxX, maxY)
		}
		for ri, r := range robots {
			if r.vx > 0 {
				for x := 0; x < r.vx; x++ {
					if robots[ri].x+1 > maxX-1 {
						robots[ri].x = 0
					} else {
						robots[ri].x++
					}
				}
			} else if r.vx < 0 {
				for x := 0; x > r.vx; x-- {
					if robots[ri].x-1 < 0 {
						robots[ri].x = maxX - 1
					} else {
						robots[ri].x--
					}
				}
			}
			if r.vy > 0 {
				for y := 0; y < r.vy; y++ {
					if robots[ri].y+1 > maxY-1 {
						robots[ri].y = 0
					} else {
						robots[ri].y++
					}
				}
			} else if r.vy < 0 {
				for y := 0; y > r.vy; y-- {
					if robots[ri].y-1 < 0 {
						robots[ri].y = maxY - 1
					} else {
						robots[ri].y--
					}
				}
			}
		}
	}
	if s.Debug {
		s.debugRobots(robots, maxX, maxY)
	}
	// count robots in each quadrant
	var quadrants [4]int
	halfX := (maxX - 1) / 2
	halfY := (maxY - 1) / 2
	for _, r := range robots {
		if r.x < halfX && r.y < halfY {
			quadrants[0]++
		}
		if r.x > halfX && r.y < halfY {
			quadrants[1]++
		}
		if r.x < halfX && r.y > halfY {
			quadrants[2]++
		}
		if r.x > halfX && r.y > halfY {
			quadrants[3]++
		}
	}
	if s.Debug {
		fmt.Println(quadrants)
	}
	sum = quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
	return sum
}

func (s *s2414) processPart2() int {
	sum := 0
	var robots []robot
	maxX := 11
	maxY := 7
	// real input has 500 lines and a different grid size
	if len(s.Input) > 100 {
		maxX = 101
		maxY = 103
	}
	for _, line := range s.Input {
		ints := ExtractSInts(line)
		var r robot
		r.x = ints[0]
		r.y = ints[1]
		r.vx = ints[2]
		r.vy = ints[3]
		robots = append(robots, r)
	}
	// pattern should repeat after maxX * maxY
	for i := 0; i < maxX*maxY; i++ {
		if s.Debug && i < 6 {
			fmt.Println(i, robots[0])
			s.debugRobots(robots, maxX, maxY)
		}
		for ri, r := range robots {
			if r.vx > 0 {
				for x := 0; x < r.vx; x++ {
					if robots[ri].x+1 > maxX-1 {
						robots[ri].x = 0
					} else {
						robots[ri].x++
					}
				}
			} else if r.vx < 0 {
				for x := 0; x > r.vx; x-- {
					if robots[ri].x-1 < 0 {
						robots[ri].x = maxX - 1
					} else {
						robots[ri].x--
					}
				}
			}
			if r.vy > 0 {
				for y := 0; y < r.vy; y++ {
					if robots[ri].y+1 > maxY-1 {
						robots[ri].y = 0
					} else {
						robots[ri].y++
					}
				}
			} else if r.vy < 0 {
				for y := 0; y > r.vy; y-- {
					if robots[ri].y-1 < 0 {
						robots[ri].y = maxY - 1
					} else {
						robots[ri].y--
					}
				}
			}
		}
		// assume the christmas tree pattern has no overlapping robots
		duplicates := make(map[int]int)
		for _, r := range robots {
			duplicates[r.x*1000+r.y]++
		}
		noDuplicates := true
		for _, d := range duplicates {
			if d > 1 {
				noDuplicates = false
			}
		}
		if noDuplicates {
			sum = i + 1
			break
		}
	}
	if s.Debug {
		s.debugRobots(robots, maxX, maxY)
	}
	return sum
}

func (s *s2414) debugRobots(robots []robot, maxX, maxY int) {
	fmt.Println("")
	for y := 0; y < maxY; y++ {
		debugString := ""
		for x := 0; x < maxX; x++ {
			count := 0
			for _, r := range robots {
				if r.x == x && r.y == y {
					count++
				}
			}
			if count > 0 {
				debugString += strconv.Itoa(count)
			} else {
				debugString += "."
			}
		}
		fmt.Println(debugString)
	}
}
