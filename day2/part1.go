package main

import (
	"fmt"
	"strings"

	"github.com/daanjo3/adventofcode2022/helper"
)

type Gesture int64

const (
	Rock    Gesture = 1
	Paper           = 2
	Scissor         = 3
)

func parseGesture(key string) Gesture {
	switch key {
	case "A":
		fallthrough
	case "X":
		return Rock
	case "B":
		fallthrough
	case "Y":
		return Paper
	case "C":
		fallthrough
	case "Z":
		return Scissor
	}
	panic("This should never be reached")
}

func doMatch(me Gesture, opponent Gesture) int {
	if me == opponent {
		return 3
	}
	if me == opponent+1 || me == opponent-2 {
		return 6
	}
	return 0
}

func getGesturePoint(gesture Gesture) int {
	switch gesture {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissor:
		return 3
	}
	panic("This should never be reached")
}

func main() {
	matches := helper.ReadLines("input.txt")

	total := 0
	for _, match := range matches {
		gestureCodes := strings.Fields(match)
		opponent := parseGesture(gestureCodes[0])
		me := parseGesture(gestureCodes[1])
		points := doMatch(me, opponent) + getGesturePoint(me)
		fmt.Println(gestureCodes, points)
		total += points
	}
	fmt.Println(total)
}
