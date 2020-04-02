package main

import (
	"fmt"
	"runtime"
)

/*
main_function
Prints the number of CPU Cores
*/
func main() {
	fmt.Println(runtime.NumCPU() / 2)
}
