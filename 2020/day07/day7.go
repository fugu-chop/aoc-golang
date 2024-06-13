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

	fmt.Println(lineage[targetBag])
}

func memoiseRelationships(rule string, relationship map[string]*bag) {
	parent := &bag{}
	relationshipComponents := strings.Split(rule, "bags contain")
	parent.name = strings.TrimSpace(relationshipComponents[0])
	relationship[parent.name] = parent
	children := relationshipComponents[1]
	if children == terminus {
		return
	}

	createChildBags(children, parent, relationship)
}

// This is not populating the parents correctly for shiny gold for sample input
func createChildBags(children string, parent *bag, relationship map[string]*bag) {
	bagNames := strings.Split(children, ",")

	re := regexp.MustCompile(`\d (.+) bag`)

	for _, bagName := range bagNames {
		matches := re.FindAllStringSubmatch(bagName, -1)
		childName := strings.TrimSpace(matches[0][1])

		_, ok := relationship[childName]
		if !ok {
			childBag := &bag{
				name: childName,
			}
			relationship[childName] = childBag
		}

		// There is some bug here - this algorithm seems to only work
		// if the child appears before the parent (i.e. bag already exists)
		relationship[childName].parents = append(relationship[childName].parents, parent)
		parent.children = append(parent.children, relationship[childName])
	}
}
