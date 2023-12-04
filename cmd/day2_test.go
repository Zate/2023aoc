/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import "testing"

func Test_processGame(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
		{
			name: "Game 1",
			args: args{
				line: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			},
			want: 1,
		},
		// Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
		{
			name: "Game 2",
			args: args{
				line: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			},
			want: 2,
		},
		// Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
		{
			name: "Game 3",
			args: args{
				line: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			},
			want: 0,
		},
		// Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
		{
			name: "Game 4",
			args: args{
				line: "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			},
			want: 0,
		},
		// Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
		{
			name: "Game 5",
			args: args{
				line: "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processGame(tt.args.line); got != tt.want {
				t.Errorf("findNums() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getPower(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
		{
			name: "Game 1",
			args: args{
				line: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			},
			want: 48,
		},
		// Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
		{
			name: "Game 2",
			args: args{
				line: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			},
			want: 12,
		},
		// Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
		{
			name: "Game 3",
			args: args{
				line: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			},
			want: 1560,
		},
		// Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
		{
			name: "Game 4",
			args: args{
				line: "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			},
			want: 630,
		},
		// Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
		{
			name: "Game 5",
			args: args{
				line: "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			},
			want: 36,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPower(tt.args.line); got != tt.want {
				t.Errorf("getPower() = %v, want %v", got, tt.want)
			}
		})
	}
}
