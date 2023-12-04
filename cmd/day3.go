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

type partNum struct {
	x         int
	y         int
	number    int
	hasSymbol bool
}

// type symbol struct {
// 	x      int
// 	y      int
// 	symbol rune
// }

// day3Cmd represents the day3 command
var (
	day3Cmd = &cobra.Command{
		Use:   "day3",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			day3(cmd, args)
		},
	}
	// schematic [][]rune
)

func init() {
	rootCmd.AddCommand(day3Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day3Cmd.PersistentFlags().String("foo", "", "A help for foo")
	day3Cmd.PersistentFlags().StringVarP(&inputFile, "input", "i", "input/3", "Input file to read from")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day3Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func day3(cmd *cobra.Command, args []string) {
	fmt.Println("day3 called")
	// var parts []partNum
	// parts = append(parts, partNum{})
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	height := 0
	width := 0
	for scanner.Scan() {
		height++
		if len(scanner.Text()) > width {
			width = len(scanner.Text())
		}
	}

	// Create the grid
	grid := make([][]rune, height)
	for i := range grid {
		grid[i] = make([]rune, width)
	}

	// Reset the file to the beginning
	file.Seek(0, 0)

	// Second pass: read the runes into the grid
	scanner = bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		for x, r := range scanner.Text() {
			grid[y][x] = r
		}
		y++
	}
	// // choose 2 random numbers between 1 and 140 and display the grid at those coordinates
	// h := 0 //rand.Intn(height)
	// w := 9 //rand.Intn(width)
	// fmt.Printf("Char at %d, %d: %s\n", h, w, string(grid[h][w]))
	// fmt.Printf("Height: %d, Width: %d\n", height, width)
	parts := findPartNumbers(grid)
	fmt.Println(len(parts))
}

func findPartNumbers(grid [][]rune) []partNum {
	var parts []partNum
	// var symbols []symbol
	for y, row := range grid {
		for x, r := range row {
			switch r {
			case '.':
				// do nothing
			// if its a number 1 to 9
			case '1', '2', '3', '4', '5', '6', '7', '8', '9':
				if x > 0 {
					if isNum(grid[y][x-1]) {
						continue
					}
					newPartNum := partNum{
						x:      x,
						y:      y,
						number: getPartNumberFromXY(x, y, grid),
					}
					parts = append(parts, newPartNum)
				}

			// if its not a number, or a ., its a symbol
			default:
				// newSymbol := symbol{
				// 	x:      x,
				// 	y:      y,
				// 	symbol: r,
				// }
				// symbols = append(symbols, newSymbol)
			}
		}
	}
	parts = validatePartNums(parts, grid)
	return parts
}

func getPartNumberFromXY(x int, y int, grid [][]rune) int {
	var num string
	if isNum(grid[y][x]) {
		// isnum = true
		num += string(grid[y][x])
		for x < len(grid[y]) && isNum(grid[y][x+1]) {
			num += string(grid[y][x+1])
			x++
		}
	}
	intNum, err := strconv.Atoi(num)
	if err != nil {
		log.Fatal(err)
	}
	return intNum
}

func isNum(r rune) bool {
	switch r {
	case '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return true
	default:
		return false
	}
}

func validatePartNums(parts []partNum, grid [][]rune) []partNum {
	var newParts []partNum
	for _, part := range parts {
		// we should first validate that the run at grid[y][x] is a number
		// for each of the runes in the number, check the surrounding runes to see if they are symbols
		// if any are symbols, we set hasSymbol to true and break

		r := grid[part.y][part.x]
		if !isNum(r) || !(r == []rune(strconv.Itoa(part.number))[0]) {
			// newParts = append(newParts, part)
			continue
		}
		// y-1, x-1
		if part.y == 0 || part.x == 0 {

		} else if grid[part.y-1][part.x-1] != '.' && !isNum(grid[part.y-1][part.x-1]) {
			part.hasSymbol = true
			newParts = append(newParts, part)
			continue
		}
		// y-1, x
		if part.y == 0 {
		} else if grid[part.y-1][part.x] != '.' && !isNum(grid[part.y-1][part.x]) {
			part.hasSymbol = true
			newParts = append(newParts, part)
			continue
		}
		// y-1, x+1
		if part.y == 0 || part.x == len(grid[part.y])-1 {
		} else if grid[part.y-1][part.x+1] != '.' && !isNum(grid[part.y-1][part.x+1]) {
			part.hasSymbol = true
			newParts = append(newParts, part)
			continue
		}
		// y, x-1
		if part.x == 0 {
		} else if grid[part.y][part.x-1] != '.' && !isNum(grid[part.y][part.x-1]) {
			part.hasSymbol = true
			newParts = append(newParts, part)
			continue
		}
		// y, x+1
		if part.x == len(grid[part.y])-1 {
		} else if grid[part.y][part.x+1] != '.' && !isNum(grid[part.y][part.x+1]) {
			part.hasSymbol = true
			newParts = append(newParts, part)
			continue
		}
		// y+1, x-1
		if part.y == len(grid)-1 || part.x == 0 {
		} else if grid[part.y+1][part.x-1] != '.' && !isNum(grid[part.y+1][part.x-1]) {
			part.hasSymbol = true
			newParts = append(newParts, part)
			continue
		}
		// y+1, x
		if part.y == len(grid)-1 {
		} else if grid[part.y+1][part.x] != '.' && !isNum(grid[part.y+1][part.x]) {
			part.hasSymbol = true
			newParts = append(newParts, part)
			continue
		}
		// y+1, x+1
		if part.y == len(grid)-1 || part.x == len(grid[part.y])-1 {
		} else if grid[part.y+1][part.x+1] != '.' && !isNum(grid[part.y+1][part.x+1]) {
			part.hasSymbol = true
			newParts = append(newParts, part)
			continue
		}
		newParts = append(newParts, part)
	}
	return newParts
}

func engineParts(line string) int {
	var partNum int

	return partNum
}
