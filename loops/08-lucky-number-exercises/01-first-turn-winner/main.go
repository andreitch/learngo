package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: First Turn Winner
//
//  If the player wins on the first turn, then display
//  a special bonus message.
//
// RESTRICTION
//  The first picked random number by the computer should
//  match the player's guess.
//
// EXAMPLE SCENARIO
//  1. The player guesses 6
//  2. The computer picks a random number and it happens
//     to be 6
//  3. Terminate the game and print the bonus message
// ---------------------------------------------------------

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Guess number")
		return
	}
	rand.Seed(time.Now().UnixNano())

	if num, err := strconv.Atoi(os.Args[1]); num < 1 || err != nil {
		fmt.Println("Guess number")
	} else {
		for i := 0; i < 5; i++ {
			if num == rand.Intn(10) {
				if i == 1 {
					fmt.Print("So fast!")
				}
				fmt.Println("You won!")
				return
			}
		}
	}
}
