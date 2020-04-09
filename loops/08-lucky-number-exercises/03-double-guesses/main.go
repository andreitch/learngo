package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Double Guesses
//
//  Let the player guess two numbers instead of one.
//
// HINT:
//  Generate random numbers using the greatest number
//  among the guessed numbers.
//
// EXAMPLES
//  go run main.go 5 6
//  Player wins if the random number is either 5 or 6.
// ---------------------------------------------------------

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Guess 2 numbers")
		return
	}
	rand.Seed(time.Now().UnixNano())
	num, err := strconv.Atoi(os.Args[1])
	num2, err2 := strconv.Atoi(os.Args[2])

	if num < 1 || err != nil || num2 < 1 || err2 != nil {
		fmt.Println("Guess 2 numbers")
	} else {
		for i := 0; i < 5; i++ {
			n := rand.Intn(num2 + 1)
			if num == n || num2 == n {
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
