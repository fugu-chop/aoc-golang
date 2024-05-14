package main

import (
	"reflect"
	"testing"
)

func TestSumToTarget(t *testing.T) {
	defer resetTarget()

	data := []struct {
		testCase string
		target   int
		operands []int
		want     bool
	}{
		{"target correct, 2 nums", 5, []int{3, 2}, true},
		{"target correct, 5 nums", 10, []int{3, 2, 2, 2, 1}, true},
		{"target incorrect", 6, []int{3, 2}, false},
	}

	for _, d := range data {
		t.Run(d.testCase, func(t *testing.T) {
			target = d.target
			got := sumToTarget(d.operands...)
			if got != d.want {
				t.Errorf("sumTo() - got: %t, expected: %t", got, d.want)
			}
		})
	}
}

func TestConvertStringToIntSlice(t *testing.T) {
	data := []struct {
		testCase   string
		inputSlice []string
		length     int
		want       []int
	}{
		{"empty slice", []string{}, 0, []int{}},
		{"fully populated slice", []string{"1", "2", "3"}, 3, []int{1, 2, 3}},
		{"slice with empty strings", []string{"1", "", "3"}, 3, []int{1, 0, 3}},
	}

	for _, d := range data {
		t.Run(d.testCase, func(t *testing.T) {
			got := convertStringToIntSlice(d.inputSlice)
			if !reflect.DeepEqual(got, d.want) {
				t.Errorf("convertStringToIntSlice(), slice does not match, got: %v, want %v", got, d.want)
			}
			if len(got) != d.length {
				t.Errorf("convertStringToIntSlice(), length of slice, got: %d, want %d", len(got), d.length)
			}
		})
	}
}

func TestParseTwoNumbers(t *testing.T) {
	data := []struct {
		testCase     string
		inputSlice   []int
		want1, want2 int
		found        bool
	}{
		{"empty slice", []int{}, 0, 0, false},
		{"populated slice with a match", []int{2, 1000, 2, 5, 1020}, 1, 5, true},
		{"populated slice, but no match", []int{5, 2}, 0, 0, false},
	}

	for _, d := range data {
		t.Run(d.testCase, func(t *testing.T) {
			got1, got2, gotFound := parseTwoNumbers(d.inputSlice)
			if got1 != d.want1 && got2 != d.want2 && gotFound != d.found {
				t.Errorf("parseTwoNumbers(): got: %d, %d, %t; wanted: %d, %d, %t", got1, got2, gotFound, d.want1, d.want2, d.found)
			}
		})
	}
}

func resetTarget() {
	target = 2020
}
