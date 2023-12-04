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
	"strings"

	"github.com/spf13/cobra"
)

// day2Cmd represents the day2 command
var (
	day2Cmd = &cobra.Command{
		Use:   "day2",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			day2(cmd, args)
		},
	}
	// 12 red cubes, 13 green cubes, and 14 blue cubes
	TestSet = gameSet{
		blue:  14,
		green: 13,
		red:   12,
	}
)

type gameSet struct {
	blue  int
	green int
	red   int
}

type game struct {
	id   int
	sets []gameSet
}

func init() {
	rootCmd.AddCommand(day2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	day2Cmd.PersistentFlags().StringVarP(&inputFile, "input", "i", "input/2", "Input file to read from")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func day2(cmd *cobra.Command, args []string) {
	fmt.Println("day2 called")

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	power := 0

	for scanner.Scan() {
		line := scanner.Text()

		// find the first and last numbers in each line
		total += processGame(line)
		power += getPower(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// print the total
	fmt.Println(total)
	fmt.Println(power)
}

func processGame(line string) int {
	var game game
	// break the string at :
	parts := strings.Split(line, ":")
	game.id, _ = strconv.Atoi(strings.TrimSpace(strings.Split(parts[0], " ")[1]))
	gameSetStrings := strings.Split(parts[1], ";")
	for _, gameSetString := range gameSetStrings {
		var gameSet gameSet
		gameSetString = strings.TrimSpace(gameSetString)
		gameSetParts := strings.Split(gameSetString, ",")
		for _, gameSetPart := range gameSetParts {
			gameSetPart = strings.TrimSpace(gameSetPart)
			if strings.Contains(gameSetPart, "blue") {
				gameSet.blue, _ = strconv.Atoi(strings.Split(gameSetPart, " ")[0])
			} else if strings.Contains(gameSetPart, "green") {
				gameSet.green, _ = strconv.Atoi(strings.Split(gameSetPart, " ")[0])
			} else if strings.Contains(gameSetPart, "red") {
				gameSet.red, _ = strconv.Atoi(strings.Split(gameSetPart, " ")[0])
			}
		}
		game.sets = append(game.sets, gameSet)
	}
	for _, set := range game.sets {
		if set.blue > TestSet.blue || set.green > TestSet.green || set.red > TestSet.red {
			return 0
		}
	}

	// fmt.Printf("%d: %s\n", id, gameSetStrings)
	return game.id
}

func getPower(line string) int {
	var power int
	var game game
	minset := gameSet{
		blue:  0,
		green: 0,
		red:   0,
	}
	// break the string at :
	parts := strings.Split(line, ":")
	game.id, _ = strconv.Atoi(strings.TrimSpace(strings.Split(parts[0], " ")[1]))
	gameSetStrings := strings.Split(parts[1], ";")
	for _, gameSetString := range gameSetStrings {
		var gameSet gameSet
		gameSetString = strings.TrimSpace(gameSetString)
		gameSetParts := strings.Split(gameSetString, ",")
		for _, gameSetPart := range gameSetParts {
			gameSetPart = strings.TrimSpace(gameSetPart)
			if strings.Contains(gameSetPart, "blue") {
				gameSet.blue, _ = strconv.Atoi(strings.Split(gameSetPart, " ")[0])
				if gameSet.blue > minset.blue {
					minset.blue = gameSet.blue
				}
			} else if strings.Contains(gameSetPart, "green") {
				gameSet.green, _ = strconv.Atoi(strings.Split(gameSetPart, " ")[0])
				if gameSet.green > minset.green {
					minset.green = gameSet.green
				}
			} else if strings.Contains(gameSetPart, "red") {
				gameSet.red, _ = strconv.Atoi(strings.Split(gameSetPart, " ")[0])
				if gameSet.red > minset.red {
					minset.red = gameSet.red
				}
			}
		}
		game.sets = append(game.sets, gameSet)
	}
	// for _, set := range game.sets {
	// 	if set.blue > TestSet.blue || set.green > TestSet.green || set.red > TestSet.red {
	// 		return 0
	// 	}
	// }
	fmt.Println(minset)
	power = minset.blue * minset.green * minset.red
	return power
}
