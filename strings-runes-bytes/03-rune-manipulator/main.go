package main

import (
	"fmt"
	"unicode/utf8"
)

// ---------------------------------------------------------
// EXERCISE: Rune Manipulator
//
//  Please read the comments inside the following code.
//
// EXPECTED OUTPUT
//  "cool"
// 			has 4 bytes and 4 runes
// 			bytes   : 63 6f 6f 6c
// 			runes   : [63 6f 6f 6c]
// 			runes   : 'c' 'o' 'o' 'l'
// 			first   : 'c' (1 bytes)
// 			last    : 'l' (1 bytes)
// 			first 2 : "co"
// 			last 2  : "ol"
// 			first 2 : "co"
// 			last  2 : "ol"
// "güzel"
// 			has 6 bytes and 5 runes
// 			bytes   : 67 c3 bc 7a 65 6c
// 			runes   : [67 fc 7a 65 6c]
// 			runes   : 'g' 'ü' 'z' 'e' 'l'
// 			first   : 'g' (1 bytes)
// 			last    : 'l' (1 bytes)
// 			first 2 : "gü"
// 			last 2  : "el"
// 			first 2 : "gü"
// 			last  2 : "el"
// "jīntiān"
// 			has 9 bytes and 7 runes
// 			bytes   : 6a c4 ab 6e 74 69 c4 81 6e
// 			runes   : [6a 12b 6e 74 69 101 6e]
// 			runes   : 'j' 'ī' 'n' 't' 'i' 'ā' 'n'
// 			first   : 'j' (1 bytes)
// 			last    : 'n' (1 bytes)
// 			first 2 : "jī"
// 			last 2  : "ān"
// 			first 2 : "jī"
// 			last  2 : "ān"
// "今天"
// 			has 6 bytes and 2 runes
// 			bytes   : e4 bb 8a e5 a4 a9
// 			runes   : [4eca 5929]
// 			runes   : '今' '天'
// 			first   : '今' (3 bytes)
// 			last    : '天' (3 bytes)
// 			first 2 : "今天"
// 			last 2  : "今天"
// 			first 2 : "今天"
// 			last  2 : "今天"
// "read 🤓"
// 			has 9 bytes and 6 runes
// 			bytes   : 72 65 61 64 20 f0 9f a4 93
// 			runes   : [72 65 61 64 20 1f913]
// 			runes   : 'r' 'e' 'a' 'd' ' ' '🤓'
// 			first   : 'r' (1 bytes)
// 			last    : '🤓' (4 bytes)
// 			first 2 : "re"
// 			last 2  : " 🤓"
// 			first 2 : "re"
// 			last  2 : " 🤓"
// ---------------------------------------------------------

func main() {
	strings := []string{
		"cool",
		"güzel",
		"jīntiān",
		"今天",
		"read 🤓",
	}

	_ = strings

	for i := range strings {
		fmt.Printf("%q\n", strings[i])
		// Print the byte and rune length of the strings
		// Hint: Use len and utf8.RuneCountInString
		fmt.Printf("\thas %d bytes and %d runes\n", len(strings[i]), utf8.RuneCountInString(strings[i]))

		// Print the bytes of the strings in hexadecimal
		// Hint: Use % x verb
		fmt.Printf("\tbytes   : % x\n", strings[i])

		// Print the runes of the strings in hexadecimal
		// Hint: Use % x verb
		fmt.Printf("\trunes   : %x\n", []rune(strings[i]))

		// Print the runes of the strings as rune literals
		// Hint: Use for range
		fmt.Printf("\trunes   : ")
		for _, r := range strings[i] {
			fmt.Printf("%q ", r)
		}
		fmt.Println()

		// Print the first rune and its byte size of the strings
		// Hint: Use utf8.DecodeRuneInString
		r, size := utf8.DecodeRuneInString(strings[i])
		fmt.Printf("\tfirst   : %q (%d bytes)\n", r, size)

		// Print the last rune of the strings
		// Hint: Use utf8.DecodeLastRuneInString
		r, size = utf8.DecodeLastRuneInString(strings[i])
		fmt.Printf("\tlast    : %q (%d bytes)\n", r, size)

		// Slice and print the first two runes of the strings
		r, size = utf8.DecodeRuneInString(strings[i])
		r2, _ := utf8.DecodeRuneInString(strings[i][size:])
		fmt.Printf("\tfirst 2 : \"%c%c\"\n", r, r2)

		// Slice and print the last two runes of the strings
		r, size = utf8.DecodeLastRuneInString(strings[i])
		r2, _ = utf8.DecodeLastRuneInString(strings[i][:len(strings[i])-size])
		fmt.Printf("\tlast 2  : \"%c%c\"\n", r2, r)

		// Convert the string to []rune
		// Print the first and last two runes
		rns := []rune(strings[i])
		fmt.Printf("\tfirst 2 : \"%c%c\"\n", rns[0], rns[1])
		fmt.Printf("\tlast  2 : \"%c%c\"\n", rns[len(rns)-2], rns[len(rns)-1])
	}
}
