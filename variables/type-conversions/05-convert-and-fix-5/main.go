package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Convert and Fix #5
//
//  Fix the code.
//
// HINTS
//   maximum of int8  can be 127
//   maximum of int16 can be 32767
//
// EXPECTED OUTPUT
//  1127
// ---------------------------------------------------------

func main() {
	min := int8(127)
	max := int16(1000)

	fmt.Println(max + int16(min))
}
