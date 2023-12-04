/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"testing"

	"github.com/spf13/cobra"
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
