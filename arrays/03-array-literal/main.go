package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Refactor to Array Literals
//
//  1. Use the 02-get-set-arrays exercise
//
//  2. Refactor the array assignments to array literals
//
//    1. You would need to change the array declarations to array literals
//
//    2. Then, you would need to move the right-hand side of the assignments,
//       into the array literals.
//
// EXPECTED OUTPUT
//   The output should be the same as the 02-get-set-arrays exercise.
// ---------------------------------------------------------

func main() {
	names := [3]string{"Einstein", "Tesla", "Shepard"}
	distances := [5]int{50, 40, 75, 30, 125}
	data := [5]uint8{'H', 'E', 'L', 'L', 'O'}
	ratios := [1]float64{3.14145}
	alives := [4]bool{true, false, true, false}
	zero := [0]uint8{}

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
