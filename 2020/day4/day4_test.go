package main

import (
	"reflect"
	"testing"
)

func Test_validBirthYear(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"valid year":         {"1920", true},
		"valid year max":     {"2002", true},
		"invalid year":       {"239", false},
		"invalid format":     {"xyz", false},
		"invalid characters": {"!@JH*E(M@)", false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			parseFunc := validBirthYear()
			got := parseFunc(tc.input)
			if got != tc.want {
				t.Errorf("validBirthYear %s err - got: %t, want: %t", name, got, tc.want)
			}
		})
	}
}

func Test_validIssueYear(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"valid year min":     {"2010", true},
		"valid year max":     {"2020", true},
		"invalid year":       {"239", false},
		"invalid format":     {"xyz", false},
		"invalid characters": {"!@JH*E(M@)", false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			parseFunc := validIssueYear()
			got := parseFunc(tc.input)
			if got != tc.want {
				t.Errorf("validIssueYear %s err - got: %t, want: %t", name, got, tc.want)
			}
		})
	}
}

func Test_validExpirationYear(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"valid year min":     {"2020", true},
		"valid year max":     {"2030", true},
		"invalid year":       {"239", false},
		"invalid format":     {"xyz", false},
		"invalid characters": {"!@JH*E(M@)", false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			parseFunc := validExpirationYear()
			got := parseFunc(tc.input)
			if got != tc.want {
				t.Errorf("validExpirationYear %s err - got: %t, want: %t", name, got, tc.want)
			}
		})
	}
}

func Test_validEyeColour(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"valid colour":   {"hzl", true},
		"invalid colour": {"red", false},
		"case sensitive": {"oTh", false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			parseFunc := validEyeColour()
			got := parseFunc(tc.input)
			if got != tc.want {
				t.Errorf("validEyeColour %s err - got: %t, want: %t", name, got, tc.want)
			}
		})
	}
}

func Test_validPassportID(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"valid passport ID": {"123456789", true},
		"too short":         {"12345678", false},
		"too long":          {"1234567891", false},
		"other characters":  {"#@!tuvwxy", false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			parseFunc := validPassportID()
			got := parseFunc(tc.input)
			if got != tc.want {
				t.Errorf("validPassportID %s err - got: %t, want: %t", name, got, tc.want)
			}
		})
	}
}

func Test_validHairColour(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"valid hair colour":         {"#123abc", true},
		"valid hair colour (order)": {"#a1b31c", true},
		"case sensitive":            {"#123ABC", false},
		"missing #":                 {"123abc", false},
		"incorrect characters":      {"#zxy890", false},
		"too long":                  {"#12345678", false},
		"too short":                 {"#12", false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			parseFunc := validHairColour()
			got := parseFunc(tc.input)
			if got != tc.want {
				t.Errorf("validHairColour %s err - got: %t, want: %t", name, got, tc.want)
			}
		})
	}
}

func Test_validHeight(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"inches":         {"70in", true},
		"centimetres":    {"175cm", true},
		"invalid unit":   {"125m", false},
		"invalid input":  {"a!bcd", false},
		"invalid cm":     {"2cm", false},
		"invalid inches": {"1000in", false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			parseFunc := validHeight()
			got := parseFunc(tc.input)
			if got != tc.want {
				t.Errorf("validHeight %s err - got: %t, want: %t", name, got, tc.want)
			}
		})
	}
}

func Test_validPassport(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"handles valid passport": {"ecl:#eef340 eyr:2023 hcl:#c0946f pid:244684338 iyr:2020 byr:1969 hgt:152cm", true},
		"handles unusual spacing": {`pid:303807545 cid:213 ecl:gry hcl:#fffffd
eyr:2038 byr:1951
hgt:171cm iyr:2011`, true},
		"ignores CID field": {"ecl:#eef340 eyr:2023 hcl:#c0946f pid:244684338 iyr:2020 cid:57 byr:1969 hgt:152cm", true},
		"handles invalid passports": {`iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929`, false},
		"short circuits short invalid passports": {`eyr:2023 pid:028048884`, false},
		"short circuits long invalid passports":  {`eyr:2023 pid:028048884 eyr:2023 eyr:2023 eyr:2023 eyr:2023 eyr:2023 eyr:2023 eyr:2023`, false},
		"short circuits invalid fields":          {"ecc:#eef340 eyr:2023 hcl:#c0946f pid:244684338 iyr:2020 byr:1969 hgt:152cm", false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := validPassport(tc.input)
			if got != tc.want {
				t.Errorf("validPassport %s err - got: %t, want: %t", name, got, tc.want)
			}
		})
	}
}

func Test_cleanedPassportFields(t *testing.T) {
	tests := map[string]struct {
		input []string
		want  []string
	}{
		"filters out values": {
			[]string{"field:value", "field1:value1", "field2:value2"},
			[]string{"field", "field1", "field2"},
		},
		"ignores duplicates": {
			[]string{"field:value", "field1:value1", "field1:value1"},
			[]string{"field", "field1"},
		},
		"handles fields with no values": {
			[]string{"field", "field1", "field2"},
			[]string{"field", "field1", "field2"},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := cleanedPassportFields(tc.input)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("cleanedPassportFields %s err - got: %+v, want: %+v", name, got, tc.want)
			}
		})
	}
}

func Test_cleanPassport(t *testing.T) {
	tests := map[string]struct {
		input string
		want  []string
	}{
		"separated by spaces only": {"ecl:gry pid:860033327 eyr:2020",
			[]string{"ecl:gry", "pid:860033327", "eyr:2020"}},
		"separated by newlines only": {"\necl:gry\npid:860033327\neyr:2020\n",
			[]string{"ecl:gry", "pid:860033327", "eyr:2020"}},
		"separated by both newlines and spaces": {"ecl:gry\npid:860033327 eyr:2020 ",
			[]string{"ecl:gry", "pid:860033327", "eyr:2020"}},
		"handles trailing and leading whitepaces, newlines and tabs": {"\tecl:gry\npid:860033327 eyr:2020 ",
			[]string{"ecl:gry", "pid:860033327", "eyr:2020"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := cleanPassport(tc.input)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("cleanPassport() %s err: got: %s, want: %s", name, got, tc.want)
			}
		})
	}
}

func Test_validField(t *testing.T) {
	tests := map[string]struct {
		want  bool
		field string
	}{
		"valid field":                          {true, "byr:1992"},
		"valid field - contains CID":           {true, "cid:hello"},
		"invalid field - no delimiter":         {false, "asdf"},
		"invalid field - no value":             {false, "hgt:"},
		"invalid field - field does not match": {false, "asdf:123"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := validField(tc.field)

			if got != tc.want {
				t.Errorf("validField() %s err: got %t, want: %t", name, got, tc.want)
			}
		})
	}
}
