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

func (p Puzzle) check1(dir int) bool {
	s := []string{}
	r := p.rowCursor
	c := p.colCursor
	// log.Printf("Calling check with row: %v, col: %v, dir: %v", r, c, dir)
	switch dir {
	case 0:
		for i := 0; i < 4; i++ {
			if !p.checkBounds(r-i, c) {
				return false
			}
			s = append(s, string(p.board[r-i][c]))
		}
	case 1:
		for i := 0; i < 4; i++ {
			if !p.checkBounds(r-i, c+i) {
				return false
			}
			s = append(s, string(p.board[r-i][c+i]))
		}
	case 2:
		for i := 0; i < 4; i++ {
			if !p.checkBounds(r, c+i) {
				return false
			}
			s = append(s, string(p.board[r][c+i]))
		}
	case 3:
		for i := 0; i < 4; i++ {
			if !p.checkBounds(r+i, c+i) {
				return false
			}
			s = append(s, string(p.board[r+i][c+i]))
		}
	case 4:
		for i := 0; i < 4; i++ {
			if !p.checkBounds(r+i, c) {
				return false
			}
			s = append(s, string(p.board[r+i][c]))
		}
	case 5:
		for i := 0; i < 4; i++ {
			if !p.checkBounds(r+i, c-i) {
				return false
			}
			s = append(s, string(p.board[r+i][c-i]))
		}
	case 6:
		for i := 0; i < 4; i++ {
			if !p.checkBounds(r, c-i) {
				return false
			}
			s = append(s, string(p.board[r][c-i]))
		}
	case 7:
		for i := 0; i < 4; i++ {
			if !p.checkBounds(r-i, c-i) {
				return false
			}
			s = append(s, string(p.board[r-i][c-i]))
		}

	}

	word := strings.Join(s, "")
	if word == "XMAS" {
		// log.Printf("GOT WORD: %v from row: %v, col: %v in direction: %v", word, r, c, dir)
		return true
	} else {
		// log.Printf("MISSED WORD: %v from row: %v, col: %v in direction: %v", word, r, c, dir)
		return false
	}
}

func (p Puzzle) check2(dir int, r int, c int) bool {
	s := []string{}
	// log.Printf("Calling check with row: %v, col: %v, dir: %v", r, c, dir)
	switch dir {
	// bottom-right
	case 0:
		for i := 0; i < 3; i++ {
			if !p.checkBounds(r+i, c+i) {
				return false
			}
			s = append(s, string(p.board[r+i][c+i]))
		}
		// bottom-left
	case 1:
		for i := 0; i < 3; i++ {
			if !p.checkBounds(r+i, c-i) {
				return false
			}
			s = append(s, string(p.board[r+i][c-i]))
		}

	}

	word := strings.Join(s, "")
	if word == "MAS" || word == "SAM" {
		log.Printf("GOT WORD: %v from row: %v, col: %v in direction: %v", word, r, c, dir)
		return true
	} else {
		// log.Printf("MISSED WORD: %v from row: %v, col: %v in direction: %v", word, r, c, dir)
		return false
	}
}

func (p Puzzle) checkBounds(row, col int) bool {
	if row >= len(p.board) || row < 0 || col < 0 || col >= len(p.board[row]) {
		return false
	}

	return true
}

func (p Puzzle) peek1() (sum int) {
	// 0 = up
	// 1 = top-right
	// 2 = right
	// 3 = bottom-right
	// 4 = down
	// 5 = bottom-left
	// 6 = left
	// 7 = top-left
	for i := 0; i < 8; i++ {
		if p.check1(i) {
			sum++
		}
	}

	return sum
}

func (p Puzzle) peek2() (sum int) {
	// 0 = bottom-right
	// 1 = bottom-left
	if p.check2(0, p.rowCursor, p.colCursor) && p.check2(1, p.rowCursor, p.colCursor+2) {
		log.Println("Got both directions")
		sum++
	}

	return sum
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
	// log.Printf("%v, %v", len(p.board), len(p.board[0]))
	for i := 0; i < len(p.board); i++ {
		p.colCursor = 0
		for j := 0; j < len(p.board[i]); j++ {
			// log.Printf("p.row: %v, p.col: %v", p.rowCursor, p.colCursor)
			sum += p.peek1()
			p.colCursor++
		}
		p.rowCursor++
	}

	// log.Printf("%v", p)
	return sum
}

func part2(lines []string) (sum int) {
	p := newPuzzle(lines)
	// log.Printf("%v, %v", len(p.board), len(p.board[0]))
	for i := 0; i < len(p.board); i++ {
		p.colCursor = 0
		for j := 0; j < len(p.board[i]); j++ {
			// log.Printf("p.row: %v, p.col: %v", p.rowCursor, p.colCursor)
			sum += p.peek2()
			p.colCursor++
		}
		p.rowCursor++
	}

	// log.Printf("%v", p)
	return sum
}

func main() {
	lines, err := input.GetInputLines("./input.txt")
	if err != nil {
		log.Fatalf("Could not read file: %v", err)
	}

	sum1 := part1(lines)
	sum2 := part2(lines)

	log.Printf("Part 1  is: %v", sum1)
	log.Printf("Part 2  is: %v", sum2)
}
