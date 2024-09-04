package main

import (
	"log"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	// To keep the expected workload reasonable, quit once we've reached a
	// predetermined number of pages.
	if len(cfg.pages) >= cfg.maxPages {
		return
	}

	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	// We're crawling withing a site's internal structure, so any traversable
	// links will have matching domains.
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		log.Println("Error parsing current URL: ", err)
		return
	}

	if cfg.baseURL.Hostname() != currentURL.Hostname() {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		log.Printf("Error normalizing URL %s : %v\n", rawCurrentURL, err)
		return
	}

	// We want to track occurrences, but we don't want to re-visit links.
	isFirst := cfg.addPageVisit(normalizedURL)
	if !isFirst {
		return
	}

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		log.Printf("Error getting HTML from %s : %v\n", rawCurrentURL, err)
		return
	}

	urls, err := getURLsFromHTML(htmlBody, cfg.baseURL)
	if err != nil {
		log.Println("Error getting URLs from HTML body : ", err)
		return
	}

	for _, url := range urls {
		cfg.wg.Add(1)
		go cfg.crawlPage(url)
	}
}
