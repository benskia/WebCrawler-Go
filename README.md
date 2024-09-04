# WebCrawler-Go

A Go CLI tool that crawls a given domain and reports on its link structure.

## Requirements

1. Install GO: `https://go.dev/doc/install`
2. Clone this repo.
3. From project root: `go run . [URL] [Concurrency Limit] [Page Limit]`

Concurrency and page limits are optional args. Respectively, they limit the
maximum live goroutines and number of traversed pages. Typically, a higher
concurrency limit will result in a faster crawl. The page limit is meant to be
a reasonable maximum for the resulting map of traversed URLs.

## TODO

* Make the script run on a timer and deploy it to a server. Have it email you every so often with a report.
* Add more robust error checking so that you can crawl larger sites without issues.
* Count external links, as well as internal links, and add them to the report
* Save the report as a CSV spreadsheet rather than printing it to the console
* Use a graphics library to create an image that shows the links between the pages as a graph visualization
* Make requests concurrently to speed up the crawling process
* Add a README.md file explaining to users how to clone your git repo and get started
