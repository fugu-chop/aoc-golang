package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	/*
		Problem
			Need to calculate the highest Seat ID in the sample
			- Seat ID is Column + (Row * 8) calculation

			Given a 10 character string, figure out the Seat ID
				- First 7 characters are either "F" or "B"
					- These are used to deduce the Row number.
					- There are 128 rows on a plane, numbered 0 to 127
					- "F" means use the 1st half (0-63), "B" means use the 2nd half (64-127)
				- Last 3 characters are "L" or "R"
					- These are used to deduce the Column number.
					- There are 8 columns on a plane, numbered 0 to 7
					- "L" means take the first half (0-3), "R" means take the 2nd half (4-7)

		Examples
			Assume all examples will have 10 characters
			- first 7 are "F" or "B" characters
			- last 3 are "L" or "B" charcters

			FBFBBFFRLR
			Row
			- "F" -> (0-63)
			- "B" -> (32-63)
			- "F" -> (32-47)
			- "B" -> (40-47)
			- "B" -> (44-47)
			- "F" -> (44-45)
			- "F" -> (44-44)

			Column
			- "R" -> (4-7)
			- "L" -> (4-5)
			- "R" -> (5)

			Seat ID
			- (Row * 8) + Column
			(44 * 8) + 5


			BBFFBBFRLL
			Row
			- "B" -> (64-127)
			- "B" -> (96-127)
			- "F" -> (96-111)
			- "F" -> (96-103)
			- "B" -> (100-103)
			- "B" -> (102-103)
			- "F" -> (102-102)

			Column
			- "R" -> (4-7)
			- "L" -> (4-5)
			- "L" -> (4-4)

			Seat ID
			- (102 * 8) + 4

		Data Structures
		- Input is a file, with strings delimited by a newline
		- Each example is a string
		- Return a single int

		- Type SeatPosition to keep track of row, column, seat ID?
		- Type Range to keep track of upper and lower limit?

		Algorithm
		- Initialise highestSeatID
		- Input string
		- Break into slice of characters
		- Create a new seatPosition type
		- Create new boundary type, with value of 0 and 127 as upper and lower
		- Iterate over first 7 characters
			- If "F/L", take 1st half
				- retain lower
				- take upper + 1, subtract lower, half, subtract 1, -= upper
			- If "B/R" take 2nd half
				- retain upper
				- take upper + 1, subtract lower, half, += lower
		- This returns the row - populate seatPosition Row

		- Iterate over last 3 characters
			Repeat above algo
		- Populate seatPosition Column

		- Calculate Seat ID
		- Compare with highestSeatID, replace if larger
		- Print highestSeatID
	*/

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rowBoundary := &boundary{
			lower: firstRow,
			upper: lastRow,
		}
		colBoundary := &boundary{
			lower: firstCol,
			upper: lastCol,
		}
		seatPosition := &seatPosition{}
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

		seatPosition.column = colBoundary.lower
		seatPosition.row = rowBoundary.lower
		seatPosition.seatID = (seatPosition.row * 8) + seatPosition.column

		if seatPosition.seatID > maxSeatID {
			maxSeatID = seatPosition.seatID
		}
	}

	fmt.Println(maxSeatID)
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
