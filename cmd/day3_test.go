/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"reflect"
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

func Test_engineParts(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := engineParts(tt.args.line); got != tt.want {
				t.Errorf("engineParts() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPartNumberFromXY(tt.args.x, tt.args.y, tt.args.grid); got != tt.want {
				t.Errorf("getPartNumberFromXY() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validatePartNums(t *testing.T) {
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

	// for _, row := range Grid {
	// 	for _, r := range row {
	// 		fmt.Print(string(r))
	// 	}
	// }
	Parts := findPartNumbers(Grid)
	type args struct {
		parts []partNum
		grid  [][]rune
	}
	tests := []struct {
		name string
		args args
		want []partNum
	}{
		{
			name: "1",
			args: args{
				parts: Parts,
				grid:  Grid,
			},
			want: []partNum{
				{
					x:         0,
					y:         0,
					number:    467,
					hasSymbol: true,
				},
				{
					x:         5,
					y:         0,
					number:    114,
					hasSymbol: false,
				},
				{
					x:         0,
					y:         2,
					number:    35,
					hasSymbol: true,
				},
				{
					x:         6,
					y:         2,
					number:    633,
					hasSymbol: true,
				},
				{
					x:         0,
					y:         4,
					number:    617,
					hasSymbol: true,
				},
				{
					x:         5,
					y:         7,
					number:    58,
					hasSymbol: false,
				},
				{
					x:         2,
					y:         8,
					number:    592,
					hasSymbol: true,
				},
				{
					x:         7,
					y:         8,
					number:    755,
					hasSymbol: true,
				},
				{
					x:         1,
					y:         9,
					number:    664,
					hasSymbol: true,
				},
				{
					x:         5,
					y:         9,
					number:    598,
					hasSymbol: true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validatePartNums(tt.args.parts, tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validatePartNums() = %v, want %v", got, tt.want)
			}
		})
	}
}
