package main

import (
	input "aocInput"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func part1(lines []string) (sum int) {
	re := regexp.MustCompile(`mul\((\d\d?\d?),(\d\d?\d?)\)`)
	input := strings.Join(lines, "")
	muls := re.FindAllStringSubmatch(input, -1)
	for _, mul := range muls {
		a, err := strconv.Atoi(mul[1])
		if err != nil {
			log.Fatalf("Error parsing '%v' to int: %v", mul[1], err)
		}

		b, err := strconv.Atoi(mul[2])
		if err != nil {
			log.Fatalf("Error parsing '%v' to int: %v", mul[2], err)
		}

		sum += a * b
	}
	return sum
}

func part2(lines []string) (sum int) {
	re := regexp.MustCompile(`mul\((\d\d?\d?),(\d\d?\d?)\)|don't\(\)|do\(\)`)
	input := strings.Join(lines, "")
	muls := re.FindAllStringSubmatch(input, -1)
	adding := true
	for _, mul := range muls {
		if mul[0] == "do()" {
			adding = true
			continue
		} else if mul[0] == "don't()" {
			adding = false
			continue
		}

		if !adding {
			log.Println("skipping")
			continue
		}
		log.Print(mul)
		a, err := strconv.Atoi(mul[1])
		if err != nil {
			log.Fatalf("Error parsing '%v' to int: %v", mul[1], err)
		}

		b, err := strconv.Atoi(mul[2])
		if err != nil {
			log.Fatalf("Error parsing '%v' to int: %v", mul[2], err)
		}

		sum += a * b
	}
	return sum
}

func main() {
	lines, err := input.GetInputLines("./input.txt")
	if err != nil {
		log.Fatalf("Could not read file: %v", err)
	}

	sum1 := part1(lines)
	sum2 := part2(lines)

	log.Printf("Part 1 is: %v", sum1)
	log.Printf("Part 2 is: %v", sum2)
}
