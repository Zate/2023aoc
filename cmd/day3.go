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

type symbol struct {
	x       int
	y       int
	numbers []int
}

type cell struct {
	x int
	y int
}

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
	Symbols map[cell]symbol
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
	Symbols = make(map[cell]symbol)
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

	parts := findPartNumbers(grid)

	total := 0
	for _, part := range parts {
		if part.hasSymbol {
			total += part.number
		}
	}
	fmt.Println(total)
	total_part1 := 0
	for _, symbol := range Symbols {
		// multuply the numbers together
		total2 := 1
		if len(symbol.numbers) >= 2 {

			for _, num := range symbol.numbers {
				total2 *= num
			}
			total_part1 += total2
		}

	}
	fmt.Println(total_part1)

}

func findPartNumbers(grid [][]rune) []partNum {
	var parts []partNum
	for y, row := range grid {
		for x, r := range row {
			switch r {
			case '.':
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				if isNum(r) && (x == 0 || !isNum(grid[y][x-1])) {
					newPartNum := partNum{
						x:      x,
						y:      y,
						number: getPartNumberFromXY(x, y, grid),
					}
					parts = append(parts, newPartNum)
				}
			default:
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
		for x+1 < len(grid[y]) && isNum(grid[y][x+1]) {
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
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return true
	default:
		return false
	}
}

func validatePartNums(parts []partNum, grid [][]rune) []partNum {
	var newParts []partNum
	for _, part := range parts {
		part.hasSymbol = partNumHasSymbol(part.x, part.y, part.number, grid)
		newParts = append(newParts, part)
	}
	return newParts
}

func partNumHasSymbol(x, y, num int, grid [][]rune) bool {
	// r := grid[y][x]
	width := len(grid[y])
	height := len(grid)
	allRunes := strconv.Itoa(num)
	for _, r := range allRunes {
		// fmt.Printf("Rune: %s\n", string(r))
		if !isNum(r) {
			return false
		}

		dx := []int{-1, 0, 1, -1, 1, -1, 0, 1}
		dy := []int{-1, -1, -1, 0, 0, 1, 1, 1}

		for i := 0; i < 8; i++ {
			nx, ny := x+dx[i], y+dy[i]
			if nx < 0 || ny < 0 || nx >= width || ny >= height {
				continue
			}
			// fmt.Printf("%d, %d is %s\n", ny, nx, string(grid[ny][nx]))
			if !isNum(grid[ny][nx]) && grid[ny][nx] != '.' {
				fmt.Printf("%d has a Symbol at %d, %d: %s\n", num, ny, nx, string(grid[ny][nx]))
				if string(grid[ny][nx]) == "*" {
					cell := cell{
						x: nx,
						y: ny,
					}
					if _, ok := Symbols[cell]; !ok {
						Symbols[cell] = symbol{
							x:       nx,
							y:       ny,
							numbers: []int{num},
						}
					} else {
						Symbols[cell] = symbol{
							x:       nx,
							y:       ny,
							numbers: append(Symbols[cell].numbers, num),
						}
					}
				}
				return true
			}
		}
		x++
	}
	return false
}
