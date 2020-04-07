package main

import (
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------
// STORY
//  You want to write a program that will manipulate a
//  given string to uppercase, lowercase, and title case.
//
// EXERCISE: String Manipulator
//
//  1. Get the operation as the first argument.
//
//  2. Get the string to be manipulated as the 2nd argument.
//
// HINT
//  You can find the manipulation functions in the strings
//  package Go documentation (ToLower, ToUpper, Title).
//
// EXPECTED OUTPUT
//
//  go run main.go
//    [command] [string]
//
//    Available commands: lower, upper and title
//
//  go run main.go lower 'OMG!'
//    omg!
//
//  go run main.go upper 'omg!'
//    OMG!
//
//  go run main.go title "mr. charles darwin"
//    Mr. Charles Darwin
//
//  go run main.go genius "mr. charles darwin"
//    Unknown command: "genius"
// ---------------------------------------------------------

func main() {
	if len(os.Args) < 3 {
		fmt.Println("[command] [string]")
		fmt.Println("")
		fmt.Println("Available commands: lower, upper and title")
		return
	}

	switch os.Args[1] {
	case "lower":
		fmt.Println(strings.ToLower(os.Args[2]))
	case "upper":
		fmt.Println(strings.ToUpper(os.Args[2]))
	case "title":
		fmt.Println(strings.Title(os.Args[2]))
	default:
		fmt.Printf("Unknown command: %q\n", os.Args[1])
	}
}
