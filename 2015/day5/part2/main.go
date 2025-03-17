package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("/Users/dean/Documents/aoc-golang/2015/day5/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var goodWords int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		word := scanner.Text()
		if doublePair(word) && letterBetween(word) {
			goodWords += 1
		}
	}

	fmt.Println(goodWords)
}

func doublePair(word string) bool {
	letterSlice := strings.Split(word, "")
	previousPair := strings.Join(letterSlice[0:2], "")
	counts := map[string]int{
		previousPair: 1,
	}
	pair := []string{letterSlice[1]}

	for _, letter := range letterSlice[2:] {
		pair = append(pair, letter)

		joinedPair := strings.Join(pair, "")

		if previousPair != joinedPair {
			counts[joinedPair]++
		}

		if counts[joinedPair] == 2 {
			return true
		}

		previousPair = joinedPair
		pair = []string{pair[1]}
	}

	return false
}

func letterBetween(word string) bool {
	letterSlice := strings.Split(word, "")
	if len(letterSlice) < 3 {
		return false
	}

	combo := letterSlice[0:3]
	if combo[0] == combo[2] {
		return true
	}

	for _, letter := range letterSlice[3:] {
		combo = combo[1:3]
		combo = append(combo, letter)
		if combo[0] == combo[2] {
			return true
		}
	}

	return false
}
