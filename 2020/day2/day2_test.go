package main

import (
	"reflect"
	"testing"
)

func TestParseCriteria(t *testing.T) {
	data := []struct {
		testCase string
		input    string
		want     *Criteria
	}{
		{"empty criteria", "", nil},
		{"malformed criteria: too many spaces", "fdha jkf ld", nil},
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
