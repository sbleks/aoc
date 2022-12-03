package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	day1, _ := day1()
	fmt.Println("Day 1", day1)
	day2, _ := day2()
	fmt.Println("Day 2", day2)
}

func day1() (int, error) {
	var elves []int
	var sum int
	file, err := os.Open("./1.txt")
	if err != nil {
		return 0, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			elves = append(elves, sum)
			sum = 0
		}

		num, _ := strconv.Atoi(line)
		sum += num
	}

	maxval, err := max(elves)
	if err != nil {
		return 0, err
	}

	return maxval, nil
}

func max(arr []int) (int, error) {
	var max int
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	if max == 0 {
		return 0, errors.New("cannot find max")
	}
	return max, nil
}

func day2() (int, error) {
	var elves []int
	var sum int
	file, err := os.Open("./1.txt")
	if err != nil {
		return 0, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			elves = append(elves, sum)
			sum = 0
		}

		num, _ := strconv.Atoi(line)
		sum += num
	}
	sort.Ints(elves[:])
	// fmt.Println(elves[len(elves)-3:])
	var topthree int
	for _, num := range elves[len(elves)-3:] {
		topthree += num
	}
	return topthree, nil
}
