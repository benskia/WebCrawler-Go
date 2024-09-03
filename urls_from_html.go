package main

import (
	"log"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	r := strings.NewReader(htmlBody)

	doc, err := html.Parse(r)
	if err != nil {
		return []string{}, err
	}

	urls := []string{}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			var href string
			for _, a := range n.Attr {
				if a.Key == "href" {
					href = a.Val
					break
				}
			}

			u, err := url.Parse(href)
			if err != nil {
				log.Println("Error parsing href to URL: ", err)
				return
			}

			// A blank Host indicates relative URL.
			if u.Host == "" {
				urls = append(urls, rawBaseURL+u.Path)
			} else {
				urls = append(urls, u.Scheme+"://"+u.Host+u.Path)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return urls, nil
}
