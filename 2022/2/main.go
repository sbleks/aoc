package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	part1, _ := part1()
	part2, _ := part2()
	fmt.Println("Part 1", part1)
	fmt.Println("Part 2", part2)
}

func readInput(path string) (lines []string, err error) {
	file, err := os.Open(path)
	if err != nil {
		return lines, err
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines, nil
}

func part1() (int, error) {
	var sum int
	lines, err := readInput("./1.txt")
	if err != nil {
		return 0, err
	}
	for _, v := range lines {

		// fmt.Println(v[0:1], v[2:3])
		if v[0:1] == "B" && v[2:3] == "Z" {
			sum += 6 + 3
		}
		if v[0:1] == "B" && v[2:3] == "Y" {
			sum += 3 + 2
		}
		if v[0:1] == "B" && v[2:3] == "X" {
			sum += 0 + 1
		}
		if v[0:1] == "A" && v[2:3] == "Z" {
			sum += 0 + 3
		}
		if v[0:1] == "A" && v[2:3] == "Y" {
			sum += 6 + 2
		}
		if v[0:1] == "A" && v[2:3] == "X" {
			sum += 3 + 1
		}
		if v[0:1] == "C" && v[2:3] == "Z" {
			sum += 3 + 3
		}
		if v[0:1] == "C" && v[2:3] == "Y" {
			sum += 0 + 2
		}
		if v[0:1] == "C" && v[2:3] == "X" {
			sum += 6 + 1
		}

	}
	return sum, nil
}

func part2() (int, error) {
	var sum int
	lines, err := readInput("./1.txt")
	if err != nil {
		return 0, err
	}
	for _, v := range lines {

		// fmt.Println(v[0:1], v[2:3])
		if v[0:1] == "A" && v[2:3] == "X" {
			sum += 0 + 3
		}
		if v[0:1] == "A" && v[2:3] == "Y" {
			sum += 3 + 1
		}
		if v[0:1] == "A" && v[2:3] == "Z" {
			sum += 6 + 2
		}

		if v[0:1] == "B" && v[2:3] == "X" {
			sum += 0 + 1
		}
		if v[0:1] == "B" && v[2:3] == "Y" {
			sum += 3 + 2
		}
		if v[0:1] == "B" && v[2:3] == "Z" {
			sum += 6 + 3
		}

		if v[0:1] == "C" && v[2:3] == "X" {
			sum += 0 + 2
		}
		if v[0:1] == "C" && v[2:3] == "Y" {
			sum += 3 + 3
		}
		if v[0:1] == "C" && v[2:3] == "Z" {
			sum += 6 + 1
		}

	}
	return sum, nil
}
