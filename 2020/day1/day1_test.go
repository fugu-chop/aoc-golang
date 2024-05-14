package main

import (
	"testing"
)

func TestSumToTarget(t *testing.T) {
	data := []struct {
		testCase string
		target   int
		operand1 int
		operand2 int
		want     bool
	}{
		{"target correct", 5, 3, 2, true},
		{"target incorrect", 6, 3, 2, false},
	}

	for _, d := range data {
		t.Run(d.testCase, func(t *testing.T) {
			target = d.target
			got := sumToTarget(d.operand1, d.operand2)
			if got != d.want {
				t.Errorf("sumTo() - got: %t, expected: %t", got, d.want)
			}
		})
	}
}
