package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	treeChar = "#"
)

type jump struct {
	horizontal int
	vertical   int
}

type coordinate struct {
	height      int
	width       int
	coordinates map[int][]string
}

func main() {
	part := flag.Int("part", 1, "define which problem part the solution should attempt to solve")
	flag.Parse()

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	coordinate := generateCoordinates(scanner)

	jumps := []jump{
		{
			horizontal: 3,
			vertical:   1,
		},
	}

	if *part == 2 {
		jumps = append(jumps,
			jump{horizontal: 1, vertical: 1},
			jump{horizontal: 5, vertical: 1},
			jump{horizontal: 7, vertical: 1},
			jump{horizontal: 1, vertical: 2},
		)
	}
	treesHit := []int{}

	for _, jump := range jumps {
		treesHit = append(treesHit, coordinate.calculateTreesHit(jump))
	}

	totalHit := 1
	for _, hits := range treesHit {
		totalHit *= hits
	}

	fmt.Printf("trees hit: %d\n", totalHit)
}

/*
countTrees parses a slice of strings, incrementing a count
when encountering the tree character.
*/
func (c *coordinate) countTrees(currentRowIdx int, row []string) int {
	var trees int
	if currentRowIdx > len(row) {
		return trees
	}

	if row[currentRowIdx] == treeChar {
		trees += 1
	}
	return trees
}

/*
updateCurrentRowIdx updates where the horizontal position of the person
is after each move down vertically. It handles the situation where a horizontal
jump would land them out of bounds by subtracting the out of bounds index from
the max row length incremented by the horizontal jump.
*/
func (c *coordinate) updateCurrentRowIdx(currentRowIdx, rowLength, jump int) int {
	if (currentRowIdx + jump) >= rowLength {
		return currentRowIdx - rowLength + jump
	}
	return currentRowIdx + jump
}

/*
calculateTreesHit iterates over the coordinate typed object and calculates
the number of trees hit using the global variables that define the horizontal
and vertical jumps. It handles scenarios where the horizontal jump exceeds
the horizontal length of a row
*/
func (c *coordinate) calculateTreesHit(jumps jump) int {
	var treesHit, currentRowIdx, currentHeight int

	for currentHeight < c.height {
		row := c.coordinates[currentHeight]
		treesHit += c.countTrees(currentRowIdx, row)
		currentRowIdx = c.updateCurrentRowIdx(currentRowIdx, c.width, jumps.horizontal)
		currentHeight += jumps.vertical
	}

	return treesHit
}

/*
generateCoordinates takes in a pointer to a bufio.Scanner interface
and returns a coordinate type
*/
func generateCoordinates(scanner *bufio.Scanner) *coordinate {
	coordinate := new(coordinate)
	coordinate.coordinates = make(map[int][]string)
	coordinateIdx := 0

	for scanner.Scan() {
		coordinate.coordinates[coordinateIdx] = strings.Split(scanner.Text(), "")
		coordinateIdx += 1
	}

	coordinate.height = len(coordinate.coordinates)
	coordinate.width = len(coordinate.coordinates[0])

	return coordinate
}
