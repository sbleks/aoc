package main

import (
	input "aocInput"
	"fmt"
	"log"
	// "strconv"
	// "strings"
)

type Puzzle struct {
	puzzleMap [][]rune
	// x, y
	pos []int
	// x, y
	dir   []int
	arrow rune
}

func NewPuzzle(lines []string) (p Puzzle) {
	out := make([][]rune, len(lines))
	for i := 0; i < len(lines); i++ {
		row := make([]rune, len(lines[i]))
		for j := 0; j < len(lines[i]); j++ {
			v := rune(lines[i][j])
			switch v {
			case 'v':
				p.pos = []int{j, i}
				p.dir = []int{0, 1}
				row[j] = 'X'
				p.arrow = v
				continue
			case '<':
				p.pos = []int{j, i}
				p.dir = []int{-1, 0}
				row[j] = 'X'
				p.arrow = v
				continue
			case '^':
				p.pos = []int{j, i}
				p.dir = []int{0, -1}
				row[j] = 'X'
				p.arrow = v
				continue
			case '>':
				p.pos = []int{j, i}
				p.dir = []int{1, 0}
				row[j] = 'X'
				p.arrow = v
				continue
			}
			row[j] = v
		}
		out[i] = row
	}
	p.puzzleMap = out
	return p
}

func (p Puzzle) String() string {
	out := fmt.Sprintf("pos (x,y): %v, dir (x,y): %v\n", p.pos, p.dir)

	for i := 0; i < len(p.puzzleMap); i++ {
		out += fmt.Sprintf("%v: %v\n", i, string(p.puzzleMap[i]))
	}

	return out
}

func (p *Puzzle) turnRight() {
	switch p.arrow {
	case 'v':
		p.dir = []int{-1, 0}
		p.arrow = '<'
	case '<':
		p.dir = []int{0, -1}
		p.arrow = '^'
	case '^':
		p.dir = []int{1, 0}
		p.arrow = '>'
	case '>':
		p.dir = []int{0, 1}
		p.arrow = 'v'
	}
}

func (p *Puzzle) walk() bool {
	if p.pos[1]+p.dir[1] >= len(p.puzzleMap) || p.pos[0]+p.dir[0] >= len(p.puzzleMap[p.pos[1]]) || p.pos[1]+p.dir[1] < 0 || p.pos[0]+p.dir[0] < 0 {
		p.puzzleMap[p.pos[1]][p.pos[0]] = 'X'
		return false
	}

	next := p.puzzleMap[p.pos[1]+p.dir[1]][p.pos[0]+p.dir[0]]

	switch next {
	case '.', 'X':
		p.puzzleMap[p.pos[1]][p.pos[0]] = 'X'
		p.puzzleMap[p.pos[1]+p.dir[1]][p.pos[0]+p.dir[0]] = p.arrow
		p.pos[0] += p.dir[0]
		p.pos[1] += p.dir[1]
		p.walk()
	case '#':
		p.turnRight()
		p.puzzleMap[p.pos[1]][p.pos[0]] = 'X'
		p.puzzleMap[p.pos[1]+p.dir[1]][p.pos[0]+p.dir[0]] = p.arrow
		p.pos[0] += p.dir[0]
		p.pos[1] += p.dir[1]
		p.walk()
	}

	return false
}

func part1(p *Puzzle) (sum int) {
	p.walk()

	for i := 0; i < len(p.puzzleMap); i++ {
		for j := 0; j < len(p.puzzleMap[i]); j++ {
			if p.puzzleMap[i][j] == 'X' {
				sum++
			}
		}
	}
	return sum
}

func part2(p Puzzle) (sum int) {
	return sum
}

func main() {
	lines, err := input.GetInputLines("./input.txt")
	if err != nil {
		log.Fatalf("Could not read file: %v", err)
	}

	p1 := NewPuzzle(lines)
	p2 := NewPuzzle(lines)

	sum1 := part1(&p1)
	sum2 := part2(p2)

	log.Printf("Part 1  is: %v", sum1)
	log.Printf("Part 2  is: %v", sum2)
}
