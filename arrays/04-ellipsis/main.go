package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Refactor to Ellipsis
//
//  1. Use the 03-array-literal exercise
//
//  2. Refactor the length of the array literals to ellipsis
//
//    This means: Use the ellipsis instead of defining the array's length
//                manually.
//
// EXPECTED OUTPUT
//   The output should be the same as the 03-array-literal exercise.
// ---------------------------------------------------------

func main() {
	names := [...]string{"Einstein", "Tesla", "Shepard"}
	distances := [...]int{50, 40, 75, 30, 125}
	data := [...]uint8{'H', 'E', 'L', 'L', 'O'}
	ratios := [...]float64{3.14145}
	alives := [...]bool{true, false, true, false}
	zero := [...]uint8{}

	fmt.Println("names")
	fmt.Println("====================")
	for i := 0; i < len(names); i++ {
		fmt.Printf("names[%d]: %q\n", i, names[i])
	}
	fmt.Println()

	fmt.Println("distances")
	fmt.Println("====================")
	for i := 0; i < len(distances); i++ {
		fmt.Printf("distances[%d]: %v\n", i, distances[i])
	}
	fmt.Println()

	fmt.Println("data")
	fmt.Println("====================")
	for i := 0; i < len(data); i++ {
		fmt.Printf("data[%d]: %v\n", i, data[i])
	}
	fmt.Println()

	fmt.Println("ratios")
	fmt.Println("====================")
	for i := 0; i < len(ratios); i++ {
		fmt.Printf("rations[%d]: %v\n", i, ratios[i])
	}
	fmt.Println()

	fmt.Println("alives")
	fmt.Println("====================")
	for i := 0; i < len(alives); i++ {
		fmt.Printf("alives[%d]: %v\n", i, alives[i])
	}
	fmt.Println()

	fmt.Println("zero")
	fmt.Println("====================")
	for i := 0; i < len(zero); i++ {
		fmt.Printf("zero[%d]: %v\n", i, zero[i])
	}
	fmt.Println()

	fmt.Printf(`
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
FOR RANGES
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

`)
	fmt.Println()

	fmt.Println("names")
	fmt.Println("====================")
	for i, v := range names {
		fmt.Printf("names[%d]: %q\n", i, v)
	}
	fmt.Println()

	fmt.Println("distances")
	fmt.Println("====================")
	for i, v := range distances {
		fmt.Printf("distances[%d]: %v\n", i, v)
	}
	fmt.Println()

	fmt.Println("data")
	fmt.Println("====================")
	for i, v := range data {
		fmt.Printf("data[%d]: %v\n", i, v)
	}
	fmt.Println()

	fmt.Println("rations")
	fmt.Println("====================")
	for i, v := range ratios {
		fmt.Printf("ratios[%d]: %v\n", i, v)
	}
	fmt.Println()

	fmt.Println("alives")
	fmt.Println("====================")
	for i, v := range alives {
		fmt.Printf("alives[%d]: %v\n", i, v)
	}
	fmt.Println()

	fmt.Println("zero")
	fmt.Println("====================")
	for i, v := range zero {
		fmt.Printf("zero[%d]: %v\n", i, v)
	}
	fmt.Println()

}
