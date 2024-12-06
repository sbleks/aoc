package main

import (
	input "aocInput"
	"log"
	"strconv"
	"strings"
)

// // Part 1 plan is to iterate through every line and if the line is deemed to be safe (meets conditions), then we add it to our sum. If it fails, just return early and move on
// func part1(lines []string) (sum int) {

// 	for _, line := range lines {
// 		items := strings.Split(line, " ")

// 		var curr, prev int = 0, 0
// 		var diffState, directionState, inc, dec bool = true, false, false, false

// 		// log.Printf("Starting row: %v", row)

// 		for col, item := range items {
// 			level, err := strconv.Atoi(item)
// 			if err != nil {
// 				log.Fatalf("Cannot parse level to number: %v", err)
// 			}
// 			// log.Printf("Starting row: %v col: %v, curr: %v, prev: %v", row, col, curr, prev)

// 			curr = level
// 			diff := curr - prev

// 			// log.Printf("[ITEM INIT] curr: %v, prev: %v, diff: %v\n", curr, prev, diff)

// 			prev = curr

// 			if col == 0 {
// 				// log.Print("first iteration, skipping")
// 				continue
// 			}

// 			if diff > 0 {
// 				inc = true
// 			}

// 			if diff < 0 {
// 				dec = true
// 				diff = -diff
// 			}

// 			if diff > 3 || diff < 1 {
// 				diffState = false
// 				// log.Printf("[ITEM STATE] Breaking. diffState: %v, inc: %v, dec: %v\n", diffState, inc, dec)
// 				break
// 			}

// 			// log.Printf("[ITEM STATE] diffState: %v, inc: %v, dec: %v\n", diffState, inc, dec)
// 		}

// 		if (inc && !dec) || (dec && !inc) {
// 			directionState = true
// 		}

// 		if diffState && directionState {
// 			sum += 1
// 		}
// 		// log.Printf("[LEVEL STATE] diffState: %v, directionState: %v, inc: %v, dec: %v, sum: %v\n", diffState, directionState, inc, dec, sum)
// 	}
// 	return sum
// }

func parseReports(line string) (report []int) {
	items := strings.Split(line, " ")

	for _, item := range items {
		level, err := strconv.Atoi(item)
		if err != nil {
			log.Fatalf("Cannot parse level to number: %v", err)
		}

		report = append(report, level)
	}

	return report
}

// func checkReportSafety(levels []int) bool {
// 	flagInc, flagDec := false, false
// 	for i := 1; i < len(levels); i++ {
// 		diff := levels[i] - levels[i-1]

// 		if diff > 0 {
// 			flagInc = true
// 		} else if diff < 0 {
// 			flagDec = true
// 		} else {
// 			return false
// 		}

// 		if flagInc && flagDec {
// 			return false
// 		}

// 		if diff > 3 || diff < -3 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func checkReportSafetyWithDeletion(levels []int) bool {

// 	for i := 0; i < len(levels); i++ {
// 		return isReportSafeWithDeletion(levels, i)
// 	}

// 	return false
// }

// func isReportSafeWithDeletion(levels []int, s int) bool {
// 	r := make([]int, len(levels))
// 	copy(r, levels)

// 	if s == len(levels)-1 {
// 		r = r[:s]
// 	} else {
// 		r = append(r[:s], r[s+1:]...)
// 	}
// 	return checkReportSafety(r)
// }

// func part2(lines []string) (sum int) {
// 	reports := parseReports(lines)
// 	countWithDeletion := 0
// 	for _, report := range reports {

// 		if checkReportSafety(report) {
// 			sum += 1
// 		} else if checkReportSafetyWithDeletion(report) {
// 			countWithDeletion += 1
// 		}
// 	}

// 	log.Printf("Count of sums: %v, Count with Deletions: %v, Total Count: %v", sum, countWithDeletion, sum+countWithDeletion)

// 	return sum + countWithDeletion
// }

// func main() {
// 	lines, err := input.GetInputLines("./input.txt")
// 	if err != nil {
// 		log.Fatalf("Could not read file: %v", err)
// 	}

// 	sum1 := part1(lines)
// 	sum2 := part2(lines)

// 	log.Printf("Sum1 is: %v", sum1)
// 	log.Printf("Sum2 is: %v", sum2)
// }

func main() {
	lines, err := input.GetInputLines("./input.txt")
	if err != nil {
		log.Fatalf("Could not read file: %v", err)
	}

	getTotalSafeReportCount(lines)
}

func isReportSafe(reportNum []int) bool {
	flagIncrease, flagDecrease := false, false

	for i := 1; i < len(reportNum); i++ {
		diff := reportNum[i] - reportNum[i-1]

		if diff > 0 {
			flagIncrease = true
		} else if diff < 0 {
			flagDecrease = true
		} else {
			return false
		}

		if flagDecrease && flagIncrease {
			return false
		}

		if diff > 3 || diff < -3 {
			return false
		}
	}

	return true
}

func checkReportSafetyWithDeletion(reportNum []int) bool {

	for i := 0; i < len(reportNum); i++ {
		isSafe := isReportSafeWithDeletion(reportNum, i)
		if isSafe {
			return true
		}
	}

	return false
}

func isReportSafeWithDeletion(report []int, deleteIndex int) bool {
	copyReport := make([]int, len(report))
	copy(copyReport, report)

	if deleteIndex == len(copyReport)-1 {
		copyReport = copyReport[:deleteIndex]
	} else {
		copyReport = append(copyReport[:deleteIndex], copyReport[deleteIndex+1:]...)
	}
	return isReportSafe(copyReport)
}

func getTotalSafeReportCount(reports []string) int {
	var count int
	var countWithDeletion int
	for _, report := range reports {
		reportNum := parseReports(report)

		if isReportSafe(reportNum) {
			count++
		} else if checkReportSafetyWithDeletion(reportNum) {
			countWithDeletion++
		}
	}
	log.Printf("answer for part 1: %d\nanswer for part 2: %d\n", count, count+countWithDeletion)
	return count
}
