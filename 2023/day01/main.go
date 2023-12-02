package main

import (
	_ "embed"
	"errors"
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
			if !isNumeric {
				continue
			}
			if first == nil {
				first = &chars[i]
			}
			last = &chars[i]
		}
		if first == nil || last == nil {
			fmt.Printf("No digits found on line <%s>\n", line)
			continue
		}
		num, err := strconv.Atoi(*first + *last)
		if err != nil {
			fmt.Println(err)
			return
		}
		total += num
	}
	fmt.Println(total)
}

func getDigits() []string {
	return []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
}

func findDigitAtPosition(str string, pos int) (*int, error) {
	digits := getDigits()
	char := str[pos]
	if char >= '0' && char <= '9' {
		num, err := strconv.Atoi(string(char))
		if err != nil {
			return nil, errors.New("Unexpected unparseable character")
		}
		return &num, nil
	}

	for i, digit := range digits {
		endPosition := pos + len(digit)
		if endPosition > len(str) {
			continue
		}
		subStr := string(str[pos:endPosition])
		if digit == subStr {
			return &i, nil
		}
	}

	return nil, errors.New("No digit found")
}

func part2() {
	total := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		var first, last *int
		length := len(line)
		for i := 0; i < length; i++ {
			num, err := findDigitAtPosition(line, i)
			if err != nil {
				continue
			}
			if first == nil {
				first = num
			}
			last = num
		}
		if first == nil || last == nil {
			fmt.Printf("No digits found on line <%s>\n", line)
			continue
		}
		total += (*first * 10) + *last
	}
	fmt.Println(total)
}
