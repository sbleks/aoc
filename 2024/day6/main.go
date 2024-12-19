package main

import (
	input "aocInput"
	"fmt"
	"log"
	"slices"
	// "strconv"
	// "strings"
)

type Puzzle struct {
	puzzleMap [][]rune
	// x, y
	pos []int
	// x, y
	dir         []int
	arrow       rune
	breadcrumb  rune
	startingPos []int
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
				p.startingPos = []int{j, i}
				p.dir = []int{0, 1}
				row[j] = 'X'
				p.arrow = v
				continue
			case '<':
				p.pos = []int{j, i}
				p.startingPos = []int{j, i}
				p.dir = []int{-1, 0}
				row[j] = 'X'
				p.arrow = v
				continue
			case '^':
				p.pos = []int{j, i}
				p.startingPos = []int{j, i}
				p.dir = []int{0, -1}
				row[j] = 'X'
				p.arrow = v
				continue
			case '>':
				p.pos = []int{j, i}
				p.startingPos = []int{j, i}
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
	p.breadcrumb = 'X'
	return p
}

func (p Puzzle) String() string {
	out := fmt.Sprintf("pos (x,y): %v, dir (x,y): %v\n", p.pos, p.dir)

	for i := 0; i < len(p.puzzleMap); i++ {
		out += fmt.Sprintf("%03d: %v\n", i, string(p.puzzleMap[i]))
	}

	return out
}

func (p Puzzle) turnRight() Puzzle {
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
	return p
}

func encPosDir(pos, dir []int) string {
	return fmt.Sprintf("%v%v", pos, dir)
}

func (op Puzzle) Copy() Puzzle {
	newPuzzleMap := [][]rune{}
	for i := 0; i < len(op.puzzleMap); i++ {
		newPuzzleMap = append(newPuzzleMap, slices.Clone(op.puzzleMap[i]))
	}

	return Puzzle{
		puzzleMap:  newPuzzleMap,
		pos:        slices.Clone(op.pos),
		dir:        slices.Clone(op.dir),
		arrow:      op.arrow,
		breadcrumb: op.breadcrumb,
	}

}

func (op Puzzle) checkForLoop() bool {
	p := op.Copy()

	done, p := p.checkForWin()
	if done {
		return false
	}

	seen := make(map[string]bool)
	multiCount := 0

	// curr := p.puzzleMap[p.pos[1]][p.pos[0]]
	// next := p.puzzleMap[p.pos[1]+p.dir[1]][p.pos[0]+p.dir[0]]

	seen[encPosDir(p.pos, p.dir)] = true
	p.puzzleMap[p.pos[1]+p.dir[1]][p.pos[0]+p.dir[0]] = 'O'
	p = p.turnRight()
	run := true
	for run {
		// log.Printf("Checking for loop: %v\n%v", seen, p)
		run, p = p.walk()
		if !run {
			return false
		}
		_, ok := seen[encPosDir(p.pos, p.dir)]
		if ok {
			// p.puzzleMap[p.pos[1]][p.pos[0]] = 'D'
			// return true
			multiCount++
			p.breadcrumb = 'H'
			// log.Printf("Got another mulitcount!\n")
		}
		// } else {
		// 	multiCount = 0
		// 	p.breadcrumb = 'X'
		// }
		seen[encPosDir(p.pos, p.dir)] = true

		if multiCount > 10000 {
			// log.Printf("%v\n\n%v\n\n", op, p)
			// log.Printf("Got 10 mulitcounts!\n")
			// log.Printf("looped at: %v\n%v\n%v\n\n", encPosDir(p.pos, p.dir), seen, p)
			return true
		}
	}

	return false
}

func (p Puzzle) checkForWin() (bool, Puzzle) {
	if p.pos[1]+p.dir[1] >= len(p.puzzleMap) || p.pos[0]+p.dir[0] >= len(p.puzzleMap[p.pos[1]]) || p.pos[1]+p.dir[1] < 0 || p.pos[0]+p.dir[0] < 0 {
		p.puzzleMap[p.pos[1]][p.pos[0]] = p.breadcrumb
		return true, p
	}
	return false, p
}

func (p Puzzle) walk() (bool, Puzzle) {
	done, p := p.checkForWin()
	if done {
		return false, p
	}
	// log.Printf("Not done, walking: %v", p)

	// curr := p.puzzleMap[p.pos[1]][p.pos[0]]
	next := p.puzzleMap[p.pos[1]+p.dir[1]][p.pos[0]+p.dir[0]]

	switch next {
	case '.', 'X', 'H':

		p.puzzleMap[p.pos[1]+p.dir[1]][p.pos[0]+p.dir[0]] = p.arrow
		p.puzzleMap[p.pos[1]][p.pos[0]] = p.breadcrumb

		p.pos[0] += p.dir[0]
		p.pos[1] += p.dir[1]
		return true, p

	case '#', 'O':
		p = p.turnRight()
		done, p := p.checkForWin()
		if done {
			return false, p
		}
		p.puzzleMap[p.pos[1]][p.pos[0]] = p.breadcrumb
		p.puzzleMap[p.pos[1]+p.dir[1]][p.pos[0]+p.dir[0]] = p.arrow
		p.pos[0] += p.dir[0]
		p.pos[1] += p.dir[1]
		return true, p
	}

	return false, p
}

func part1(p Puzzle) (sum int) {
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
	// log.Printf("%v", p)
	run := true
	for run {
		run, p = p.walk()
		if !run {
			break
		}
		if p.pos[1]+p.dir[1] == p.startingPos[1]+p.dir[1] && p.pos[0]+p.dir[0] == p.startingPos[0]+p.dir[0] {
			log.Printf("Skipping starting pos, pos: %v, startingPos: %v, dir: %v \n", p.pos, p.startingPos, p.dir)
			continue
		} else {
			if p.checkForLoop() {
				sum++
			}
		}
	}
	// log.Printf("%v", p)

	// for i := 0; i < len(p.puzzleMap); i++ {
	// 	for j := 0; j < len(p.puzzleMap[i]); j++ {
	// 		if p.puzzleMap[i][j] == 'O' {
	// 			sum++
	// 		}
	// 	}
	// }
	return sum
}

func main() {
	lines, err := input.GetInputLines("./input.txt")
	if err != nil {
		log.Fatalf("Could not read file: %v", err)
	}
	// f, err := os.OpenFile("out.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatalf("error opening file: %v", err)
	// }
	// defer f.Close()

	// log.SetOutput(f)

	// p1 := NewPuzzle(lines)
	p2 := NewPuzzle(lines)

	// sum1 := part1(p1)
	sum2 := part2(p2)

	// log.Printf("Part 1  is: %v", sum1)
	log.Printf("Part 2  is: %v", sum2)
}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"slices"
// )

// type index struct {
// 	r, c int
// }

// func main() {
// 	puzzleInput, startIndex := parseInput("input.txt")
// 	visitedIndices, res := partOne(puzzleInput, startIndex)

// 	fmt.Println("Part One: ", res)
// 	fmt.Println("Part Two: ", partTwo(puzzleInput, startIndex, visitedIndices))
// }

// func parseInput(fileName string) ([][]rune, index) {
// 	file, err := os.Open(fileName)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)

// 	var grid [][]rune
// 	var startIndex index
// 	for scanner.Scan() {
// 		grid = append(grid, []rune(scanner.Text()))
// 		if c := slices.Index(grid[len(grid)-1], '^'); c != -1 {
// 			startIndex = index{r: len(grid) - 1, c: c}
// 			grid[startIndex.r][startIndex.c] = '.'
// 		}
// 	}

// 	return grid, startIndex
// }

// func partOne(grid [][]rune, startIndex index) ([]index, int) {
// 	directions := []index{
// 		{-1, 0},
// 		{0, 1},
// 		{1, 0},
// 		{0, -1},
// 	}
// 	maxRow := len(grid)
// 	maxCol := len(grid[0])

// 	visited := make(map[index]bool)

// 	at := startIndex
// 	facing := 0

// 	for isValidIndex(at, maxRow, maxCol) {
// 		if grid[at.r][at.c] == '#' {
// 			at.r -= directions[facing].r
// 			at.c -= directions[facing].c
// 			facing = (facing + 1) % len(directions)
// 			continue
// 		}
// 		visited[at] = true
// 		grid[at.r][at.c] = 'X'
// 		at.r += directions[facing].r
// 		at.c += directions[facing].c
// 	}

// 	visitedIndices := make([]index, 0, len(visited))

// 	for idx := range visited {
// 		visitedIndices = append(visitedIndices, idx)
// 	}

// 	return visitedIndices, len(visitedIndices)
// }

// func partTwo(grid [][]rune, startIndex index, visitedIndices []index) int {
// 	cycleCount := 0
// 	for _, index := range visitedIndices {
// 		if (index == startIndex) || grid[index.r][index.c] == '#' {
// 			continue
// 		}
// 		grid[index.r][index.c] = '#'
// 		if hasCycle(grid, startIndex) {
// 			cycleCount++
// 		}
// 		grid[index.r][index.c] = '.'
// 	}

// 	return cycleCount
// }

// func hasCycle(grid [][]rune, startIndex index) bool {
// 	visited := make(map[index]index)

// 	directions := []index{
// 		{-1, 0},
// 		{0, 1},
// 		{1, 0},
// 		{0, -1},
// 	}
// 	maxRow := len(grid)
// 	maxCol := len(grid[0])

// 	at := startIndex
// 	facing := 0

// 	for isValidIndex(at, maxRow, maxCol) {
// 		if visited[at] == directions[facing] {
// 			return true
// 		}

// 		visited[at] = directions[facing]

// 		if grid[at.r][at.c] == '#' {
// 			at.r -= directions[facing].r
// 			at.c -= directions[facing].c
// 			facing = (facing + 1) % len(directions)
// 			continue
// 		}

// 		grid[at.r][at.c] = 'X'
// 		at.r += directions[facing].r
// 		at.c += directions[facing].c
// 	}

// 	return false
// }

// func isValidIndex(idx index, maxRow, maxCol int) bool {
// 	return idx.r >= 0 && idx.c >= 0 && idx.r < maxRow && idx.c < maxCol
// }
