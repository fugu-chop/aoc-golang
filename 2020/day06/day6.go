package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

var (
	fileLocation = "./input.txt"
)

func main() {
	part := flag.Int("part", 1, "which part of the solution is being solved")
	flag.Parse()

	file, err := os.ReadFile(fileLocation)
	if err != nil {
		log.Fatal(err)
	}

	stringFile := string(file)
	groups := strings.Split(stringFile, "\n\n")

	var totalUniqueAnswers int

	if *part == 1 {
		for _, resps := range groups {
			uniqueAnswers := []string{}

			answers := strings.Split(resps, "")
			for _, answer := range answers {
				cleanedAnswer := strings.TrimSpace(answer)

				if !slices.Contains(uniqueAnswers, cleanedAnswer) && len(cleanedAnswer) > 0 {
					uniqueAnswers = append(uniqueAnswers, cleanedAnswer)
				}
			}

			totalUniqueAnswers += len(uniqueAnswers)
		}
	}

	if *part == 2 {
		for _, group := range groups {
			counts, numAnswers := memoiseGroup(group)
			for _, value := range counts {
				if value == numAnswers {
					totalUniqueAnswers++
				}
			}
		}
	}

	fmt.Println(totalUniqueAnswers)
}

/*
memoiseGroup takes in a string with multiple newlines and returns
a map of letters and the number of times they occur in the group
*/
func memoiseGroup(group string) (map[string]int, int) {
	counts := map[string]int{}
	groups := cleanGroup(group)

	// concatenate the groups into a single string
	answerSlice := strings.Split(strings.Join(groups, ""), "")
	numAnswers := len(groups)

	for _, answer := range answerSlice {
		counts[answer]++
	}

	return counts, numAnswers
}

func cleanGroup(group string) []string {
	output := []string{}
	clean := strings.Split(group, "\n")

	for _, item := range clean {
		trimmedItem := strings.TrimSpace(item)
		if len(trimmedItem) > 0 {
			output = append(output, trimmedItem)
		}
	}

	return output
}
