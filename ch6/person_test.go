package main

import (
	"testing"
)

func TestMakePersonSlice(t *testing.T) {
	people := make([]Person, 0, 10_000_000)
	for i := 0; i < 10_000_000; i++ {
		people = append(people, Person{
			FirstName: "Abby",
			LastName:  "Deng",
			Age:       19,
		})
	}
}
