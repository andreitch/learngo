package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Swapper #2
//
//  1. Swap the values of `red` and `blue` variables
//
//  2. Print them
//
// EXPECTED OUTPUT
//  blue red
// ---------------------------------------------------------

func main() {
	red, blue := "red", "blue"

	red, blue = blue, red

	fmt.Println(red, blue)
}
