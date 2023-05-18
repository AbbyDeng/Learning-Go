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
}
