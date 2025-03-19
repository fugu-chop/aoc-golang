package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type instructions struct {
	Action  string
	FirstX  int
	FirstY  int
	SecondX int
	SecondY int
}

func main() {
	f, err := os.Open("/Users/dean/Documents/aoc-golang/2015/day6/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	grid := generateLightGrid()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		instruction := generateInstruction(scanner.Text())
		alterGrid(instruction, grid)
	}

	lightsOn := calculateLightsOn(grid)

	fmt.Println(lightsOn)
}

func calculateLightsOn(grid [][]int) int {
	var lightsOn int

	for i := 0; i <= 999; i++ {
		for j := 0; j <= 999; j++ {
			lightsOn += grid[i][j]
		}
	}

	return lightsOn
}

func alterGrid(instruction instructions, grid [][]int) {
	for i := instruction.FirstX; i <= instruction.SecondX; i++ {
		for j := instruction.FirstY; j <= instruction.SecondY; j++ {
			switch instruction.Action {
			case "toggle":
				grid[i][j] += 2
			case "on":
				grid[i][j] += 1
			case "off":
				if grid[i][j] >= 1 {
					grid[i][j] -= 1
				}
			}
		}
	}
}

func generateLightGrid() [][]int {
	grid := make([][]int, 1000)
	for i := range grid {
		grid[i] = make([]int, 1000)
	}

	return grid
}

func generateInstruction(input string) instructions {
	coordinateIdx := 1
	actionIdx := 0

	words := strings.Split(input, " ")

	var (
		firstCoordX  int
		firstCoordY  int
		secondCoordX int
		secondCoordY int
	)

	if len(words) == 5 {
		actionIdx = 1
		coordinateIdx = 2
	}

	action := words[actionIdx]
	_, err := fmt.Sscanf(
		strings.Join(words[coordinateIdx:], " "),
		"%d,%d through %d,%d",
		&firstCoordX, &firstCoordY,
		&secondCoordX, &secondCoordY)
	if err != nil {
		panic(err)
	}

	return instructions{
		Action:  action,
		FirstX:  firstCoordX,
		FirstY:  firstCoordY,
		SecondX: secondCoordX,
		SecondY: secondCoordY,
	}
}
