package main

import (
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
