package main

import (
	input "aocInput"
	"log"
	"strconv"
	"strings"
)

// Part 1 plan is to iterate through every line and if the line is deemed to be safe (meets conditions), then we add it to our sum. If it fails, just return early and move on
func part1(lines []string) (sum int) {

	for _, line := range lines {
		items := strings.Split(line, " ")

		var curr, prev int = 0, 0
		var diffState, directionState, inc, dec bool = true, false, false, false

		// log.Printf("Starting row: %v", row)

		for col, item := range items {
			level, err := strconv.Atoi(item)
			if err != nil {
				log.Fatalf("Cannot parse level to number: %v", err)
			}
			// log.Printf("Starting row: %v col: %v, curr: %v, prev: %v", row, col, curr, prev)

			curr = level
			diff := curr - prev

			// log.Printf("[ITEM INIT] curr: %v, prev: %v, diff: %v\n", curr, prev, diff)

			prev = curr

			if col == 0 {
				// log.Print("first iteration, skipping")
				continue
			}

			if diff > 0 {
				inc = true
			}

			if diff < 0 {
				dec = true
				diff = -diff
			}

			if diff > 3 || diff < 1 {
				diffState = false
				// log.Printf("[ITEM STATE] Breaking. diffState: %v, inc: %v, dec: %v\n", diffState, inc, dec)
				break
			}

			// log.Printf("[ITEM STATE] diffState: %v, inc: %v, dec: %v\n", diffState, inc, dec)
		}

		if (inc && !dec) || (dec && !inc) {
			directionState = true
		}

		if diffState && directionState {
			sum += 1
		}
		// log.Printf("[LEVEL STATE] diffState: %v, directionState: %v, inc: %v, dec: %v, sum: %v\n", diffState, directionState, inc, dec, sum)
	}
	return sum
}

type Report = struct {
	levels               []int
	diffs                []int
	unsafeLevelPositions []int
	increasing           bool
	decreasing           bool
	diffIssue            bool
	directionIssue       bool
	safe                 bool
}

func parseReports(lines []string) (reports []Report) {
	for _, line := range lines {
		report := Report{}
		items := strings.Split(line, " ")

		for _, item := range items {
			level, err := strconv.Atoi(item)
			if err != nil {
				log.Fatalf("Cannot parse level to number: %v", err)
			}

			report.levels = append(report.levels, level)

		}
		calculateDiffs(&report)

		reports = append(reports, report)

	}
	return reports
}

func checkReportSafety(report *Report) {
	for i, diff := range report.diffs {
		if diff > 0 {
			if report.decreasing {
				report.directionIssue = true
				report.unsafeLevelPositions = append(report.unsafeLevelPositions, i)
			}

			report.increasing = true
		}

		if diff < 0 {
			diff = -diff

			if report.increasing {
				report.directionIssue = true
				report.unsafeLevelPositions = append(report.unsafeLevelPositions, i)
			}

			report.decreasing = true
		}

		if diff > 3 || diff < 1 {
			report.diffIssue = true
			report.unsafeLevelPositions = append(report.unsafeLevelPositions, i)
		}

	}

	if report.decreasing && report.increasing {
		report.directionIssue = true
	}

	if !report.diffIssue && !report.directionIssue && len(report.unsafeLevelPositions) == 0 {
		report.safe = true
	}
}

func calculateDiffs(report *Report) {
	var curr, prev int = 0, 0

	for i, item := range report.levels {
		curr = item
		diff := curr - prev
		prev = curr
		if i == 0 {
			continue
		}
		report.diffs = append(report.diffs, diff)
	}
}

func part2(lines []string) (sum int) {
	reports := parseReports(lines)
	for _, report := range reports {
		checkReportSafety(&report)
		var newReport Report

		if len(report.unsafeLevelPositions) == 1 {
			log.Printf("%v", report)
			s := report.unsafeLevelPositions[0] + 1
			newReport = Report{
				levels: append(report.levels[:s], report.levels[s+1:]...),
			}
			calculateDiffs(&newReport)
			checkReportSafety(&newReport)
			log.Printf("%v", newReport)
		}

		if report.safe {
			sum += 1
		}

		if newReport.safe {
			sum += 1
		}

		// log.Printf("%v\n", report)
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

	log.Printf("Sum1 is: %v", sum1)
	log.Printf("Sum2 is: %v", sum2)
}
