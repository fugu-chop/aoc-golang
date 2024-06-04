package main

import (
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
	file, err := os.ReadFile(fileLocation)
	if err != nil {
		log.Fatal(err)
	}

	stringFile := string(file)
	groups := strings.Split(stringFile, "\n\n")

	var totalUniqueAnswers int

	for _, resps := range groups {
		uniqueAnswers := []string{}
		answers := strings.Split(resps, "")

		for _, answer := range answers {
			cleanedAnswer := strings.TrimSpace(answer)

			/*
				new requirement - only increment totalUniqueAnswers
				if a letter is present on every row

				We need to be smarter than just iterating over every single entry
				as this will have O(n!) time complexity - not great

				Potential approach is to select an entry with the shortest length -
				this is our lowest common denominator

				we can then iterate over each letter of that entry
				We don't need to retain which letter it is, only that it was present

				Algorithm
					- Given a group, e.g.

						heqznia
						cipkn
						gvsitwynrxb

					- Since the strings are all English ASCII, we can safely use len(string) to find
						the shortest slice, shortestAnswers
					- We can then iterate over shorestAnswers, checking if each other answer contains
					all of the letters in shortestAnswers
					- Increment counter if yes
					- Print counter
			*/
			if !slices.Contains(uniqueAnswers, cleanedAnswer) && len(cleanedAnswer) > 0 {
				uniqueAnswers = append(uniqueAnswers, cleanedAnswer)
			}
		}

		totalUniqueAnswers += len(uniqueAnswers)
	}

	fmt.Println(totalUniqueAnswers)
}
