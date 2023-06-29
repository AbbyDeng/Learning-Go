package main

import (
	"fmt"
)

// ----- 6.1 -----

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName string, lastName string, age int) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func MakePersonPointer(firstName string, lastName string, age int) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

// ----- 6.2 -----

func UpdateSlice(slice []string, toReplace string) {
	slice[len(slice)-1] = toReplace
	fmt.Println(slice)
}

func GrowSlice(slice []string, toAppend string) {
	slice = append(slice, toAppend)
	fmt.Println(slice)
}

func main() {
	// ----- 6.1 -----

	fmt.Println("----- 6.1 -----")

	abby := MakePerson("Abby", "Deng", 19)
	pointerAbby := MakePersonPointer("Abby", "Deng", 19)
	fmt.Println(abby)
	fmt.Println(pointerAbby)

	// go build -gcflags="-m"

	// ----- 6.2 -----

	fmt.Println("----- 6.2 -----")

	slice := []string{"abby", "is", "infra's", "most", "chaotic", "kid", "right?"}

	fmt.Println("calling UpdateSlice()")
	fmt.Println(slice)
	UpdateSlice(slice, "indeed!")
	fmt.Println(slice)

	fmt.Println()

	fmt.Println("calling GrowSlice()")
	fmt.Println(slice)
	GrowSlice(slice, "absolutely!")
	fmt.Println(slice)
}
