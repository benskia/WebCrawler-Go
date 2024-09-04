package main

import (
	"fmt"
	"strconv"
)

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

	maxPages, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Error converting arg to int: ", err)
		fmt.Println("Using default page limit: ", defaultMaxPages)
		return defaultMaxPages
	}

	return maxPages
}
