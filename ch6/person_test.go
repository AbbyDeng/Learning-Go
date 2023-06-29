package main

import (
	"testing"
)

func TestMakePersonSlice(t *testing.T) {
	people := [10_000_000]Person{}
	for i, _ := range people {
		people[i] = Person{
			FirstName: "Abby",
			LastName:  "Deng",
			Age:       19,
		}
	}
}
