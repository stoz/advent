package main

import (
	"fmt"
	"strconv"
)

type s2417 Puzzle

func (s *s2417) SetDebug(debug bool) error {
	s.Debug = debug
	return nil
}

func (s *s2417) SetInput(input []string) error {
	s.Input = input
	return nil
}

func (s *s2417) SetPart(part int) error {
	s.Part = part
	return nil
}

func (s *s2417) Solve() (string, error) {
	if s.Part != 2 {
		ints := s.processPart1()
		a := ""
		for i, v := range ints {
			if i > 0 {
				a += ","
			}
			a += strconv.Itoa(v)
		}
		return a, nil
	} else {
		return strconv.Itoa(s.processPart2()), nil
	}
}

func (s *s2417) processPart1() []int {
	ints := ExtractInts(s.Input[0])
	regA := ints[0]
	ints = ExtractInts(s.Input[1])
	regB := ints[0]
	ints = ExtractInts(s.Input[2])
	regC := ints[0]

	program := ExtractInts(s.Input[4])
	if s.Debug {
		fmt.Println(regA, regB, regC, program)
	}
	return s.compute(regA, regB, regC, program)
}

func (s *s2417) processPart2() int {
	ints := ExtractInts(s.Input[1])
	regB := ints[0]
	ints = ExtractInts(s.Input[2])
	regC := ints[0]

	program := ExtractInts(s.Input[4])

	i := 1
	digits := 1
	do := true
	for do {
		result := s.compute(i, regB, regC, program)
		if len(result) > 0 {
			match := true
			pCount := len(program)
			for x := len(result) - 1; x >= 0; x-- {
				pCount--
				if result[x] != program[pCount] {
					match = false
				}
			}
			if match {
				if s.Debug {
					fmt.Println("match", i, result)
				}
				if digits == len(program) {
					return i
				}
				digits++
				i *= 8
			} else {
				i++
			}
		}
	}
	return 0
}

func (s *s2417) compute(regA, regB, regC int, program []int) []int {
	ptr := 0
	var output []int
	for ptr < len(program) {
		jump := true
		operand := program[ptr+1]
		switch program[ptr] {
		case 0: // adv
			regA = FloorDivision(regA, IntPow(2, s.getOperand(regA, regB, regC, operand)))
		case 1: // bxl
			bits1 := s.to3bits(regB)
			bits2 := s.to3bits(operand)
			bits := [3]int{(bits1[0] + bits2[0]) % 2, (bits1[1] + bits2[1]) % 2, (bits1[2] + bits2[2]) % 2}
			regB = bits[0] + bits[1]*2 + bits[2]*4
		case 2: // bst
			regB = s.getOperand(regA, regB, regC, operand) % 8
		case 3: // jnz
			if regA != 0 {
				ptr = operand
				jump = false
			}
		case 4: // bxc
			bits1 := s.to3bits(regB)
			bits2 := s.to3bits(regC)
			bits := [3]int{(bits1[0] + bits2[0]) % 2, (bits1[1] + bits2[1]) % 2, (bits1[2] + bits2[2]) % 2}
			regB = bits[0] + bits[1]*2 + bits[2]*4
		case 5: // out
			out := s.getOperand(regA, regB, regC, operand) % 8
			output = append(output, out)
		case 6: // bdv
			regB = FloorDivision(regA, IntPow(2, s.getOperand(regA, regB, regC, operand)))
		case 7: // cdv
			regC = FloorDivision(regA, IntPow(2, s.getOperand(regA, regB, regC, operand)))
		}
		if jump {
			ptr += 2
		}
	}
	// this breaks part 1
	//output = append(output, regA)
	return output
}

func (s *s2417) getOperand(regA, regB, regC, operand int) int {
	switch operand {
	case 4:
		return regA
	case 5:
		return regB
	case 6:
		return regC
	}
	return operand
}

// convert int to 3 bits
func (s *s2417) to3bits(i int) []int {
	bits := make([]int, 3)
	for j := 0; j < 3; j++ {
		bits[j] = i % 2
		i /= 2
	}
	return bits
}
