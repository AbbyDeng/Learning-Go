package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
)

//go:embed english_rights.txt
var englishRights string

//go:embed chinese_rights.txt
var chineseRights string

// ----- 11.1 -----

func main() {
	process(os.Args)
}

func process(args []string) {
	if len(args) == 1 {
		fmt.Println("no language provided")
		return
	}

	switch strings.ToLower(args[1]) {
	case "english":
		fmt.Println(englishRights)
	case "chinese":
		fmt.Println(chineseRights)
	default:
		fmt.Println("unknown language:", os.Args[1])
	}
}
