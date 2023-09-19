package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	part1 := part1()
	fmt.Println("Part 1", part1)
	// part2 := part2()
	// fmt.Println("Part 2", part2)
}

func readInput(path string) (lines []string, err error) {
	file, err := os.Open(path)
	defer file.Close()
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

func part1() (sum int){

	lines, _ := readInput("./3.txt")

	lines.

	return sum
}
