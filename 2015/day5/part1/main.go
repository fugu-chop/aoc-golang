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
		if !hasForbiddenCombo(scanner.Text()) &&
			hasConsecutiveLetters(scanner.Text()) &&
			hasThreeVowels(scanner.Text()) {
			goodWords += 1
		}
	}

	fmt.Println(goodWords)
}

func hasConsecutiveLetters(word string) bool {
	wordSlice := strings.Split(word, "")

	slidingWindow := []string{wordSlice[0]}
	for _, letter := range wordSlice[1:] {
		slidingWindow = append(slidingWindow, letter)
		if slidingWindow[0] == slidingWindow[1] {
			return true
		}
		slidingWindow = []string{letter}
	}

	return false
}

func hasForbiddenCombo(word string) bool {
	forbiddenCombos := []string{"ab", "cd", "pq", "xy"}
	for _, combo := range forbiddenCombos {
		if strings.Contains(word, combo) {
			return true
		}
	}
	return false
}

func hasThreeVowels(word string) bool {
	vowels := "aeiou"
	containedVowels := []string{}

	wordSlice := strings.Split(word, "")
	for _, letter := range wordSlice {
		if strings.Contains(vowels, letter) {
			containedVowels = append(containedVowels, letter)
			if len(containedVowels) >= 3 {
				return true
			}
		}
	}

	return false
}
