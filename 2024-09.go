package main

import (
	"fmt"
	"strconv"
)

type s2409 Puzzle

func (s *s2409) SetDebug(debug bool) error {
	s.Debug = debug
	return nil
}

func (s *s2409) SetInput(input []string) error {
	s.Input = input
	return nil
}

func (s *s2409) SetPart(part int) error {
	s.Part = part
	return nil
}

func (s *s2409) Solve() (string, error) {
	var number int
	if s.Part != 2 {
		number = s.processPart1()
	} else {
		number = s.processPart2()
	}
	return strconv.Itoa(number), nil
}

func (s *s2409) processPart1() int {
	sum := 0
	disk := make(map[int]int)
	index := 0
	id := 0
	for i, d := range s.Input[0] {
		num, _ := strconv.Atoi(string(d))
		if i%2 == 0 {
			for j := 0; j < num; j++ {
				disk[index] = id
				index++
			}
			id++
		} else {
			for j := 0; j < num; j++ {
				disk[index] = -1
				index++
			}
		}
	}
	if s.Debug {
		fmt.Println(s.diskToString(disk))
	}

	// loop through the disk backwards and transpose any data to the empty space at the front of the disk
	for i := len(disk) - 1; i >= 0; i-- {
		if disk[i] != -1 {
			for j := 0; j < i; j++ {
				if disk[j] == -1 {
					disk[i], disk[j] = disk[j], disk[i]
					break
				}
			}
		}
	}

	if s.Debug {
		fmt.Println(s.diskToString(disk))
	}

	// calculate checksum
	for i := 0; i < len(disk); i++ {
		if disk[i] != -1 {
			sum += disk[i] * i
		}
	}
	return sum
}

func (s *s2409) processPart2() int {
	sum := 0
	disk := make(map[int]int)
	index := 0
	id := 0
	for i, d := range s.Input[0] {
		num, _ := strconv.Atoi(string(d))
		if i%2 == 0 {
			for j := 0; j < num; j++ {
				disk[index] = id
				index++
			}
			id++
		} else {
			for j := 0; j < num; j++ {
				disk[index] = -1
				index++
			}
		}
	}
	if s.Debug {
		fmt.Println(s.diskToString(disk))
	}

	// loop through the disk backwards and transpose any data to the empty space at the front of the disk
	fileBuffer := -1
	fileLength := 0
	for i := len(disk) - 1; i >= 0; i-- {
		if disk[i] != -1 && (fileBuffer == -1 || fileBuffer == disk[i]) {
			fileBuffer = disk[i]
			fileLength++
		} else {
			// attempt the swap
			emptyLength := 0
			for j := 0; j <= i; j++ {
				if disk[j] == -1 {
					emptyLength++
					if emptyLength == fileLength {
						for k := 0; k < fileLength; k++ {
							disk[i+1+k], disk[j-k] = disk[j-k], disk[i+1+k]
						}
						break
					}
				} else {
					emptyLength = 0
				}
			}

			if disk[i] == -1 {
				fileBuffer = -1
				fileLength = 0
			} else {
				fileBuffer = disk[i]
				fileLength = 1
			}

			if s.Debug {
				fmt.Println(s.diskToString(disk))
			}
		}
	}

	if s.Debug {
		fmt.Println(s.diskToString(disk))
	}

	// calculate checksum
	for i := 0; i < len(disk); i++ {
		if disk[i] != -1 {
			sum += disk[i] * i
		}
	}
	return sum
}

func (s *s2409) diskToString(disk map[int]int) string {
	a := ""
	for i := 0; i < len(disk); i++ {
		if disk[i] == -1 {
			a += "."
		} else {
			a += strconv.Itoa(disk[i])
		}
	}
	return a
}
