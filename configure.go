package main

import (
	"log"
	"net/url"
	"sync"
)

const maxConcurrency int = 1

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

func (cfg *config) addPageVisit(normalizedURL string) bool {
	log.Println("Adding page visit at ", normalizedURL)
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	_, ok := cfg.pages[normalizedURL]

	if ok {
		log.Println("Incrementing ", normalizedURL)
		cfg.pages[normalizedURL]++
	} else {
		log.Println("First visit to ", normalizedURL)
		cfg.pages[normalizedURL] = 1
	}

	return ok
}
