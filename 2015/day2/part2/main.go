package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	f, err := os.Open("/Users/dean/Documents/aoc-golang/2015/day2/input.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var totalRibbon int

	for scanner.Scan() {
		var (
			length int
			width  int
			height int
		)
		_, err := fmt.Sscanf(scanner.Text(), "%dx%dx%d", &length, &width, &height)
		if err != nil {
			panic(err)
		}

		ordered := []int{length, width, height}
		slices.Sort(ordered)

		perimeter := 2*ordered[0] + 2*ordered[1]
		ribbon := length * width * height

		totalRibbon += (perimeter + ribbon)
	}

	fmt.Println(totalRibbon)
}
