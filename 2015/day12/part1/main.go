package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var sum int

	f, err := os.ReadFile("/Users/dean/Documents/aoc-golang/2015/day12/input.txt")
	if err != nil {
		panic(err)
	}
	// You will not encounter any strings containing numbers.
	re := regexp.MustCompile(`-?\d+`)
	for _, sliceB := range re.FindAll(f, -1) {
		var strSlice []string
		for _, b := range sliceB {
			strSlice = append(strSlice, string(b))
		}
		strVal := strings.Join(strSlice, "")
		i, err := strconv.Atoi(strVal)
		if err != nil {
			panic(err)
		}

		sum += i
	}

	fmt.Println(sum)
}
