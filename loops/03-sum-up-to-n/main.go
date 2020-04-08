package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Sum up to N
//
//  1. Get two numbers from the command-line: min and max
//  2. Convert them to integers (using Atoi)
//  3. By using a loop, sum the numbers between min and max
//
// RESTRICTIONS
//  1. Be sure to handle the errors. So, if a user doesn't
//     pass enough arguments or she passes non-numerics,
//     then warn the user and exit from the program.
//
//  2. Also, check that, min < max.
//
// HINT
//  For converting the numbers, you can use `strconv.Atoi`.
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10 = 55
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
		if i == max {
			fmt.Print(i)
		} else {
			fmt.Print(i, " + ")
		}
		sum += i
	}

	fmt.Printf(" = %d\n", sum)
}
