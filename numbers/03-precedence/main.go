package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Precedence
//
//  Change the expressions to produce the expected outputs
//
//  1. 10 + 5 - 5 - 10 should print 20
//  2. -10 + 0.5 - 1 + 5.5 should print -16
//  3. 5 + 10*2 - 5 should print -25
//  4. 0.5*2 - 1should print 0.5
//  5. 3 + 1/2*10 + 4 should print 24
//  6. 10 / 2 * 10 % 7 should print 15
//  7. 100 / 5 / 2 should print 40 (Note that you may need to use floats to solve this)
//
// RESTRICTION
//  Use parentheses to change the precedence
// ---------------------------------------------------------

func main() {
	// This expression should print 20
	fmt.Println(10 + 5 - (5 - 10))

	// This expression should print -16
	fmt.Println(-10 + 0.5 - (1 + 5.5))

	// This expression should print -25
	fmt.Println(5 + 10*(2-5))

	// This expression should print 0.5
	fmt.Println(0.5 * (2 - 1))

	// This expression should print 24
	fmt.Println((3+1)/2*10 + 4)

	// This expression should print 15
	fmt.Println(10 / 2 * (10 % 7))

	// This expression should print 40
	fmt.Println(100 / (5.0 / 2))
}
