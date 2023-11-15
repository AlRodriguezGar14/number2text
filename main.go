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

// func (t *Triad) WriteTriad(dic map[string]string) {
//
// 	if t.MaxValue > 0 {
// 		printMax(strconv.Itoa(t.MaxValue), dic)
// 	}
//
// 	if t.MidValue+t.MinValue <= 20 && t.MinValue+t.MinValue != 0 {
// 		printMidMin(strconv.Itoa(t.MidValue+t.MinValue), "", dic)
//
// 	} else if t.MinValue == 0 {
// 		printMidMin(strconv.Itoa(t.MidValue), "", dic)
// 	} else if t.MidValue == 0 {
// 		printMidMin(strconv.Itoa(t.MinValue), "", dic)
// 	} else {
// 		printMidMin(strconv.Itoa(t.MidValue), strconv.Itoa(t.MinValue), dic)
// 	}
//
// 	if t.TriadBlock != "1" {
// 		fmt.Printf("%s ", dic[t.TriadBlock])
// 	}
//
// }

func (t *Triad) WriteTriad(dic map[string]string) {

	if t.MaxValue > 0 {
		max := strconv.Itoa(t.MaxValue)
		fmt.Printf("%s %s ", dic[max], dic["100"])
	}

	if t.MidValue+t.MinValue <= 20 && t.MinValue+t.MinValue != 0 {
		fmt.Printf("%s ", dic[strconv.Itoa(t.MidValue+t.MinValue)])

	} else if t.MinValue == 0 {
		fmt.Printf("%s ", dic[strconv.Itoa(t.MidValue)])
	} else if t.MidValue == 0 {
		fmt.Printf("%s ", dic[strconv.Itoa(t.MinValue)])
	} else {
		if t.TriadBlock == "1" {
			fmt.Printf("and ")
		}
		fmt.Printf("%s-%s", dic[strconv.Itoa(t.MidValue)], dic[strconv.Itoa(t.MinValue)])
		if t.TriadBlock != "1" {
			fmt.Printf(" ")
		}
	}

	if t.TriadBlock != "1" {
		fmt.Printf("%s ", dic[t.TriadBlock])
	}

}

func printMax(i string, dic map[string]string) {
	fmt.Printf("%s %s ", dic[i], dic["100"])
}

func printMidMin(i string, opt string, dic map[string]string) {
	if opt == "" {
		fmt.Printf("%s ", dic[i])
	} else {
		fmt.Printf("%s", dic[i])
		fmt.Printf("-%s ", dic[opt])
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

func reverseTriadContainer(container *[]Triad) {
	l := len(*container)

	for i := 0; i < l/2; i++ {
		j := l - i - 1
		(*container)[i], (*container)[j] = (*container)[j], (*container)[i]
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

	reverseTriadContainer(&triadContainer)

	for i := 0; i < len(triadContainer); i++ {
		triadContainer[i].WriteTriad(Dictionary)
	}
}
