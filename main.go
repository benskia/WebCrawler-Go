package main

import (
	"fmt"
	"os"
	"strconv"
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

	for url, count := range config.pages {
		fmt.Printf("%s : %d\n", url, count)
	}
}

func parseMaxConcurrency(args []string, defaultMaxConcurrency int) int {
	if len(args) < 2 {
		fmt.Println("No arg found for max concurrency. Using default concurrency limit: ", defaultMaxConcurrency)
		return defaultMaxConcurrency
	}

	maxConcurrency, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Error converting arg to int: ", err)
		fmt.Println("Using default concurrency limit: ", defaultMaxConcurrency)
		return defaultMaxConcurrency
	}

	return maxConcurrency
}

func parseMaxPages(args []string, defaultMaxPages int) int {
	if len(args) < 3 {
		fmt.Println("No arg found for max pages. Using default page limit: ", defaultMaxPages)
		return defaultMaxPages
	}

	maxPages, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Error converting arg to int: ", err)
		fmt.Println("Using default page limit: ", defaultMaxPages)
		return defaultMaxPages
	}

	return maxPages
}
