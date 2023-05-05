package main

import "fmt"

const value = 19

func main() {
	i := 20
	f := float64(i)

	fmt.Println(i)
	fmt.Println(f)

	i = value
	f = value

	fmt.Println(i)
	fmt.Println(f)

	var b byte = 255
	var smallI int32 = 2_147_483_647
	var bigI uint64 = 18_446_744_073_709_551_615

	b += 1
	smallI += 1
	bigI += 1

	fmt.Println(b)
	fmt.Println(smallI)
	fmt.Println(bigI)
}
