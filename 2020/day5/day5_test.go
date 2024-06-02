package main

import "testing"

func Test_recalculateUpperBound(t *testing.T) {
	tests := map[string]struct {
		input *boundary
		want  *boundary
	}{}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := recalculateUpperBound(test.input)
			if got != test.want {
				t.Errorf("recalculateUpperBound %s err: got: %+v, want: %+v", name, got, test.want)
			}
		})
	}
}

func Test_recalculateLowerBound(t *testing.T) {
	tests := map[string]struct {
		input *boundary
		want  *boundary
	}{}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := recalculateLowerBound(test.input)
			if got != test.want {
				t.Errorf("recalculateLowerBound %s err: got: %+v, want: %+v", name, got, test.want)
			}
		})
	}
}
