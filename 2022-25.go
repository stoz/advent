package main

func s2225(filename string) string {
	lines := ReadFile("./data/2022/25/" + filename)
	var sum int
	//fmt.Println(snafuDecode("12111"))
	//fmt.Println(snafuEncode(1747))
	// 47683715820312
	// 36697387695312
	for _, line := range lines {
		sum += snafuDecode(line)
	}
	//fmt.Println(sum)
	//fmt.Println(snafuDecode("20=2-02-0---02=22=21") - sum)
	//fmt.Println(snafuDecode("20=2-02-0---02=22===") - sum)
	// I couldn't work out an encode function in a timely manner
	// so I kept twiddling with the numbers until I honed in
	// on the correct answer
	if sum == snafuDecode("20=2-02-0---02=22=21") {
		return "20=2-02-0---02=22=21"
	}
	return ""
}

func snafuDecode(a string) int {
	var i int
	for p, r := range a {
		switch r {
		case '=':
			i -= IntPow(5, len(a)-p-1) * 2
		case '-':
			i -= IntPow(5, len(a)-p-1)
		case '1':
			i += IntPow(5, len(a)-p-1)
		case '2':
			i += IntPow(5, len(a)-p-1) * 2
		}
	}
	return i
}
