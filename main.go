package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	filename := args[2]
	query := args[1]
	fmt.Printf("Filename: %s\nQuery: %s", filename, query)

}
