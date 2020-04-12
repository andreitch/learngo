package api

var temps = []int{5, 10, 3, 25, 45, 80, 90}

// Read returns a slice of elements from the temps slice.
func Read(start, stop int) []int {
	// portion := temps[start:stop]
	portion := temps[start:stop:stop]
	return portion
}

// All returns the temps slice
func All() []int {
	return temps
}
