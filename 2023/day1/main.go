package main

import (
	"aocInput"
	"fmt"
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

var numbers = [9]numMap{numMap{word: "one", value: "1"}, numMap{word: "two", value: "2"}, numMap{word: "three", value: "3"}, numMap{word: "four", value: "4"}, numMap{word: "five", value: "5"}, numMap{word: "six", value: "6"}, numMap{word: "seven", value: "7"}, numMap{word: "eight", value: "8"}, numMap{word: "nine", value: "9"}}

type calNum struct {
	value     string
	stringNum string
	num       int
	start     int
	end       int
}

type calNumArr []calNum

func (c calNum) String() string {
	return fmt.Sprintf("{value: \"%s\", stringNum: \"%s\", num: %d start: %d, end: %d}\n", c.value, c.stringNum, c.num, c.start, c.end)
}

func (c calNumArr) Len() int      { return len(c) }
func (c calNumArr) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c calNumArr) Less(i, j int) bool {
	if c[i].start < c[j].start {
		return true
	} else {
		return false
	}
}

/*
xtwone3four
 ^ ^
 j k


*/

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
			// 	for _, num := range numbers {
			// 		if rune(num.word[j]) == ch {
			// 			// peek next character and collect until it fails?
			// 			collector = collector + string(ch)
			// 			continue
			// 		} else{
			// 			break
			// 		}
			// 	}
			// 	// for strings.Contains(line, num.word) {
			// 	// 	fmt.Printf("Before: %s\n", line)
			// 	// 	newLine := strings.Replace(line, num.word, num.value, 1)
			// 	// 	lines[i] = newLine
			// 	// 	line = newLine
			// 	// 	fmt.Printf("After: %s\n", line)
			// 	// }
		}

		// // create a row in the calNums array to store data from that line
		// calNums = append(calNums, []calNum{})
		//
		// for j, ch := range line {
		// 	_, err := strconv.Atoi(string(ch))
		// 	if err == nil {
		// 		val, err := strconv.Atoi(string(ch))
		// 		if err != nil {
		// 			log.Panicf("Could not convert %c to int: %v", ch, err)
		// 		}
		// 		calNums[i] = append(calNums[i], calNum{
		// 			value:     string(ch),
		// 			stringNum: string(ch),
		// 			num:       val,
		// 			start:     j,
		// 			end:       j + 1,
		// 		})
		// 	}
		// }
		//
		// // iterate over each of the numbers
		// // TODO: if the line has two of the same numbers, like twotwo, this would only return one number
		// for j, number := range numbers {
		// 	if strings.Contains(line, number) {
		// 		start := strings.Index(line, number)
		// 		calNums[i] = append(calNums[i], calNum{
		// 			value:     number,
		// 			stringNum: fmt.Sprintf("%d", j+1),
		// 			num:       j + 1,
		// 			start:     start,
		// 			end:       start + len(number),
		// 		})
		// 		rest := line[start+len(number):]
		// 		if i == 712 {
		// 			log.Printf("%v: %s", calNums[i], rest)
		// 		}
		// 		if strings.Contains(rest, number) {
		// 			start2 := strings.Index(rest, number)
		// 			calNums[i] = append(calNums[i], calNum{
		// 				value:     number,
		// 				stringNum: fmt.Sprintf("%d", j+1),
		// 				num:       j + 1,
		// 				start:     start + start2,
		// 				end:       start2 + len(number),
		// 			})
		// 			// log.Printf("%d: %v", i, calNums[i])
		// 			if i == 712 {
		// 				log.Printf("%v: %s", calNums[i], rest)
		// 			}
		//
		// 			rest2 := line[start2+len(number):]
		// 			if strings.Contains(rest2, number) {
		// 				start3 := strings.Index(rest2, number)
		// 				calNums[i] = append(calNums[i], calNum{
		// 					value:     number,
		// 					stringNum: fmt.Sprintf("%d", j+1),
		// 					num:       j + 1,
		// 					start:     start + start3,
		// 					end:       start3 + len(number),
		// 				})
		// 				// log.Printf("%d: %v", i, calNums[i])
		// 				if i == 712 {
		// 					log.Printf("%v: %s", calNums[i], rest)
		// 				}
		// 			}
		// 		}
		// 	}
		// }
		//
		// if len(calNums[i]) < 2 {
		// 	log.Printf("Line %d doesn't have enough numbers parsed: %v", i+1, calNums[i])
		// }
		// fmt.Printf("%v\n", lines)
	}

	return part1(lines)
	// for k, v := range calNums {
	// 	k += 1
	// 	sort.Sort(calNumArr(v))
	// 	// log.Printf("%v", v)
	// 	// if len(v) > 1 {
	// 	num1 := v[0].num * 10
	// 	num2 := v[len(v)-1].num
	// 	sum += (num1 + num2)
	// 	if k == 713 {
	// 		log.Printf("%d: (%d, %d). %d + %d = %d. sum: %d", k, num1, num2, num1, num2, num1+num2, sum)
	// 	}
	// }
	// log.Printf("Post-sort: %v", calNums)
	// log.Printf("%v", calNums[712])
	// return 0
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
