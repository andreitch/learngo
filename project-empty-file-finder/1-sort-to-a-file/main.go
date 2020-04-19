package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

// ---------------------------------------------------------
// EXERCISE: Sort and write items to a file
//
//  1. Get arguments from command-line
//
//  2. Sort them
//
//  3. Write the sorted slice to a file
//
//
// EXPECTED OUTPUT
//
//   go run main.go
//     Send me some items and I will sort them
//
//   go run main.go orange banana apple
//
//   cat sorted.txt
//     apple
//     banana
//     orange
//
//
// HINTS
//
//   + REMEMBER: os.Args is a []string
//
//   + String slices are sortable using `sort.Strings`
//
//   + Use ioutil.WriteFile to write to a file.
//
//   + But you need to convert []string to []byte to be able to
//     write it to a file using the ioutil.WriteFile.
//
//   + To do that, create a new []byte and append the elements of your
//     []string.
//
// ---------------------------------------------------------

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Send me some items and I will sort them")
		return
	}

	items := os.Args[1:]
	sort.Strings(items)
	text := []byte{}
	for _, item := range items {
		text = append(text, item...)
		text = append(text, "\n"...)
	}
	err := ioutil.WriteFile("sorted.txt", text, 0666)
	if err != nil {
		fmt.Println(err)
	}
}
