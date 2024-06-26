package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	target     = 2020
	stringNums = []string{}
)

func main() {
	fileLocation := flag.String("inputLocation", "input.txt", "define a file location for the input file")
	testVariant := flag.Int("testPart", 1, "which part should be tested")
	flag.Parse()

	file, err := os.ReadFile(*fileLocation)
	if err != nil {
		log.Fatal(err)
	}

	fileContent := string(file)
	stringNums = strings.Split(fileContent, "\n")
	// Handle newline added to file
	intNums := convertStringToIntSlice(stringNums[:len(stringNums)-1])

	if *testVariant == 1 {
		startIdx, endIdx, found := parseTwoNumbers(intNums)
		if !found {
			log.Fatalf("Could not find numbers that sum to %d", target)
		}

		fmt.Println(intNums[startIdx] * intNums[endIdx])
	}

	if *testVariant == 2 {
		startIdx, midIdx, endIdx, found := parseThreeNumbers(intNums)
		if !found {
			log.Fatalf("Could not find numbers that sum to %d", target)
		}

		fmt.Println(intNums[startIdx] * intNums[midIdx] * intNums[endIdx])
	}
}

func sumToTarget(a ...int) bool {
	sum := 0

	for _, num := range a {
		sum += num
	}

	return sum == target
}

func parseTwoNumbers(numbers []int) (int, int, bool) {
	if len(numbers) < 2 {
		return 0, 0, false
	}

	for startIdx := 0; startIdx < len(numbers)-1; startIdx++ {
		for endIdx := startIdx + 1; endIdx < len(numbers); endIdx++ {
			if sumToTarget(numbers[startIdx], numbers[endIdx]) {
				return startIdx, endIdx, true
			}
		}
	}
	return 0, 0, false
}

func parseThreeNumbers(numbers []int) (int, int, int, bool) {
	if len(numbers) < 3 {
		return 0, 0, 0, false
	}

	for startIdx := 0; startIdx < len(numbers)-2; startIdx++ {
		for midIdx := startIdx + 1; midIdx < len(numbers)-1; midIdx++ {
			for endIdx := startIdx + 2; endIdx < len(numbers); endIdx++ {
				if sumToTarget(numbers[startIdx], numbers[midIdx], numbers[endIdx]) {
					return startIdx, endIdx, midIdx, true
				}
			}
		}
	}
	return 0, 0, 0, false
}

func convertStringToIntSlice(nums []string) []int {
	intSlice := make([]int, len(nums))

	for i, element := range nums {
		if len(element) > 0 {
			num, err := strconv.Atoi(element)
			if err != nil {
				log.Fatal(err)
			}
			intSlice[i] = num
		}
	}

	return intSlice
}
