package main

import (
	"log"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	r := strings.NewReader(htmlBody)

	doc, err := html.Parse(r)
	if err != nil {
		return []string{}, err
	}

	var urls []string
	var traverseNodes func(*html.Node)
	traverseNodes = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					href, err := url.Parse(a.Val)
					if err != nil {
						log.Println("Error parsing HREF: ", err)
						continue
					}

					resolvedURL := baseURL.ResolveReference(href)
					urls = append(urls, resolvedURL.String())
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverseNodes(c)
		}
	}
	traverseNodes(doc)

	return urls, nil
}
