package main

import (
	"testing"
)

func TestSumTo(t *testing.T) {
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
		target = d.target
		got := sumTo(d.operand1, d.operand2)
		if got != d.want {
			t.Errorf("sumTo() - got: %t, expected: %t", got, d.want)
		}
	}
}
