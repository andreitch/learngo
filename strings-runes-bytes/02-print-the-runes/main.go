package main

import (
	"fmt"
)

// ---------------------------------------------------------
// EXERCISE: Print the runes
//
//  1. Loop over the "console" word and print its runes one by one,
//     in decimals, hexadecimals and binary.
//
//  2. Manually put the runes of the "console" word to a byte slice, one by one.
//
//     As the elements of the byte slice use only the rune literals.
//
//     Print the byte slice.
//
//  3. Repeat the step 2 but this time, as the elements of the byte slice,
//     use only decimal numbers.
//
//  4. Repeat the step 2 but this time, as the elements of the byte slice,
//     use only hexadecimal numbers.
//
//
// EXPECTED OUTPUT
// c
// 		decimal: 99
// 		hex    : 0x63
// 		binary : 0b01100011
// o
// 		decimal: 111
// 		hex    : 0x6f
// 		binary : 0b01101111
// n
// 		decimal: 110
// 		hex    : 0x6e
// 		binary : 0b01101110
// s
// 		decimal: 115
// 		hex    : 0x73
// 		binary : 0b01110011
// o
// 		decimal: 111
// 		hex    : 0x6f
// 		binary : 0b01101111
// l
// 		decimal: 108
// 		hex    : 0x6c
// 		binary : 0b01101100
// e
// 		decimal: 101
// 		hex    : 0x65
// 		binary : 0b01100101
// with runes       : console
// with decimals    : console
// with hexadecimals: console
// ---------------------------------------------------------

func main() {
	const word = "console"
	for i := range word {
		fmt.Println(string(word[i]))
		fmt.Printf("\t%7s: %d\n", "decimal", word[i])
		fmt.Printf("\t%7s: 0x%x\n", "hex", word[i])
		fmt.Printf("\t%7s: 0b%b\n", "binary", word[i])
	}
	fmt.Printf("%-17s: %s\n", "with runes", []byte{'c', 'o', 'n', 's', 'o', 'l', 'e'})
	fmt.Printf("%-17s: %s\n", "with decimal", []byte{99, 111, 110, 115, 111, 108, 101})
	fmt.Printf("%-17s: %s\n", "with hexadecimal", []byte{0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65})
}
