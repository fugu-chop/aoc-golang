package main

import (
	"reflect"
	"testing"
)

func TestCheckCountsCompliance(t *testing.T) {
	tests := []struct {
		testCase string
		input    string
		want     bool
	}{
		{"compliant password", "3-4 c: cctc", true},
		{"non-compliant password", "1-3 b: cdefg", false},
	}

	for _, tc := range tests {
		t.Run(tc.testCase, func(t *testing.T) {
			got := checkCountsCompliance(tc.input)
			if got != tc.want {
				t.Errorf("checkCompliance(), got: %t, want: %t", got, tc.want)
			}
		})
	}
}

func TestParseCriteria(t *testing.T) {
	tests := []struct {
		testCase string
		input    string
		want     *CountCriteria
	}{
		{"empty criteria", "", nil},
		{"malformed criteria: too many spaces", "fdha jkf ld", nil},
		{"malformed criteria: too many dashes", "3-7-z f", nil},
		{"malformed criteria: not enough elements", "x", nil},
		{"valid criteria", "3-7 z", &CountCriteria{
			letter:   "z",
			minCount: 3,
			maxCount: 7,
		}},
	}

	for _, tc := range tests {
		t.Run(tc.testCase, func(t *testing.T) {
			got := parseCountCriteria(tc.input)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("parseCriteria(), got: %v, want: %v", got, tc.want)
			}
		})
	}
}

func TestMemoisePassword(t *testing.T) {
	tests := []struct {
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

	for _, tc := range tests {
		t.Run(tc.testCase, func(t *testing.T) {
			got := memoisePassword(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("memoisePassword(), got: %v, want: %v", got, tc.want)
			}
		})
	}
}
