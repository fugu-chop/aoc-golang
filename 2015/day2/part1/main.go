package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("/Users/dean/Documents/aoc-golang/2015/day2/input.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	var total int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		dimensions := strings.Split(scanner.Text(), "x")

		var dimensionsSlice []int
		for _, dimension := range dimensions {
			dimensionInt, err := strconv.Atoi(dimension)
			if err != nil {
				panic(err)
			}
			dimensionsSlice = append(dimensionsSlice, dimensionInt)
		}

		// calculate the total area
		area := 2*(dimensionsSlice[0]*dimensionsSlice[1]) +
			2*(dimensionsSlice[1]*dimensionsSlice[2]) +
			2*(dimensionsSlice[0]*dimensionsSlice[2])

		// find the smallest two elements and multiply
		slices.Sort(dimensionsSlice)
		slack := dimensionsSlice[0] * dimensionsSlice[1]

		total += (area + slack)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(total)
}
