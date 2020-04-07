package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Days in a Month
//
//  Print the number of days in a given month.
//
// RESTRICTIONS
//  1. On a leap year, february should print 29. Otherwise, 28.
//
//     Set your computer clock to 2020 to see whether it works.
//
//  2. It should work case-insensitive. See below.
//
//     Search on Google: golang pkg strings ToLower
//
//  3. Get the current year using the time.Now()
//
//     Search on Google: golang pkg time now year
//
//
// EXPECTED OUTPUT
//
//  -----------------------------------------
//  Your solution should not accept invalid months
//  -----------------------------------------
//  go run main.go
//    Give me a month name
//
//  go run main.go sheep
//    "sheep" is not a month.
//
//  -----------------------------------------
//  Your solution should handle the leap years
//  -----------------------------------------
//  go run main.go january
//    "january" has 31 days.
//
//  go run main.go february
//    "february" has 28 days.
//
//  go run main.go march
//    "march" has 31 days.
//
//  go run main.go april
//    "april" has 30 days.
//
//  go run main.go may
//    "may" has 31 days.
//
//  go run main.go june
//    "june" has 30 days.
//
//  go run main.go july
//    "july" has 31 days.
//
//  go run main.go august
//    "august" has 31 days.
//
//  go run main.go september
//    "september" has 30 days.
//
//  go run main.go october
//    "october" has 31 days.
//
//  go run main.go november
//    "november" has 30 days.
//
//  go run main.go december
//    "december" has 31 days.
//
//  -----------------------------------------
//  Your solution should be case insensitive
//  -----------------------------------------
//  go run main.go DECEMBER
//    "DECEMBER" has 31 days.
//
//  go run main.go dEcEmBeR
//    "dEcEmBeR" has 31 days.
// ---------------------------------------------------------

func main() {
	const (
		january   = 31
		february  = 28
		march     = 31
		april     = 30
		may       = 31
		june      = 30
		july      = 31
		august    = 31
		september = 30
		october   = 31
		november  = 30
		december  = 31
	)

	if len(os.Args) != 2 {
		fmt.Println("Give me a month name")
	} else if strings.ToLower(os.Args[1]) == "january" {
		fmt.Printf("%q has %d days\n", os.Args[1], january)
	} else if strings.ToLower(os.Args[1]) == "february" {
		days, year := february, time.Now().Year()
		if year%4 != 0 {
		} else if year%100 != 0 {
			days++
		} else if year%400 != 0 {
		} else {
			days++
		}
		fmt.Printf("%q has %d days\n", os.Args[1], days)
	} else if strings.ToLower(os.Args[1]) == "march" {
		fmt.Printf("%q has %d days\n", os.Args[1], march)
	} else if strings.ToLower(os.Args[1]) == "april" {
		fmt.Printf("%q has %d days\n", os.Args[1], april)
	} else if strings.ToLower(os.Args[1]) == "may" {
		fmt.Printf("%q has %d days\n", os.Args[1], may)
	} else if strings.ToLower(os.Args[1]) == "june" {
		fmt.Printf("%q has %d days\n", os.Args[1], june)
	} else if strings.ToLower(os.Args[1]) == "july" {
		fmt.Printf("%q has %d days\n", os.Args[1], july)
	} else if strings.ToLower(os.Args[1]) == "august" {
		fmt.Printf("%q has %d days\n", os.Args[1], august)
	} else if strings.ToLower(os.Args[1]) == "september" {
		fmt.Printf("%q has %d days\n", os.Args[1], september)
	} else if strings.ToLower(os.Args[1]) == "october" {
		fmt.Printf("%q has %d days\n", os.Args[1], october)
	} else if strings.ToLower(os.Args[1]) == "november" {
		fmt.Printf("%q has %d days\n", os.Args[1], november)
	} else if strings.ToLower(os.Args[1]) == "december" {
		fmt.Printf("%q has %d days\n", os.Args[1], december)
	} else {
		fmt.Printf("%q is not a month.\n", os.Args[1])
	}
}
