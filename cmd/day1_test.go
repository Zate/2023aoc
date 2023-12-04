/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import "testing"

func Test_findNums(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// 1abc2
		// pqr3stu8vwx
		// a1b2c3d4e5f
		// treb7uchet
		{
			name: "1abc2",
			args: args{
				line: "1abc2",
			},
			want: 12,
		},
		{
			name: "pqr3stu8vwx",
			args: args{
				line: "pqr3stu8vwx",
			},
			want: 38,
		},
		{
			name: "a1b2c3d4e5f",
			args: args{
				line: "a1b2c3d4e5f",
			},
			want: 15,
		},
		{
			name: "treb7uchet",
			args: args{
				line: "treb7uchet",
			},
			want: 77,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNums(tt.args.line); got != tt.want {
				t.Errorf("findNums() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringsToNum(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// two1nine
		// eightwothree
		// abcone2threexyz
		// xtwone3four
		// 4nineeightseven2
		// zoneight234
		// 7pqrstsixteen
		{
			name: "two1nine",
			args: args{
				line: "two1nine",
			},
			want: 29,
		},
		{
			name: "eightwothree",
			args: args{
				line: "eightwothree",
			},
			want: 83,
		},
		{
			name: "abcone2threexyz",
			args: args{
				line: "abcone2threexyz",
			},
			want: 13,
		},
		{
			name: "xtwone3four",
			args: args{
				line: "xtwone3four",
			},
			want: 24,
		},
		{
			name: "4nineeightseven2",
			args: args{
				line: "4nineeightseven2",
			},
			want: 42,
		},
		{
			name: "zoneight234",
			args: args{
				line: "zoneight234",
			},
			want: 14,
		},
		{
			name: "7pqrstsixteen",
			args: args{
				line: "7pqrstsixteen",
			},
			want: 76,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringsToNum(tt.args.line); got != tt.want {
				t.Errorf("findNums() = %v, want %v", got, tt.want)
			}
		})
	}
}
