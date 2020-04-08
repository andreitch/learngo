package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Math Tables
//
//  Create division, addition and subtraction tables
//
//  1. Get the math operation and
//     the size of the table from the user
//
//  2. Print the table accordingly
//
//  3. Correctly handle the division by zero error
//
//
// BONUS #1
//
//  Use strings.IndexAny function to detect
//    the valid operations.
//
//  Search on Google for: golang pkg strings IndexAny
//
// BONUS #2
//
//  Add remainder operation as well (remainder table using %).
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//    Usage: [op=*/+-] [size]
//
//  go run main.go "*"
//    Size is missing
//    Usage: [op=*/+-] [size]
//
//  go run main.go "%" 4
//    Invalid operator.
//    Valid ops one of: */+-
//
//  go run main.go "*" 4
//    *    0    1    2    3    4
//    0    0    0    0    0    0
//    1    0    1    2    3    4
//    2    0    2    4    6    8
//    3    0    3    6    9   12
//    4    0    4    8   12   16
//
//  go run main.go "/" 4
//    /    0    1    2    3    4
//    0    0    0    0    0    0
//    1    0    1    0    0    0
//    2    0    2    1    0    0
//    3    0    3    1    1    0
//    4    0    4    2    1    1
//
//  go run main.go "+" 4
//    +    0    1    2    3    4
//    0    0    1    2    3    4
//    1    1    2    3    4    5
//    2    2    3    4    5    6
//    3    3    4    5    6    7
//    4    4    5    6    7    8
//
//  go run main.go "-" 4
//    -    0    1    2    3    4
//    0    0   -1   -2   -3   -4
//    1    1    0   -1   -2   -3
//    2    2    1    0   -1   -2
//    3    3    2    1    0   -1
//    4    4    3    2    1    0
//
// BONUS:
//
//  go run main.go "%" 4
//    %    0    1    2    3    4
//    0    0    0    0    0    0
//    1    0    0    1    1    1
//    2    0    0    0    2    2
//    3    0    0    1    0    3
//    4    0    0    0    1    0
//
// NOTES:
//
//   When running the program in Windows Git Bash, you might need
//   to escape the characters like this.
//
//   This happens because those characters have special meaning.
//
//   Division:
//     go run main.go // 4
//
//   Multiplication and others:
//   (this is also necessary for non-Windows bashes):
//
//     go run main.go "*" 4
// ---------------------------------------------------------

func main() {
	switch len(os.Args) {
	case 1:
		fmt.Println("Usage: [op=*/+-] [size]")
		return
	case 2:
		fmt.Println("Size is missing")
		fmt.Println("Usage: [op=*/+-] [size]")
		return
	}

	if num, err := strconv.Atoi(os.Args[2]); num < 1 || err != nil || strings.IndexAny(os.Args[1], "*/+-%") == -1 {
		fmt.Println("Invalid operator.")
		fmt.Println("Valid ops one of: */+-%")
		return
	} else {
		if num > 10 {
			num = 10
		}
		fmt.Printf("%2s", os.Args[1])
		for f := 0; f <= num; f++ {
			fmt.Printf("%5d", f)
		}

		for x := 0; x <= num; x++ {
			fmt.Printf("\n%2d", x)
			for j := 0; j <= num; j++ {
				switch os.Args[1] {
				case "*":
					fmt.Printf("%5d", x*j)
				case "/":
					if j == 0 {
						fmt.Printf("%5s", "0")
					} else {
						fmt.Printf("%5d", x/j)
					}
				case "+":
					fmt.Printf("%5d", x+j)
				case "-":
					fmt.Printf("%5d", x-j)
				case "%":
					if j == 0 {
						fmt.Printf("%5s", "0")
					} else {
						fmt.Printf("%5d", x%j)
					}
				default:
					fmt.Printf("%5s", "-")
				}

			}
		}
		fmt.Println()
	}
}
