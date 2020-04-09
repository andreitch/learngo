package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Dynamic Difficulty
//
//  Current game picks only 5 numbers (5 turns).
//
//  Make sure that the game adjust its own difficulty
//  depending on the guess number.
//
// RESTRICTION
//  Do not make the game to easy. Only adjust the
//  difficulty if the guess is above 10.
//
// EXPECTED OUTPUT
//  Suppose that the player runs the game like this:
//    go run main.go 5
//
//    Then the computer should pick 5 random numbers.
//
//  Or, if the player runs it like this:
//    go run main.go 25
//
//    Then the computer may pick 11 random numbers
//    instead.
//
//  Or, if the player runs it like this:
//    go run main.go 100
//
//    Then the computer may pick 30 random numbers
//    instead.
//
//  As you can see, greater guess number causes the
//  game to increase the game turns, which in turn
//  adjust the game's difficulty dynamically.
//
//  Because, greater guess number makes it harder to win.
//  But, this way, game's difficulty will be dynamic.
//  It will adjust its own difficulty depending on the
//  guess number.
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
		for i := 0; i < (5 + max/3); i++ {
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
