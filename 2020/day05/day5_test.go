package main

import (
	"testing"
)

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

func Test_createSeatPosition(t *testing.T) {
	tests := map[string]struct {
		row, col, seatID int
		want             *seatPosition
	}{
		"returns new object correctly": {
			row:    44,
			col:    5,
			seatID: 357,
			want: &seatPosition{
				row:    44,
				column: 5,
				seatID: 357,
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := createSeatPosition(tc.row, tc.col)
			if got.column != tc.col || got.row != tc.row || got.seatID != tc.seatID {
				t.Errorf("createSeatPosition %s err: got: %+v, want: %+v", name, got, tc.want)
			}
		})
	}
}

func Test_findMissingSeat(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"sorts and finds missing seat": {
			input: []int{717, 719, 716, 720},
			want:  718,
		},
	}

	for name, tc := range tests {
		got := findMissingSeat(tc.input)
		if got != tc.want {
			t.Errorf("findMissingSeat %s err: got: %d, want: %d", name, got, tc.want)
		}
	}
}
