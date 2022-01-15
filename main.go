package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"sort"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	dice := flag.String("d", "d20", "The type of dice to roll. Format:dX where is an integer. Default: d20	")
	numroll := flag.Int("n", 1, "The number of die to roll. Default: 1")
	sum := flag.Bool("s", false, "Get the sum of all the dice rolls")
	advantage := flag.Bool("a", false, "Roll the dice advantage")
	disadvantage := flag.Bool("dis", false, "Roll the dice disadvantage")
	flag.Parse()

	fmt.Printf("You roll a %s\n", *dice)

	matched, _ := regexp.Match("d\\d+", []byte(*dice))

	if matched == true {
		rolls := rollDice(dice, numroll)
		printDice(rolls)
		if *sum == true {
			dicesum := sumDice(rolls)
			fmt.Printf("The sum of the dice was %d\n", dicesum)
		}
		if *advantage == true {
			roll := rollWithAdvantage(rolls)
			fmt.Printf("The roll with advantage was %d\n", roll)
		}
		if *disadvantage == true {
			roll := rollWithDisadvantage(rolls)
			fmt.Printf("The roll with disadvantage was %d\n", roll)
		}
	} else {
		log.Fatalf("Improper format for the dice. The format should be dX where X is an integer")
	}
}

func rollDice(dice *string, times *int) []int {
	var rolls []int
	diceSides := (*dice)[1:] // With paranthesis, d reference happens first and we access the stuff within the string.
	d, err := strconv.Atoi(diceSides)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < *times; i++ {
		rolls = append(rolls, rand.Intn(d)+1)
	}
	return rolls
}

func printDice(rolls []int) {
	for i, dice := range rolls {
		fmt.Printf("Roll %d was %d\n", i+1, dice)
	}
}

func sumDice(rolls []int) int {
	sum := 0
	for _, dice := range rolls {
		sum += dice
	}
	return sum
}

func rollWithAdvantage(rolls []int) int {
	sort.Ints(rolls)
	return rolls[len(rolls)-1]
}

func rollWithDisadvantage(rolls []int) int {
	sort.Ints(rolls)
	return rolls[0]
}
