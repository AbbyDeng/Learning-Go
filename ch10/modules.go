// Package main provides functions for mathematical addition.
package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// ----- 10.1 & 10.2 -----

// Add implements mathematical addition of two integers.
// See [Addition]: https://www.mathsisfun.com/numbers/addition.html
func Add(i int, j int) int {
	return i + j
}

type Number interface {
	constraints.Integer | constraints.Float
}

func AddGeneric[T Number](i T, j T) T {
	return i + j
}

func main() {
	fmt.Println(AddGeneric(5, 19))
}
