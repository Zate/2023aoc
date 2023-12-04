/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// day1Cmd represents the day1 command
var (
	day1Cmd = &cobra.Command{
		Use:   "day1",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			day1(cmd, args)
			day11(cmd, args)
		},
	}
)

func init() {
	rootCmd.AddCommand(day1Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	day1Cmd.PersistentFlags().StringVarP(&inputFile, "input", "i", "input/1", "Input file to read from")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day1Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func day1(cmd *cobra.Command, args []string) {
	fmt.Println("day1 called")

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		// find the first and last numbers in each line
		number := findNums(line)

		total += number
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// print the total
	fmt.Println(total)

}

func findNums(line string) int {
	var first, last rune
	// break this string into a slice of strings
	// chars := []rune(line)
	for _, char := range line {
		// if the char is a number, add it to the first string
		if char >= '0' && char <= '9' {
			if first == 0 && last == 0 {
				first = char
				last = char
			} else {
				last = char
			}
		}
	}
	num := string(first) + string(last)
	intNum, err := strconv.Atoi(num)
	if err != nil {
		log.Fatal(err)
	}
	return intNum

}

func day11(cmd *cobra.Command, args []string) {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		// find the first and last numbers in each line
		number := stringsToNum(line)

		total += number
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// print the total
	fmt.Println(total)
}

func stringsToNum(line string) int {
	found := []string{}
	max := len(line) - 1
	for i, char := range line {
		switch char {
		// case between 1 and 9
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			found = append(found, string(char))
			continue
		case 'o':
			if max-i >= 2 {
				if line[i+1] == 'n' && line[i+2] == 'e' {
					found = append(found, "1")
					continue
				}
			}
		case 't':
			if max-i >= 2 {
				if line[i+1] == 'w' && line[i+2] == 'o' {
					found = append(found, "2")
					continue
				} else if max-i >= 4 {
					if line[i+1] == 'h' && line[i+2] == 'r' && line[i+3] == 'e' && line[i+4] == 'e' {
						found = append(found, "3")
						continue
					}
				}
			}
		case 'f':
			if max-i >= 3 {
				if line[i+1] == 'o' && line[i+2] == 'u' && line[i+3] == 'r' {
					found = append(found, "4")
					continue
				} else if line[i+1] == 'i' && line[i+2] == 'v' && line[i+3] == 'e' {
					found = append(found, "5")
					continue
				}
			}
		case 's':
			if max-i >= 2 {
				if line[i+1] == 'i' && line[i+2] == 'x' {
					found = append(found, "6")
					continue
				} else if max-i >= 4 {
					if line[i+1] == 'e' && line[i+2] == 'v' && line[i+3] == 'e' && line[i+4] == 'n' {
						found = append(found, "7")
						continue
					}
				}
			}
		case 'e':
			if max-i >= 4 {
				if line[i+1] == 'i' && line[i+2] == 'g' && line[i+3] == 'h' && line[i+4] == 't' {
					found = append(found, "8")
					continue
				}
			}
		case 'n':
			if max-i >= 3 {
				if line[i+1] == 'i' && line[i+2] == 'n' && line[i+3] == 'e' {
					found = append(found, "9")
					continue
				}
			}
		default:
		}
	}
	num := found[0] + found[len(found)-1]
	intNum, err := strconv.Atoi(num)
	if err != nil {
		log.Fatal(err)
	}
	return intNum
}
