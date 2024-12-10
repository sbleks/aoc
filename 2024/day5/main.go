package main

import (
	input "aocInput"
	"fmt"
	"log"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Node struct {
	before []int
	after  []int
}

func (n Node) String() string {
	return fmt.Sprintf("before: %v after: %v", n.before, n.after)
}

func parseInput(in string) (map[int]Node, [][]int) {
	// Parse page ordering
	parts := strings.Split(in, "\n\n")
	orderInput, pageNumInput := parts[0], parts[1]
	pageNumSlice := strings.Split(pageNumInput, "\n")
	orders := strings.Split(orderInput, "\n")
	graph := make(map[int]Node)
	for _, orderPair := range orders {
		order := strings.Split(orderPair, "|")
		left, err := strconv.Atoi(order[0])
		if err != nil {
			log.Panicf("Could not parse left side of order instruction: %v", err)
		}
		right, err := strconv.Atoi(order[1])
		if err != nil {
			log.Panicf("Could not parse right side of order instruction: %v", err)
		}

		valLeft, okLeft := graph[left]
		if okLeft {
			if slices.Contains(valLeft.after, right) {
				continue
			}
			valLeft.after = append(valLeft.after, right)
			graph[left] = valLeft
		} else {
			newNode := Node{after: []int{right}}
			graph[left] = newNode
		}

		valRight, okRight := graph[right]
		if okRight {
			if slices.Contains(valRight.before, left) {
				continue
			}
			valRight.before = append(valRight.before, left)
			graph[right] = valRight
		} else {
			newNode := Node{before: []int{left}}
			graph[right] = newNode
		}

	}

	// Parse Page Nums
	pageNums := [][]int{}
	for _, pageNum := range pageNumSlice {
		// skip if empty line
		if pageNum == "" {
			continue
		}
		numStrSlice := strings.Split(pageNum, ",")
		numSlice := []int{}
		for _, numStr := range numStrSlice {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Panicf("Cannot parse page num to int: %v", err)
			}
			numSlice = append(numSlice, num)
		}

		pageNums = append(pageNums, numSlice)
	}

	return graph, pageNums
}

func checkOrder(g map[int]Node, curr, next int) bool {
	if slices.Contains(g[curr].after, next) {
		return true
	}

	return false
}

func Less(g map[int]Node, curr, next int) bool {
	if slices.Contains(g[curr].before, next) {
		return true
	}

	return false
}

// func part1(in string) (sum int) {
// 	graph, pageNums := parseInput(in)

// 	for _, pageNum := range pageNums {
// 		inOrder := false
// 		for i := 1; i < len(pageNum); i++ {
// 			inOrder = checkOrder(graph, pageNum[i-1], pageNum[i])
// 			// log.Printf("%v", inOrder)
// 			if !inOrder {
// 				break
// 			}
// 		}
// 		if inOrder {
// 			mid := pageNum[len(pageNum)/2]
// 			sum += mid
// 		}
// 	}

// 	return sum
// }

func checkPageNumOrders(g map[int]Node, pageNum []int, part2 bool) bool {
	if !part2 {
		part2 = false
	}

	inOrder := false
	for i := 1; i < len(pageNum); i++ {
		inOrder = checkOrder(g, pageNum[i-1], pageNum[i])
		if !inOrder {
			if part2 {
				sort.SliceStable(pageNum, func(i, j int) bool {
					return Less(g, pageNum[i], pageNum[j])
				})
				return false
			} else {
				return false
			}
		}
	}
	return true
}

func part1(in string) (sum int) {
	graph, pageNums := parseInput(in)

	for _, pageNum := range pageNums {
		inOrder := checkPageNumOrders(graph, pageNum, false)
		if inOrder {
			mid := pageNum[len(pageNum)/2]
			sum += mid
		}
	}

	return sum
}

func part2(in string) (sum int) {
	graph, pageNums := parseInput(in)

	for _, pageNum := range pageNums {
		inOrder := checkPageNumOrders(graph, pageNum, true)
		if !inOrder {
			mid := pageNum[len(pageNum)/2]
			sum += mid
		}
	}

	return sum
}

func main() {
	in, err := input.GetRawInput("./input.txt")
	if err != nil {
		log.Fatalf("Could not read file: %v", err)
	}

	sum1 := part1(in)
	sum2 := part2(in)

	log.Printf("Part 1  is: %v", sum1)
	log.Printf("Part 2  is: %v", sum2)
}
