package main

import "testing"

func Test_countTrees(t *testing.T) {
	tests := map[string]struct {
		currentRowIdx int
		row           []string
		want          int
	}{
		"correctly iterates":           {2, []string{"", "#", "#"}, 1},
		"skips unoccupied spaces":      {2, []string{"", "#", ""}, 0},
		"handles out of bounds inputs": {4, []string{"", "#", "#"}, 0},
	}

	for test, tc := range tests {
		c := coordinate{}
		t.Run(test, func(t *testing.T) {
			got := c.countTrees(tc.currentRowIdx, tc.row)
			if got != tc.want {
				t.Errorf("countTrees: got: %d, want: %d", got, tc.want)
			}
		})
	}
}

func Test_updateCurrentRowIdx(t *testing.T) {
	tests := map[string]struct {
		currentRowIdx int
		rowLength     int
		want          int
	}{
		"updates when not out of bounds":              {0, 10, 3},
		"updates when out of bounds":                  {8, 10, 1},
		"handles when currentRowIdx is out of bounds": {11, 10, 4},
	}

	for name, tc := range tests {
		c := coordinate{}
		t.Run(name, func(t *testing.T) {
			got := c.updateCurrentRowIdx(tc.currentRowIdx, tc.rowLength)
			if got != tc.want {
				t.Errorf("updateCurrentRowIdx(): got: %d, want: %d", got, tc.want)
			}
		})
	}
}

func Test_calculateTreesHit(t *testing.T) {
	tests := map[string]struct {
		coordinate coordinate
		want       int
	}{
		"calculates trees hit": {
			coordinate: coordinate{
				height: 2,
				width:  4,
				coordinates: map[int][]string{
					0: {".", ".", ".", ".", "."},
					1: {".", ".", ".", "#", "."},
				},
			},
			want: 1,
		},
		"handles extra width": {
			coordinate: coordinate{
				height: 2,
				width:  10,
				coordinates: map[int][]string{
					0: {".", ".", ".", ".", ".", "#"},
					1: {".", ".", ".", "#", ".", "#"},
				},
			},
			want: 1,
		},
		"handles less width": {
			coordinate: coordinate{
				height: 2,
				width:  1,
				coordinates: map[int][]string{
					0: {".", ".", ".", ".", ".", "#"},
					1: {".", ".", ".", "#", ".", "#"},
				},
			},
			want: 0,
		},
		"handles zero height": {
			coordinate: coordinate{
				height: 0,
				width:  5,
				coordinates: map[int][]string{
					0: {".", ".", ".", ".", ".", "#"},
					1: {".", ".", ".", "#", ".", "#"},
				},
			},
			want: 0,
		},
	}

	for name, tc := range tests {
		c := tc.coordinate
		t.Run(name, func(t *testing.T) {
			got := c.calculateTreesHit()
			if got != tc.want {
				t.Errorf("calculateTreesHit(): got: %d, want: %d", got, tc.want)
			}
		})
	}
}
