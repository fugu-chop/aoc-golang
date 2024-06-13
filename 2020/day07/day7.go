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
	fileLocation = "./sample_input.txt"
	lineage      = map[string]*bag{}
	terminus     = " no other bags."
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
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		memoiseRelationships(scanner.Text(), lineage)
	}

	for _, c := range lineage[targetBag].parents {
		fmt.Println("parent name: " + c.name)
	}

	// total := countParents(lineage, targetBag)
	// fmt.Println(total)

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

func memoiseRelationships(rule string, relationship map[string]*bag) {
	parent := &bag{}
	relationshipComponents := strings.Split(rule, "bags contain")
	parent.name = relationshipComponents[0]
	relationship[parent.name] = parent
	children := relationshipComponents[1]
	if children == terminus {
		return
	}

	createChildBags(children, parent, relationship)
}

func createChildBags(children string, parent *bag, relationship map[string]*bag) {
	bagNames := strings.Split(children, ",")

	re := regexp.MustCompile(`\d (.+) bag`)

	for _, bagName := range bagNames {
		matches := re.FindAllStringSubmatch(bagName, -1)
		childName := matches[0][1]

		_, ok := relationship[childName]
		if !ok {
			childBag := &bag{
				name: childName,
			}
			relationship[childName] = childBag
		}

		relationship[childName].parents = append(relationship[childName].parents, parent)
		parent.children = append(parent.children, relationship[childName])
	}
}

func countParents(relationship map[string]*bag, node string) int {
	fmt.Println(relationship[node].name)

	for _, parent := range relationship[node].parents {
		fmt.Println(parent.name)
		// countParents(relationship, parent.name)
	}

	return 1
}
