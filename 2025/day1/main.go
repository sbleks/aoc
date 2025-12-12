package main

import (
	input "aocInput"
	"log"
	"strconv"
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

		startNum := num
		zerosInRotation := 0

		// Count zeros during the rotation by checking each position
		// Note: we start from i=1 because i=0 is the starting position,
		// which was already counted as the end of the previous rotation (if it was 0)
		switch dir {
		case 'R':
			// For R: positions are (startNum + i) % 100 for i = 1 to dist
			// Count how many times this equals 0
			for i := 1; i <= dist; i++ {
				pos := (startNum + i) % 100
				if pos == 0 {
					zerosInRotation++
				}
			}
			num = (startNum + dist) % 100
		case 'L':
			// For L: positions are (startNum - i + 100) % 100 for i = 1 to dist
			// Count how many times this equals 0
			for i := 1; i <= dist; i++ {
				pos := (startNum - i + 100) % 100
				if pos == 0 {
					zerosInRotation++
				}
			}
			// Ensure positive modulo result
			num = ((startNum-dist)%100 + 100) % 100
		}

		sum += zerosInRotation

		if num < 0 || num > 99 {
			log.Fatalf("num: %d is out of bounds", num)
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
