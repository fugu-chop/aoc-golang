package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var (
	fileLocation   = "./input.txt"
	requiredFields = map[string]func(string) bool{
		"byr": validBirthYear(),
		"iyr": validIssueYear(),
		"eyr": validExpirationYear(),
		"hgt": validHeight(),
		"hcl": validHairColour(),
		"ecl": validEyeColour(),
		"pid": validPassportID(),
	}
)

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

	parseFunc, ok := requiredFields[k]
	if ok && parseFunc(v) {
		return true
	}

	return false
}

/*
validBirthYear checks if the value provided to the byr field is valid
A valid birth year is between 1920 and 2002.
*/
func validBirthYear() func(string) bool {
	return func(year string) bool {
		yearInt, err := strconv.Atoi(year)
		if err != nil {
			return false
		}

		return yearInt >= 1920 && yearInt <= 2002
	}
}

/*
validIssueYear checks if the value provided to the iyr field is valid.
A valid issue year is between 2010 and 2020.
*/
func validIssueYear() func(string) bool {
	return func(year string) bool {
		yearInt, err := strconv.Atoi(year)
		if err != nil {
			return false
		}
		return yearInt >= 2010 && yearInt <= 2020
	}
}

/*
validExpirationYear checks if the value provided to the eyr field is valid
A valid expiration year is between 2020 and 2030.
*/
func validExpirationYear() func(string) bool {
	return func(year string) bool {
		yearInt, err := strconv.Atoi(year)
		if err != nil {
			return false
		}
		return yearInt >= 2020 && yearInt <= 2030
	}
}

/*
validEyeColour checks if the value provided to the ecl field is valid.
A valid eye colour is exactly one of: amb blu brn gry grn hzl oth.
*/
func validEyeColour() func(string) bool {
	return func(colour string) bool {
		validColours := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
		return slices.Contains(validColours, colour)
	}
}

/*
validPassportID checks if the value provided to the pid field is valid.
A valid passport ID is a nine-digit number, including leading zeroes.
*/
func validPassportID() func(string) bool {
	return func(id string) bool {
		re := regexp.MustCompile(`^\d{9}$`)
		return re.Match([]byte(id))
	}
}

/*
validHairColour checks if the value provided to the hcl field is valid
A valid hair colour is a # followed by exactly six characters 0-9 or a-f.
*/
func validHairColour() func(string) bool {
	return func(colour string) bool {
		re := regexp.MustCompile(`^#(\d|[a-f]){6}$`)
		return re.Match([]byte(colour))
	}
}

/*
validHeight checks if the value provided to the hgt field is valid
A height is valid if a number is followed by either cm or in:

	If cm, the number must be at least 150 and at most 193.
	If in, the number must be at least 59 and at most 76.
*/
func validHeight() func(string) bool {
	return func(value string) bool {
		unitPosition := len(value) - 2
		unit := string([]byte(value)[unitPosition:])
		heightStr := string([]byte(value)[:unitPosition])
		height, err := strconv.Atoi(heightStr)
		if err != nil {
			return false
		}

		switch unit {
		case "cm":
			return height >= 150 && height <= 193
		case "in":
			return height >= 59 && height <= 76
		default:
			return false
		}
	}
}
