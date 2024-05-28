package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

var (
	fileLocation   = "./input.txt"
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
			- If all fields are present, CONTINUE
		- Also need to iterate over entries in requiredFields
			- Handle basic scenario where we might have 6 fields that match
		- increment validPassports
	*/

	// Ideally we would use scanner, but the presence of a newline at
	// end of the file breaks the iteration (will skip the last entry)
	file, err := os.ReadFile(fileLocation)
	if err != nil {
		log.Fatal(err)
	}
	stringFile := string(file)
	// Break up passports into units
	passportList := strings.Split(stringFile, "\n\n")
	validPassports := 0

	for _, entry := range passportList {
		if validPassport(entry) {
			validPassports++
		}
	}

	fmt.Println(validPassports)
}

// TODO - add tests
/*
validPassport is a wrapper function that performs checks on the validity
of a passport entry provided in an input file, returning a boolean that
indicates whether a passport entry is valid per the AoC requirements.
It performs string cleaning and parsing via several helper functions.
*/
func validPassport(passport string) bool {
	var valid bool
	passportFields := cleanPassport(passport)
	if len(passportFields) < 7 || len(passportFields) > 8 {
		return valid
	}
	// Check if fields all exist in requiredFields
	for _, field := range passportFields {
		if !validField(field) {
			fmt.Println("invalid field: " + field)
			return valid
		}
	}
	// Check if all fields in requiredFields exist within passportFields
	cleanedFields := cleanedPassportFields(passportFields)
	for field := range requiredFields {
		if !slices.Contains(cleanedFields, field) {
			return valid
		}
	}

	valid = true

	return valid
}

// TODO - add tests
/*
cleanedPassportFields takes a slice of strings, where each string takes
the form of "field:value" (delimited by colon) and returns a slice of strings
fields. cleanedPassport expects that each entry in cleanedPassport is already
trimmed of whitespaces.
*/
func cleanedPassportFields(cleanedPassport []string) []string {
	cleanedFields := []string{}

	for _, entry := range cleanedPassport {
		stringSlice := strings.Split(entry, ":")
		cleanedFields = append(cleanedFields, stringSlice[0])
	}

	return cleanedFields
}

/*
cleanPassport takes in a string of unformatted field-value pairs where
each field-value pair is delimited by a space or newline between the next
field-value pair. It returns a slice of field-value pairs with any leading
or trailing whitespaces (including tabs) removed.
*/
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

/*
validField checks for the validity of field-value pair by looking at
whether the field exists within the accepted passport fields and if
the value is a non-empty string. It skips the check for the CID field
as this is a non-mandatory field.
*/
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

	val, ok := requiredFields[k]

	if !ok || !val {
		return valid
	}

	valid = true

	return valid
}
