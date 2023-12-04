/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// day4Cmd represents the day4 command
var day4Cmd = &cobra.Command{
	Use:   "day4",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		day4(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(day4Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day4Cmd.PersistentFlags().String("foo", "", "A help for foo")
	day4Cmd.PersistentFlags().StringVarP(&inputFile, "input", "i", "input/4", "Input file to read from")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day4Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func day4(cmd *cobra.Command, args []string) {
	fmt.Println("day4 called")
}
