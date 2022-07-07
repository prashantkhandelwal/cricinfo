package main

import (
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

	t20matchresult.ToString()

	// c.OnRequest(func(r *colly.Request) {
	// 	log.Println("visiting", r.URL.String())
	// })

	// c.OnHTML(".engineTable thead", func(e *colly.HTMLElement) {
	// 	e.ForEach("th", func(i int, h *colly.HTMLElement) {
	// 		fmt.Println(h.Text)
	// 	})

	// })

}
