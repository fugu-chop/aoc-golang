package main

import (
	"reflect"
	"testing"
)

func TestCheckCompliance(t *testing.T) {
	data := []struct {
		testCase string
		input    string
		want     bool
	}{
		{"compliant password", "3-4 c: cctc", true},
		{"non-compliant password", "1-3 b: cdefg", false},
	}

	for _, d := range data {
		t.Run(d.testCase, func(t *testing.T) {
			got := checkCompliance(d.input)
			if got != d.want {
				t.Errorf("checkCompliance(), got: %t, want: %t", got, d.want)
			}
		})
	}
}

func TestParseCriteria(t *testing.T) {
	data := []struct {
		testCase string
		input    string
		want     *Criteria
	}{
		{"empty criteria", "", nil},
		{"malformed criteria: too many spaces", "fdha jkf ld", nil},
		{"malformed criteria: too many dashes", "3-7-z f", nil},
		{"malformed criteria: not enough elements", "x", nil},
		{"valid criteria", "3-7 z", &Criteria{
			letter:   "z",
			minCount: 3,
			maxCount: 7,
		}},
	}

	for _, d := range data {
		t.Run(d.testCase, func(t *testing.T) {
			got := parseCriteria(d.input)

			if !reflect.DeepEqual(got, d.want) {
				t.Errorf("parseCriteria(), got: %v, want: %v", got, d.want)
			}
		})
	}
}

func TestMemoisePassword(t *testing.T) {
	data := []struct {
		testCase string
		input    string
		want     map[string]int
	}{
		{"empty password", "", map[string]int{}},
		{"normal password", "zzzxxy", map[string]int{
			"z": 3,
			"x": 2,
			"y": 1,
		}},
		{"handles numbers", "11122", map[string]int{
			"1": 3,
			"2": 2,
		}},
		{"handles spaces", "   111x2    ", map[string]int{
			"1": 3,
			"2": 1,
			"x": 1,
		}},
	}

	for _, d := range data {
		t.Run(d.testCase, func(t *testing.T) {
			got := memoisePassword(d.input)
			if !reflect.DeepEqual(got, d.want) {
				t.Errorf("memoisePassword(), got: %v, want: %v", got, d.want)
			}
		})
	}
}
