package main

import (
	"github.com/gocolly/colly/v2"
)

func main() {

	// Initialize Colly
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
		colly.CacheDir("cache"),
		colly.Async(true),
	)

	t20_downloader := c.Clone()
	t20matchresult := T20MatchResult{}

	for _, m := range T20_URLS {
		t20matchresult.DownloadMatchResults(t20_downloader, m)
	}
}
