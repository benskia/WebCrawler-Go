package main

import (
	"fmt"
)

func printReport(pages map[string]int, baseURL string) {
	fmt.Print("==============================\n\n")
	fmt.Println("REPORT for ", baseURL)
	fmt.Print("==============================\n\n")

	entries := sortCrawlEntries(pages)

	for _, entry := range entries {
		fmt.Printf("Found %d internal links to %s", entry.Occurrence, entry.URL)
	}
}
