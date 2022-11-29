package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	part1 := part1()
	part2 := part2()
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)

}

func part1() int {

	var changes int
	var prev int

	file, _ := os.Open("./day1.txt")

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		cur, _ := strconv.Atoi(fileScanner.Text())

		if prev == 0 {
			prev = cur
			continue
		}

		if cur > prev {
			changes++
		}

		prev = cur

	}

	return changes
}

func part2() int {
	var changes int
	var lines []int
	var j int

	file, _ := os.Open("./day1.txt")

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		cur, _ := strconv.Atoi(fileScanner.Text())
		lines = append(lines, cur)
	}

	var wins [][]int
	wins = append(wins, make([]int, 0, 666))

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		wins = append(wins, make([]int, 0, 3))

		if i == 0 {
			wins[j] = append(wins[j], line)
			continue
		}
		if i == 1 {
			wins[j] = append(wins[j], line)
			wins[j+1] = append(wins[j+1], line)
			continue
		}
		if i == 2 {
			wins[j] = append(wins[j], line)
			wins[j+1] = append(wins[j+1], line)
			wins[j+2] = append(wins[j+2], line)
			j++
			continue
		}

		if len(wins[j]) < 3 {
			wins[j] = append(wins[j], line)
		}

		if len(wins[j+1]) < 3 {
			wins[j+1] = append(wins[j+1], line)
		}

		if len(wins[j+2]) < 3 {
			wins[j+2] = append(wins[j+2], line)
		}

		if len(wins[j]) == 3 {

			j++
		}
	}

	var a int
	var b int
	for i, v := range wins {
		if i == 0 {
			a = v[0] + v[1] + v[2]
			continue
		}

		if len(v) < 3 {
			continue
		}

		b = v[0] + v[1] + v[2]
		if b > a {
			changes++
		}

		a = b
		b = 0

	}

	return changes
}
