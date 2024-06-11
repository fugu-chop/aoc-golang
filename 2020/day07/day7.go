package main

import (
	"log"
	"os"
	"regexp"
	"strings"
)

var (
	targetBag    = "shiny gold"
	fileLocation = "./input.txt"
	lineage      = []string{}
)

func main() {
	file, err := os.ReadFile(fileLocation)
	if err != nil {
		log.Fatal(err)
	}

	fileString := string(file)
	rules := strings.Split(fileString, "\n")
	for _, rule := range rules {
		if hasTargetBagChildren(rule) {
			parent := getParent(rule)
			if len(parent) > 0 {
				lineage = append(lineage, parent)
			}
		}
	}

	/*
		Problem
			Input: A file of strings separated by newline
			Output: Int, representing how many types of bags can carry a shiny gold bad

			We are given a large list of rules.
			These rules specify different types of bags and what other types of bags they can carry

			_ASSUMPTION_: A bag cannot carry itself

			Given a ruleset, we want to know how many types of bags can carry a shiny gold bag
			either directly (holds a shiny gold back itself) or
			indirectly (holds a bag that can hold shiny gold bag)

		Example
			light red bags contain 1 bright white bag, 2 muted yellow bags.
			dark orange bags contain 3 bright white bags, 4 muted yellow bags.
			bright white bags contain 1 shiny gold bag.
			muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
			shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
			dark olive bags contain 3 faded blue bags, 4 dotted black bags.
			vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
			faded blue bags contain no other bags.
			dotted black bags contain no other bags.

			Bright white and Muted yellow bags directly carry shiny gold bags
			Dark orange can hold bright white or muted yellow
			Light red can also carry bright white or muted yellow

			Answer is `4`

		Data Structures
			A tree-like type - node that can have parent or child nodes
			Within the type, we will need the name (colour), children (slice of nodes) and parents (slice of nodes)

		Algorithm
			We can establish 'shiny gold' as the root node
			We read the whole file into memory
			Break up into a slice by new lines
			Iterate over it
			Apply our regex to find 'shiny gold' after 'bags contain'
				Populate the bag type with name = 'shiny gold'
			We then know the immediate parents of the 'shiny gold'
				Create bags for the parents
				Populate the shiny gold bag's parents with the parent bags
			We can repeat the above steps for the immediate parents until there are no more parents
			We can then start counting with our shiny gold bags, and iterate over the parents
				Will likely involve some recursion here - `parents` might have multiple parents
				Stop when there are no more parents
			Print the count
	*/
}

func getParent(rule string) string {
	parent := strings.Split(rule, "bags contain")[0]

	if parent == targetBag {
		return ""
	}

	return parent
}

func hasTargetBagChildren(rule string) bool {
	re := regexp.MustCompile(`\d+ ` + targetBag + ` bag`)
	matches := re.FindStringSubmatch(rule)

	return len(matches) > 0
}
