package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Slice the numbers
//
//   We've a string that contains even and odd numbers.
//
//   1. Convert the string to an []int
//
//   2. Print the slice
//
//   3. Slice it for the even numbers and print it (assign it to a new slice variable)
//
//   4. Slice it for the odd numbers and print it (assign it to a new slice variable)
//
//   5. Slice it for the two numbers at the middle
//
//   6. Slice it for the first two numbers
//
//   7. Slice it for the last two numbers (use the len function)
//
//   8. Slice the evens slice for the last one number
//
//   9. Slice the odds slice for the last two numbers
//
//
// EXPECTED OUTPUT
//   go run main.go
//
//     nums        : [2 4 6 1 3 5]
//     evens       : [2 4 6]
//     odds        : [1 3 5]
//     middle      : [6 1]
//     first 2     : [2 4]
//     last 2      : [3 5]
//     evens last 1: [6]
//     odds last 2 : [3 5]
//
//
// NOTE
//
//  You can also use my prettyslice package for printing the slices.
//
//
// HINT
//
//  Find a function in the strings package for splitting the string into []string
//
// ---------------------------------------------------------

func main() {
	var (
		nums, evens, odds []int
	)
	data := "2 4 6 1 3 5"
	for _, el := range strings.Split(data, " ") {
		e, _ := strconv.Atoi(el)
		nums = append(nums, e)
		if e%2 == 0 {
			evens = append(evens, e)
		} else {
			odds = append(odds, e)
		}
	}
	fmt.Printf("%-12s: %v\n", "nums", nums)
	fmt.Printf("%-12s: %v\n", "evens", evens)
	fmt.Printf("%-12s: %v\n", "odds", odds)
	fmt.Printf("%-12s: %v\n", "middle", nums[len(nums)/3:len(nums)/3*2])
	fmt.Printf("%-12s: %v\n", "first 2", nums[:2])
	fmt.Printf("%-12s: %v\n", "last 2", nums[len(nums)-2:])
	fmt.Printf("%-12s: %v\n", "evens last 1", evens[len(evens)-1:])
	fmt.Printf("%-12s: %v\n", "odds last 2", odds[len(odds)-2:])
}
