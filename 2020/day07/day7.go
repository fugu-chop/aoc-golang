package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

var (
	targetBag    = "shiny gold"
	fileLocation = "./input.txt"
)

type bag struct {
	name     string
	children []*bag
	parents  []*bag
}

func main() {
	file, err := os.Open(fileLocation)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parseRule(scanner.Text())
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
			Read file, iterate over each line
			Parse the line - need to split out the bag (colour) and children bag (quantities can be ignored)
	*/
}

/*
parseRule takes in an input string, such as "pale chartreuse bags
contain 3 faded orange bags." and converts it into a pointer to a bag
type, specifying the colour of parent bag and it's child bags.
*/
func parseRule(rule string) *bag {
	output := bag{}

	re := regexp.MustCompile(`(\D+)( bags contain )(.+)`)
	strings := re.FindAllStringSubmatch(rule, -1)
	fmt.Printf("%q\n", strings)

	return &output
}
