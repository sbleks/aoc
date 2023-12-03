package main

import (
	"aocInput"
	"log"
	"strconv"
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

func main() {

	lines, err := input.GetInputLines("./input.txt")
	if err != nil {
		log.Fatalf("Could not read file: %v", err)
	}

	sum := part1(lines)

	log.Printf("%v", sum)

}
