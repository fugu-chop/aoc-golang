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

type Criteria struct {
	letter   string
	minCount int
	maxCount int
}

func main() {
	/*
		Problem
			get an input of passwords and policies in the format
			3-11 z: zzzzzdzzzzlzz

			Format is x-y z where:
				x is the 'positive' position - MUST contain the letter
				y is the 'negative' position - MUST NOT contain the letter
				z is the letter of the alphabet

			NOTE position != index:
				index 0 == position 1

			Password must meet both criteria to be compliant.

			Our job is to return an int, representing the total number of compliant passwords

		Edge Cases
			- Based on manually reading the input file, no uppercase
			- will always have a lower and upper bound represented as "integer"

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
	variant := flag.Int("part", 1, "which part of the solution is attempted")
	flag.Parse()

	file, err := os.Open(*fileLocation)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var validCounts int

	scanner := bufio.NewScanner(file)

	if *variant == 1 {
		for scanner.Scan() {
			if checkCountsCompliance(scanner.Text()) {
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
checkCompliance expects a string in the form of 9-14 d: ddddbdddddddxfdd
it will handle trimming of whitespaces
*/
func checkCountsCompliance(line string) bool {
	components := strings.Split(line, ":")
	criteria := parseCriteria(components[0])
	memoPassword := memoisePassword(components[1])

	if memoPassword[criteria.letter] >= criteria.minCount &&
		memoPassword[criteria.letter] <= criteria.maxCount {
		return true
	}

	return false
}

/*
parseCriteria expects a string in the form of "3-11 z"
i.e. trimmed of leading and trailing whitespaces
*/
func parseCriteria(criteria string) *Criteria {
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

	return &Criteria{
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
