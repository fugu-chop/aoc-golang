package main

import (
	"reflect"
	"testing"
)

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
				t.Errorf("cleanedPassportFields err - got: %+v, want: %+v", got, tc.want)
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
				t.Errorf("cleanPassport() err: got: %s, want: %s", got, tc.want)
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
				t.Errorf("validField() err: got %t, want: %t", got, tc.want)
			}
		})
	}
}
