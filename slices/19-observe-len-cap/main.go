package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Observe the length and capacity
//
//  Follow the instructions inside the code below to
//  gain more intuition about the length and capacity of a slice.
//
// ---------------------------------------------------------

func main() {
	// --- #1 ---
	// var games = []string{"Ultima", "Doom"}
	// fmt.Printf("games len: %d cap: %d", len(games), cap(games))

	// var games []string
	// fmt.Printf("games len: %d cap: %d\n", len(games), cap(games))

	// games = append(games, "pacman", "mario", "tetris", "doom")
	// fmt.Printf("games len: %d cap: %d\n", len(games), cap(games))

	games := []string{"pacman", "mario", "tetris", "doom"}
	// fmt.Printf("games's len    : %d cap: %d\n", len(games), cap(games))

	// --- #2 ---
	fmt.Println()

	for i := 0; i < 4; i++ {
		fmt.Printf("games[:%d]'s len: %d cap: %d\n", i, len(games[:i]), cap(games[:i]))
	}

	// --- #3 ---
	fmt.Println()

	zero := games[:0]
	fmt.Printf("games's     len: %d cap: %d\n", len(games), cap(games))
	fmt.Printf("zero's      len: %d cap: %d\n", len(zero), cap(zero))

	for _, e := range []string{"ultima", "dagger", "pong", "coldspot", "zetra"} {
		zero = append(zero, e)
		fmt.Printf("zero's      len: %d cap: %d\n", len(zero), cap(zero))
	}

	// --- #4 ---
	fmt.Println()

	for i, _ := range zero {
		s := zero[:i]
		fmt.Printf("zero[:%d]'s  len: %d cap: %d\n", i, len(s), cap(s))
	}

	// --- #5 ---
	fmt.Println()

	zero = zero[:cap(zero)]
	for i, _ := range zero {
		fmt.Printf("zero[:%d]'s  len: %d cap: %d - %q\n", cap(zero[:i]), len(zero[:i]), cap(zero[:i]), zero[:i])
	}

	// --- #6 ---
	fmt.Println()

	zero[3] = "booma"
	games[3] = "cooma"
	fmt.Printf("zero  : %q\n", zero)
	fmt.Printf("games : %q\n", games)

	// --- #7 ---
	// _ = games[:cap(games)+1]
}
