package main

import (
	"fmt"
	"strconv"
)

type Triad struct {
	TriadBlock string
	MinValue   int
	MidValue   int
	MaxValue   int
}

func (t *Triad) WriteTriad(dic map[string]string) {

	// fmt.Println("cent:", t.MaxValue)
	if t.MaxValue > 0 {
		printMax(strconv.Itoa(t.MaxValue), dic)
	}

	if t.MidValue+t.MinValue <= 20 && t.MinValue+t.MinValue != 0 {
		printMidMin(strconv.Itoa(t.MidValue+t.MinValue), "", dic)

	} else if t.MinValue == 0 {
		printMidMin(strconv.Itoa(t.MidValue), "", dic)
	} else if t.MidValue == 0 {
		printMidMin(strconv.Itoa(t.MinValue), "", dic)
	} else {
		printMidMin(strconv.Itoa(t.MidValue), strconv.Itoa(t.MinValue), dic)
	}

	// fmt.Println("dec:", t.MidValue)
	// fmt.Println("num:", t.MinValue)
}

func printMax(i string, dic map[string]string) {
	fmt.Println(dic[i])
	fmt.Println(dic["100"])
}

func printMidMin(i string, opt string, dic map[string]string) {
	if opt == "" {
		fmt.Println(dic[i])
	} else {
		fmt.Println(dic[i])
		fmt.Println(dic[opt])
	}
}

func triadConstructor(triadBlock string, minValue int, midValue int, maxValue int) *Triad {
	return &Triad{TriadBlock: triadBlock, MinValue: minValue, MidValue: midValue, MaxValue: maxValue}
}

func createTriads(i string, container *[]Triad) {
	idx := len(i) - 1
	triadBlock := "1"
	for idx >= 0 {
		num, _ := strconv.Atoi(string(i[idx]))
		dec := 0
		if idx >= 1 {
			dec, _ = strconv.Atoi(string(i[idx-1]))
		}
		cent := 0
		if idx >= 2 {
			cent, _ = strconv.Atoi(string(i[idx-2]))
		}

		minValue := num
		midValue := dec * 10
		maxValue := cent

		*container = append(*container, *triadConstructor(triadBlock, minValue, midValue, maxValue))

		triadBlock += "000"
		idx -= 3
	}
}

func main() {
	var input string
	var triadContainer []Triad
	fmt.Println("Write a number:")
	fmt.Scanln(&input)

	if input == "" {
		fmt.Println("Incorrect input")
		return
	}
	createTriads(input, &triadContainer)

	for i := 0; i < len(triadContainer); i++ {
		triadContainer[i].WriteTriad(Dictionary)
	}
	fmt.Println(triadContainer)
}
