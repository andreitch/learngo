package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Refactor Feet to Meter
//
//  Remember the feet to meters program?
//  Now, it's time to refactor it.
//  Define your own Feet and Meters types.
//
//  Follow the steps inside the code.
// ---------------------------------------------------------

func main() {
	type (
		Meters float64
		Feet   float64
	)

	var (
		feet   Feet
		meters Meters
	)

	f, _ := strconv.ParseFloat(os.Args[1], 64)
	feet = Feet(f)

	meters = Meters(feet) * 0.3048

	fmt.Printf("%g feet is %g meters.\n", feet, meters)
}
