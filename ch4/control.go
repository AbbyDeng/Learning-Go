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
			fmt.Println(num, " ", "Six!")
		case num%2 == 0:
			fmt.Println(num, " ", "Two!")
		case num%3 == 0:
			fmt.Println(num, " ", "Three!")
		default:
			fmt.Println(num, " ", "Never mind")
		}
	}

	fmt.Println("----- 4.3 -----")

	var total int
	for i := 0; i < 10; i++ {
		total := total + i // this total declared in for shadows the one declared on line 35
		fmt.Println(total)
	}
	fmt.Println(total) // this total is the one declared on line 35, which is 0
}
