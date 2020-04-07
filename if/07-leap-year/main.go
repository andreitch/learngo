package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Leap Year
//
//  Find out whether a given year is a leap year or not.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me a year number
//
//  go run main.go eighties
//    "eighties" is not a valid year.
//
//  go run main.go 2018
//    2018 is not a leap year.
//
//  go run main.go 2100
//    2100 is not a leap year.
//
//  go run main.go 2019
//    2019 is not a leap year.
//
//  go run main.go 2020
//    2020 is a leap year.
//
//  go run main.go 2024
//    2024 is a leap year.
// ---------------------------------------------------------

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println("Give me a year number")
	} else if i, err := strconv.Atoi(args[1]); err != nil {
		fmt.Printf("%q is not a valid year.\n", args[1])
	} else if i%4 != 0 {
		fmt.Println(i, "is not a leap year.")
	} else if i%100 != 0 {
		fmt.Println(i, "is a leap year.")
	} else if i%400 != 0 {
		fmt.Println(i, "is not a leap year.")
	} else {
		fmt.Println(i, "is a leap.year.")
	}
}
