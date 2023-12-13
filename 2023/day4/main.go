package main

import (
	"aocInput"
	"log"
	// "strconv"
	// "strings"
)

func part1(lines []string) (sum int) {
	return 0
}


func part2(lines []string) (sum int) {
	return 0
}

func main() {
	lines, err := input.GetInputLines("./input.txt")
	if err != nil {
		log.Fatalf("Could not read file: %v", err)
	}

	sum := part1(lines)
	// sum := part2(lines)

	log.Printf("Sum is: %v", sum)
}
