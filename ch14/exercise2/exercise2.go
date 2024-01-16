package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// ----- 14.2 -----

func main() {
	// ----- 14.2 -----
	fmt.Println("----- 14.2 -----")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	total := 0
	count := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("total:", total, "number of iterations:", count, ctx.Err())
			return
		default:
		}

		num := rand.Intn(100_000_000)
		if num == 1_234 {
			fmt.Println("total:", total, "number of iterations:", count, "got 1,234")
		}
		total += num
		count++
	}
}
