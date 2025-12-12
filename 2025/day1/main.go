package main

import (
	input "aocInput"
	"log"
	"math"
	"strconv"
	// "strings"
)

func part1(lines []string) (sum int) {
	num := 50
	for _, line := range lines {
		dir := line[0]
		dist, err := strconv.Atoi(line[1:])
		dist = dist % 100
		if err != nil {
			log.Fatalf("Could not parse distance: %v", err)
		}
		log.Printf("dir: %c, dist: %d, num: %d", dir, dist, num)
		switch dir {
		case 'L':
			// log.Printf("L: num-dist: %d", num-dist)
			if num-dist < 0 {
				num = 100 - (dist - num)
			} else {
				num -= dist
			}
		case 'R':
			// log.Printf("R: num+dist: %d", num+dist)
			if num+dist > 99 {
				num = (num + dist) - 100
			} else {
				num += dist
			}
		}

		if num < 0 || num > 99 {
			log.Fatalf("num: %d is out of bounds", num)
		}

		if num == 0 {
			log.Printf("num: %d", num)
			sum += 1
			log.Printf("sum: %d", sum)
		}

	}
	return sum
}

func part2(lines []string) (sum int) {
	num := 50

	for _, line := range lines {
		dir := line[0]
		dist, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatalf("Could not parse distance: %v", err)
		}

		rotations := int(math.Floor(float64(dist / 100)))
		distance := dist % 100

		log.Printf("dir: %c, dist: %d, num: %d, distance: %d, rotations: %d", dir, dist, num, distance, rotations)

		switch dir {
		case 'L':
			// log.Printf("L: num-distance: %d", num-distance)
			if num-distance < 0 {
				num = 100 - (distance - num)
				sum += 1
				log.Printf("Passing 0, sum: %d, num: %d, distance: %d", sum, num, distance)
			} else {
				num -= distance
			}
		case 'R':
			// log.Printf("R: num+distance: %d", num+distance)
			if num+distance > 99 {
				num = (num + distance) - 100
				sum += 1
				log.Printf("Passing 0, sum: %d, num: %d, distance: %d", sum, num, distance)
			} else {
				num += distance
			}
		}

		if num < 0 || num > 99 {
			log.Fatalf("num: %d is out of bounds", num)
		}

		if rotations > 0 {
			sum += rotations
			log.Printf("Adding %d rotations", rotations)
		}

	}
	return sum
}

func main() {
	lines, err := input.GetInputLines("./input.txt")
	if err != nil {
		log.Fatalf("Could not read file: %v", err)
	}

	// sum1 := part1(lines)
	sum2 := part2(lines)

	// log.Printf("Part 1  is: %v", sum1)
	log.Printf("Part 2  is: %v", sum2)
}
