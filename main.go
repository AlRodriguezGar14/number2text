package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type Triad struct {
	TriadBlock string
	MinValue   int
	MidValue   int
	MaxValue   int
}

func (t *Triad) WriteTriadEnglish(dic map[string]string) {
	var printableTriad string
	addSpace := false

	if addSpace {
		fmt.Printf(" ")
		addSpace = false
	}
	if t.MaxValue > 0 {
		max := strconv.Itoa(t.MaxValue)
		printableTriad += fmt.Sprintf("%s %s ", dic[max], dic["100"])
	}

	if t.MidValue+t.MinValue <= 20 && t.MinValue+t.MinValue != 0 {
		printableTriad += fmt.Sprintf("%s ", dic[strconv.Itoa(t.MidValue+t.MinValue)])

	} else if t.MinValue == 0 && t.MidValue != 0 {
		printableTriad += fmt.Sprintf("%s ", dic[strconv.Itoa(t.MidValue)])

	} else if t.MidValue == 0 && t.MinValue != 0 {
		printableTriad += fmt.Sprintf("%s ", dic[strconv.Itoa(t.MinValue)])

	} else if t.MidValue != 0 && t.MinValue != 0 {
		if t.TriadBlock == "1" {
			printableTriad += fmt.Sprintf("and ")
		}
		printableTriad += fmt.Sprintf("%s-%s ", dic[strconv.Itoa(t.MidValue)], dic[strconv.Itoa(t.MinValue)])

	} else if t.MinValue == 0 && t.MidValue == 0 && t.TriadBlock == "1" {
		printableTriad += fmt.Sprintf("%s ", dic[strconv.Itoa(t.MaxValue)])
	}

	if t.TriadBlock != "1" && t.MaxValue+t.MidValue+t.MinValue > 0 {
		printableTriad += fmt.Sprintf("%s ", dic[t.TriadBlock])
	}
	if len(printableTriad) > 0 {
		printableTriad = printableTriad[:len(printableTriad)-1]
		fmt.Printf(printableTriad)
		addSpace = true
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

func isValidInput(input string) bool {
	for _, char := range input {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func main() {
	var input string
	var triadContainer []Triad

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Write a number: ")

	if scanner.Scan() {
		input = scanner.Text()
	}

	if isValidInput(input) == false || input == "" {
		fmt.Println("Incorrect input")
		return
	}
	createTriads(input, &triadContainer)

	reverseTriadContainer(&triadContainer)

	for i := 0; i < len(triadContainer); i++ {
		triadContainer[i].WriteTriadEnglish(Dictionary)
	}
	// fmt.Println()
}
