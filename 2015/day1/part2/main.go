package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.ReadFile("/Users/dean/Documents/aoc-golang/2015/day1/input.txt")
	if err != nil {
		panic(err)
	}

	var floorCount int

	for idx, floor := range f {
		if string(floor) == "(" {
			floorCount += 1
		} else {
			floorCount -= 1
		}

		if floorCount == -1 {
			fmt.Println(idx + 1)
			return
		}
	}
}
