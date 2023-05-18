package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("----- 4.1 -----")

	randoms := make([]int, 100)
	for i := 0; i < 100; i++ {
		randoms[i] = rand.Intn(100)
	}
	fmt.Println(len(randoms))
	fmt.Println(randoms)

	fmt.Println("----- 4.2 -----")

	for _, num := range randoms {
		switch {
		case num%2 == 0 && num%3 == 0:
			fmt.Println("Six!")
		case num%2 == 0:
			fmt.Println("Two!")
		case num%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Never mind")
		}
	}
}
