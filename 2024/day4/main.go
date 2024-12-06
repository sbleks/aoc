package main

import (
	input "aocInput"
	"fmt"
	"log"

	// "strconv"
	"strings"
)

type PuzzleBoard [][]rune

type Puzzle struct {
	board     PuzzleBoard
	rowCursor int
	colCursor int
}

func (pb PuzzleBoard) String() string {
	rows := []string{}
	for _, lines := range pb {
		row := []string{}
		for _, line := range lines {
			row = append(row, string(line))
		}
		rows = append(rows, fmt.Sprintf("%v", row))
	}
	return strings.Join(rows, "\n")
}

func (p Puzzle) String() string {
	return fmt.Sprintf("\n{ rowCursor: %v\ncolCursor: %v\nboard:\n%v }\n", p.rowCursor, p.colCursor, p.board)
}

func (p Puzzle) peek(row, col int) bool {
	ch := p.board[col][row]

	return false
}

func newPuzzle(lines []string) *Puzzle {
	board := [][]rune{}
	for _, line := range lines {
		row := []rune{}
		for _, ch := range line {
			row = append(row, ch)
		}
		board = append(board, row)
	}

	return &Puzzle{board: board}
}

func part1(lines []string) (sum int) {
	p := newPuzzle(lines)
	log.Printf("%v", p)
	return sum
}

func part2(lines []string) (sum int) {
	return sum
}

func main() {
	lines, err := input.GetInputLines("./sample.txt")
	if err != nil {
		log.Fatalf("Could not read file: %v", err)
	}

	sum1 := part1(lines)
	sum2 := part2(lines)

	log.Printf("Part 1  is: %v", sum1)
	log.Printf("Part 2  is: %v", sum2)
}
