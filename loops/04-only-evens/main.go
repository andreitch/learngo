package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Only Evens
//
//  1. Extend the "Sum up to N" exercise
//  2. Sum only the even numbers
//
// RESTRICTIONS
//  Skip odd numbers using the `continue` statement
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    2 + 4 + 6 + 8 + 10 = 30
// ---------------------------------------------------------

func main() {
	var sum int
	if len(os.Args) != 3 {
		fmt.Println("Input two valid numbers")
		return
	}

	min, errMin := strconv.Atoi(os.Args[1])
	max, errMax := strconv.Atoi(os.Args[2])

	if errMin != nil || errMax != nil {
		fmt.Println("Input two valid numbers")
		return
	}
	if min >= max {
		fmt.Println("Input two valid numbers")
		return
	}

	for i := min; i <= max; i++ {
		if i%2 != 0 {
			continue
		}
		if i == max {
			fmt.Print(i)
		} else {
			fmt.Print(i, " + ")
		}
		sum += i
	}

	fmt.Printf(" = %d\n", sum)
}
