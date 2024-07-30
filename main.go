package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args

	filename := args[2]
	query := args[1]
	fmt.Printf("Filename: %s\nQuery: %s", filename, query)
	//search

	results := search(filename, query)

	fmt.Printf("results: %s\n", results)

}

func search(filename string, query string) []string {

	contents, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Failed to read file", err)
	}
	fmt.Printf("Contents: %s\n", contents)
	lines := strings.Split(string(contents), "\n")
	var results []string
	for _, line := range lines {
		if strings.Contains(line, query) {
			results = append(results, line)
		}
	}
	return results
}

func SearchCaseSensitive(filename string, query string) []string {
	query = strings.ToLower(query)
	contents, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Failed to read file", err)
	}
	lines := strings.Split(string(contents), "\n")
	var results []string
	for _, line := range lines {
		if strings.Contains(strings.ToLower(line), query) {
			results = append(results, line)
		}
	}
	return results
}
