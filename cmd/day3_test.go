/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"
	"testing"
)

var (
	testData = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
	Grid [][]rune
)

func Test_getPartNumberFromXY(t *testing.T) {
	lines := strings.Split(testData, "\n")

	height := len(lines)
	width := len(lines[0]) // Assuming all lines have the same length

	Grid := make([][]rune, height)
	for i := range Grid {
		Grid[i] = make([]rune, width)
	}

	for y, line := range lines {
		for x, r := range line {
			Grid[y][x] = r
		}
	}

	for _, row := range Grid {
		for _, r := range row {
			fmt.Print(string(r))
		}
		fmt.Println()
	}
	type args struct {
		x    int
		y    int
		grid [][]rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				x:    0,
				y:    0,
				grid: Grid,
			},
			want: 467,
		},
		{
			name: "2",
			args: args{
				x:    5,
				y:    0,
				grid: Grid,
			},
			want: 114,
		},
		{
			name: "3",
			args: args{
				x:    2,
				y:    2,
				grid: Grid,
			},
			want: 35,
		},
		{
			name: "4",
			args: args{
				x:    6,
				y:    2,
				grid: Grid,
			},
			want: 633,
		},
		{
			name: "5",
			args: args{
				x:    0,
				y:    4,
				grid: Grid,
			},
			want: 617,
		},
		{
			name: "6",
			args: args{
				x:    7,
				y:    5,
				grid: Grid,
			},
			want: 58,
		},
		{
			name: "7",
			args: args{
				x:    2,
				y:    6,
				grid: Grid,
			},
			want: 592,
		},
		{
			name: "8",
			args: args{
				x:    6,
				y:    7,
				grid: Grid,
			},
			want: 755,
		},
		{
			name: "9",
			args: args{
				x:    1,
				y:    9,
				grid: Grid,
			},
			want: 664,
		},
		{
			name: "10",
			args: args{
				x:    5,
				y:    9,
				grid: Grid,
			},
			want: 598,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPartNumberFromXY(tt.args.x, tt.args.y, tt.args.grid); got != tt.want {
				t.Errorf("getPartNumberFromXY() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_validatePartNums(t *testing.T) {
// 	lines := strings.Split(testData, "\n")

// 	height := len(lines)
// 	width := len(lines[0]) // Assuming all lines have the same length

// 	Grid := make([][]rune, height)
// 	for i := range Grid {
// 		Grid[i] = make([]rune, width)
// 	}

// 	for y, line := range lines {
// 		for x, r := range line {
// 			Grid[y][x] = r
// 		}
// 	}

// 	Parts := findPartNumbers(Grid)
// 	type args struct {
// 		parts []partNum
// 		grid  [][]rune
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want []partNum
// 	}{
// 		{
// 			name: "1",
// 			args: args{
// 				parts: Parts,
// 				grid:  Grid,
// 			},
// 			want: []partNum{
// 				{
// 					x:         0,
// 					y:         0,
// 					number:    467,
// 					hasSymbol: true,
// 				},
// 				{
// 					x:         5,
// 					y:         0,
// 					number:    114,
// 					hasSymbol: false,
// 				},
// 				{
// 					x:         2,
// 					y:         2,
// 					number:    35,
// 					hasSymbol: true,
// 				},
// 				{
// 					x:         6,
// 					y:         2,
// 					number:    633,
// 					hasSymbol: true,
// 				},
// 				{
// 					x:         0,
// 					y:         4,
// 					number:    617,
// 					hasSymbol: true,
// 				},
// 				{
// 					x:         7,
// 					y:         5,
// 					number:    58,
// 					hasSymbol: false,
// 				},
// 				{
// 					x:         2,
// 					y:         6,
// 					number:    592,
// 					hasSymbol: true,
// 				},
// 				{
// 					x:         6,
// 					y:         7,
// 					number:    755,
// 					hasSymbol: true,
// 				},
// 				{
// 					x:         1,
// 					y:         9,
// 					number:    664,
// 					hasSymbol: true,
// 				},
// 				{
// 					x:         5,
// 					y:         9,
// 					number:    598,
// 					hasSymbol: true,
// 				},
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := validatePartNums(tt.args.parts, tt.args.grid); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("validatePartNums() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func Test_partNumHasSymbol(t *testing.T) {
	lines := strings.Split(testData, "\n")

	height := len(lines)
	width := len(lines[0]) // Assuming all lines have the same length

	Grid := make([][]rune, height)
	for i := range Grid {
		Grid[i] = make([]rune, width)
	}

	for y, line := range lines {
		for x, r := range line {
			Grid[y][x] = r
		}
	}

	// Parts := findPartNumbers(Grid)
	type args struct {
		x      int
		y      int
		number int
		grid   [][]rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				x:      0,
				y:      0,
				number: 467,
				grid:   Grid,
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				x:      5,
				y:      0,
				number: 114,
				grid:   Grid,
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				x:      2,
				y:      2,
				number: 35,
				grid:   Grid,
			},
			want: true,
		},
		{
			name: "4",
			args: args{
				x:      6,
				y:      2,
				number: 633,
				grid:   Grid,
			},
			want: true,
		},
		{
			name: "5",
			args: args{
				x:      0,
				y:      4,
				number: 617,
				grid:   Grid,
			},
			want: true,
		},
		{
			name: "6",
			args: args{
				x:      7,
				y:      5,
				number: 58,
				grid:   Grid,
			},
			want: false,
		},
		{
			name: "7",
			args: args{
				x:      2,
				y:      6,
				number: 592,
				grid:   Grid,
			},
			want: true,
		},
		{
			name: "8",
			args: args{
				x:      6,
				y:      7,
				number: 755,
				grid:   Grid,
			},
			want: true,
		},
		{
			name: "9",
			args: args{
				x:      1,
				y:      9,
				number: 664,
				grid:   Grid,
			},
			want: true,
		},
		{
			name: "10",
			args: args{
				x:      5,
				y:      9,
				number: 598,
				grid:   Grid,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partNumHasSymbol(tt.args.x, tt.args.y, tt.args.number, tt.args.grid); got != tt.want {
				t.Errorf("partNumHasSymbol() = %v, want %v", got, tt.want)
			}
		})
	}
}
