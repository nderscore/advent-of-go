package main

import (
	"flag"

	"github.com/nderscore/advent-of-go/2023/day01"
)

func main() {
	var day, part int
	flag.IntVar(&day, "d", 1, "Day (1-25)")
	flag.IntVar(&part, "p", 1, "Part (1 or 2)")
	flag.Parse()

	switch day {
	case 1:
		day01.Run(part)
	default:
		println("Day not implemented")
	}
}
