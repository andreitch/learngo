package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Sum the Numbers
//
//  1. By using a loop, sum the numbers between 1 and 10.
//  2. Print the sum.
//
// EXPECTED OUTPUT
//  Sum: 55
// ---------------------------------------------------------

func main() {
	var sum int
	var i int

	for {
		if i > 10 {
			break
		}
		sum += i
		i++
	}

	fmt.Println("Sum:", sum)
}
