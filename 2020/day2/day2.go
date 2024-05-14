package main

import "strings"

func main() {
	/*
		Problem
			get an input of passwords and policies in the format
			3-11 z: zzzzzdzzzzlzz

			Format is x-y z where:
				x is the lower bound (must contain at least)
				y is the upper bound (must contain no more than)
				z is the letter of the alphabet

			Password must meet both criteria to be compliant.

			Our job is to return an int, representing the total number of compliant passwords

		Edge Cases
			- Based on manually reading the input file, no uppercase
			- will always have a lower and upper bound represented as "integer"

		Data Structures
			- Throwaway slices for string splitting
			- Map for the password calculation
			- int Var for tracking compliant passwords
			- Custom type for criteria is easiest

		Algo
			We will need to parse an input file, row by row, ideally (use scanner)
			For each line (iterate):
				Split on `:` character, this will give a slice of two elements, a 'guide' and pw

				Password:
					- Trim whitespace
					- split by letters, memoise into a map with letters
					- key of letter, values of counts

				Criteria:
					- Split by space, maybe turn into a type
					- we can then easily compare a Criteria with the map of letter counts?
					- E.g. letterCount[Criteria.letter] > Criteria.min && > Criteria.max
					- if yes, increment our int Var by one

					Return intVar
	*/
}

func memoisePassword(password string) map[string]int {
	counts := map[string]int{}
	letterSlice := strings.Split(password, "")

	for _, letter := range letterSlice {
		counts[letter] += 1
	}

	return counts
}
