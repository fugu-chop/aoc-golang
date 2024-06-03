package main

import "testing"

func Test_recalculateUpperBound(t *testing.T) {
	boundaryObj := &boundary{
		lower: 0,
		upper: 127,
	}

	tests := map[string]struct {
		input *boundary
		want  int
	}{
		"recalculates upper boundary": {
			input: boundaryObj,
			want:  63,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := recalculateUpperBound(test.input)
			if got.upper != test.want {
				t.Errorf("recalculateUpperBound %s err: got: %+v, want: %+v", name, got, test.want)
			}
		})
	}
}

func Test_recalculateLowerBound(t *testing.T) {
	boundaryObj := &boundary{
		lower: 0,
		upper: 127,
	}
	tests := map[string]struct {
		input *boundary
		want  int
	}{
		"recalculates lower boundary": {
			input: boundaryObj,
			want:  64,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := recalculateLowerBound(boundaryObj)
			if got.lower != test.want {
				t.Errorf("recalculateLowerBound %s err: got: %+v, want: %+v", name, got.lower, test.want)
			}
		})
	}
}
