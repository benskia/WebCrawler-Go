package main

import "sort"

type crawlEntry struct {
	URL        string
	Occurrence int
}

func sortCrawlEntries(pages map[string]int) []crawlEntry {
	var entries []crawlEntry

	for url, occurrence := range pages {
		entries = append(entries, crawlEntry{url, occurrence})
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Occurrence > entries[j].Occurrence
	})

	return entries
}
