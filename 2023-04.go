package main

func s2304(filename string, part, debug bool) int {
	lines := ReadFile("./data/2023/04/" + filename)
	div := 0
	total := 0
	var cards [204]int
	for c := range cards {
		cards[c] = 1
	}
	if filename == "sample.txt" {
		div = 6
	} else {
		div = 11
	}
	for lineIndex, line := range lines {
		score := 0
		a := ExtractInts(line)
		for i := 1; i < div; i++ {
			for j := div; j < len(a); j++ {
				if a[i] == a[j] {
					if part {
						score++
						if lineIndex+score <= len(cards) {
							cards[lineIndex+score] += cards[lineIndex]
						}
					} else if score == 0 {
						score = 1
					} else {
						score = score * 2
					}
				}
			}
		}
		if part {
			total += cards[lineIndex]
		} else {
			total += score
		}
	}
	return total
}
