package main

import (
	input "aocInput"
	"log"
	"sort"
	"strconv"
	"strings"
	// "strconv"
	// "strings"
)

func part1(left, right []int) (sum int) {

	sort.Ints(left)
	sort.Ints(right)

	// Assert that we have parsed both sides correctly
	if len(left) != len(right) {
		log.Fatalln("Parsing error: left and right slices are not equal.")
	}

	for i := 0; i < len(left); i++ {
		diff := left[i] - right[i]
		if diff < 0 {
			diff = -diff
		}

		// log.Printf("left: %v, right: %v, diff: %v", left[i], right[i], diff)

		sum += diff
	}
	return sum
}

func part2(left, right []int) (sum int) {
	m := make(map[int]int)

	for _, val := range right {
		m[val] += 1
	}

	for _, leftVal := range left {
		sum += leftVal * m[leftVal]
	}

	return sum
}

func main() {
	lines, err := input.GetInputLines("./input.txt")
	if err != nil {
		log.Fatalf("Could not read file: %v", err)
	}

	var left, right []int
	for _, line := range lines {
		values := strings.Split(line, "   ")
		leftVal, err := strconv.Atoi(values[0])
		if err != nil {
			log.Panicln("ERROR: Could not convert leftVal to int", err)
		}
		rightVal, err := strconv.Atoi(values[1])
		if err != nil {
			log.Panicln("ERROR: Could not convert rightVal to int", err)
		}

		left = append(left, leftVal)
		right = append(right, rightVal)
	}

	// sum := part1(left,right)
	sum := part2(left, right)

	log.Printf("Sum is: %v", sum)
}
