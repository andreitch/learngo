package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Enough Picks
//
//  If the player's guess number is below 10;
//  then make sure that the computer generates a random
//  number between 0 and 10.
//
//  However, if the guess number is above 10; then the
//  computer should generate the numbers
//  between 0 and the guess number.
//
// WHY?
//  This way the game will be harder.
//
//  Because, in the current version of the game, if
//  the player's guess number is for example 3; then the
//  computer picks a random number between 0 and 3.
//
//  So, currently a player can easily win the game.
//
// EXAMPLE
//  Suppose that the player runs the game like this:
//    go run main.go 9
//
//  Or like this:
//    go run main.go 2
//
//    Then the computer should pick a random number
//    between 0-10.
//
//  Or, if the player runs it like this:
//    go run main.go 15
//
//    Then the computer should pick a random number
//    between 0-15.
// ---------------------------------------------------------

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Guess number")
		return
	}
	rand.Seed(time.Now().UnixNano())
	max := 10

	if num, err := strconv.Atoi(os.Args[1]); num < 1 || err != nil {
		fmt.Println("Guess number")
	} else {
		if num > max {
			max = num + 1
		}
		for i := 0; i < 5; i++ {
			if num == rand.Intn(max) {
				if i == 1 {
					fmt.Print("So fast!")
				}
				switch rand.Intn(2) {
				case 0:
					fmt.Println("You won!")
				case 1:
					fmt.Println("YOU'RE AWESOME")
				}
				return
			}
		}
		switch rand.Intn(2) {
		case 0:
			fmt.Println("LOSER!")
		case 1:
			fmt.Println("YOU LOST. TRY AGAIN?")
		}
	}
}
