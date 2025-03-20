package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "3113322113"

	for i := 0; i < 40; i++ {
		input = lookAndSay(input)
	}

	fmt.Println(len(input))
}

func lookAndSay(input string) string {
	var (
		count       int
		currentChar string
		results     []string
	)

	chars := strings.Split(input, "")

	for i := 0; i < len(chars); i++ {
		currentChar = chars[i]
		count++

		if i < len(chars)-1 {
			if chars[i+1] != currentChar {
				results = append(results, strconv.Itoa(count))
				results = append(results, currentChar)
				count = 0
			}
		}
		if i == len(chars) {
			if chars[i-1] == currentChar {
				count++
			}
		}
	}

	results = append(results, strconv.Itoa(count))
	results = append(results, currentChar)

	return strings.Join(results, "")
}
