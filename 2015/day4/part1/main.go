package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

const input = "bgvyzdsv"

func main() {
	for i := 0; i < 1000000; i++ {
		h := md5.New()
		/*
			Goal is to find what the second string should be
			such that the first 5 digits of hexSolution is 00000
		*/
		io.WriteString(h, fmt.Sprintf("%s%d", input, i))

		hexSolution := fmt.Sprintf("%x", h.Sum(nil))
		if string(hexSolution[:5]) == "00000" {
			fmt.Println(hexSolution)
			fmt.Println(i)
			return
		}
	}
}
