package main

import (
	"aocInput"
	"errors"
	"fmt"
	"strconv"

	// "fmt"
	"log"
	// "strconv"
	"strings"
)

type cubeSet struct {
	red   int
	green int
	blue  int
}

type game struct {
	id   int
	sets []cubeSet
}

var p1win = cubeSet{red: 12, green: 13, blue: 14}

func parseGame(line string) (game game, err error) {
	gameInfo := strings.Split(line, ":")
	_, gameId, found := strings.Cut(gameInfo[0], " ")
	id, err := strconv.Atoi(gameId)
	if err != nil {
		return game, err
	}
	if found == false {
		err = errors.New("Could not parse game id")
		return game, err
	}

	game.id = id
	sets := strings.Split(strings.Trim(gameInfo[1], " "), "; ")
	for _, set := range sets {
		colors := strings.Split(set, ", ")
		cubeSet := cubeSet{}
		for _, colorSet := range colors {
			parsedColor := strings.Split(colorSet, " ")
			color := parsedColor[1]
			amount, err := strconv.Atoi(parsedColor[0])
			if err != nil {
				return game, err
			}

			switch color {
			case "red":
				cubeSet.red = amount
			case "green":
				cubeSet.green = amount
			case "blue":
				cubeSet.blue = amount
			}

		}
		game.sets = append(game.sets, cubeSet)
	}

	return game, nil
}

func part1(lines []string) (sum int) {
	for _, line := range lines {
		game, err := parseGame(line)
		if err != nil {
			log.Panicf("Could not parse game: %s\n", err)
		}
		possible := true
		for _, set := range game.sets {
			if set.blue > p1win.blue || set.green > p1win.green || set.red > p1win.red {
				possible = false
			}
		}
		if possible {
			sum += game.id
		}
	}
	return
}

func part2(lines []string) (sum int) {
	for _, line := range lines {
		game, err := parseGame(line)
		if err != nil {
			log.Panicf("Could not parse game: %s\n", err)
		}
		mini := cubeSet{}
		for _, set := range game.sets {
			for i := 0; i < 3; i++ {
				if set.blue > mini.blue {
					mini.blue = set.blue
				}

				if set.red > mini.red {
					mini.red = set.red
				}

				if set.green > mini.green {
					mini.green = set.green
				}
			}
		}
		power := mini.red * mini.green * mini.blue
		sum += power
	}
	return
}

func main() {
	lines, err := input.GetInputLines("./input.txt")
	if err != nil {
		log.Panicf("Could not read file")
	}

	// sum := part1(lines)
	sum := part2(lines)
	fmt.Printf("%d\n", sum)
}
