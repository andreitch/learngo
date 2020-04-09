package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Random Messages
//
//  Display a few different won and lost messages "randomly".
//
// HINTS
//  1. You can use a switch statement to do that.
//  2. Set its condition to the random number generator.
//  3. I would use a short switch.
//
// EXAMPLES
//  The Player wins: then randomly print one of these:
//
//  go run main.go 5
//    YOU WON
//  go run main.go 5
//    YOU'RE AWESOME
//
//  The Player loses: then randomly print one of these:
//
//  go run main.go 5
//    LOSER!
//  go run main.go 5
//    YOU LOST. TRY AGAIN?
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
