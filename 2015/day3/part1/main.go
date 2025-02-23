package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("/Users/dean/Documents/aoc-golang/2015/day3/input.txt")
	if err != nil {
		panic(err)
	}

	baseCoordinate := "[0,0]"
	previousCoordinate := baseCoordinate
	coordinatesVisited := map[string]int{
		baseCoordinate: 1,
	}
	moves := strings.Split(string(data), "")

	for _, move := range moves {
		intPreviousCoordinate := unstringifyCoordinate(previousCoordinate)
		intNewCoordinate := incrementCoordinate(intPreviousCoordinate, move)
		newCoordinate := stringifyCoordinate(intNewCoordinate)

		if _, ok := coordinatesVisited[newCoordinate]; !ok {
			coordinatesVisited[newCoordinate] += 1
		}
		previousCoordinate = newCoordinate
	}

	fmt.Println(len(coordinatesVisited))
}

func unstringifyCoordinate(coordinate string) []int {
	var (
		x int
		y int
	)
	_, err := fmt.Sscanf(coordinate, "[%d,%d]", &x, &y)
	if err != nil {
		panic(err)
	}

	return []int{x, y}
}

func incrementCoordinate(intPreviousCoordinate []int, move string) []int {
	switch move {
	case "v":
		intPreviousCoordinate[1] -= 1
	case "^":
		intPreviousCoordinate[1] += 1
	case ">":
		intPreviousCoordinate[0] += 1
	case "<":
		intPreviousCoordinate[0] -= 1
	}

	return intPreviousCoordinate
}

func stringifyCoordinate(intCoordinate []int) string {
	return fmt.Sprintf("[%d,%d]", intCoordinate[0], intCoordinate[1])
}
