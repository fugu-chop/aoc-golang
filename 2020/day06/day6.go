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

type group struct {
	respondents   int
	uniqueAnswers []string
}

func main() {
	/*
		Problem
			We're given a txt file where meaningful entries are separated by two newlines (\n\n)
			An entry consists of one or many lines of strings
			Each line will contain unique letters of the alphabet - this represents a 'group' of people
				that have answered 'yes' to a question, where each letter represents a question

			We need to return the number of unique questions for each group that have a 'yes' answer
			i.e. show up - a distinct count of letters per two newlines (group)

			For now, we don't need to consider the number of groups for now, but it seems the problem
			will require us to think about it

		Examples & Edge Cases
			abc

			a
			b
			c

			ab
			ac

			a
			a
			a
			a

			b

			In the above examples, we have:
				- 5 groups
				- The first group has 1 person
					- this person has answered 'yes' to 3 questions
					- There are 3 unique 'yes'
				- The 2nd group has 3 people
					- Each person has answered 'yes' to a different question
					- There are 3 unique 'yes'
				- The 3rd group has two people
					- Each person has answered 'yes' to two different questions
					- There are 3 unique 'yes'
				- The 4th group has 4 people
					- Each person has answered 'yes' to 1 question
					- There are 1 unique 'yes'
				- The 5th group has 1 person
					- Each person answered 'yes' to 1 question
					- There are 1 unique 'yes'

			In this example, the answer we want is 3 + 3 + 3 + 1 + 1 => 11 (unique questions)

		Data Structures
			Slice of strings (each group, each line)
			A type to capture information - e.g. 'group' -> number of people, number of unique questions
			A map to aggregate this info -> key = index, value = group type
			A variable to capture number of unique 'yes' (var int)

		Algorithm
		- Ideally we would iterate over newlines via Scanner.Scan(), but the newlines will
		probably break the Scan() method (will leave out the last group)
		- Read entire file into memory (not great)
		- Split into a slice via \n\n character
			- This is our number of groups (var _groups_)
		- Iterate over _groups_
			- As we encounter each group
				- Create new _group_ type
				- populate the size of the group (len(group))
				- Calculate the number of unique answers within group
					- create a slice (uniqueChars)
					- break up a group into a slice of strings
					- Iterate over the broken up group, appending characters to uniqueChars
					if uniqueChars does not contain the character
					- Take the length of uniqueChars
				- Increment our _counts_ variable by len(uniqueChars)
		- Return uniqueChars
	*/

	file, err := os.ReadFile(fileLocation)
	if err != nil {
		log.Fatal(err)
	}

	stringFile := string(file)
	groups := strings.Split(stringFile, "\n\n")

	var totalUniqueAnswers int
	ds := map[int]*group{}

	for idx, resps := range groups {
		tempGroup := &group{
			respondents: len(resps),
		}

		// I think whitespace is somehow making it's way in here?
		uniqueAnswers := []string{}
		answers := strings.Split(resps, "")

		for _, answer := range answers {
			cleanedAnswer := strings.TrimSpace(answer)

			if !slices.Contains(uniqueAnswers, cleanedAnswer) && len(cleanedAnswer) > 0 {
				uniqueAnswers = append(uniqueAnswers, cleanedAnswer)
			}
		}

		tempGroup.uniqueAnswers = uniqueAnswers
		ds[idx] = tempGroup
		totalUniqueAnswers += len(uniqueAnswers)
	}

	fmt.Println(totalUniqueAnswers)
}
