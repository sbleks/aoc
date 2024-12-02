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

type Report struct {
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
		report = calculateDiffs(report)

		reports = append(reports, report)

	}
	return reports
}

func checkReportSafety(report Report) Report {
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

	if !report.diffIssue && !report.directionIssue {
		report.safe = true
	}
	return report
}

func calculateDiffs(report Report) Report {
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
	return report
}

func removeLevel(report Report, s int) Report {
	r := report
	newReport := Report{
		levels: append(r.levels[:s], r.levels[s+1:]...),
	}

	// log.Printf("removeLevel New Report: %v", newReport)

	return newReport
}

func (r Report) Copy() Report {
	levels := make([]int, len(r.levels))
	diffs := make([]int, len(r.diffs))
	unsafeLevelPositions := make([]int, len(r.unsafeLevelPositions))
	copy(levels, r.levels)
	copy(diffs, r.diffs)
	copy(unsafeLevelPositions, r.unsafeLevelPositions)

	return Report{
		levels:               levels,
		diffs:                diffs,
		unsafeLevelPositions: unsafeLevelPositions,
		increasing:           r.increasing,
		decreasing:           r.decreasing,
		diffIssue:            r.diffIssue,
		directionIssue:       r.directionIssue,
		safe:                 r.safe,
	}
}

func part2(lines []string) (sum int) {
	reports := parseReports(lines)
	for _, report := range reports {
		report = checkReportSafety(report)

		if len(report.unsafeLevelPositions) != 0 {
			for _, issueIdx := range report.unsafeLevelPositions {
				log.Printf("%v", report)

				r := report.Copy()
				newReport1 := removeLevel(r, issueIdx)
				// log.Printf("after newReport1 created\nr: %v \n report: %v\n newReport1: %v", r, report, newReport1)
				r = report.Copy()
				newReport2 := removeLevel(r, issueIdx+1)
				// log.Printf("after newReport2 created\nr: %v \n report: %v\n newReport1: %v", r, report, newReport1)
				newReport1 = calculateDiffs(newReport1)
				newReport1 = checkReportSafety(newReport1)
				// log.Printf("newReport1: %v", newReport1)

				newReport2 = calculateDiffs(newReport2)
				newReport2 = checkReportSafety(newReport2)
				// log.Printf("newReport2: %v", newReport2)

				if newReport1.safe {
					report = newReport1
					log.Printf("[UPDATE]: Setting newReport1 to report: %v", report)
					break
				} else if newReport2.safe {
					report = newReport2
					log.Printf("[UPDATE]: Setting newReport2 to report: %v", report)
					break
				}
			}
		}

		if report.safe {
			log.Printf("[COUNTING]: %v", report)
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
