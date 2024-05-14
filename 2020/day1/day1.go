package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	target = 2020
)

func main() {
	fileLocation := flag.String("inputLocation", "input.txt", "define a file location for thje input file")
	flag.Parse()

	file, err := os.Open(*fileLocation)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func sumTo(a, b int) bool {
	return a+b == target
}
