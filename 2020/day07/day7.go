package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var (
	targetBag    = "shiny gold"
	fileLocation = "./simple_input.txt"
)

type bag struct {
	name     string
	children []string
	parents  []string
}

func main() {
	file, err := os.Open(fileLocation)
	if err != nil {
		log.Fatal(err)
	}

	collection := map[string]*bag{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// Build our collection
		parseRule(scanner.Text(), collection)
	}

	fmt.Printf("%+v", collection)

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
func parseRule(rule string, collection map[string]*bag) []*bag {
	output := []*bag{}

	name, children := generateBagOutputsFromRule(rule)
	createBag(name, children, collection)

	return output
}

// TODO: Split this function out - createBag currently does too much
// have the children creation occur elsewhere
func createBag(name string, children []string, collection map[string]*bag) {
	result := &bag{}
	result.name = name
	if _, ok := collection[name]; !ok {
		collection[name] = result
	}

	// Split this out
	for _, child := range children {
		childBag := &bag{}
		childBag.name = child
		collection[child] = childBag
		childBag.parents = append(childBag.parents, result.name)
		result.children = append(result.children, childBag.name)
	}
}

/*
generateBagOutputsFromRule takes in a string of the format
"posh black bags contain 3 dark lavender bags, 3 mirrored coral bags, 1 dotted chartreuse bag."
and returns a string with the name and a slice of strings with children
*/
func generateBagOutputsFromRule(rule string) (string, []string) {
	re := regexp.MustCompile(`(\D+)( bags contain )(.+)`)

	/*
		This will return a nested slice of strings. E.g.
		[
			[
				"posh black bags contain 3 dark lavender bags, 3 mirrored coral bags, 1 dotted chartreuse bag.",
				"posh black",
				" bags contain ",
				"3 dark lavender bags, 3 mirrored coral bags, 1 dotted chartreuse bag."
			]
		]

		The second element of the inner slice will contain the colour of the bag.
		The 4th elements onward contain the children bags.
	*/
	rules := re.FindAllStringSubmatch(rule, -1)
	name := rules[0][1]
	children := cleanChildren(rules[0][3:])

	return name, children
}

/*
cleanChildren accepts a string slice with a single element
E.g. ["3 dark lavender bags, 3 mirrored coral bags, 1 dotted chartreuse bag."]
and returns a slice of the colours only
(excluding counts, punctuation and the word 'bag').
*/
func cleanChildren(children []string) []string {
	cleanedChildren := []string{}
	splitChildren := strings.Split(children[0], ", ")
	re := regexp.MustCompile(`\d+ (.+) bag`)

	for _, child := range splitChildren {
		cleanedChildren = append(cleanedChildren, re.FindStringSubmatch(child)[1])
	}

	return cleanedChildren
}
