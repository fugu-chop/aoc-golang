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
type criteria struct {
	letter string
	first  int
	second int
}

func main() {
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
		switch *variant {
		case 1:
			if checkCountsCompliance(scanner.Text()) {
				validCounts += 1
			}
		case 2:
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
	criteria := parseIndexCriteria(components[0])
	positions := findPasswordPositions(components[1], criteria.first, criteria.second)

	if (positions[criteria.first] == criteria.letter && positions[criteria.second] != criteria.letter) ||
		(positions[criteria.first] != criteria.letter && positions[criteria.second] == criteria.letter) {
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
	criteria := parseCriteria(components[0])
	memoPassword := memoisePassword(components[1])

	if memoPassword[criteria.letter] >= criteria.first &&
		memoPassword[criteria.letter] <= criteria.second {
		return true
	}

	return false
}

/*
parseCriteria expects a string in the form of "3-11 z"
i.e. trimmed of leading and trailing whitespaces
*/
func parseCriteria(stringCriteria string) *criteria {
	splitCriteria := strings.Split(stringCriteria, " ")
	if len(splitCriteria) != 2 {
		return nil
	}

	counts := strings.Split(splitCriteria[0], "-")
	if len(counts) != 2 {
		return nil
	}

	first, err := strconv.Atoi(counts[0])
	if err != nil {
		log.Fatal(err)
	}
	second, err := strconv.Atoi(counts[1])
	if err != nil {
		log.Fatal(err)
	}

	return &criteria{
		letter: splitCriteria[1],
		first:  first,
		second: second,
	}
}

func parseIndexCriteria(criteria string) *criteria {
	parsedCriteria := parseCriteria(criteria)
	parsedCriteria.first -= 1
	parsedCriteria.second -= 1
	return parsedCriteria
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
