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
			want: card{
				id:          1,
				winners:     []int{41, 48, 83, 86, 17},
				myNums:      []int{83, 86, 6, 31, 17, 9, 48, 53},
				winningNums: []int{48, 83, 86, 17},
				points:      8,
			},
		},
		{
			name: "Card 2",
			args: args{
				line: "Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
			},
			want: card{
				id:          2,
				winners:     []int{13, 32, 20, 16, 61},
				myNums:      []int{61, 30, 68, 82, 17, 32, 24, 19},
				winningNums: []int{32, 61},
				points:      2,
			},
		},
		{
			name: "Card 3",
			args: args{
				line: "Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
			},
			want: card{
				id:          3,
				winners:     []int{1, 21, 53, 59, 44},
				myNums:      []int{69, 82, 63, 72, 16, 21, 14, 1},
				winningNums: []int{1, 21},
				points:      2,
			},
		},
		{
			name: "Card 4",
			args: args{
				line: "Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
			},
			want: card{
				id:          4,
				winners:     []int{41, 92, 73, 84, 69},
				myNums:      []int{59, 84, 76, 51, 58, 5, 54, 83},
				winningNums: []int{84},
				points:      1,
			},
		},
		{
			name: "Card 5",
			args: args{
				line: "Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
			},
			want: card{
				id:          5,
				winners:     []int{87, 83, 26, 28, 32},
				myNums:      []int{88, 30, 70, 12, 93, 22, 82, 36},
				winningNums: []int{},
				points:      0,
			},
		},
		{
			name: "Card 6",
			args: args{
				line: "Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
			},
			want: card{
				id:          6,
				winners:     []int{31, 18, 13, 56, 72},
				myNums:      []int{74, 77, 10, 23, 35, 67, 36, 11},
				winningNums: []int{},
				points:      0,
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

func Test_processCardList(t *testing.T) {
	cards := []card{
		{
			id:          1,
			winners:     []int{41, 48, 83, 86, 17},
			myNums:      []int{83, 86, 6, 31, 17, 9, 48, 53},
			winningNums: []int{48, 83, 86, 17},
			points:      8,
		},
		{
			id:          2,
			winners:     []int{13, 32, 20, 16, 61},
			myNums:      []int{61, 30, 68, 82, 17, 32, 24, 19},
			winningNums: []int{32, 61},
			points:      2,
		},
		{
			id:          3,
			winners:     []int{1, 21, 53, 59, 44},
			myNums:      []int{69, 82, 63, 72, 16, 21, 14, 1},
			winningNums: []int{1, 21},
			points:      2,
		},
		{
			id:          4,
			winners:     []int{41, 92, 73, 84, 69},
			myNums:      []int{59, 84, 76, 51, 58, 5, 54, 83},
			winningNums: []int{84},
			points:      1,
		},

		{
			id:          5,
			winners:     []int{87, 83, 26, 28, 32},
			myNums:      []int{88, 30, 70, 12, 93, 22, 82, 36},
			winningNums: []int{},
			points:      0,
		},
		{
			id:          6,
			winners:     []int{31, 18, 13, 56, 72},
			myNums:      []int{74, 77, 10, 23, 35, 67, 36, 11},
			winningNums: []int{},
			points:      0,
		},
	}
	cardList := make(map[int]card)
	for card := range cards {
		cardList[cards[card].id] = cards[card]

	}
	type args struct {
		c map[int]card
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				c: cardList,
			},
			want: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processCardList(tt.args.c); got != tt.want {
				t.Errorf("processCardList() = %v, want %v", got, tt.want)
			}
		})
	}
}
