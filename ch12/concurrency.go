package main

import (
	"fmt"
	"math"
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

// ----- 12.2 -----

func twoGoroutines() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for i := 10; i < 20; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	count := 2
	for count != 0 {
		select {
		case v, ok := <-ch1:
			if !ok {
				ch1 = nil
				count--
				break
			}
			fmt.Println("from ch1: ", v)
		case v, ok := <-ch2:
			if !ok {
				ch2 = nil
				count--
				break
			}
			fmt.Println("from ch2: ", v)
		}
	}
}

// ----- 12.3 -----

func buildSquareRootMap() map[int]float64 {
	out := make(map[int]float64, 100_000)
	for i := 0; i < 100_000; i++ {
		out[i] = math.Sqrt(float64(i))
	}

	return out
}

var squareRootMapCache = sync.OnceValue(buildSquareRootMap)

func main() {
	// ----- 12.1 -----
	fmt.Println("----- 12.1 -----")
	threeGoroutines()
	fmt.Println()

	// ----- 12.2 -----
	fmt.Println("----- 12.2 -----")
	twoGoroutines()

	// ----- 12.3 -----
	for i := 0; i < 100_000; i += 1_000 {
		fmt.Println(i, " ", squareRootMapCache()[i])
	}
}
