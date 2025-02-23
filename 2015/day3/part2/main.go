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
	previousCoordinateSanta := baseCoordinate
	previousCoordinateRobot := baseCoordinate
	coordinatesVisited := map[string]int{
		baseCoordinate: 1,
	}
	moves := strings.Split(string(data), "")
	pairedCoordinates := []string{}

	for idx, move := range moves {
		pairedCoordinates = append(pairedCoordinates, move)
		if (idx+1)%2 == 0 {
			intPreviousCoordinateSanta,
				intPreviousCoordinateRobot := unstringifyCoordinate(previousCoordinateSanta, previousCoordinateRobot)

			intNewCoordinateSanta := incrementCoordinate(intPreviousCoordinateSanta, pairedCoordinates[0])
			newCoordinateSanta := stringifyCoordinate(intNewCoordinateSanta)
			intNewCoordinateRobot := incrementCoordinate(intPreviousCoordinateRobot, pairedCoordinates[1])
			newCoordinateRobot := stringifyCoordinate(intNewCoordinateRobot)

			if _, ok := coordinatesVisited[newCoordinateSanta]; !ok {
				coordinatesVisited[newCoordinateSanta] += 1
			}

			if _, ok := coordinatesVisited[newCoordinateRobot]; !ok {
				coordinatesVisited[newCoordinateRobot] += 1
			}

			previousCoordinateSanta = newCoordinateSanta
			previousCoordinateRobot = newCoordinateRobot

			pairedCoordinates = []string{}
		}
	}

	fmt.Println(len(coordinatesVisited))
}

func unstringifyCoordinate(coordinateSanta string, coordinateRobot string) ([]int, []int) {
	var (
		santaX int
		santaY int
		robotX int
		robotY int
	)
	_, err := fmt.Sscanf(coordinateSanta, "[%d,%d]", &santaX, &santaY)
	if err != nil {
		panic(err)
	}
	_, err = fmt.Sscanf(coordinateRobot, "[%d,%d]", &robotX, &robotY)
	if err != nil {
		panic(err)
	}

	return []int{santaX, santaY}, []int{robotX, robotY}
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
