package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// ----- 5.1 -----

func add(i int, j int) (int, error) {
	return i + j, nil
}

func sub(i int, j int) (int, error) {
	return i - j, nil
}

func mul(i int, j int) (int, error) {
	return i * j, nil
}

func div(i int, j int) (int, error) {
	if j == 0 {
		return 0, errors.New("division by zero")
	}
	return i / j, nil
}

var opMap = map[string]func(int, int) (int, error){
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

// ----- 5.2 -----

func fileLen(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return 0, err
	}

	return int(stat.Size()), nil
}

// ----- 5.3 -----

func prefixer(prefix string) func(string) string {
	return func(s string) string {
		return prefix + " " + s
	}
}

func main() {
	// ----- 5.1 -----

	fmt.Println("----- 5.1 -----")

	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "/", "0"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression:", expression)
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operator:", op)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}

		// changed
		result, err := opFunc(p1, p2)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(result)
	}

	// ----- 5.2 -----

	fmt.Println("----- 5.2 -----")

	files := []string{
		"functions.go",
		"go.mod",
		"abby.go",
	}

	for _, file := range files {
		length, err := fileLen(file)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(length)
	}

	// ----- 5.3 -----

	fmt.Println("----- 5.3 -----")

	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Abby"))
	fmt.Println(helloPrefix("Gopher"))
}
