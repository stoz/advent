package main

import (
	"fmt"
	"strings"
	"sync"
)

type Valve struct {
	p int
	n []string
}

func s22161(filename string, debug bool) int {
	// expected result: 2077
	lines := ReadFile("./data/2022/16/" + filename)
	valves := make(map[string]Valve)
	grid := make(map[string]StrPoint)
	var visit []string
	for _, line := range lines {
		var valve Valve
		var point StrPoint
		words := strings.Fields(line)
		ints := ExtractInts(line)
		v := words[1]
		valve.p = ints[0]
		point.c = 999999
		for i, w := range words {
			if i > 8 {
				valve.n = append(valve.n, w[:2])
				point.n = append(point.n, w[:2])
			}
		}
		if valve.p > 0 {
			visit = append(visit, v)
		}
		valves[v] = valve
		grid[v] = point
	}
	visit2 := visit
	visit2 = append(visit2, "AA")
	quickest := make(map[string]map[string]int)
	for _, v := range visit2 {
		quickest[v] = make(map[string]int)
		for _, t := range visit {
			if v != t {
				gridCopy := make(map[string]StrPoint)
				for a, b := range grid {
					gridCopy[a] = b
				}
				start := gridCopy[v]
				start.c = 0
				gridCopy[v] = start
				quickest[v][t] = DijkstraPaths(gridCopy, t)
			}
		}
	}
	if debug {
		for s, q := range quickest {
			fmt.Println(s, q)
		}
		fmt.Println(visit)
	}

	// permutations
	perms := Per("AA", visit, quickest, false)

	// Threads
	if debug {
		fmt.Println("Threads:", len(perms))
	}
	max := 0
	var wg sync.WaitGroup
	wg.Add(len(perms))
	var results []int
	for i := 0; i < len(perms); i++ {
		go func(i int) {
			defer wg.Done()
			test := simValves("AA", 30, perms[i], quickest, valves)
			results = append(results, test)
		}(i)
	}
	wg.Wait()
	for _, i := range results {
		if i > max {
			max = i
		}
	}
	if debug {
		fmt.Println(len(results))
	}
	return max
}

func s22162(filename string, debug bool) int {
	// expected result: 2741
	lines := ReadFile("./data/2022/16/" + filename)
	valves := make(map[string]Valve)
	grid := make(map[string]StrPoint)
	var visit []string
	for _, line := range lines {
		var valve Valve
		var point StrPoint
		words := strings.Fields(line)
		ints := ExtractInts(line)
		v := words[1]
		valve.p = ints[0]
		point.c = 999999
		for i, w := range words {
			if i > 8 {
				valve.n = append(valve.n, w[:2])
				point.n = append(point.n, w[:2])
			}
		}
		if valve.p > 0 {
			visit = append(visit, v)
		}
		valves[v] = valve
		grid[v] = point
	}
	visit2 := visit
	visit2 = append(visit2, "AA")
	quickest := make(map[string]map[string]int)
	for _, v := range visit2 {
		quickest[v] = make(map[string]int)
		for _, t := range visit {
			if v != t {
				gridCopy := make(map[string]StrPoint)
				for a, b := range grid {
					gridCopy[a] = b
				}
				start := gridCopy[v]
				start.c = 0
				gridCopy[v] = start
				quickest[v][t] = DijkstraPaths(gridCopy, t)
			}
		}
	}

	if debug {
		for s, q := range quickest {
			fmt.Println(s, q)
		}
		fmt.Println(visit)
	}

	// "compress" to ints
	lookup := make(map[string]int)
	lookup["AA"] = 0
	var visitInt []int
	for compressi, compressv := range visit {
		lookup[compressv] = compressi + 1
		visitInt = append(visitInt, compressi+1)
	}
	quickestInt := make(map[int]map[int]int)
	for quickestA, quickestB := range quickest {
		quickestAInt := lookup[quickestA]
		quickestInt[quickestAInt] = make(map[int]int)
		for quickestC, quickestD := range quickestB {
			quickestCInt := lookup[quickestC]
			quickestInt[quickestAInt][quickestCInt] = quickestD
		}
	}
	valveInt := make(map[int]Valve)
	for valvea, valveb := range valves {
		valveInt[lookup[valvea]] = valveb
	}

	if debug {
		fmt.Println(visitInt)
	}

	// permutations
	perms := PerInt(0, visitInt, quickestInt, true)

	// Threads
	if debug {
		fmt.Println("Threads:", len(perms))
	}
	max := 0
	var wg sync.WaitGroup
	wg.Add(len(perms))
	var results []int
	for i := 0; i < len(perms); i++ {
		go func(i int) {
			defer wg.Done()
			test := simValvesInt(0, 26, perms[i], quickestInt, valveInt)
			var test2 int
			var visit3 []int
			for _, v := range visitInt {
				inc := true
				for _, p := range perms[i] {
					if v == p {
						inc = false
					}
				}
				if inc {
					visit3 = append(visit3, v)
				}
			}
			perms2 := PerInt(0, visit3, quickestInt, false)
			for _, xx := range perms2 {
				xxresult := simValvesInt(0, 26, xx, quickestInt, valveInt)
				if xxresult > test2 {
					test2 = xxresult
				}
			}
			test += test2
			results = append(results, test)
		}(i)
	}
	wg.Wait()
	for _, i := range results {
		if i > max {
			max = i
		}
	}
	if debug {
		fmt.Println(len(results))
	}
	return max
}

func PerInt(start int, visit []int, quickest map[int]map[int]int, split bool) [][]int {
	var r [][]int
	len := len(visit)
	if split {
		len = len / 2
	}
	for _, a := range visit {
		alim := quickest[start][a] + 1
		for _, b := range visit {
			if a != b {
				blim := alim + quickest[a][b] + 1
				for _, c := range visit {
					if a != c && b != c {
						clim := blim + quickest[b][c] + 1
						if clim > 29 || len == 3 {
							r = append(r, []int{a, b, c})
						} else {
							for _, d := range visit {
								if a != d && b != d && c != d {
									dlim := clim + quickest[c][d] + 1
									if dlim > 29 || len == 4 {
										r = append(r, []int{a, b, c, d})
									} else {
										for _, e := range visit {
											if a != e && b != e && c != e && d != e {
												elim := dlim + quickest[d][e] + 1
												if elim > 29 || len == 5 {
													r = append(r, []int{a, b, c, d, e})
												} else {
													for _, f := range visit {
														if a != f && b != f && c != f && d != f && e != f {
															flim := elim + quickest[e][f] + 1
															if flim > 29 || len == 6 {
																r = append(r, []int{a, b, c, d, e, f})
															} else {
																for _, g := range visit {
																	if a != g && b != g && c != g && d != g && e != g && f != g {
																		glim := flim + quickest[f][g] + 1
																		if glim > 29 || len == 7 {
																			r = append(r, []int{a, b, c, d, e, f, g})
																		} else {
																			for _, h := range visit {
																				if a != h && b != h && c != h && d != h && e != h && f != h && g != h {
																					r = append(r, []int{a, b, c, d, e, f, g, h})
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return r
}

func Per(start string, visit []string, quickest map[string]map[string]int, split bool) [][]string {
	var r [][]string
	len := len(visit)
	if split {
		len = len / 2
	}
	for _, a := range visit {
		alim := quickest[start][a] + 1
		for _, b := range visit {
			if a != b {
				blim := alim + quickest[a][b] + 1
				for _, c := range visit {
					if a != c && b != c {
						clim := blim + quickest[b][c] + 1
						if clim > 29 || len == 3 {
							r = append(r, []string{a, b, c})
						} else {
							for _, d := range visit {
								if a != d && b != d && c != d {
									dlim := clim + quickest[c][d] + 1
									if dlim > 29 || len == 4 {
										r = append(r, []string{a, b, c, d})
									} else {
										for _, e := range visit {
											if a != e && b != e && c != e && d != e {
												elim := dlim + quickest[d][e] + 1
												if elim > 29 || len == 5 {
													r = append(r, []string{a, b, c, d, e})
												} else {
													for _, f := range visit {
														if a != f && b != f && c != f && d != f && e != f {
															flim := elim + quickest[e][f] + 1
															if flim > 29 || len == 6 {
																r = append(r, []string{a, b, c, d, e, f})
															} else {
																for _, g := range visit {
																	if a != g && b != g && c != g && d != g && e != g && f != g {
																		glim := flim + quickest[f][g] + 1
																		if glim > 29 || len == 7 {
																			r = append(r, []string{a, b, c, d, e, f, g})
																		} else {
																			for _, h := range visit {
																				if a != h && b != h && c != h && d != h && e != h && f != h && g != h {
																					r = append(r, []string{a, b, c, d, e, f, g, h})
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return r
}

func simValvesInt(start int, cycles int, visit []int, quickest map[int]map[int]int, valves map[int]Valve) int {
	//fmt.Println(visit)
	var pressure, total, wait int
	var flip bool
	for i := 0; i < cycles; i++ {
		total += pressure
		if wait > 0 {
			wait--
		} else if flip {
			pressure += valves[start].p
			flip = false
		} else if len(visit) > 0 {
			var v int
			v, visit = visit[0], visit[1:]
			wait = quickest[start][v] - 1
			start = v
			flip = true
		}
	}
	//fmt.Println(total)
	return total
}

func simValves(start string, cycles int, visit []string, quickest map[string]map[string]int, valves map[string]Valve) int {
	//fmt.Println(visit)
	var pressure, total, wait int
	var flip bool
	for i := 0; i < cycles; i++ {
		total += pressure
		if wait > 0 {
			wait--
		} else if flip {
			pressure += valves[start].p
			flip = false
		} else if len(visit) > 0 {
			var v string
			v, visit = visit[0], visit[1:]
			wait = quickest[start][v] - 1
			start = v
			flip = true
		}
	}
	//fmt.Println(total)
	return total
}
