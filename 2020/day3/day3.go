package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	horizontalJumps = 3
	verticalJumps   = 1
)

func main() {
	/*
		Problem
			We're given a 'map' as an input (txt file)
			A map represents clear spaces with the `.` character and trees with with the `#` character
			A map is fixed in height (Y axis), but repeats exactly the same in width, to the right (X axis)

			We start at the top left hand corner
			We have a travel pattern, which is 3 spaces right, one space down
			Given a map, we need to calculate how many trees we hit before we exit the bottom limit of the map

		Examples
			The map only contains `.` and `#` characters, no need to handle other characters
			Each 'row' of the map is the same width as every other row
			The horizontal and vertical jumps will always fit within the width and height of the map
			The map will repeat exactly to the right, no deviation/overwriting between maps
			The map can expand to the right as many times as we need, so long as we hit the bottom
			One row's pattern says nothing about the pattern of another row beneath or above it

		Data Structures
			Order of 'cells' in a row is important, so we can represent it using a Slice
			Capture all rows as slice of a slice?
				In terms of speed, reading array by index is fast, writes are slow
				Alternatively, we could use a map, where the key is the index
					Fast writes, fast reads

		Algorithm
			We can read the file, line by line
			As we read the file in, we take the 'line', split the individual chars into a slice
			We append the slice to our map, inserting the value and the index as the key

			Once this is finished, we need to capture:
				the height of the map as a variable to know at which point we 'exit' the map
				the width of a row to know when an OOB error might occur

			We then start iterating over our map, limiting to the height of the map
			We use a variable to capture what 'X' coordinate we have
				On each iteration, we increment X by 3
				Check if this is OOB with our width variable
				Check if the element in the slice is a `.` or a `#`
				Increment a counter variable if it's a `#`.

			What to do if incrementing X by 3 would OOB?
				Using the AoC example, the map is 11 units wide
				This means after 3 iterations (i.e. index 10),
				we would be OOB on the next jump

				We count how many 'free' spots there are until the end of the row
				Subtract from three, add the result to the next row to get our
				new starting index

			Print counter variable
	*/

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	coordinates := map[int][]string{}
	coordinateIdx := 0
	coordinatesHeight := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		coordinates[coordinateIdx] = strings.Split(scanner.Text(), "")
		coordinateIdx += 1
	}

	coordinatesHeight = len(coordinates)

	fmt.Println(coordinatesHeight)
}
