package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Combine these to a single type
type CountCriteria struct {
	letter   string
	minCount int
	maxCount int
}

type IndexCriteria struct {
	letter      string
	firstIndex  int
	secondIndex int
}

func main() {
	/*
		Problem
			get an input of passwords and policies in the format
			3-11 z: zzzzzdzzzzlzz

			Format is x-y z where:
				x and y are 'positions' the n-th character of the string
					EITHER BUT NOT BOTH must be true
					e.g. for our example, idx: 2 OR idx: 10 must be "z", but not both
				z is the letter of the alphabet

			NOTE position != index:
				index 0 == position 1

			Password must meet both criteria to be compliant.

			Our job is to return an int, representing the total number of compliant passwords

		Edge Cases
			- Based on manually reading the input file, no uppercase
			- will always have a lower and upper bound represented as "integer"
			- min is 1, max is no larger than string length

		Data Structures
			- Throwaway slices for string splitting
			- Map for the password calculation
			- int Var for tracking compliant passwords
			- Custom type for criteria is easiest

		Algo
			We will need to parse an input file, row by row, ideally (use scanner)
			For each line (iterate):
				Split on `:` character, this will give a slice of two elements, a 'guide' and pw

				Password:
					- Trim whitespace
					- We don't even need to memoise, just need to access an index of a string
						- Can probably convert string(password[idx])
					- Need an out of bounds index check on the string to avoid panics

				Criteria:
					- Split by space, turn into a type
					- We need to convert our 'positions' into indexes (position - 1)

			Return intVar
	*/

	fileLocation := flag.String("inputLocation", "input.txt", "the location where the input file is")
	variant := flag.Int("part", 2, "which part of the solution is attempted")
	flag.Parse()

	file, err := os.Open(*fileLocation)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var validCounts int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// switch statement
		if *variant == 1 {
			if checkCountsCompliance(scanner.Text()) {
				validCounts += 1
			}
		}

		if *variant == 2 {
			if checkPositionCompliance(scanner.Text()) {
				validCounts += 1
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(validCounts)
}

/*
checkPositionCompliance expects a string in the form of 9-14 d: ddddbdddddddxfdd
it will handle trimming of whitespaces
*/
func checkPositionCompliance(line string) bool {
	components := strings.Split(line, ":")
	// can probably pass as func args
	criteria := parseIndexCriteria(components[0])
	positions := findPasswordPositions(components[1], criteria.firstIndex, criteria.secondIndex)

	if (positions[criteria.firstIndex] == criteria.letter && positions[criteria.secondIndex] != criteria.letter) ||
		(positions[criteria.firstIndex] != criteria.letter && positions[criteria.secondIndex] == criteria.letter) {
		return true
	}

	return false
}

/*
checkCountsCompliance expects a string in the form of 9-14 d: ddddbdddddddxfdd
it will handle trimming of whitespaces
*/
func checkCountsCompliance(line string) bool {
	components := strings.Split(line, ":")
	criteria := parseCountCriteria(components[0])
	memoPassword := memoisePassword(components[1])

	if memoPassword[criteria.letter] >= criteria.minCount &&
		memoPassword[criteria.letter] <= criteria.maxCount {
		return true
	}

	return false
}

func parseIndexCriteria(criteria string) *IndexCriteria {
	// extract to method
	splitCriteria := strings.Split(criteria, " ")
	if len(splitCriteria) != 2 {
		return nil
	}

	counts := strings.Split(splitCriteria[0], "-")
	if len(counts) != 2 {
		return nil
	}

	minCount, err := strconv.Atoi(counts[0])
	if err != nil {
		log.Fatal(err)
	}
	maxCount, err := strconv.Atoi(counts[1])
	if err != nil {
		log.Fatal(err)
	}

	return &IndexCriteria{
		letter:      splitCriteria[1],
		firstIndex:  minCount - 1,
		secondIndex: maxCount - 1,
	}
}

/*
parseCriteria expects a string in the form of "3-11 z"
i.e. trimmed of leading and trailing whitespaces
*/
func parseCountCriteria(criteria string) *CountCriteria {
	splitCriteria := strings.Split(criteria, " ")
	if len(splitCriteria) != 2 {
		return nil
	}

	counts := strings.Split(splitCriteria[0], "-")
	if len(counts) != 2 {
		return nil
	}

	minCount, err := strconv.Atoi(counts[0])
	if err != nil {
		log.Fatal(err)
	}
	maxCount, err := strconv.Atoi(counts[1])
	if err != nil {
		log.Fatal(err)
	}

	return &CountCriteria{
		letter:   splitCriteria[1],
		minCount: minCount,
		maxCount: maxCount,
	}
}

func memoisePassword(password string) map[string]int {
	counts := map[string]int{}
	cleanedPassword := strings.TrimSpace(password)
	letterSlice := strings.Split(cleanedPassword, "")

	for _, letter := range letterSlice {
		counts[letter] += 1
	}

	return counts
}

func findPasswordPositions(password string, firstIndex, secondIndex int) map[int]string {
	positions := map[int]string{}
	cleanedPassword := strings.TrimSpace(password)
	letterSlice := strings.Split(cleanedPassword, "")

	positions[firstIndex] = letterSlice[firstIndex]
	positions[secondIndex] = letterSlice[secondIndex]

	return positions
}
