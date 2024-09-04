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

	baseURL := args[0]
	fmt.Println("starting crawl of: ", baseURL)

	config, err := NewConfig(baseURL, maxConcurrency)
	if err != nil {
		fmt.Println("Error creating configuration: ", err)
		os.Exit(1)
	}

	config.wg.Add(1)
	config.crawlPage(baseURL)
	config.wg.Wait()

	for url, count := range config.pages {
		fmt.Printf("%s : %d\n", url, count)
	}
}
