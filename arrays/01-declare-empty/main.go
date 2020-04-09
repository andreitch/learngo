package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Declare empty arrays
//
//  1. Declare and print the following arrays with their types:
//
//    1. The names of your best three friends (names array)
//
//    2. The distances to five different locations (distances array)
//
//    3. A data buffer with five bytes of capacity (data array)
//
//    4. Currency exchange ratios only for a single currency (ratios array)
//
//    5. Up/Down status of four different web servers (alives array)
//
//    6. A byte array that doesn't occupy memory space (zero array)
//
//  2. Print only the types of the same arrays.
//
//  3. Print only the elements of the same arrays.
//
// HINT
//   When printing the elements of an array, you can use the usual Printf verbs.
//
//   For example:
//     When printing a string array, you can use "%q" verb as usual.
//
// EXPECTED OUTPUT
//  names    : [3]string{"", "", "", ""}
//  distances: [5]int{0, 0, 0, 0, 0}
//  data     : [5]uint8{0x0, 0x0, 0x0, 0x0, 0x0}
//  ratios   : [1]float64{0}
//  alives   : [4]bool{false, false, false, false}
//  zero     : [0]uint8{}
//
//  names    : [3]string
//  distances: [5]int
//  data     : [5]uint8
//  ratios   : [1]float64
//  alives   : [4]bool
//  zero     : [0]uint8
//
//  names    : ["" "" ""]
//  distances: [0 0 0 0 0]
//  data     : [0 0 0 0 0]
//  ratios   : [0.00]
//  alives   : [false false false false]
//  zero     : []
// ---------------------------------------------------------

func main() {
	var (
		names     [3]string
		distances [5]int
		data      [5]byte
		ratios    [1]float64
		alives    [4]bool
		zero      [0]byte
	)

	fmt.Printf("%-10s: %#v\n", "names", names)
	fmt.Printf("%-10s: %#v\n", "distances", distances)
	fmt.Printf("%-10s: %#v\n", "data", data)
	fmt.Printf("%-10s: %#v\n", "rations", ratios)
	fmt.Printf("%-10s: %#v\n", "alives", alives)
	fmt.Printf("%-10s: %#v\n", "zero", zero)
	fmt.Println()
	fmt.Printf("%-10s: %T\n", "names", names)
	fmt.Printf("%-10s: %T\n", "distances", distances)
	fmt.Printf("%-10s: %T\n", "data", data)
	fmt.Printf("%-10s: %T\n", "rations", ratios)
	fmt.Printf("%-10s: %T\n", "alives", alives)
	fmt.Printf("%-10s: %T\n", "zero", zero)
	fmt.Println()
	fmt.Printf("%-10s: %v\n", "names", names)
	fmt.Printf("%-10s: %v\n", "distances", distances)
	fmt.Printf("%-10s: %v\n", "data", data)
	fmt.Printf("%-10s: %v\n", "rations", ratios)
	fmt.Printf("%-10s: %v\n", "alives", alives)
	fmt.Printf("%-10s: %v\n", "zero", zero)
}
