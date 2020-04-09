package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Break Up
//
//  1. Extend the "Only Evens" exercise
//  2. This time, use an infinite loop.
//  3. Break the loop when it reaches to the `max`.
//
// RESTRICTIONS
//  You should use the `break` statement once.
//
// HINT
//  Do not forget incrementing the `i` before the `continue`
//  statement and at the end of the loop.
//
// EXPECTED OUTPUT
//    go run main.go 1 10
//    2 + 4 + 6 + 8 + 10 = 30
// ---------------------------------------------------------

func main() {
	var sum int
	var i int
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

	for {
		i++
		if i%2 != 0 {
			continue
		}
		if i > max {
			break
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
