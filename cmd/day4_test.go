/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"reflect"
	"testing"

	"github.com/spf13/cobra"
)

var (
	testData4 = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
	Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
	Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
	Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
	Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
	Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`
)

func Test_day4(t *testing.T) {
	type args struct {
		cmd  *cobra.Command
		args []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			day4(tt.args.cmd, tt.args.args)
		})
	}
}

func Test_processCard(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want card
	}{
		{
			name: "Card 1",
			args: args{
				line: "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			},
			// id          int
			// winners     []int
			// myNums      []int
			// winningNums []int
			// points      int
			want: card{
				id:          1,
				winners:     []int{41, 48, 83, 86, 17},
				myNums:      []int{83, 86, 6, 31, 17, 9, 48, 53},
				winningNums: []int{48, 83, 86, 17},
				points:      8,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processCard(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processCard() = %v, want %v", got, tt.want)
			}
		})
	}
}
