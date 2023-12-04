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

type card struct {
	id          int
	winners     []int
	myNums      []int
	winningNums []int
	points      int
}

type scratchGame struct {
	cards       []card
	totalPoints int
}

// day4Cmd represents the day4 command
var (
	day4Cmd = &cobra.Command{
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
)

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
	var Game scratchGame
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	Game.totalPoints = 0

	for scanner.Scan() {
		line := scanner.Text()

		// find the first and last numbers in each line
		Game.cards = append(Game.cards, processCard(line))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for card := range Game.cards {
		Game.totalPoints += Game.cards[card].points
	}
	fmt.Printf("Total: %d\n", Game.totalPoints)
}

func processCard(line string) card {
	var c card
	cardArray := strings.Split(line, ":")
	idArray := strings.Split(cardArray[0], " ")
	id, err := strconv.Atoi(idArray[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	c.id = id
	numArray := strings.Split(cardArray[1], "|")
	winnersArray := strings.Split(numArray[0], " ")
	myNumbersArray := strings.Split(numArray[1], " ")
	c.myNums = make([]int, len(myNumbersArray))
	for i, num := range myNumbersArray {
		c.myNums[i], err = strconv.Atoi(num)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	c.winners = make([]int, len(winnersArray))
	for i, num := range winnersArray {
		c.winners[i], err = strconv.Atoi(num)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	c.winningNums = make([]int, len(c.winners))
	for _, winner := range c.winners {
		for _, num := range c.myNums {
			if winner == num {
				c.winningNums = append(c.winningNums, winner)
			}
		}
	}
	points := 1
	for i := 0; i < len(c.winningNums); i++ {
		points *= 2
	}
	c.points = points

	return c
}
