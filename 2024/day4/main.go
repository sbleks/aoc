package main

import (
	"aocInput"
	"log"
	// "strconv"
	// "strings"
)

func part1(lines []string) (sum int) {
	return sum
}


func part2(lines []string) (sum int) {
	return sum
}

func main() {
	lines, err := input.GetInputLines("./input.txt")
	if err != nil {
		log.Fatalf("Could not read file: %v", err)
	}

	sum1 := part1(lines)
	sum2 := part2(lines)

	log.Printf("Part 1  is: %v", sum1)
	log.Printf("Part 2  is: %v", sum2)
}
