package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: TAU
//
//  Fix the following program and print the TAU number.
//
// HINT
//  You can use %g verb for printing tau.
//
// EXPECTED OUTPUT
//  tau = 6.283185307179586
// ---------------------------------------------------------

func main() {
	// What's the problem with this code?
	// Why it doesn't work?

	const pi = 3.14159265358979323846264
	const tau = pi * 2

	fmt.Printf("tau = %g\n", tau)
}
