package main

import (
	"fmt"
	"minigrep/lib"

	"os"
)

func main() {
	args := os.Args

	filename := args[2]
	query := args[1]
	fmt.Printf("Filename: %s\nQuery: %s", filename, query)
	//search

	results := lib.Search(filename, query)

	fmt.Printf("results: %s\n", results)

}

