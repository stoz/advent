package main

func s22031(filename string) int {
	// expected result: 7597
	lines := ReadFile("./data/2022/03/" + filename)
	count := 0
	var chunks []string
	for _, line := range lines {
		chunks = Chunks(line, len(line)/2)
		matched := false
		for _, a := range chunks[0] {
			for _, b := range chunks[1] {
				if a == b && !matched {
					matched = true
					// Go runes, A-Z = 65-90, a-z=97-122
					if a < 91 {
						// uppercase letter
						count += (int(a) - 38)
					} else {
						// lowercase letter
						count += (int(a) - 96)
					}
				}
			}
		}
	}
	return count
}

func s22032(filename string) int {
	// expected result: 2607
	lines := ReadFile("./data/2022/03/" + filename)
	count := 0
	i := 0
	var elves [3]string
	for _, line := range lines {
		elves[i] = line
		if i > 1 {
			matched := false
			for _, a := range elves[0] {
				for _, b := range elves[1] {
					for _, c := range elves[2] {
						if a == b && a == c && !matched {
							matched = true
							// Go runes, A-Z = 65-90, a-z=97-122
							if a < 91 {
								// uppercase letter
								count += (int(a) - 38)
							} else {
								// lowercase letter
								count += (int(a) - 96)
							}
						}
					}
				}
			}
			i = 0
		} else {
			i++
		}
	}
	return count
}
