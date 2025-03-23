package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	var (
		valid bool
		input = "vzbxxyzz"
	)

	for !valid {
		input = increment(input)
		if !hasForbiddenLetters(input) &&
			hasTwoDoubles(input) &&
			hasIncreasingStraights(input) {
			valid = true
			fmt.Println(input)
		}
	}
}

func increment(password string) string {
	passwordSlice := strings.Split(password, "")
	slices.Reverse(passwordSlice)

	for i, letter := range passwordSlice {
		if strings.ContainsAny(letter, "hko") {
			passwordSlice[i] = string([]byte(letter)[0] + 2)
			break
		}
		if letter == "z" {
			passwordSlice[i] = "a"
			continue
		}

		passwordSlice[i] = string([]byte(letter)[0] + 1)
		break
	}

	slices.Reverse(passwordSlice)
	return strings.Join(passwordSlice, "")
}

func hasForbiddenLetters(password string) bool {
	return strings.ContainsAny(password, "iol")
}

func hasTwoDoubles(password string) bool {
	var doubles int

	passwordSlice := strings.Split(password, "")

	for i := 1; i < len(passwordSlice); i++ {
		if passwordSlice[i] == passwordSlice[i-1] {
			doubles++
			i++
		}
		if doubles == 2 {
			return true
		}
	}

	return false
}

// Does zab count?
func hasIncreasingStraights(password string) bool {
	passwordSlice := strings.Split(password, "")
	for i := 0; i < len(passwordSlice)-2; i++ {
		if []byte(passwordSlice[i])[0] == []byte(passwordSlice[i+1])[0]-1 &&
			[]byte(passwordSlice[i])[0] == []byte(passwordSlice[i+2])[0]-2 {
			return true
		}
	}

	return false
}
