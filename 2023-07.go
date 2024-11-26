package main

import (
	"strconv"
	"strings"
)

type s2307 Puzzle

func (s *s2307) SetDebug(debug bool) error {
	s.Debug = debug
	return nil
}

func (s *s2307) SetInput(input []string) error {
	s.Input = input
	return nil
}

func (s *s2307) SetPart(part int) error {
	s.Part = part
	return nil
}

func (s *s2307) Solve() (string, error) {
	return strconv.Itoa(s.Process()), nil
}

func (s *s2307) Process() int {
	score := 0
	for i, hand1 := range s.Input {
		rank := 1
		var inputs [2]string
		words1 := strings.Fields(hand1)
		inputs[0] = words1[0]
		for j, hand2 := range s.Input {
			if i != j {
				words2 := strings.Fields(hand2)
				inputs[1] = words2[0]
				if s.Part == 2 {
					if s.camelCard2(inputs) {
						rank++
					}
				} else {
					if s.camelCard(inputs) {
						rank++
					}
				}
			}
		}
		bid, _ := strconv.Atoi(words1[1])
		score += rank * bid
	}
	return score
}

func (s *s2307) camelCard(inputs [2]string) bool {
	var hands [2][5]int
	var ranks [2]int
	for i, hand := range inputs {
		for j, card := range hand {
			var num int
			switch card {
			case '2':
				num = 2
			case '3':
				num = 3
			case '4':
				num = 4
			case '5':
				num = 5
			case '6':
				num = 6
			case '7':
				num = 7
			case '8':
				num = 8
			case '9':
				num = 9
			case 'T':
				num = 10
			case 'J':
				num = 11
			case 'Q':
				num = 12
			case 'K':
				num = 13
			case 'A':
				num = 15
			}
			hands[i][j] = num
		}
	}
	// rank each hand
	for i, hand := range hands {
		var groups [16]int
		for _, card := range hand {
			groups[card]++
		}
		pairCount := 0
		for _, group := range groups {
			if group == 5 {
				ranks[i] = 6
			} else if ranks[i] < 5 && group == 4 {
				ranks[i] = 5
			} else if ranks[i] < 3 && group == 3 {
				ranks[i] = 3
			} else if ranks[i] < 1 && group == 2 {
				ranks[i] = 1
			}
			if group == 2 {
				pairCount++
			}
		}
		// check for full house and two pair
		if ranks[i] == 3 && pairCount == 1 {
			ranks[i] = 4
		} else if pairCount == 2 {
			ranks[i] = 2
		}
		//fmt.Println(hand, ranks[i], pairCount)
	}
	if ranks[0] > ranks[1] {
		return true
	}
	if ranks[0] < ranks[1] {
		return false
	}
	// card by card
	for i, card := range hands[0] {
		if card > hands[1][i] {
			return true
		}
		if card < hands[1][i] {
			return false
		}
	}
	return true
}

// part 2
func (s *s2307) camelCard2(inputs [2]string) bool {
	var hands [2][5]int
	var ranks [2]int
	for i, hand := range inputs {
		for j, card := range hand {
			var num int
			switch card {
			case '2':
				num = 2
			case '3':
				num = 3
			case '4':
				num = 4
			case '5':
				num = 5
			case '6':
				num = 6
			case '7':
				num = 7
			case '8':
				num = 8
			case '9':
				num = 9
			case 'T':
				num = 10
			case 'J':
				num = 1
			case 'Q':
				num = 12
			case 'K':
				num = 13
			case 'A':
				num = 15
			}
			hands[i][j] = num
		}
	}
	// rank each hand
	for i, hand := range hands {
		var groups [16]int
		jokers := 0
		for _, card := range hand {
			if card == 1 {
				jokers++
			} else {
				groups[card]++
			}
		}
		pairCount := 0
		for _, group := range groups {
			if group+jokers == 5 {
				ranks[i] = 6
			} else if ranks[i] < 5 && group+jokers == 4 {
				ranks[i] = 5
			} else if ranks[i] < 3 && group+jokers == 3 {
				ranks[i] = 3
			} else if ranks[i] < 1 && group+jokers == 2 {
				ranks[i] = 1
			}
			if group == 2 {
				pairCount++
			}
		}
		// check for full house and two pair
		switch jokers {
		case 0:
			if ranks[i] == 3 && pairCount == 1 {
				ranks[i] = 4
			} else if pairCount == 2 {
				ranks[i] = 2
			}
		case 1:
			if pairCount == 2 {
				ranks[i] = 4
			} else if ranks[i] < 1 {
				ranks[i] = 1
			}
		}
		//fmt.Println(hand, ranks[i], pairCount)
	}
	if ranks[0] > ranks[1] {
		return true
	}
	if ranks[0] < ranks[1] {
		return false
	}
	// card by card
	for i, card := range hands[0] {
		if card > hands[1][i] {
			return true
		}
		if card < hands[1][i] {
			return false
		}
	}
	return true
}
