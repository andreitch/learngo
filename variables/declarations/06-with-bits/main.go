package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Declare with bits
//
//  1. Declare a few variables using the following types
//    int
//    int8
//    int16
//    int32
//    int64
//    float32
//    float64
//    complex64
//    complex128
//    bool
//    string
//    rune
//    byte
//
// 2. Observe their output
//
// EXPECTED OUTPUT
//  0 0 0 0 0 0 0 (0+0i) (0+0i) false 0 0
//  ""
// ---------------------------------------------------------

func main() {
	var (
		// integer types
		i   int
		i8  int8
		i16 int16
		i32 int32
		i64 int64

		// float types
		f32 float32
		f64 float64

		// complex types
		c64  complex64
		c128 complex128

		// bool type
		b bool

		// string types
		s  string
		r  rune
		by byte
	)

	fmt.Println(
		i, i8, i16, i32, i64,
		f32, f64,
		c64, c128,
		b, r, by,
	)

	fmt.Println(s)
}
