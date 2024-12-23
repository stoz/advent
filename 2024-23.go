package main

import (
	"fmt"
	"slices"
	"strconv"
)

type s2423 Puzzle

func (s *s2423) SetDebug(debug bool) error {
	s.Debug = debug
	return nil
}

func (s *s2423) SetInput(input []string) error {
	s.Input = input
	return nil
}

func (s *s2423) SetPart(part int) error {
	s.Part = part
	return nil
}

func (s *s2423) Solve() (string, error) {
	var a string
	if s.Part != 2 {
		a = strconv.Itoa(s.processPart1())
	} else {
		a = s.processPart2()
	}
	return a, nil
}

func (s *s2423) processPart1() int {
	sum := 0
	pairs := [][2]string{}
	sets := [][3]string{}
	unique := map[string]struct{}{}
	for _, line := range s.Input {
		pair := [2]string{line[:2], line[3:]}
		pairs = append(pairs, pair)
		unique[pair[0]] = struct{}{}
		unique[pair[1]] = struct{}{}
	}
	for i, pair := range pairs {
		for u, _ := range unique {
			if u != pair[0] && u != pair[1] {
				zero := false
				one := false
				for j, check := range pairs {
					if i != j {
						if (pair[0] == check[0] && u == check[1]) || pair[0] == check[1] && u == check[0] {
							zero = true
						}
						if (pair[1] == check[0] && u == check[1]) || pair[1] == check[1] && u == check[0] {
							one = true
						}
					}
				}
				if zero && one {
					set := [3]string{pair[0], pair[1], u}
					sets = append(sets, set)
				}
			}
		}
	}
	for _, set := range sets {
		if set[0][0] == 't' || set[1][0] == 't' || set[2][0] == 't' {
			sum++
		}
	}
	sum = sum / 3
	return sum
}

func (s *s2423) processPart2() string {
	pairs := [][2]string{}
	sets := [][]string{}
	unique := map[string]map[string]struct{}{}
	for _, line := range s.Input {
		pair := [2]string{line[:2], line[3:]}
		pairs = append(pairs, pair)
		set := []string{pair[0], pair[1]}
		sets = append(sets, set)
		_, ok := unique[pair[0]]
		if !ok {
			unique[pair[0]] = make(map[string]struct{})
		}
		unique[pair[0]][pair[1]] = struct{}{}
		_, ok1 := unique[pair[1]]
		if !ok1 {
			unique[pair[1]] = make(map[string]struct{})
		}
		unique[pair[1]][pair[0]] = struct{}{}
	}
	if s.Debug {
		fmt.Println(unique)
	}
	do := true
	maxSetLength := 2
	for do {
		for u, links := range unique {
			for i, set := range sets {
				linkedToAllMembersOfSet := true
				for _, member := range set {
					_, ok := links[member]
					if !ok {
						linkedToAllMembersOfSet = false
					}
				}
				if linkedToAllMembersOfSet {
					sets[i] = append(sets[i], u)
				}
			}
		}
		oldMaxSetLength := maxSetLength
		for _, set := range sets {
			if len(set) > maxSetLength {
				maxSetLength = len(set)
			}
		}
		if oldMaxSetLength == maxSetLength {
			do = false
		}
	}
	for _, set := range sets {
		if len(set) == maxSetLength {
			slices.Sort(set)
			buf := ""
			for i, a := range set {
				if i > 0 {
					buf += ","
				}
				buf += a
			}
			return buf
		}
	}
	return ""
}
