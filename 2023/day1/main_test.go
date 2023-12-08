package main

import (
	"aocInput"
	"log"
	"testing"
)

func Test1(t *testing.T) {
	lines, err := input.GetInputLines("./testinput.txt")
	if err != nil {
		t.Fail()
	}

	sum := part2(lines)
	log.Printf("%d", sum)
	if sum != 249 {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	lines, err := input.GetInputLines("./example2.txt")
	if err != nil {
		t.Fail()
	}

	sum := part2(lines)
	log.Printf("%d", sum)
	if sum != 281 {
		t.Fail()
	}
}
func Test3(t *testing.T) {
	lines, err := input.GetInputLines("./input.txt")
	if err != nil {
		t.Fail()
	}

	sum := part2(lines)
	log.Printf("%d", sum)
	if sum != 54208 {
		t.Fail()
	}
}
