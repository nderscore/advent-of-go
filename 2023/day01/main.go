package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "Part 1 or 2")
	flag.Parse()

	switch part {
	case 1:
		fmt.Println("Part 1")
		part1()
	case 2:
		fmt.Println("Part 2")
		part2()
	default:
		fmt.Println("Invalid part")
	}
}

func part1() {
	total := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		chars := strings.Split(line, "")
		var first, last *string
		for i, char := range chars {
			isNumeric := char >= "0" && char <= "9"
			if isNumeric {
				if first == nil {
					first = &chars[i]
				}
				last = &chars[i]
			}
		}
		num, err := strconv.Atoi(strings.Join([]string{*first, *last}, ""))
		if err != nil {
			fmt.Println(err)
			return
		}
		total += num
	}
	fmt.Println(total)
}

func getDigits() []string {
	return []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
}

func findDigitAtPosition(str string, pos int) int {
	digits := getDigits()
	char := str[pos]
	if char >= '0' && char <= '9' {
		num, err := strconv.Atoi(string(char))
		if err != nil {
			return -1
		}
		// fmt.Println("Digit", num, pos)
		return num
	}

	for i, digit := range digits {
		endPosition := pos + len(digit)
		if endPosition > len(str) {
			continue
		}
		subStr := string(str[pos:endPosition])
		if digit == subStr {
			// fmt.Println("Word", i+1, pos)
			return i + 1
		}
	}

	return -1
}

func part2() {
	total := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		// fmt.Println(line)
		first := -1
		last := -1
		length := len(line)
		for i := 0; i < length; i++ {
			num := findDigitAtPosition(line, i)
			if num != -1 {
				if first == -1 {
					first = num
				}
				last = num
			}
		}
		// fmt.Println(first, last)
		total += (first * 10) + last
	}
	fmt.Println(total)
}
