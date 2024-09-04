package main

import (
	"log"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	// We're crawling withing a site's internal structure, so any traversable
	// links will have matching domains.
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		log.Println("Error parsing current URL: ", err)
		return
	}

	baseHostname := cfg.baseURL.Hostname()
	currentHostname := currentURL.Hostname()
	if baseHostname != currentHostname {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		log.Printf("Error normalizing URL %s : %v\n", rawCurrentURL, err)
		return
	}

	// We want to track occurrences, but we don't want to re-visit links.
	if cfg.addPageVisit(normalizedURL) {
		return
	}

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		log.Printf("Error getting HTML from %s : %v\n", rawCurrentURL, err)
		return
	}

	urls, err := getURLsFromHTML(htmlBody, baseHostname)
	if err != nil {
		log.Println("Error getting URLs from HTML body : ", err)
		return
	}

	for _, url := range urls {
		cfg.wg.Add(1)
		defer cfg.wg.Done()
		go cfg.crawlPage(url)
	}

	<-cfg.concurrencyControl
}
