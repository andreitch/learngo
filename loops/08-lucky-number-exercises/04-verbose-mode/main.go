package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Verbose Mode
//
//  When the player runs the game like this:
//
//    go run main.go -v 4
//
//  Display each generated random number:

//    1 3 4 🎉  YOU WIN!
//
//  In this example, computer picks 1, 3, and 4. And the
//  player wins.
//
// HINT
//  You need to get and interpret the command-line arguments.
// ---------------------------------------------------------

const usage = `Usage:
go run main.go -v 4
go run main.go 4
`

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage)
		return
	}

	rand.Seed(time.Now().UnixNano())

	if len(os.Args) < 3 {
		num, err := strconv.Atoi(os.Args[1])
		if num < 0 || err != nil {
			fmt.Println(usage)
			return
		}
		for turn := 0; turn < 3; turn++ {
			if num == rand.Intn(10) {
				fmt.Println("YOU WIN!")
				return
			}
		}
	} else {
		num, err := strconv.Atoi(os.Args[2])
		if os.Args[1] != "-v" || err != nil || num < 0 {
			fmt.Println(usage)
			return
		}

		for turn := 0; turn < 3; turn++ {
			n := rand.Intn(10)
			fmt.Print(n, " ")
			if num == n {
				fmt.Println("YOU WIN!")
				return
			}
		}
	}
	fmt.Println("YOU LOST!")
}
