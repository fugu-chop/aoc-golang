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
		t.Run(test, func(t *testing.T) {
			got := countTrees(tc.currentRowIdx, tc.row)
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
		"updates when not out of bounds": {0, 10, 3},
		"updates when out of bounds":     {8, 10, 1},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := updateCurrentRowIdx(tc.currentRowIdx, tc.rowLength)
			if got != tc.want {
				t.Errorf("updateCurrentRowIdx(): got: %d, want: %d", got, tc.want)
			}
		})
	}
}