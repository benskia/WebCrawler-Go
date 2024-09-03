package main

import (
	"net/url"
)

// Normalizing the URLs will make it easier to compare the links we find while
// crawling a website.
func normalizeURL(rawURL string) (string, error) {
	url, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	hostPath := url.Host + url.Path

	return hostPath, nil
}
