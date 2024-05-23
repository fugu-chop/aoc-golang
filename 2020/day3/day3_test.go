package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"testing"
)

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

func Test_generateCoordinate(t *testing.T) {
	fileLocation := "./test_golden_file.txt"

	wantHeight, wantWidth := 5, 11
	wantCoordinate := map[int][]string{
		0: strings.Split("..##.......", ""),
		1: strings.Split("#...#...#..", ""),
		2: strings.Split(".#....#..#.", ""),
		3: strings.Split("..#.#...#.#", ""),
		4: strings.Split(".#...##..#.", ""),
	}

	file, err := os.Open(fileLocation)
	if err != nil {
		setGoldenFile(fileLocation)
		file, err = os.Open(fileLocation)
		if err != nil {
			t.Fatalf("generateCoordinate(): could not open file at %s: %+v", fileLocation, err)
		}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	got := generateCoordinates(scanner)

	if got.height != wantHeight {
		t.Errorf("generateCoordinate(): incorrect height: %d, want: %d", got.height, wantHeight)
	}
	if got.width != wantWidth {
		t.Errorf("generateCoordinate(): incorrect width: %d, want: %d", got.width, wantWidth)
	}
	if !reflect.DeepEqual(got.coordinates, wantCoordinate) {
		t.Errorf("generateCoordinate(): incorrect map generated: %+v, want: %+v", got.coordinates, wantCoordinate)
	}
}

func setGoldenFile(fileLocation string) {
	goldenSample := []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
	}

	file, err := os.OpenFile(fileLocation, os.O_APPEND|os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for _, row := range goldenSample {
		file.WriteString(fmt.Sprintf("%+v\n", row))
	}
}
