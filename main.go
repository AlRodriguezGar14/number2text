package main

import (
	"fmt"
	"strconv"
)

type TriadInfo struct {
	TriadBlock string
	MinValue   int
	MidValue   int
	MaxValue   int
}
type Triad [3]int

func triadInfoConstructor(triadBlock string, minValue int, midValue int, maxValue int) *TriadInfo {
	return &TriadInfo{TriadBlock: triadBlock, MinValue: minValue, MidValue: midValue, MaxValue: maxValue}
}

func triadConstructor(cent int, dec int, num int) *Triad {
	return &Triad{cent, dec, num}
}

func createTriads(i string, container *[]TriadInfo) {
	idx := len(i) - 1
	triadBlock := "1"
	for idx >= 0 {
		num, _ := strconv.Atoi(string(i[idx]))
		dec, _ := strconv.Atoi(string(i[idx-1]))
		cent := 0
		if idx >= 2 {
			cent, _ = strconv.Atoi(string(i[idx-2]))
		}

		minValue := num
		midValue := dec * 10
		maxValue := cent * 100

		*container = append(*container, *triadInfoConstructor(triadBlock, minValue, midValue, maxValue))

		triadBlock += "000"
		idx -= 3
	}
}

func main() {
	var input string
	var triadContainer []TriadInfo
	fmt.Println("Write a number:")
	fmt.Scanln(&input)

	if input == "" {
		fmt.Println("Incorrect input")
		return
	}
	createTriads(input, &triadContainer)
	fmt.Println(triadContainer)
}
