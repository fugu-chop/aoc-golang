package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("/Users/dean/Documents/aoc-golang/2015/day8/input.txt")
	if err != nil {
		panic(err)
	}

	var (
		codeChars   int
		memoryChars int
	)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		codeChars += len(scanner.Text())
		// Account for additional quotation marks at start and end for each word
		memoryChars += 2
		for _, x := range scanner.Text() {
			switch string(x) {
			// " -> \"
			case `"`:
				memoryChars += 2
			// \ -> \\
			case `\`:
				memoryChars += 2
			default:
				memoryChars += 1
			}
		}
	}

	fmt.Println(memoryChars - codeChars)
}
