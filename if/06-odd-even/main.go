package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Odd or Even
//
//  1. Get a number from the command-line.
//
//  2. Find whether the number is odd, even and divisible by 8.
//
// RESTRICTION
//  Handle the error. If the number is not a valid number,
//  or it's not provided, let the user know.
//
// EXPECTED OUTPUT
//  go run main.go 16
//    16 is an even number and it's divisible by 8
//
//  go run main.go 4
//    4 is an even number
//
//  go run main.go 3
//    3 is an odd number
//
//  go run main.go
//    Pick a number
//
//  go run main.go ABC
//    "ABC" is not a number
// ---------------------------------------------------------

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println("Pick a number")
	} else if i, err := strconv.Atoi(args[1]); err != nil {
		fmt.Printf("%q is not a number\n", args[1])
	} else if i%2 == 1 {
		fmt.Println(i, "is an odd number")
	} else if i%8 == 0 {
		fmt.Println(i, "is an even number and it's divisible by 8")
	} else {
		fmt.Println(i, "is an even number")
	}
}
