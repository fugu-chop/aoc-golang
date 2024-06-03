package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

var (
	firstRow   = 0
	lastRow    = 127
	firstCol   = 0
	lastCol    = 7
	front      = "F"
	back       = "B"
	left       = "L"
	right      = "R"
	lastRowIdx = 7
	maxSeatID  = 0
)

type seatPosition struct {
	row    int
	column int
	seatID int
}

type boundary struct {
	lower int
	upper int
}

func main() {
	part := flag.Int("part", 1, "which part should be attempted")
	flag.Parse()

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	seatIDs := []int{}

	for scanner.Scan() {
		rowBoundary := &boundary{
			lower: firstRow,
			upper: lastRow,
		}
		colBoundary := &boundary{
			lower: firstCol,
			upper: lastCol,
		}
		boardingPassSlice := strings.Split(scanner.Text(), "")

		for _, row := range boardingPassSlice[:lastRowIdx] {
			if row == front {
				recalculateUpperBound(rowBoundary)
			}
			if row == back {
				recalculateLowerBound(rowBoundary)
			}
		}
		for _, column := range boardingPassSlice[lastRowIdx:] {
			if column == left {
				recalculateUpperBound(colBoundary)
			}
			if column == right {
				recalculateLowerBound(colBoundary)
			}
		}

		seatPosition := createSeatPosition(rowBoundary.lower, colBoundary.lower)

		if *part == 1 {
			if seatPosition.seatID > maxSeatID {
				maxSeatID = seatPosition.seatID
			}
		} else {
			seatIDs = append(seatIDs, seatPosition.seatID)
		}
	}

	if *part == 1 {
		fmt.Println(maxSeatID)
	} else {
		missingSeat := findMissingSeat(seatIDs)
		fmt.Println(missingSeat)
	}

}

func recalculateUpperBound(boundary *boundary) *boundary {
	newUpper := boundary.upper + 1
	newRange := (newUpper - boundary.lower) / 2
	boundary.upper -= newRange
	return boundary
}

func recalculateLowerBound(boundary *boundary) *boundary {
	newLower := boundary.upper + 1
	boundary.lower += (newLower - boundary.lower) / 2
	return boundary
}

func createSeatPosition(row, column int) *seatPosition {
	return &seatPosition{
		row:    row,
		column: column,
		seatID: (row * 8) + column,
	}
}

func findMissingSeat(seats []int) int {
	slices.Sort(seats)
	current := seats[1]

	for _, seat := range seats {
		if seat+1 != current {
			return current - 1
		}
		current += 1
	}

	return current
}
