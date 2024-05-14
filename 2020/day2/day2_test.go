package main

import (
	"reflect"
	"testing"
)

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
