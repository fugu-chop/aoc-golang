package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var (
	fileLocation = "./input.txt"
)

func main() {
	file, err := os.Open(fileLocation)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
