package main

import (
	"encoding/json"
	"log"
	"sort"
)

func s22131(filename string) int {
	lines := ReadFile("./data/2022/13/" + filename)
	var lline, rline string
	var count int

	for i, line := range lines {
		switch i % 3 {
		case 0:
			lline = line
		case 1:
			rline = line
			// comparison
			var ldata, rdata any
			err := json.Unmarshal([]byte(lline), &ldata)
			if err != nil {
				log.Fatal("Cannot unmarshal the json ", err)
			}
			err = json.Unmarshal([]byte(rline), &rdata)
			if err != nil {
				log.Fatal("Cannot unmarshal the json ", err)
			}
			if CompareDistress(ldata, rdata) <= 0 {
				count += ((i - i%3) / 3) + 1
			}
		}
	}
	return count
}

func s22132(filename string) int {
	lines := ReadFile("./data/2022/13/" + filename)
	var linesInterface []any
	add := []string{"2", "6"}
	for _, a := range add {
		lines = append(lines, "[["+a+"]]")
	}
	for _, line := range lines {
		if line != "" {
			var data any
			err := json.Unmarshal([]byte(line), &data)
			if err != nil {
				log.Fatal("Cannot unmarshal the json ", err)
			}
			linesInterface = append(linesInterface, data)
		}
	}

	// sort using the CompareDistress function
	sort.Slice(linesInterface, func(i, j int) bool {
		return CompareDistress(linesInterface[i], linesInterface[j]) < 0
	})

	//fmt.Println(linesInterface)

	sum := 1
	for i, line := range linesInterface {
		data, _ := json.Marshal(line)
		for _, a := range add {
			if string(data) == "[["+a+"]]" {
				//fmt.Println(sum, i+1)
				sum *= i + 1
			}
		}
	}
	return sum
}

func CompareDistress(left, right any) int {
	// float64 type assertion
	l, lok := left.(float64)
	r, rok := right.(float64)
	// if both are float64, return the difference
	if lok && rok {
		return int(l) - int(r)
	}

	var llist, rlist []any
	switch left.(type) {
	case []any, []float64:
		llist = left.([]any)
	case float64:
		llist = []any{left}
	}
	switch right.(type) {
	case []any, []float64:
		rlist = right.([]any)
	case float64:
		rlist = []any{right}
	}

	for i := range llist {
		if len(rlist) <= i {
			return 1
		}
		if ret := CompareDistress(llist[i], rlist[i]); ret != 0 {
			return ret
		}
	}
	if len(llist) == len(rlist) {
		return 0
	}
	return -1
}
