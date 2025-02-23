package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"math"
)

const (
	input         = "bgvyzdsv"
	leadingZeroes = "000000"
)

func main() {
	for i := 0; i < math.MaxInt; i++ {
		h := md5.New()
		io.WriteString(h, fmt.Sprintf("%s%d", input, i))

		hexSolution := fmt.Sprintf("%x", h.Sum(nil))
		if string(hexSolution[:6]) == leadingZeroes {
			fmt.Println(i)
			return
		}
	}
}
