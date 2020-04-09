package main

import (
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Case Insensitive Search
//
//  Allow for case-insensitive searching
//
// EXAMPLE
//  Let's say that the user runs the program like this:
//    go run main.go LAZY
//
//  Or like this: go run main.go lAzY
//  Or like this: go run main.go lazy
//
//  For all cases above, the program should find
//  the "lazy" keyword.
// ---------------------------------------------------------

func main() {
	const corpus = "The lazy cat is sleeping slpeepy"

	if len(os.Args) < 2 {
		fmt.Println("Specify search word")
		return
	}

	words := strings.Fields(corpus)
	search := os.Args[1:]

	for i := 0; i < len(words); i++ {
	queries:
		for j := 0; j < len(search); j++ {
			if strings.ToLower(words[i]) == strings.ToLower(search[j]) {
				fmt.Printf("#%-2d %q\n", i+1, search[j])
				break queries
			}
		}
	}
}
