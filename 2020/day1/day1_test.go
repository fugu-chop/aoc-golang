package main

import (
	"reflect"
	"testing"
)

func TestSumToTarget(t *testing.T) {
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
