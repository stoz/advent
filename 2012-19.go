package main

import (
	"fmt"
)

type RobotState [9]uint8 // 0-3 robots, 4-7 resources, 8 buy mask

func s2219(filename string, part2, debug bool) int {
	lines := ReadFile("./data/2022/19/" + filename)
	var blueprints [][6]uint8
	var sum int
	minutes := 23 // 23 because we don't buy robots on final state 24
	for i, line := range lines {
		ints := ExtractUint8s(line)
		//                       ore      ore       ore + clay     ore + obsidian
		appender := [6]uint8{ints[1], ints[2], ints[3], ints[4], ints[5], ints[6]}
		blueprints = append(blueprints, appender)
		// part 2 only uses first 3 blueprints
		if part2 && i == 2 {
			break
		}
	}
	// part 2 does multiplication instead of addition, and 32 minutes instead of 24
	if part2 {
		minutes = 31
		sum = 1
	}
	for blueindex, blue := range blueprints {
		max := doBlueprint(blue, minutes, debug)
		if part2 {
			sum *= max
		} else {
			sum += (blueindex + 1) * max
		}
		if debug {
			fmt.Println("Blueprint", blueindex+1, max)
		}
	}
	return sum
}

func doBlueprint(blue [6]uint8, steps int, debug bool) int {
	var max uint8
	states := []RobotState{{1}}
	for i := 0; i < steps; i++ {
		if debug {
			fmt.Println("State depth:", i, "States:", len(states))
		}
		var newstates []RobotState
		var newstates2 []RobotState
		var geodeBots, remain uint8
		var best int
		remain = uint8(steps - i)
		for _, s := range states {
			result := simRobot(blue, s, remain)
			for _, r := range result {
				newstates = append(newstates, r)
			}
		}
		// Throw away states with less than the max geode bots
		for _, s := range newstates {
			if s[3] > geodeBots {
				geodeBots = s[3]
			}
			futureGeodes := int(s[3])*int(remain) + int(s[7])
			if futureGeodes > best {
				best = futureGeodes
			}
		}
		if geodeBots > 1 {
			for _, s := range newstates {
				futureGeodes := int(s[3])*int(remain) + int(s[7]) + (int(remain) * (int(remain) + 1) / 2)
				if s[3] >= geodeBots-2 && futureGeodes >= best {
					newstates2 = append(newstates2, s)
				}
			}
			states = newstates2
		} else {
			states = newstates
		}
	}
	for _, s := range states {
		if s[7]+s[3] > max {
			max = s[7] + s[3]
		}
	}
	return int(max)
}

func simRobot(blue [6]uint8, state RobotState, remain uint8) []RobotState {
	var results []RobotState
	// collect resources
	rob := [4]uint8{state[0], state[1], state[2], state[3]}
	res := [4]uint8{state[4], state[5], state[6], state[7]}
	// collect resources
	for r, c := range rob {
		res[r] += c
	}
	ore := rob[0]
	clay := rob[1]
	obsidian := rob[2]
	// always buy robot 3
	if blue[4] <= state[4] && blue[5] <= state[6] {
		app := RobotState{rob[0], rob[1], rob[2], rob[3] + 1, res[0] - blue[4], res[1], res[2] - blue[5], res[3]}
		results = append(results, app)
		return results
	} else {
		if state[8] < 4 && blue[2] <= state[4] && blue[3] <= state[5] && obsidian < blue[5] {
			app := RobotState{rob[0], rob[1], rob[2] + 1, rob[3], res[0] - blue[2], res[1] - blue[3], res[2], res[3]}
			results = append(results, app)
		}
		if state[8] != 2 && state[8] != 3 && state[8] != 6 && blue[1] <= state[4] && clay < blue[3] {
			app := RobotState{rob[0], rob[1] + 1, rob[2], rob[3], res[0] - blue[1], res[1], res[2], res[3]}
			results = append(results, app)
		}
		if state[8] != 1 && state[8] != 3 && state[8] != 6 && blue[0] <= state[4] && (ore < blue[0] || ore < blue[1] || ore < blue[2] || ore < blue[4]) {
			app := RobotState{rob[0] + 1, rob[1], rob[2], rob[3], res[0] - blue[0], res[1], res[2], res[3]}
			results = append(results, app)
		}
		// no point to delay robot purchasing. If we can afford something but skip it, we can't buy it next
		// 1 = ore, 2 = clay, 3 = ore + clay, 4 = obsidian, 5 = obsidian + ore, 6 = obsidian + clay + ore
		var mask uint8
		if blue[0] <= state[4] {
			mask++
		}
		if blue[1] <= state[4] {
			mask += 2
		}
		if blue[2] <= state[4] && blue[3] <= state[5] {
			mask += 4
		}
		app := RobotState{rob[0], rob[1], rob[2], rob[3], res[0], res[1], res[2], res[3], mask}
		results = append(results, app)
	}
	return results
}
