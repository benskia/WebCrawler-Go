package main

import (
	"fmt"
	"os"
)

func main() {
	// Default configuration limits.
	defaultMaxConcurrency := 10
	defaultMaxPages := 100

	// This webcrawler expects a single URL arg.
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(args) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	maxConcurrency := parseMaxConcurrency(args, defaultMaxConcurrency)
	maxPages := parseMaxPages(args, defaultMaxPages)

	baseURL := args[0]

	config, err := NewConfig(baseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Println("Error creating configuration: ", err)
		os.Exit(1)
	}

	fmt.Println("starting crawl of: ", baseURL)
	config.wg.Add(1)
	config.crawlPage(baseURL)
	config.wg.Wait()

	printReport(config.pages, baseURL)
}
