package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type seatPosition struct {
	row    int
	column int
	seatID int
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
