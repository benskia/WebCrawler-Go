package main

import (
	"fmt"
	"os"
)

func main() {
	// This webcrawler expects a single URL arg.
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	fmt.Println("starting crawl of: ", args[0])

	html, err := getHTML(args[0])
	if err != nil {
		fmt.Println("Error getting HTML: ", err)
		os.Exit(1)
	}
	fmt.Print(html)
}
