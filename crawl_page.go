package main

import (
	"log"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	log.Printf("Crawling...BaseURL: %s - CurrentURL: %s", rawBaseURL, rawCurrentURL)

	// We're crawling withing a site's internal structure, so any traversable
	// links will have matching domains.
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		log.Println("Error parsing base URL: ", err)
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		log.Println("Error parsing current URL: ", err)
	}

	if baseURL.Hostname() != currentURL.Hostname() {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		log.Printf("Error normalizing URL %s : %v\n", rawCurrentURL, err)
		return
	}

	// Part of our tracking involves occurrence metrics.
	if _, ok := pages[normalizedURL]; ok {
		pages[normalizedURL]++
	} else {
		pages[normalizedURL] = 1
	}

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		log.Printf("Error getting HTML from %s : %v\n", rawCurrentURL, err)
		return
	}

	urls, err := getURLsFromHTML(htmlBody, rawBaseURL)
	if err != nil {
		log.Println("Error getting URLs from HTML body : ", err)
		return
	}

	for _, url := range urls {
		crawlPage(rawBaseURL, url, pages)
	}
}
