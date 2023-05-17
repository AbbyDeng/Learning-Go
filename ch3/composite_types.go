package main

import "fmt"

func main() {
	fmt.Println("----- 3.1 -----")

	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	subSlice1 := greetings[0:2]
	subSlice2 := greetings[1:4]
	subSlice3 := greetings[3:]

	fmt.Println(greetings)
	fmt.Println(subSlice1)
	fmt.Println(subSlice2)
	fmt.Println(subSlice3)

	fmt.Println("----- 3.2 -----")

	message := "Hi 👩 and 👨"
	fmt.Println(message[3])

	runes := []rune(message)
	fmt.Println(runes[3])
	fmt.Println(string(runes[3]))
}
