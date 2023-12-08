package main

import (
	"aocInput"
	// "fmt"
	"log"
	"strconv"
	"strings"
)

func part1(lines []string) (sum int) {

	for _, line := range lines {
		var nums []string
		var lineNum string
		for _, ch := range line {
			_, err := strconv.Atoi(string(ch))
			if err == nil {
				nums = append(nums, string(ch))
			}
		}
		lineNum = nums[0] + nums[len(nums)-1]
		num, err := strconv.Atoi(lineNum)
		if err != nil {
			log.Fatalf("Couldn't parse the lineNum %s: %v", lineNum, err)
		}
		sum += num
	}
	return
}

type numMap struct {
	word  string
	value string
}

var numbers = [9]numMap{numMap{word: "one", value: "o1e"}, numMap{word: "two", value: "t2o"}, numMap{word: "three", value: "t3e"}, numMap{word: "four", value: "f4r"}, numMap{word: "five", value: "f5e"}, numMap{word: "six", value: "s6x"}, numMap{word: "seven", value: "s7n"}, numMap{word: "eight", value: "e8t"}, numMap{word: "nine", value: "n9e"}}

func part2(lines []string) (sum int) {
	// iterate over lines in the file
	for i, line := range lines {
	Outerloop:
		for j, _ := range line {
			for _, num := range numbers {
				if len(line[j:]) < len(num.word) {
					// fmt.Printf("Comparing len(line[j:]) < len(num.word): %d < %d\n", len(line[j:]), len(num.word))
					continue
				}
				// fmt.Printf("Precheck: num: %s, slice: %s, i: %d, j: %d line: %s, lines[i]: %s\n", num.word, line[j:j+len(num.word)], i, j, line, lines[i])
				if strings.Contains(line[j:j+len(num.word)], num.word) {
					// fmt.Printf("Before: line: %s, lines[i]: %s\n", line, lines[i])
					newLine := strings.Replace(line, num.word, num.value, 1)
					lines[i] = newLine
					line = newLine
					// fmt.Printf("After: %s, lines[i]: %s\n", line, lines[i])
					goto Outerloop
				}
			}
		}

	}

	return part1(lines)
}

func main() {

	lines, err := input.GetInputLines("./input.txt")
	if err != nil {
		log.Fatalf("Could not read file: %v", err)
	}

	// sum := part1(lines)
	sum := part2(lines)

	log.Printf("Sum is: %v", sum)

}
