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

// not sure, try!
type criteria interface {
	validate(interface{}) bool
}

func main() {
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

/*
validPassport is a wrapper function that performs checks on the validity
of a passport entry provided in an input file, returning a boolean that
indicates whether a passport entry is valid per the AoC requirements.
It performs string cleaning and parsing via several helper functions.
*/
func validPassport(passport string) bool {
	passportFields := cleanPassport(passport)
	if len(passportFields) < 7 || len(passportFields) > 8 {
		return false
	}
	// Check if fields all exist in requiredFields
	for _, field := range passportFields {
		if !validField(field) {
			return false
		}
	}
	// Check if all fields in requiredFields exist within passportFields
	cleanedFields := cleanedPassportFields(passportFields)
	for field := range requiredFields {
		if !slices.Contains(cleanedFields, field) {
			return false
		}
	}

	return true
}

/*
cleanedPassportFields takes a slice of strings, where each string takes
the form of "field:value" (delimited by colon) and returns a slice of strings
fields (excluding the values). cleanedPassport expects that each entry in cleanedPassport is already
trimmed of whitespaces.
*/
func cleanedPassportFields(cleanedPassport []string) []string {
	cleanedFields := []string{}

	for _, entry := range cleanedPassport {
		stringSlice := strings.Split(entry, ":")
		if !slices.Contains(cleanedFields, stringSlice[0]) {
			cleanedFields = append(cleanedFields, stringSlice[0])
		}
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
	fieldPair := strings.Split(field, ":")
	if len(fieldPair) != 2 {
		return false
	}

	k, v := fieldPair[0], fieldPair[1]

	if len(v) < 1 {
		return false
	}

	if k == "cid" {
		return true
	}

	if _, ok := requiredFields[k]; !ok {
		return false
	}

	return true
}

// Maybe some interface nonsense can be used?
func validBirthYear(year int) func() bool {
	return func() bool {
		return year >= 1920 && year <= 2002
	}
}

func validIssueYear(year int) func() bool {
	return func() bool {
		return year >= 2010 && year <= 2020
	}
}

func validExpirationYear(year int) func() bool {
	return func() bool {
		return year >= 2020 && year <= 2030
	}
}

func validEyeColour(colour string) func() bool {
	validColours := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	return func() bool {
		return slices.Contains(validColours, colour)
	}
}

// valid hair colour probably needs regex

// passport id - check entirely via regex?

func validHeight(height string) func() bool {
	// parse last two chars, check for in or cm

	return func() bool {
		return true
	}
}
