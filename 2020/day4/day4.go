package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	fileLocation   = "./example.txt"
	requiredFields = map[string]bool{
		"byr": true,
		"iyr": true,
		"eyr": true,
		"hgt": true,
		"hcl": true,
		"ecl": true,
		"pid": true,
	}
)

func main() {
	/*
		Problem
			We're given a batch file (.txt format) that contains 'passports'
			We need to count the number of valid passports

			A valid passport needs to contain the following fields:
			- byr (Birth Year)
			- iyr (Issue Year)
			- eyr (Expiration Year)
			- hgt (Height)
			- hcl (Hair Color)
			- ecl (Eye Color)
			- pid (Passport ID)

			A passport that is missing ANY of these fields is INVALID.

			A final field, CID, is OPTIONAL
			- cid (Country ID)

		Examples, Edge Cases
			- See problem description
			- It APPEARS that we need a value for each field
			- There does not appear to be a requirement for a type for each field
				- As long as the field is populated, it's fine

			Passport Structure
			- A passport counts where there is a newline AT THE END of the entries (includes last entry)
			- Entries for a passport are separated by spaces or a newline
			- There are NO TRAILING SPACES for passport entries

		Data Structures
			- map of passport fields with `true` as values (values don't really matter) for faster reads
			- We have a `.txt` file, with uneven formatting (see `Passport Structures`)
				- Given the uneven formatting, it may make sense to collect all of the fields
				in a passport in a SLICE before converting to a map?

		Algorithm
		VALIDITY CHECKS
		- Split line into slice based on spaces
		- Can shortcut NOT VALID if the number of entries in passportBuffer is less than 7
			- If longer than 8, NOT VALID
				*assumption that extra fields are not valid*
		- Otherwise, iterate over entries in passport buffer
			- Split via `:` character into two element slice (`splitSlice`)
			- If only one element or length of second element == 0, NOT VALID
					*this is an assumption based on file contents*
			- Take first element of `splitSlice`, check if exists in our map
				- If not, NOT VALID
			- If all fields are present, VALID
				- increment validPassports
	*/

	// Ideally we would use scanner, but the presence of a newline at
	// end of the file breaks the iteration (will skip the last entry)
	file, err := os.ReadFile(fileLocation)
	if err != nil {
		log.Fatal(err)
	}
	stringFile := string(file)
	// Correctly break up passports into units
	passportList := strings.Split(stringFile, "\n\n")
	validPassports := 0
	for _, entry := range passportList {
		if validatePassport(entry) {
			fmt.Println(entry + "\n")
			validPassports++
		}
	}

	fmt.Println(validPassports)
}

/*
ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm
*/
func validatePassport(passport string) bool {
	var valid bool
	passportFields := cleanPassport(passport)
	if len(passportFields) < 7 || len(passportFields) > 8 {
		return valid

	}
	for _, field := range passportFields {

		// ALSO NEED TO CHECK NEGATIVE SCENARIO
		// IF THE FIELD IS MISSING, WE WILL NEVER ITERATE OVER IT

		if !validField(field) {
			return valid
		}
	}

	valid = true

	return valid
}

func cleanPassport(passport string) []string {
	cleanedPassports := []string{}

	passport = strings.TrimSpace(passport)
	trimmedSpacePassport := strings.Split(passport, " ")

	for _, entry := range trimmedSpacePassport {
		cleanedEntry := strings.Split(entry, "\n")
		cleanedPassports = append(cleanedPassports, cleanedEntry...)
	}

	return cleanedPassports
}

// Flip the logic - iterate over the fields in requiredFields
func validField(field string) bool {
	var valid bool

	fieldPair := strings.Split(field, ":")
	if len(fieldPair) != 2 {
		return valid
	}

	k, v := fieldPair[0], fieldPair[1]

	if len(v) < 1 {
		return valid
	}

	if k == "cid" {
		return true
	}

	// iterate over requiredFields
	// maybe change to a slice comparison
	val, ok := requiredFields[k]

	if !ok || !val {
		return valid
	}

	valid = true

	return valid
}
