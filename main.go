package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {

	// Initialize Colly
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
		colly.CacheDir("cache"),
	)

	t20_downloader := c.Clone()

	t20matchresult := T20MatchResult{}
	t20matchresult.Download(t20_downloader)

	b, err := t20matchresult.Save("D:\\T_20_2008\\1.json", JSON)
	if err != nil {
		fmt.Println(err)
	}

	if b {
		fmt.Println("File saved successfully!")
	}
}
