package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Iota Months #2
//
//  1. Initialize multiple constants using iota.
//
// EXPECTED OUTPUT
//  1 2 3
// ---------------------------------------------------------

func main() {
	const (
		_ = iota
		Jan
		Feb
		Mar
	)

	fmt.Println(Jan, Feb, Mar)
}
