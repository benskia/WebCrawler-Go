package main

import (
	"fmt"
)

func printReport(pages map[string]int, baseURL string) {
	fmt.Print("\n==============================\n\n")
	fmt.Printf("REPORT for %s\n\n", baseURL)
	fmt.Print("==============================\n\n")

	entries := sortCrawlEntries(pages)

	for _, entry := range entries {
		fmt.Printf("Found %d internal links to %s\n", entry.Occurrence, entry.URL)
	}
}
