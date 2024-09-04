package main

import (
	"log"
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func NewConfig(rawBaseURL string, maxConcurrency int) *config {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		log.Printf("Error parsing %s for config URL: %v", rawBaseURL, err)
		return &config{}
	}

	return &config{
		pages:              make(map[string]int),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
	}
}

func (cfg *config) addPageVisit(normalizedURL string) (exists bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	_, ok := cfg.pages[normalizedURL]

	if ok {
		cfg.pages[normalizedURL]++
	} else {
		cfg.pages[normalizedURL] = 1
	}

	return ok
}

func (cfg *config) crawlPage(rawCurrentURL string) {
	// log.Printf("Crawling...BaseURL: %s - CurrentURL: %s", rawBaseURL, rawCurrentURL)

	// We're crawling withing a site's internal structure, so any traversable
	// links will have matching domains.
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		log.Println("Error parsing current URL: ", err)
	}

	if cfg.baseURL.Hostname() != currentURL.Hostname() {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		log.Printf("Error normalizing URL %s : %v\n", rawCurrentURL, err)
		return
	}

	// Part of our tracking involves occurrence metrics, and we want to avoid
	// re-visiting links.
	if cfg.addPageVisit(normalizedURL) {
		return
	}

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		log.Printf("Error getting HTML from %s : %v\n", rawCurrentURL, err)
		return
	}

	urls, err := getURLsFromHTML(htmlBody, cfg.baseURL.Hostname())
	if err != nil {
		log.Println("Error getting URLs from HTML body : ", err)
		return
	}

	for _, url := range urls {
		cfg.wg.Add(1)
		go func(url string) {
			cfg.concurrencyControl <- struct{}{}
			defer cfg.wg.Done()
			defer func() {
				<-cfg.concurrencyControl
			}()
			cfg.crawlPage(url)
		}(url)
	}
}
