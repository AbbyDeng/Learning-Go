package main

import (
	"fmt"
	"sync"
)

// ----- 12.1 -----

func threeGoroutines() {
	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	go func() {
		defer wg.Done()
		for i := 10; i < 20; i++ {
			ch <- i
		}
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	var wg2 sync.WaitGroup
	wg2.Add(1)

	go func() {
		defer wg2.Done()
		for i := range ch {
			fmt.Print(i, " ")
		}
	}()

	wg2.Wait()
}

func main() {
	// ----- 12.1 -----
	fmt.Println("----- 12.1 -----")
	threeGoroutines()
}
