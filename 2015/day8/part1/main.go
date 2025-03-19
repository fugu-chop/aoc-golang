package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		// The string from scanner.Text() includes escaped quotes around text
		codeChars += len(scanner.Text())
		x, err := strconv.Unquote(scanner.Text())
		if err != nil {
			panic(err)
		}
		memoryChars += len(x)
	}

	fmt.Println(codeChars - memoryChars)
}
