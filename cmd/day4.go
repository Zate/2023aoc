/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"runtime"
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
	cardList := make(map[int]card)
	for i, card := range Game.cards {
		Game.totalPoints += Game.cards[i].points
		cardList[Game.cards[i].id] = card

	}
	fmt.Printf("Total: %d\n", Game.totalPoints)
	totalcards := processCardList(cardList)
	fmt.Printf("Total: %d\n", totalcards)

}

func processCardList(c map[int]card) int {
	totalCards := len(c)
	processedCards := make(map[int]bool)
	cardStats := make(map[int]int)
	total := 0
	iteration := 0
	for {
		iteration++
		fmt.Printf("Iteration: %d\n", iteration)
		newCardsWon := false
		for i := 0; i < totalCards; i++ {
			if processedCards[i] {
				fmt.Printf("Skipping card %d\n", i)
				continue
			}
			card := c[i]
			numCards := len(card.winningNums)
			if numCards == 0 {
				fmt.Printf("Card %d has no winners\n", i)
				processedCards[i] = true
				continue
			}
			for j := i + 1; j < i+1+numCards && j < totalCards; j++ {
				fmt.Printf("Card %d won by card %d\n", i, j)
				cardStats[j]++
				processedCards[j] = true
				newCardsWon = true
			}
		}
		fmt.Printf("newCardsWon: %v\n", newCardsWon)
		fmt.Printf("processedCards: %v\n", processedCards)
		if !newCardsWon {
			break
		}
	}
	for _, v := range cardStats {
		total += v
	}
	return total
}

func processCard(line string) card {
	var c card
	cardArray := strings.Split(line, ":")
	idArray := strings.Split(cardArray[0], " ")

	id, err := strconv.Atoi(idArray[len(idArray)-1])
	if err != nil {
		pc, _, line, _ := runtime.Caller(0)
		fmt.Printf("Error in %s[%d]: %v\n", runtime.FuncForPC(pc).Name(), line, err)
		os.Exit(1)
	}
	c.id = id
	numArray := strings.Split(cardArray[1], "|")
	winnersArray := strings.Split(numArray[0], " ")
	myNumbersArray := strings.Split(numArray[1], " ")
	c.myNums = make([]int, 0, len(myNumbersArray)) // make it with length 0 but enough capacity
	j := 0
	for _, num := range myNumbersArray {
		num = strings.Trim(num, " ")
		if num == "" {
			continue
		}
		var parsedNum int
		parsedNum, err = strconv.Atoi(num)
		if err != nil {
			pc, _, line, _ := runtime.Caller(0)
			fmt.Printf("Error in %s[%d]: %v\n", runtime.FuncForPC(pc).Name(), line, err)
			os.Exit(1)
		}
		c.myNums = append(c.myNums, parsedNum) // append the parsed number to the slice
		j++
	}
	c.winners = make([]int, 0, len(winnersArray)) // make it with length 0 but enough capacity
	j = 0
	for _, num := range winnersArray {
		num = strings.Trim(num, " ")
		if num == "" {
			continue
		}
		var parsedNum int
		parsedNum, err = strconv.Atoi(num)
		if err != nil {
			pc, _, line, _ := runtime.Caller(0)
			fmt.Printf("Error in %s[%d]: %v\n", runtime.FuncForPC(pc).Name(), line, err)
			os.Exit(1)
		}
		c.winners = append(c.winners, parsedNum) // append the parsed number to the slice
		j++
	}
	c.winningNums = make([]int, 0, len(c.winners)) // make it with length 0 but enough capacity
	for _, winner := range c.winners {
		for _, num := range c.myNums {
			if winner == num {
				c.winningNums = append(c.winningNums, num) // append the number to the slice
				break
			}
		}
	}
	c.points = int(math.Pow(2, float64(len(c.winningNums)-1)))
	// for i := 0; i < len(c.winningNums)-1; i++ {
	// 	points *= 2
	// }
	// c.points = points

	return c
}
